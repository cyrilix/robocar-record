package main

import (
	"flag"
	"github.com/cyrilix/robocar-base/cli"
	"github.com/cyrilix/robocar-record/part"
	"log"
	"os"
)

const (
	DefaultClientId = "robocar-record"
)

func main() {
	var mqttBroker, username, password, clientId string
	var cameraTopic, steeringTopic, recordTopic, switchRecordTopic string

	mqttQos := cli.InitIntFlag("MQTT_QOS", 0)
	_, mqttRetain := os.LookupEnv("MQTT_RETAIN")

	cli.InitMqttFlags(DefaultClientId, &mqttBroker, &username, &password, &clientId, &mqttQos, &mqttRetain)

	flag.StringVar(&cameraTopic, "mqtt-topic-frame", os.Getenv("MQTT_TOPIC_CAMERA"), "Mqtt topic that contains camera frame, use MQTT_TOPIC_CAMERA if args not set")
	flag.StringVar(&steeringTopic, "mqtt-topic-steering", os.Getenv("MQTT_TOPIC_STEERING"), "Mqtt topic that contains steering raw values, use MQTT_TOPIC_STEERING if args not set")
	flag.StringVar(&switchRecordTopic, "mqtt-topic-switch-record", os.Getenv("MQTT_TOPIC_SWITCH_RECORD"), "Mqtt topic that contain switch record state, use MQTT_TOPIC_SWITCH_RECORD if args not set")
	flag.StringVar(&recordTopic, "mqtt-topic-records", os.Getenv("MQTT_TOPIC_RECORDS"), "Mqtt topic to publish record messages, use MQTT_TOPIC_RECORDS if args not set")

	flag.Parse()
	if len(os.Args) <= 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	client, err := cli.Connect(mqttBroker, username, password, clientId)
	if err != nil {
		log.Fatalf("unable to connect to mqtt bus: %v", err)
	}
	defer client.Disconnect(50)

	r := part.NewRecorder(client, recordTopic, cameraTopic, steeringTopic, switchRecordTopic)
	defer r.Stop()

	cli.HandleExit(r)

	if err := r.Start(); err != nil {
		log.Fatalf("unable to start service: %v", err)
	}
}
