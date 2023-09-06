import { LoginRequest } from "@/src/gen/proto/backend/v1/messages_pb";
import { authRepo } from "@/src/repository/auth";
import { userLoginSchema } from "@/src/zschema/user";
import { Code, ConnectError } from "@bufbuild/connect";
import { Button, Card, CardBody, CardFooter, CardHeader, Container, Heading, Input, useToast } from "@chakra-ui/react";
import { zodResolver } from "@hookform/resolvers/zod";
import { useRouter } from "next/router";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { InputOrganism } from "../../ui/InputOrganism";
import { Link } from "../../ui/Link";
import { PasswordInput } from "../../ui/PasswordInput";

type FormFields = z.infer<typeof userLoginSchema>;

export const Login = () => {
  const router = useRouter();
  const toast = useToast({
    position: "bottom",
  });

  const {
    register,
    handleSubmit,
    setError: setFormError,
    formState: { errors, isSubmitting },
  } = useForm<FormFields>({
    mode: "onChange",
    resolver: zodResolver(userLoginSchema),
  });

  const onSubmit = handleSubmit(async (values) => {
    console.log("onSubmit(): values=", values);
    try {
      const resp = await authRepo.login(new LoginRequest(values));
      console.log("onSubmit(): resp=", resp);
      toast({
        title: `${resp.username} にログインしました`,
        status: "success",
      });
      router.push("/");
    } catch (e) {
      toast({
        title: "ログインに失敗しました",
        status: "error",
      });
      console.error(typeof e, e);

      if (e instanceof ConnectError && e.code === Code.Unauthenticated) {
        console.log("e.code:", e.code);
        console.log("e.name:", e.name);
        console.log("e.message:", e.message);
        console.log("e.details:", e.details);
        console.log("e.metadata:", e.metadata);
        const err = { type: "custom", message: "ユーザ名またはパスワードが間違っています" } as const;
        setFormError("username", err);
        setFormError("password", err);
      }
    }
  });

  const formId = "login-form";

  return (
    <Container maxW="48rem" py="7rem" px="2rem">
      <Card p="2rem">
        <CardHeader textAlign="center">
          <Heading as="h1" size="lg" mb={4}>ログイン</Heading>
          <p>
            ユーザ登録をしていない方はまず <Link href="/register">こちら</Link> で登録してください
          </p>
        </CardHeader>
        <CardBody>
          <form
            id={formId}
            onSubmit={onSubmit}
          >
            <InputOrganism
              schema={userLoginSchema}
              label="ユーザ名"
              name="username"
              errors={errors}
            >
              <Input {...register("username")} />
            </InputOrganism>
            <InputOrganism
              schema={userLoginSchema}
              label="パスワード"
              name="password"
              errors={errors}
            >
              <PasswordInput {...register("password")} />
            </InputOrganism>
          </form>
        </CardBody>
        <CardFooter flexDirection="column" alignItems="center" gap={2}>
          <Button type="submit" form={formId} size="lg" colorScheme="teal" isLoading={isSubmitting}>
            ログイン
          </Button>
        </CardFooter>
      </Card>
    </Container>
  );
};
