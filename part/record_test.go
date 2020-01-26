package part

import (
	"github.com/cyrilix/robocar-base/testtools"
	"github.com/cyrilix/robocar-protobuf/go/events"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"sync"
	"testing"
	"time"
)

func TestRecorder_RecordOff(t *testing.T) {
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
		{recordMsg: &events.SwitchRecordMessage{Enabled: true}, frameMsg: frame1, steeringMsg: &steeringRight, expectedRecordMsg: &events.RecordMessage{Frame: frame1, Steering: &steeringRight}, wait: 5 * time.Millisecond},
		{recordMsg: &events.SwitchRecordMessage{Enabled: true}, frameMsg: nil, steeringMsg: &steeringLeft, expectedRecordMsg: nil, wait: 5 * time.Millisecond},
		{recordMsg: &events.SwitchRecordMessage{Enabled: true}, frameMsg: frame2, steeringMsg: &steeringLeft, expectedRecordMsg: &events.RecordMessage{Frame: frame2, Steering: &steeringLeft}, wait: 5 * time.Millisecond},
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
