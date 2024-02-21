"use client";
import { Container, Table } from "@mantine/core";
import { useAuth } from "@/contexts/authContext";
import { formatTimeToLocaleString } from "@/utils/time.util";

const DevicePage = () => {
  const { loggedIn, devices } = useAuth();

  if (!loggedIn) {
    return <div>You are not logged in</div>;
  }

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
