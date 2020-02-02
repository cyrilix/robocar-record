package part

import (
	"github.com/cyrilix/robocar-base/testtools"
	"github.com/cyrilix/robocar-protobuf/go/events"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"regexp"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestRecorder_Record(t *testing.T) {
	oldRegister := registerCallBacks
	oldPublish := publish
	defer func() {
		registerCallBacks = oldRegister
		publish = oldPublish
	}()

	registerCallBacks = func(_ *Recorder) {}

	recordTopic := "topic/record"
	cameraTopic := "topic/camera"
	steeringTopic := "topic/steeringMsg"
	switchRecord := "topic/switch/record"

	var muEventsPublished sync.Mutex
	var eventsPublished *events.RecordMessage
	publish = func(client mqtt.Client, topic string, payload *[]byte) {
		if topic != recordTopic {
			t.Errorf("event published on bad topic: %v, wants %v", topic, recordTopic)
			return
		}
		muEventsPublished.Lock()
		defer muEventsPublished.Unlock()
		var msg events.RecordMessage
		err := proto.Unmarshal(*payload, &msg)
		if err != nil {
			t.Errorf("unable to record plublished event: %v", err)
		}
		eventsPublished = &msg
	}

	recorder := NewRecorder(nil, recordTopic, cameraTopic, steeringTopic, switchRecord)
	recorder.idGenerator = &DateBasedGenerator{
		muCpt:      sync.Mutex{},
		cpt:        0,
		idTemplate: "%s-%d",
		start:      "record",
	}


	go func() {
		if err := recorder.Start(); err == nil {
			t.Fatalf("unable to start recorder: %v", err)
		}
	}()

	frame1 := loadImage(t, "testdata/img.jpg", "01")
	frame2 := loadImage(t, "testdata/img.jpg", "02")
	steeringRight := events.SteeringMessage{Steering: 0.5, Confidence: 1.0}
	steeringLeft := events.SteeringMessage{Steering: -0.5, Confidence: 1.0}

	cases := []struct {
		recordMsg         *events.SwitchRecordMessage
		frameMsg          *events.FrameMessage
		steeringMsg       *events.SteeringMessage
		expectedRecordMsg *events.RecordMessage
		wait              time.Duration
	}{
		{recordMsg: &events.SwitchRecordMessage{Enabled: false}, frameMsg: nil, steeringMsg: nil, expectedRecordMsg: nil, wait: 5 * time.Millisecond},
		{recordMsg: &events.SwitchRecordMessage{Enabled: true}, frameMsg: nil, steeringMsg: nil, expectedRecordMsg: nil, wait: 5 * time.Millisecond},
		{recordMsg: &events.SwitchRecordMessage{Enabled: true}, frameMsg: frame1, steeringMsg: nil, expectedRecordMsg: nil, wait: 5 * time.Millisecond},
		{recordMsg: &events.SwitchRecordMessage{Enabled: true}, frameMsg: nil, steeringMsg: &steeringRight, expectedRecordMsg: nil, wait: 5 * time.Millisecond},
		{recordMsg: &events.SwitchRecordMessage{Enabled: true}, frameMsg: frame1, steeringMsg: &steeringRight, expectedRecordMsg: &events.RecordMessage{RecordSet: "record-1", Frame: frame1, Steering: &steeringRight}, wait: 5 * time.Millisecond},
		{recordMsg: &events.SwitchRecordMessage{Enabled: true}, frameMsg: nil, steeringMsg: &steeringLeft, expectedRecordMsg: nil, wait: 5 * time.Millisecond},
		{recordMsg: &events.SwitchRecordMessage{Enabled: true}, frameMsg: frame2, steeringMsg: &steeringLeft, expectedRecordMsg: &events.RecordMessage{RecordSet: "record-1", Frame: frame2, Steering: &steeringLeft}, wait: 5 * time.Millisecond},
		{recordMsg: &events.SwitchRecordMessage{Enabled: false}, frameMsg: nil, steeringMsg: nil, expectedRecordMsg: nil, wait: 5 * time.Millisecond},
		{recordMsg: &events.SwitchRecordMessage{Enabled: false}, frameMsg: nil, steeringMsg: nil, expectedRecordMsg: nil, wait: 5 * time.Millisecond},
		{recordMsg: &events.SwitchRecordMessage{Enabled: true}, frameMsg: frame1, steeringMsg: &steeringRight, expectedRecordMsg: &events.RecordMessage{RecordSet: "record-2", Frame: frame1, Steering: &steeringLeft}, wait: 5 * time.Millisecond},
	}

	for _, c := range cases {
		muEventsPublished.Lock()
		eventsPublished = nil
		muEventsPublished.Unlock()

		if c.recordMsg != nil {
			recorder.onSwitchRecord(nil, testtools.NewFakeMessageFromProtobuf(recordTopic, c.recordMsg))
		}
		if c.frameMsg != nil {
			recorder.onFrame(nil, testtools.NewFakeMessageFromProtobuf(cameraTopic, c.frameMsg))
		}
		if c.steeringMsg != nil {
			recorder.onSteering(nil, testtools.NewFakeMessageFromProtobuf(steeringTopic, c.steeringMsg))
		}

		time.Sleep(c.wait)

		if c.expectedRecordMsg == nil && eventsPublished != nil {
			t.Errorf("unexpected published event: %v", eventsPublished)
		}
		if c.expectedRecordMsg.String() != eventsPublished.String() {
			t.Errorf("bad message published: %v, wants %v", eventsPublished, c.expectedRecordMsg)
		}
	}

}

func loadImage(t *testing.T, imgPath string, id string) *events.FrameMessage {
	jpegContent, err := ioutil.ReadFile(imgPath)
	if err != nil {
		t.Fatalf("unable to load image: %v", err)
	}
	msg := &events.FrameMessage{
		Id: &events.FrameRef{
			Name: imgPath,
			Id:   id,
		},
		Frame: jpegContent,
	}
	return msg
}

func TestDateBasedGenerator_Next(t *testing.T) {
	log.SetLevel(log.InfoLevel)

	expectedFmt := "[0-9]{4}[01][0-9][0-2][0-9][0-5][0-9]-[0-9]+"
	r, err := regexp.Compile(expectedFmt)
	if err != nil {
		t.Fatalf("unable to compile expected regex: %v", err)
	}
	d := NewDateBasedGenerator()
	id1 := d.Next()
	log.Debugf("first id: %v", id1)
	if !r.MatchString(id1) {
		t.Errorf("Unexpected id format: %v, wants: %s", id1, expectedFmt)
	}

	id2 := d.Next()
	log.Debugf("2nd id: %v", id2)

	if strings.Split(id1, "-")[0] != strings.Split(id2, "-")[0] {
		t.Errorf("ids are differentt prefixes: %v - %v", strings.Split(id1, "-")[0], strings.Split(id2, "-")[0])
	}

	if strings.Split(id1, "-")[1] != "1" {
		t.Errorf("unexpected suffix: %v, wants %v", strings.Split(id1, "-")[1], "1")
	}
	if strings.Split(id2, "-")[1] != "2" {
		t.Errorf("unexpected suffix: %v, wants %v", strings.Split(id2, "-")[1], "2")
	}
}
