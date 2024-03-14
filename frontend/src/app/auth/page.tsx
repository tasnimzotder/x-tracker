"use client";

import { upperFirst, useToggle } from "@mantine/hooks";
import { useForm } from "@mantine/form";
import {
  Anchor,
  Button,
  Checkbox,
  Divider,
  Group,
  Paper,
  PasswordInput,
  Stack,
  Text,
  TextInput,
} from "@mantine/core";
import {
  authRequest_t,
  handleLogin,
  handleRegister,
} from "@/services/auth.service";
import { useAuth } from "@/contexts/authContext";

// import { GoogleButton } from "./GoogleButton";
// import { TwitterButton } from "./TwitterButton";

const AuthenticationForm = () => {
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

  const { login, register } = useAuth();

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
      // await handleLogin(req);
      await login(req);
    } else if (type === "register") {
      await register(req);
    }
  };

  return (
    <Paper radius="md" w={"50%"} m={"auto"} p="xl" withBorder>
      <Text size="lg" fw={500}>
        Welcome to xTracker, {type} with
      </Text>

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
};

export default AuthenticationForm;
