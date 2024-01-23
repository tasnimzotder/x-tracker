"use client";

import { useEffect, useState } from "react";
import styles from "./page.module.scss";
import { getLastLocations } from "@/services/location.service";
import BaseMap from "@/components/baseMap/BaseMap";
import DeviceListCmp from "@/components/devices/DeviceList";
import { getUserData } from "@/services/auth.service";
import { getDeviceList } from "@/services/device.service";

export default function Home() {
  const [locations, setLocations] = useState<
    Array<GeoJSON.FeatureCollection<GeoJSON.LineString>>
  >([
    {
      type: "FeatureCollection",
      features: [
        {
          type: "Feature",
          properties: {
            time: "",
            device_id: "",
          },
          geometry: {
            type: "LineString",
            coordinates: [],
          },
        },
      ],
    },
  ]);
  const [devices, setDevices] = useState<Array<any> | any>(null);

  useEffect(() => {
    const fetchData = async () => {
      let userData = await getUserData();

      if (userData == "") {
        setDevices(null);
      } else {
        let devices = await getDeviceList(userData.id);
        // console.log({ devices });

        // sort in asc order
        devices.sort((a: any, b: any) => a.id - b.id);

        await setDevices(devices);
      }
    };

    fetchData();
  }, []);

  const handleClick = async () => {
    // todo: implement
    // const res = await getLastLocations("x_tracker_t1", 10);
    // setLocations(res);

    if (devices.length == 0) {
      return;
    }

    for (let i = 0; i < 1; i++) {
      let device_id = devices[i].id as number;
      console.log({ device_id });

      const res = await getLastLocations(device_id, 10);

      console.log({ res });

      setLocations([res]);
    }
  };

  return (
    <main className={styles.main}>
      {/*<div>xTracker</div>*/}
      {devices && <DeviceListCmp devices={devices} />}

      <button onClick={handleClick}>Click me</button>

      {locations[0].features[0].geometry.coordinates.length > 0 && (
        <div>
          <BaseMap data={locations[0]} />
        </div>
      )}
    </main>
  );
}
