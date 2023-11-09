//
// Created by Tasnim Zotder on 07/11/2023.
//

#include "comm.h"

Communications::Communications() {
    Serial.println("Communications constructor");

    strcpy(this->_ssid, SSID);
    strcpy(this->_password, PASSWORD);
//    strcpy(this->_server, SERVER);
}

void Communications::setup() {
    this->_net.setCACert(AWS_IOT_ROOT_CA);
    this->_net.setCertificate(AWS_IOT_CERTIFICATE);
    this->_net.setPrivateKey(AWS_IOT_PRIVATE_KEY);

    this->_mqttClient.begin(
            AWS_IOT_ENDPOINT,
            8883,
            this->_net
    );
    this->_mqttClient.onMessage(messageHandler);
}


//void Communications::setDeviceID(uint32_t deviceID) {
//    this->_deviceID = deviceID;
//}

esp_err_t Communications::connectToWiFi() {
//    Serial.println("Connecting to WiFi...");
    ESP_LOGI(TAG, "Connecting to WiFi...");

    WiFiClass::mode(WIFI_STA);
    WiFi.begin(this->_ssid, this->_password);

    uint8_t retries = 0;

    while (WiFiClass::status() != WL_CONNECTED) {
        ESP_LOGW(TAG, ". %d", retries + 1);

        if (retries == 10) {
            ESP_LOGE(TAG, "WiFi connection failed!");
            return ESP_FAIL;
        }

        retries++;
        delay(500);
    }

    ESP_LOGI(TAG, "WiFi connected");

    ESP_LOGI(TAG, "IP address: ");
    char ip[16];
    sprintf(ip, "%s", WiFi.localIP().toString().c_str());
    ESP_LOGI(TAG, "%s", ip);

    return ESP_OK;
}

esp_err_t Communications::connectToMQTT() {
    ESP_LOGI(TAG, "Connecting to MQTT...");
    esp_err_t err;

    uint8_t retries = 0;

    while (!this->_mqttClient.connect(AWS_IOT_THING_NAME)) {
        ESP_LOGW(TAG, "MQTT connection failed, retrying... %d", retries + 1);


        if (retries == 20) {
            ESP_LOGE(TAG, "MQTT connection failed!");
            return ESP_FAIL;
        }

        retries++;
        delay(100);
    }

    if (!this->_mqttClient.connected()) {
        ESP_LOGE(TAG, "AWS IoT Timeout!");
        return ESP_FAIL;
    }

    bool res = this->_mqttClient.subscribe(AWS_IOT_SUB_TOPIC);
    if (!res) {
        ESP_LOGE(TAG, "MQTT subscription failed!");
        return ESP_FAIL;
    }

    ESP_LOGI(TAG, "MQTT Connected!");

    return ESP_OK;
}

esp_err_t Communications::publishMQTTData(MQTTUploadData mqttUploadData) {
    esp_err_t err;
    ESP_LOGI(TAG, "Publishing MQTT data...");

    StaticJsonDocument<256> doc;
    doc["latitude"] = mqttUploadData.latitude;
    doc["longitude"] = mqttUploadData.longitude;
//    doc["date"] = mqttUploadData.date;
//    doc["time"] = mqttUploadData.time;
    doc["deviceID"] = mqttUploadData.deviceID;

    String json;
    serializeJson(doc, json);

    bool res = this->_mqttClient.publish(AWS_IOT_PUB_TOPIC, json.c_str());
    if (!res) {
        ESP_LOGE(TAG, "MQTT publish failed!");
        return ESP_FAIL;
    }

    return ESP_OK;
}

void Communications::messageHandler(String &topic, String &payload) {
    Serial.println("incoming: " + topic + " - " + payload);
//    ESP_LOGI(TAG, "%s - %s", topic.c_str(), payload.c_str());
}
