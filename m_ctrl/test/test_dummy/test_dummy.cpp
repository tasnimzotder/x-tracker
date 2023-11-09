//
// Created by Tasnim Zotder on 07/11/2023.
//
#include <gtest/gtest.h>

#if defined(ARDUINO)

#include <Arduino.h>

#include "./../src/gps/gps.h"

GPS gps;

void setup() {
    Serial.begin(115200);

    ::testing::InitGoogleTest();
}

void loop() {
    if (RUN_ALL_TESTS()) {}


//    check the sum function
//    EXPECT_EQ(gps.sum(1, 2), 3);

    delay(1000);
}

#else

int main(int argc, char **argv)
{
    ::testing::InitGoogleTest(&argc, argv);
    // if you plan to use GMock, replace the line above with
    // ::testing::InitGoogleMock(&argc, argv);

    if (RUN_ALL_TESTS())
    ;

    // Always return zero-code and allow PlatformIO to parse results
    return 0;
}

#endif