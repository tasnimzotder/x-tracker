package api

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/tasnimzotder/x-tracker/utils"
	"log"
	"os"
)

type MQTTSOSMessage struct {
	DeviceID           int64 `json:"DeviceID"`
	FallDetected       bool  `json:"FallDetected"`
	PanicButtonPressed bool  `json:"PanicButtonPressed"`
}

func (s *Server) mqttListener() {
	config, err := utils.LoadConfig("./../")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("cannot get current directory: %v", err)
	}

	keyPath := dir + "/certs/key.pem"
	certPath := dir + "/certs/cert.pem"
	caPath := dir + "/certs/ca.pem"

	tlsCert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		log.Fatalf("cannot load key pair: %v", err)
	}

	certs := x509.NewCertPool()
	caPem, err := os.ReadFile(caPath)
	if err != nil {
		log.Fatalf("cannot read ca file: %v", err)
	}

	certs.AppendCertsFromPEM(caPem)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		RootCAs:      certs,
	}

	options := mqtt.NewClientOptions()
	options.AddBroker(fmt.Sprintf("tcps://%s:%s/mqtt", config.MQTTEndpoint, config.MQTTPort))
	options.SetTLSConfig(tlsConfig)
	options.SetClientID(config.MQTTClientID)

	mqttClient := mqtt.NewClient(options)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("error connecting to mqtt: %v", token.Error())
	}

	fmt.Printf("MQTT client connected\n")

	mqttClient.Subscribe("tracker/sos", 0, s.messageHandler)
}

func (s *Server) messageHandler(client mqtt.Client, msg mqtt.Message) {
	var sosMessage MQTTSOSMessage

	err := json.Unmarshal(msg.Payload(), &sosMessage)
	if err != nil {
		log.Fatalf("error unmarshalling mqtt message: %v", err)
	}

	fmt.Printf("received sos message: %v\n", sosMessage)

	if sosMessage.PanicButtonPressed || sosMessage.FallDetected {
		// todo
	}

}
