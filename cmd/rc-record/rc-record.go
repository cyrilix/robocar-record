package main

import (
	"flag"
	"github.com/cyrilix/robocar-base/cli"
	"github.com/cyrilix/robocar-record/pkg/part"
	"go.uber.org/zap"
	"log"
	"os"
)

const (
	DefaultClientId = "robocar-record"
)

func main() {
	var mqttBroker, username, password, clientId string
	var cameraTopic, rcSteeringTopic, recordTopic, tfSteeringTopic, driveModeTopic, switchRecordTopic string
	var debug bool

	mqttQos := cli.InitIntFlag("MQTT_QOS", 0)
	_, mqttRetain := os.LookupEnv("MQTT_RETAIN")

	cli.InitMqttFlags(DefaultClientId, &mqttBroker, &username, &password, &clientId, &mqttQos, &mqttRetain)

	flag.StringVar(&cameraTopic, "mqtt-topic-frame", os.Getenv("MQTT_TOPIC_CAMERA"), "Mqtt topic that contains camera frame, use MQTT_TOPIC_CAMERA if args not set")
	flag.StringVar(&rcSteeringTopic, "mqtt-topic-steering", os.Getenv("MQTT_TOPIC_STEERING"), "Mqtt topic that contains steering raw values, use MQTT_TOPIC_STEERING if args not set")
	flag.StringVar(&tfSteeringTopic, "mqtt-topic-autopilot-steering", os.Getenv("MQTT_TOPIC_AUTOPILOT_STEERING"), "Mqtt topic that contains steering computed by autopilot, use MQTT_TOPIC_AUTOPILOT_STEERING if args not set")
	flag.StringVar(&driveModeTopic, "mqtt-topic-drive-mode", os.Getenv("MQTT_TOPIC_DRIVE_MODE"), "Mqtt topic that contains drive mode instruction, use MQTT_TOPIC_DRIVE_MODE if args not set")
	flag.StringVar(&switchRecordTopic, "mqtt-topic-switch-record", os.Getenv("MQTT_TOPIC_SWITCH_RECORD"), "Mqtt topic that contain switch record state, use MQTT_TOPIC_SWITCH_RECORD if args not set")
	flag.StringVar(&recordTopic, "mqtt-topic-records", os.Getenv("MQTT_TOPIC_RECORDS"), "Mqtt topic to publish record messages, use MQTT_TOPIC_RECORDS if args not set")
	flag.BoolVar(&debug, "debug", false, "Display raw value to debug")

	flag.Parse()
	if len(os.Args) <= 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	config := zap.NewDevelopmentConfig()
	if debug {
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	} else {
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}
	lgr, err := config.Build()
	if err != nil {
		log.Fatalf("unable to init logger: %v", err)
	}
	defer func() {
		if err := lgr.Sync(); err != nil {
			log.Printf("unable to Sync logger: %v\n", err)
		}
	}()
	zap.ReplaceGlobals(lgr)

	client, err := cli.Connect(mqttBroker, username, password, clientId)
	if err != nil {
		zap.S().Fatalf("unable to connect to mqtt bus: %v", err)
	}
	defer client.Disconnect(50)

	r := part.NewRecorder(client,
		recordTopic,
		cameraTopic,
		rcSteeringTopic,
		tfSteeringTopic,
		driveModeTopic,
		switchRecordTopic,
	)
	defer r.Stop()

	cli.HandleExit(r)

	if err := r.Start(); err != nil {
		zap.S().Fatalf("unable to start service: %v", err)
	}
}
