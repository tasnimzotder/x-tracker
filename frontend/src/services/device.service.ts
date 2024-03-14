"use server";

const getDeviceList = async (user_id: number) => {
  const url: string = `${process.env.NEXT_PUBLIC_API_URL}/v1/devices/user/${user_id}`;

  const reqOptions = {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  };

  const response = await fetch(url, reqOptions).catch((err) => {
    console.error(err);
    return null;
  });

  if (!response || response.status !== 200) {
    return null;
  }

  return await response.json();
};

interface deviceCreateRequest_t {
  deviceName: string;
  userID: number;
}

const createDevice = async (req: deviceCreateRequest_t) => {
  const url: string = `${process.env.NEXT_PUBLIC_API_URL}/v1/devices/create`;

  const reqOptions = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(req),
  };

  const response = await fetch(url, reqOptions).catch((err) => {
    console.error(err);
  });

  // console.log(response);

  if (!response || response.status !== 200) {
    return false;
  }

  let data = await response.json();

  return data;
};

export { getDeviceList, createDevice };

export type { deviceCreateRequest_t };
