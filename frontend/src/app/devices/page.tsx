"use client";
import { useEffect, useState } from "react";
import { getUserData } from "@/services/auth.service";
import { getDeviceList } from "@/services/device.service";
import { formatTimeToLocaleString } from "@/utils/time.util";
import { Container, Table } from "@mantine/core";

const DevicePage = () => {
  const [devices, setDevices] = useState<Array<any> | any>(null);

  useEffect(() => {
    const fetchData = async () => {
      let userData = await getUserData();

      if (userData == "") {
        setDevices(null);
      } else {
        let devices = await getDeviceList(userData.id);
        console.log({ devices });

        await setDevices(devices);
      }
    };

    fetchData();
  }, []);

  return (
    <Container>
      <h1>Device List</h1>
      {devices && (
        <Table miw={700} maw={900} m={"auto"}>
          <Table.Thead>
            <Table.Tr>
              <Table.Th>ID</Table.Th>
              <Table.Th>Device Name</Table.Th>
              <Table.Th>Device Type</Table.Th>
              <Table.Th>Status</Table.Th>
              <Table.Th>Created At</Table.Th>
              <Table.Th>Last Seen</Table.Th>
            </Table.Tr>
          </Table.Thead>
          <Table.Tbody>
            {devices.map((device: any) => (
              <Table.Tr key={device.id}>
                <Table.Td>{device.id}</Table.Td>
                <Table.Td>{device.device_name}</Table.Td>
                <Table.Td>{device.device_type}</Table.Td>
                <Table.Td>{device.status}</Table.Td>
                {/*<Table.Td>{device.created_at}</Table.Td>*/}
                {/*<Table.Td>{device.last_seen}</Table.Td>*/}
                <Table.Td>
                  {formatTimeToLocaleString(device?.created_at as string)}
                </Table.Td>
                <Table.Td>
                  {formatTimeToLocaleString(device?.last_seen)}
                </Table.Td>
              </Table.Tr>
            ))}
          </Table.Tbody>
        </Table>
      )}
    </Container>
  );
};

export default DevicePage;
