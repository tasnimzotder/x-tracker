//
// Created by Tasnim Zotder on 07/11/2023.
//

#ifndef X_TRACKER_ALPHA_COMM_H
#define X_TRACKER_ALPHA_COMM_H

#include <Arduino.h>
#include "secrets.h"
#include "esp_log.h"
#include <WiFi.h>
#include <MQTTClient.h>
#include <ArduinoJson.h>
#include <WiFiClientSecure.h>


struct MQTTUploadData {
    double latitude;
    double longitude;
    uint32_t date;
    uint32_t time;
    const char *deviceID;
};


class Communications {
private:
    char _ssid[32]{};
    char _password[32]{};
    char _server[32]{};
    int _port = 8883;

    uint32_t _deviceID{};

    WiFiClientSecure _net = WiFiClientSecure();
    MQTTClient _mqttClient = MQTTClient(256);

    const char *TAG = "COMM";

public:
    Communications();

    void setup();

//    void setDeviceID(uint32_t deviceID);

    esp_err_t connectToWiFi();

    esp_err_t connectToMQTT();

    esp_err_t publishMQTTData(MQTTUploadData mqttUploadData);

    static void messageHandler(String &topic, String &payload);

};


#endif //X_TRACKER_ALPHA_COMM_H
