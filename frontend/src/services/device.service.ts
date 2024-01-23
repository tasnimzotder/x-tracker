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
  });

  console.log(response);

  if (!response || response.status !== 200) {
    return false;
  }

  let data = await response.json();

  return data;
};

export { getDeviceList };
