import styles from "./DevicePinMap.module.scss";

type DevicePinMap_t = {
  DeviceName?: string;
  DeviceID?: number;
  DeviceType?: string;
};

const DevicePinMap = (props: DevicePinMap_t) => {
  const { DeviceName, DeviceID, DeviceType } = props;

  return (
    <div className={styles.device_pin}>
      {/*{DeviceName && <div>{DeviceName}</div>}*/}
      {DeviceID && <div>{DeviceID}</div>}
      {/*{DeviceType && <div>{DeviceType}</div>}*/}
    </div>
  );
};

export default DevicePinMap;
