"use client";

import styles from "./page.module.scss";
import { locationData_t } from "@/services/location.service";
import { Button, Card, Container, Grid, Table } from "@mantine/core";
import { useAuth } from "@/contexts/authContext";
import LocationTable from "@/components/locations/LocationTable";

export default function Home() {
  // const [devices, setDevices] = useState<Array<any> | any>(null);
  // const [locations, setLocations] = useState<any>([]);

  const {
    loggedIn,
    userData,
    devices,
    locations,
    loadLocations,
    setLoadLocations,
  } = useAuth();

  const handleLiveLocation = () => {
    if (loggedIn) {
      setLoadLocations(!loadLocations);
    }
  };

  if (!loggedIn) {
    return <div>Please login!</div>;
  }

  return (
    <main className={styles.main}>
      <Container>
        <Button onClick={handleLiveLocation} m={"16px"}>
          {loadLocations ? "Stop Live Location" : "Start Live Location"}
        </Button>

        <Grid>
          <Card>
            <span>User</span>
            {JSON.stringify(userData)}
          </Card>

          <Container>
            <LocationTable />
          </Container>
        </Grid>
      </Container>

      {/*<div>xTracker</div>*/}
      {/*{devices && <DeviceListCmp devices={devices} />}*/}

      {/*/!* <button onClick={handleClick}>Click me</button> *!/*/}

      {/*/!* {locations[0].features[0].geometry.coordinates.length > 0 && (*/}
      {/*  <div>*/}
      {/*    <BaseMap data={locations[0]} />*/}
      {/*  </div>*/}
      {/*)} *!/*/}

      {/*{locations.length > 0 && (*/}
      {/*  <div>{<LocationTable locations={locations[0]} />}</div>*/}
      {/*)}*/}
    </main>
  );
}
