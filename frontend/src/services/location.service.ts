interface locationService_t {
  device_id: string;
  lat: number;
  lon: number;
  time: string;
}

const getLastLocations = async (device_id: number, limit: number) => {
  const url: string = `${process.env.NEXT_PUBLIC_API_URL}/v1/locations/get`;

  const reqBody = {
    device_id: device_id,
    limit: limit,
  };

  const reqOptions = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(reqBody),
  };

  const response = (await fetch(url, reqOptions).catch((err) => {
    console.error(err);
  })) as Response;

  if (response.status !== 200) {
    throw new Error("Failed to fetch locations");
  }

  let data =
    (await response.json()) as GeoJSON.FeatureCollection<GeoJSON.LineString>;

  console.log({ data });

  return data;
};

type location_t = {
  lat: number;
  lon: number;
  timestamp: string;
};

type locationData_t = {
  device_id: number;
  locations: location_t[];
};

const getLocationsByUserID = async (
  ws: WebSocket,
  user_id: number,
  device_id: number,
) => {
  ws = new WebSocket(`ws://${process.env.NEXT_PUBLIC_API_URL}/v1/ws/location`);

  ws.onopen = () => {
    console.log("Connected to WS");
    ws.send(
      JSON.stringify({
        user_id: user_id,
        device_id: device_id,
      }),
    );
  };

  ws.onmessage = (event) => {
    // console.log("Message received", event);
  };

  ws.onerror = (event) => {
    console.error("Error", event);
  };

  ws.onclose = (event) => {
    console.log("Closed", event);
  };
};

export { getLastLocations, getLocationsByUserID };

export type { locationService_t, location_t, locationData_t };
