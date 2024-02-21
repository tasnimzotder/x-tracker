import { useAuth } from "@/contexts/authContext";
import { Container, Table } from "@mantine/core";
import { formatTimeToLocaleString } from "@/utils/time.util";

const LocationTable = () => {
  const { loggedIn, locations } = useAuth();

  console.log({ loggedIn });
  console.log({ locations });

  if (!(loggedIn && locations && locations.length != 0)) {
    return <Container>No locations found</Container>;
  }

  return (
    <Container>
      <Table>
        <Table.Thead>
          <Table.Th>Device ID</Table.Th>
          <Table.Th>Lat</Table.Th>
          <Table.Th>Lng</Table.Th>
          <Table.Th>Timestamp</Table.Th>
        </Table.Thead>
        <Table.Tbody>
          {locations &&
            locations.map((location: any, index: number) => (
              <Table.Tr key={index}>
                <Table.Td>{location.device_id}</Table.Td>
                <Table.Td>{location.lat}</Table.Td>
                <Table.Td>{location.lng}</Table.Td>
                <Table.Td>
                  {formatTimeToLocaleString(location.timestamp)}
                </Table.Td>
              </Table.Tr>
            ))}
        </Table.Tbody>
      </Table>

      {/* {JSON.stringify({ locations })} */}
    </Container>
  );
};

export default LocationTable;
