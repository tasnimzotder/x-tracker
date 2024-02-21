import { useState } from "react";
// import { getUserData } from "@/services/auth.service";
import styles from "./DeviceList.module.scss";

const DeviceListCmp = (props: { devices: string[] }) => {
  const [devices, setDevices] = useState<Array<any> | any>(props.devices);

  return (
    <div>
      <div>Devices</div>

      <div className={styles.deviceLists}>
        {devices &&
          devices.map((device: any, index: number) => (
            <div key={index}>{device.device_name}</div>
          ))}
      </div>
    </div>
  );
};

export default DeviceListCmp;
