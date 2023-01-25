package part

import (
	"fmt"
	"github.com/cyrilix/robocar-base/service"
	"github.com/cyrilix/robocar-protobuf/go/events"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"sync"
	"time"
)

func NewRecorder(client mqtt.Client, recordTopic, cameraTopic, rcSteeringTopic, tfSteeringTopic, driveModeTopic,
	switchRecordTopic string) *Recorder {
	return &Recorder{
		client:            client,
		recordTopic:       recordTopic,
		cameraTopic:       cameraTopic,
		rcSteeringTopic:   rcSteeringTopic,
		tfSteeringTopic:   tfSteeringTopic,
		driveModeTopic:    driveModeTopic,
		switchRecordTopic: switchRecordTopic,
		enabled:           false,
		idGenerator:       NewDateBasedGenerator(),
		recordSet:         "",
		cancel:            make(chan interface{}),
	}
}

type Recorder struct {
	client                         mqtt.Client
	recordTopic                    string
	cameraTopic, switchRecordTopic string

	driveModeTopic, rcSteeringTopic, tfSteeringTopic string

	muRcSteeringMsg   sync.Mutex
	currentRcSteering *events.SteeringMessage

	muTfSteeringMsg   sync.Mutex
	currentTfSteering *events.SteeringMessage

	muDriveModeMsg   sync.Mutex
	currentDriveMode *events.DriveModeMessage

	muEnabled sync.RWMutex
	enabled   bool

	idGenerator IdGenerator
	recordSet   string

	cancel chan interface {
	}
}

func (r *Recorder) Start() error {
	registerCallBacks(r)

	for {
		select {
		case <-r.cancel:
			zap.S().Info("Stop service")
			return nil
		}
	}
}

func (r *Recorder) Stop() {
	close(r.cancel)
	service.StopService("record", r.client, r.cameraTopic, r.rcSteeringTopic, r.tfSteeringTopic, r.driveModeTopic)
}

func (r *Recorder) onSwitchRecord(_ mqtt.Client, message mqtt.Message) {
	var msg events.SwitchRecordMessage
	err := proto.Unmarshal(message.Payload(), &msg)
	if err != nil {
		zap.S().Errorf("unable to unmarshal protobuf %T: %v", msg, err)
		return
	}

	r.muEnabled.Lock()
	defer r.muEnabled.Unlock()

	if !r.enabled && msg.GetEnabled() {
		r.recordSet = r.idGenerator.Next()
	}

	r.enabled = msg.GetEnabled()
}

func (r *Recorder) onRcSteering(_ mqtt.Client, message mqtt.Message) {
	var msg events.SteeringMessage
	err := proto.Unmarshal(message.Payload(), &msg)
	if err != nil {
		zap.S().Errorf("unable to unmarshal protobuf %T: %v", msg, err)
		return
	}

	r.muRcSteeringMsg.Lock()
	defer r.muRcSteeringMsg.Unlock()
	r.currentRcSteering = &msg
}

func (r *Recorder) onTfSteering(_ mqtt.Client, message mqtt.Message) {
	var msg events.SteeringMessage
	err := proto.Unmarshal(message.Payload(), &msg)
	if err != nil {
		zap.S().Errorf("unable to unmarshal protobuf %T: %v", msg, err)
		return
	}

	r.muTfSteeringMsg.Lock()
	defer r.muTfSteeringMsg.Unlock()
	r.currentTfSteering = &msg
}

func (r *Recorder) onDriveMode(_ mqtt.Client, message mqtt.Message) {
	var msg events.DriveModeMessage
	err := proto.Unmarshal(message.Payload(), &msg)
	if err != nil {
		zap.S().Errorf("unable to unmarshal protobuf %T: %v", msg, err)
		return
	}

	r.muDriveModeMsg.Lock()
	defer r.muDriveModeMsg.Unlock()
	r.currentDriveMode = &msg
}

