package api

import (
	db "backend/db/sqlc"
	"backend/services"
	"backend/utils"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
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
		log.Fatal("cannot load config:", err)
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("error getting working directory: %v", err)
	}

	keyPath := dir + "/certs/key.pem"
	certPath := dir + "/certs/cert.pem"
	caCertPath := dir + "/certs/ca.crt"

	tlsCert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		log.Fatalf("error loading x509 key pair: %v", err)
	}

	certs := x509.NewCertPool()
	caPem, err := os.ReadFile(caCertPath)
	if err != nil {
		log.Fatalf("error reading ca cert: %v", err)
	}

	certs.AppendCertsFromPEM(caPem)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		RootCAs:      certs,
	}

	options := mqtt.NewClientOptions()
	options.AddBroker(fmt.Sprintf("tcps://%s:%s/mqtt", config.MQTTEndpoint, config.MQTTPort))
	options.SetClientID(config.MQTTClientID)
	options.SetTLSConfig(tlsConfig)

	mqttClient := mqtt.NewClient(options)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("error connecting to mqtt: %v", token.Error())
	}

	fmt.Printf("connected to mqtt broker\n")

	mqttClient.Subscribe("tracker/sos", 0, s.messageHandler)
}

func (s *Server) messageHandler(client mqtt.Client, message mqtt.Message) {
	var msg MQTTSOSMessage

	err := json.Unmarshal(message.Payload(), &msg)
	if err != nil {
		fmt.Printf("error unmarshalling message: %v\n", err)
		return
	}

	fmt.Printf("received message: %v\n", msg)

	if msg.PanicButtonPressed || msg.FallDetected {
		// todo
		lastLocation, err := s.GetLastLocation()
		if err != nil {
			fmt.Printf("error getting last location: %v\n", err)
			return
		}

		fmt.Printf("last location: %v\n", lastLocation)

		arg := db.CreateDeviceActivityParams{
			DeviceID: msg.DeviceID,
			Panic:    msg.PanicButtonPressed,
			Fall:     msg.FallDetected,
		}

		activity, err := s.queries.CreateDeviceActivity(context.Background(), arg)
		if err != nil {
			fmt.Printf("error creating device activity: %v\n", err)
			return
		}

		fmt.Printf("created device activity: %v\n", activity)

		var message string

		if msg.PanicButtonPressed {
			// multiple messages
			message = `🆘 PANIC BUTTON PRESSED 🆘`
			message += ` - `
			message += `Device ID: ` + fmt.Sprintf("%v", msg.DeviceID) + ` - `
			message += `Time: ` + fmt.Sprintf("%v", activity.CreatedAt) + ` - `
			message += fmt.Sprintf("Last Location: https://www.google.com/maps/search/?api=1&query=%v,%v", lastLocation.Latitude, lastLocation.Longitude)
		}

		services.SendMessages(message)
	}

	//fmt.Printf("received message: %s from topic: %s\n", message.Payload(), message.Topic())
}
