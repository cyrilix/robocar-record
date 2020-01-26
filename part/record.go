package part

import (
	"github.com/cyrilix/robocar-base/service"
	"github.com/cyrilix/robocar-protobuf/go/events"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"sync"
)

func NewRecorder(client mqtt.Client, recordTopic, cameraTopic, steeringTopic, switchRecordTopic string) *Recorder {
	return &Recorder{
		client:        client,
		recordTopic:   recordTopic,
		cameraTopic:   cameraTopic,
		steeringTopic: steeringTopic,
		switchRecordTopic: switchRecordTopic,
		enabled:       false,
		cancel:        make(chan interface{}),
	}
}

type Recorder struct {
	client                     mqtt.Client
	recordTopic                string
	cameraTopic, steeringTopic, switchRecordTopic string

	muSteeringMsg   sync.Mutex
	currentSteering *events.SteeringMessage

	muEnabled sync.RWMutex
	enabled   bool

	cancel chan interface {
	}
}

func (r *Recorder) Start() error {
	registerCallBacks(r)

	for {
		select {
		case <-r.cancel:
			log.Infof("Stop service")
			return nil
		}
	}
}

func (r *Recorder) Stop() {
	close(r.cancel)
	service.StopService("record", r.client, r.cameraTopic, r.steeringTopic)
}

func (r *Recorder) onSwitchRecord(_ mqtt.Client, message mqtt.Message) {
	var msg events.SwitchRecordMessage
	err := proto.Unmarshal(message.Payload(), &msg)
	if err != nil {
		log.Errorf("unable to unmarshal protobuf %T: %v", msg, err)
		return
	}

	r.muEnabled.Lock()
	defer r.muEnabled.Unlock()
	r.enabled = msg.GetEnabled()
}

func (r *Recorder) onSteering(_ mqtt.Client, message mqtt.Message) {
	var msg events.SteeringMessage
	err := proto.Unmarshal(message.Payload(), &msg)
	if err != nil {
		log.Errorf("unable to unmarshal protobuf %T: %v", msg, err)
		return
	}

	r.muSteeringMsg.Lock()
	defer r.muSteeringMsg.Unlock()
	r.currentSteering = &msg
}

func (r *Recorder) onFrame(_ mqtt.Client, message mqtt.Message) {
	if !r.Enabled() {
		return
	}

	var msg events.FrameMessage
	err := proto.Unmarshal(message.Payload(), &msg)
	if err != nil {
		log.Errorf("unable to unmarshal protobuf FrameMessage: %v", err)
		return
	}

	steering := r.CurrentSteering()
	if steering == nil {
		log.Warningf("no current steeringMsg, skip frameMsg %v", msg.GetId().Id)
		return
	}

	record := events.RecordMessage{
		Frame:    &msg,
		Steering: steering,
	}

	payload, err := proto.Marshal(&record)
	if err != nil {
		log.Errorf("unable to marshal message %v: %v", record, err)
		return
	}
	publish(r.client, r.recordTopic, &payload)
}

var publish = func(client mqtt.Client, topic string, payload *[]byte) {
	client.Publish(topic, 0, false, *payload)
}

func (r *Recorder) CurrentSteering() *events.SteeringMessage {
	r.muSteeringMsg.Lock()
	defer r.muSteeringMsg.Unlock()
	steering := r.currentSteering
	return steering
}

func (r *Recorder) Enabled() bool {
	r.muEnabled.RLock()
	defer r.muEnabled.RUnlock()
	return r.enabled
}

var registerCallBacks = func(r *Recorder) {
	err := service.RegisterCallback(r.client, r.cameraTopic, r.onFrame)
	if err != nil {
		log.Panicf("unable to register callback to %v:%v", r.cameraTopic, err)
	}

	err = service.RegisterCallback(r.client, r.steeringTopic, r.onSteering)
	if err != nil {
		log.Panicf("unable to register callback to %v:%v", r.steeringTopic, err)
	}

	err = service.RegisterCallback(r.client, r.switchRecordTopic, r.onSwitchRecord)
	if err != nil {
		log.Panicf("unable to register callback to %v:%v", r.switchRecordTopic, err)
	}
}
