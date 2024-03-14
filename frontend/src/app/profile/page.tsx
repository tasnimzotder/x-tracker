"use client";

import {
  Button,
  Card,
  Container,
  Flex,
  Group,
  Image,
  Modal,
  Notification,
  Text,
  TextInput,
} from "@mantine/core";
import { useEffect, useState } from "react";
import { usePathname } from "next/navigation";
import { authRequest_t } from "@/services/auth.service";

import { useDisclosure } from "@mantine/hooks";
import { useForm } from "@mantine/form";
import { createDevice, deviceCreateRequest_t } from "@/services/device.service";
import { useAuth } from "@/contexts/authContext";
import { formatTimeToLocaleString } from "@/utils/time.util";

const UpdateProfile = ({
  userData,
  closeProfile,
}: {
  userData: any;
  closeProfile: any;
}) => {
  const [successfull, setSuccessfull] = useState(false);
  const form = useForm({
    initialValues: {
      firstName: userData.firstName,
      lastName: userData.lastName,
      email: userData.email,
      username: userData.username,
      // password: "XXX",
    },

    validate: {
      email: (value) => (/^\S+@\S+$/.test(value) ? null : "Invalid email"),
      // password: (value) =>
      //   value.length < 6 ? "Password must be at least 6 characters" : null,
      username: (value) =>
        value.length < 6 ? "Username must be at least 6 characters" : null,
      firstName: (value) =>
        value.length < 3 ? "First name must be at least 3 characters" : null,
      lastName: (value) =>
        value.length < 3 ? "Last name must be at least 3 characters" : null,
    },
  });

  const handleSubmit = async () => {
    const req: authRequest_t = {
      firstName: form.values.firstName,
      lastName: form.values.lastName,
      email: form.values.email,
      username: form.values.username,
    };

    // await handleProfileUpdate(req).then(() => {
    //   refetchUserData();
    //   // alert("Profile updated successfully");
    //   // form.reset();
    //   setSuccessfull(true);
    //
    //   setTimeout(() => {
    //     setSuccessfull(false);
    //   }, 2000);
    // });
  };

  return (
    <Container>
      <Container>
        {successfull && (
          <Notification title="Profile updated successfully" color="teal" />
        )}
      </Container>
      <form onSubmit={form.onSubmit(handleSubmit)}>
        <Group>
          <TextInput
            type="text"
            label="First Name"
            placeholder="First Name"
            {...form.getInputProps("firstName")}
          />
          <TextInput
            type="text"
            label="Last Name"
            placeholder="Last Name"
            {...form.getInputProps("lastName")}
          />
          <TextInput
            type="email"
            label="Email"
            placeholder="Email"
            {...form.getInputProps("email")}
          />
          <TextInput
            type="text"
            label="Username"
            placeholder="Username"
            {...form.getInputProps("username")}
          />
        </Group>
        <Button type="submit">Update</Button>
      </form>
    </Container>
  );
};

const AddDevice = ({
  userData,
  closeProfile,
}: {
  userData: any;
  closeProfile: any;
}) => {
  const [successfull, setSuccessfull] = useState(false);
  const form = useForm({
    initialValues: {
      deviceName: "",
    },

    validate: {
      deviceName: (value) =>
        value.length < 3 ? "Device name must be at least 3 characters" : null,
    },
  });

  const handleSubmit = async () => {
    const req: deviceCreateRequest_t = {
      deviceName: form.values.deviceName,
      userID: userData.id,
    };

    await createDevice(req).then(() => {
      // refetchUserData();
      // alert("Profile updated successfully");
      // form.reset();
      setSuccessfull(true);

      setTimeout(() => {
        setSuccessfull(false);
      }, 2000);
    });
  };

  return (
    <Container>
      <Container>
        {successfull && (
          <Notification title="Device Added successfully" color="teal" />
        )}
      </Container>
      <form onSubmit={form.onSubmit(handleSubmit)}>
        <Group>
          <TextInput
            type="text"
            label="Device Name"
            placeholder="Device Name"
            {...form.getInputProps("deviceName")}
          />
        </Group>
        <Button type="submit">Update</Button>
      </form>
    </Container>
  );
};

const ProfleUpdateAndAddDeviceButton = () => {
  const [openedProfile, { open: openProfile, close: closeProfile }] =
    useDisclosure(false);
  const [openedDevice, { open: openDevice, close: closeDevice }] =
    useDisclosure(false);

  const [userData, setUserData] = useState<any>(null);

  useEffect(() => {
    // const fetchData = async () => {
    //   let userData = await getUserData();
    //
    //   if (userData == "") {
    //     setUserData(null);
    //   } else {
    //     setUserData(userData);
    //   }
    // };
    //
    // fetchData();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [openedProfile, openedDevice, usePathname().toString()]);

  return (
    <Group>
      <Modal opened={openedProfile} onClose={closeProfile} title="Profile">
        <UpdateProfile userData={userData} closeProfile={closeProfile} />
      </Modal>

      <Modal opened={openedDevice} onClose={closeDevice} title="Device">
        <AddDevice userData={userData} closeProfile={closeDevice} />
      </Modal>

      <Group mt={"25px"}>
        <Button onClick={openProfile}>Update Profile</Button>
        <Button onClick={openDevice}>Add Device</Button>
      </Group>
    </Group>
  );
};

const UserProfilePage = () => {
  const { loggedIn, userData, devices, locations, loadLocations } = useAuth();

  if (!loggedIn || !userData) {
    return (
      <Container>
        <Text>Please log in!</Text>
      </Container>
    );
  }

  return (
    <Container>
      <h1>User Profile</h1>

      <Group>
        <Card mx={"auto"}>
          <Flex direction={"row"} gap={"md"} wrap={"wrap"}>
            <Image
              src={"https://avatars.githubusercontent.com/u/44049528mm?v=4"}
              alt={"user profile picture"}
              width={200}
              height={200}
              radius={"md"}
            />

            <Container>
              <Text>
                Name: {userData.firstName} {userData.lastName}
              </Text>
              <Text>Username: {userData.username}</Text>
              <Text>Email: {userData.email}</Text>
              <Text>
                Last profile updated:{" "}
                {formatTimeToLocaleString(userData.updatedAt)}
              </Text>

              <ProfleUpdateAndAddDeviceButton />
            </Container>
            <Button
              variant={"outline"}
              onClick={async () => {
                // await refetchUserData().then(async () => {
                //   // await fetchData();
                // });
              }}
            >
              Reload
            </Button>
          </Flex>
        </Card>

        <Group>{/*<DevicePage />*/}</Group>
      </Group>
    </Container>
  );
};

export default UserProfilePage;
