package utils

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"io/ioutil"
	"log"
	"time"
)

func newTLSConfig() (config *tls.Config, err error) {
	certsPath := GetEnvVariable("CERTS_PATH")

	certPool := x509.NewCertPool()
	//permCerts, err := ioutil.ReadFile("certs/AmazonRootCA1.pem")
	permCerts, err := ioutil.ReadFile(certsPath + "/AmazonRootCA1.pem")
	if err != nil {
		log.Fatal(err)
	}

	certPool.AppendCertsFromPEM(permCerts)

	//cert, err := tls.LoadX509KeyPair("certs/certificate.pem.crt", "certs/private.pem.key")
	cert, err := tls.LoadX509KeyPair(certsPath+"/certificate.pem.crt", certsPath+"/private.pem.key")
	if err != nil {
		log.Fatal(err)
	}

	config = &tls.Config{
		RootCAs:      certPool,
		Certificates: []tls.Certificate{cert},
		// ClientCAs:    nil,
		// ClientAuth:   tls.NoClientCert,
	}

	return
}

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func GetMQTTopts() *mqtt.ClientOptions {
	tlsConfig, err := newTLSConfig()
	if err != nil {
		log.Fatalf("Failed to create TLS config for MQTT: %v", err)
	}

	mqttOpts := mqtt.NewClientOptions()
	mqttOpts.AddBroker(fmt.Sprintf(
		"tls://%s:%s",
		GetEnvVariable("MQTT_ENDPOINT"),
		GetEnvVariable("MQTT_PORT"),
	))
	mqttOpts.SetClientID(GetEnvVariable("MQTT_CLIENT_ID")).SetTLSConfig(tlsConfig)
	mqttOpts.SetDefaultPublishHandler(f)
	mqttOpts.SetKeepAlive(30 * time.Second)

	return mqttOpts
}
