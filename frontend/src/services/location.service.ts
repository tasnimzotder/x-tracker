interface locationService_t {
  device_id: string;
  lat: number;
  lon: number;
  time: string;
}

const getLastLocations = async (device_id: string, limit: number) => {
  const url: string = "http://localhost:8080/v1/locations";

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

export { getLastLocations };

export type { locationService_t };