func (r *Recorder) onFrame(_ mqtt.Client, message mqtt.Message) {
	if !r.Enabled() {
		return
	}

	var msg events.FrameMessage
	err := proto.Unmarshal(message.Payload(), &msg)
	if err != nil {
		zap.S().Errorf("unable to unmarshal protobuf FrameMessage: %v", err)
		return
	}

	steering := r.CurrentSteering()
	if steering == nil {
		zap.S().Warnf("no current steeringMsg, skip frameMsg %v", msg.GetId().Id)
		return
	}

	autopilot := r.CurrentAutopilotSteering()
	if autopilot == nil {
		zap.S().Warnf("no current autopilot steeringMsg")
	}

	driveMode := r.CurrentDriveMode()
	if driveMode == nil {
		zap.S().Warnf("no current driveModeMsg")
	}

	record := events.RecordMessage{
		Frame:             &msg,
		Steering:          steering,
		AutopilotSteering: autopilot,
		DriveMode:         driveMode,
		RecordSet:         r.recordSet,
	}

	payload, err := proto.Marshal(&record)
	if err != nil {
		zap.S().Errorf("unable to marshal message %v: %v", record, err)
		return
	}
	publish(r.client, r.recordTopic, &payload)
}

var publish = func(client mqtt.Client, topic string, payload *[]byte) {
	client.Publish(topic, 0, false, *payload)
}

func (r *Recorder) CurrentSteering() *events.SteeringMessage {
	r.muRcSteeringMsg.Lock()
	defer r.muRcSteeringMsg.Unlock()
	steering := r.currentRcSteering
	return steering
}

func (r *Recorder) CurrentAutopilotSteering() *events.SteeringMessage {
	r.muTfSteeringMsg.Lock()
	defer r.muTfSteeringMsg.Unlock()
	steering := r.currentTfSteering
	return steering
}

func (r *Recorder) CurrentDriveMode() *events.DriveModeMessage {
	r.muDriveModeMsg.Lock()
	defer r.muDriveModeMsg.Unlock()
	driveMode := r.currentDriveMode
	return driveMode
}

func (r *Recorder) Enabled() bool {
	r.muEnabled.RLock()
	defer r.muEnabled.RUnlock()
	return r.enabled
}

var registerCallBacks = func(r *Recorder) {
	err := service.RegisterCallback(r.client, r.cameraTopic, r.onFrame)
	if err != nil {
		zap.S().Panicf("unable to register callback to %v:%v", r.cameraTopic, err)
	}

	err = service.RegisterCallback(r.client, r.rcSteeringTopic, r.onRcSteering)
	if err != nil {
		zap.S().Panicf("unable to register callback to %v:%v", r.rcSteeringTopic, err)
	}

	err = service.RegisterCallback(r.client, r.tfSteeringTopic, r.onTfSteering)
	if err != nil {
		zap.S().Panicf("unable to register callback to %v:%v", r.tfSteeringTopic, err)
	}

	err = service.RegisterCallback(r.client, r.driveModeTopic, r.onDriveMode)
	if err != nil {
		zap.S().Panicf("unable to register callback to %v:%v", r.driveModeTopic, err)
	}

	err = service.RegisterCallback(r.client, r.switchRecordTopic, r.onSwitchRecord)
	if err != nil {
		zap.S().Panicf("unable to register callback to %v:%v", r.switchRecordTopic, err)
	}
}

type IdGenerator interface {
	Next() string
}

func NewDateBasedGenerator() *DateBasedGenerator {
	return &DateBasedGenerator{
		muCpt:      sync.Mutex{},
		cpt:        0,
		idTemplate: "%s_%d",
		start:      time.Now().Format("2006-01-02T15-04"),
	}
}

type DateBasedGenerator struct {
	muCpt      sync.Mutex
	cpt        int
	idTemplate string
	start      string
}

func (d *DateBasedGenerator) Next() string {
	d.muCpt.Lock()
	defer d.muCpt.Unlock()
	d.cpt += 1
	return fmt.Sprintf(d.idTemplate, d.start, d.cpt)
}
