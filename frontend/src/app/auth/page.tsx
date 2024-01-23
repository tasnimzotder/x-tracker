"use client";

import { useToggle, upperFirst } from "@mantine/hooks";
import { useForm } from "@mantine/form";
import {
  TextInput,
  PasswordInput,
  Text,
  Paper,
  Group,
  PaperProps,
  Button,
  Divider,
  Checkbox,
  Anchor,
  Stack,
} from "@mantine/core";
import { authRequest_t, handleLogin } from "@/services/auth.service";
import { useRouter } from "next/navigation";
// import { GoogleButton } from "./GoogleButton";
// import { TwitterButton } from "./TwitterButton";

export function AuthenticationForm(props: PaperProps) {
  const [type, toggle] = useToggle(["login", "register"]);
  const form = useForm({
    initialValues: {
      email: "",
      username: "",
      password: "",
      terms: true,
    },

    validate: {
      // email: (val) => (/^\S+@\S+$/.test(val) ? null : "Invalid email"),
      password: (val) =>
        val.length < 6 ? "Password should include at least 6 characters" : null,
    },
  });

  const router = useRouter();

  const handleSubmit = async () => {
    const req: authRequest_t = {
      username: form.values.username,
      password: form.values.password,
    };

    if (type === "register") {
      req.email = form.values.email;
    }

    // console.log({ req });

    if (type === "login") {
      let res = await handleLogin(req);
      //
      // if (res) {
      //   router.push("/");
      // }
    }
  };

  return (
    <Paper radius="md" w={"50%"} m={"auto"} p="xl" withBorder {...props}>
      <Text size="lg" fw={500}>
        Welcome to xTracker, {type} with
      </Text>

      {/*<Group grow mb="md" mt="md">*/}
      {/*  <GoogleButton radius="xl">Google</GoogleButton>*/}
      {/*  <TwitterButton radius="xl">Twitter</TwitterButton>*/}
      {/*</Group>*/}

      <Divider label="Or continue with email" labelPosition="center" my="lg" />

      <form
        onSubmit={form.onSubmit(() => {
          handleSubmit();
        })}
      >
        <Stack>
          {type === "register" && (
            <TextInput
              required
              label="Email"
              placeholder="hello@mantine.dev"
              value={form.values.email}
              onChange={(event) =>
                form.setFieldValue("email", event.currentTarget.value)
              }
              error={form.errors.email && "Invalid email"}
              radius="md"
            />
          )}

          <TextInput
            label="Username"
            placeholder="Your username"
            value={form.values.username}
            onChange={(event) =>
              form.setFieldValue("username", event.currentTarget.value)
            }
            radius="md"
          />

          <PasswordInput
            required
            label="Password"
            placeholder="Your password"
            value={form.values.password}
            onChange={(event) =>
              form.setFieldValue("password", event.currentTarget.value)
            }
            error={
              form.errors.password &&
              "Password should include at least 6 characters"
            }
            radius="md"
          />

          {type === "register" && (
            <Checkbox
              label="I accept terms and conditions"
              checked={form.values.terms}
              onChange={(event) =>
                form.setFieldValue("terms", event.currentTarget.checked)
              }
            />
          )}
        </Stack>

        <Group justify="space-between" mt="xl">
          <Anchor
            component="button"
            type="button"
            c="dimmed"
            onClick={() => toggle()}
            size="xs"
          >
            {type === "register"
              ? "Already have an account? Login"
              : "Don't have an account? Register"}
          </Anchor>
          <Button type="submit" radius="xl">
            {upperFirst(type)}
          </Button>
        </Group>
      </form>
    </Paper>
  );
}

export default AuthenticationForm;
