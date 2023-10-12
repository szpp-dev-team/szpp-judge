import { useLogin } from "@/src/usecases/auth";
import { userLoginSchema } from "@/src/zschema/user";
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
  // router.query.redirecturi を配列にしてはいけない. 例えば redirecturi[]=/hogehoge はだめ
  const redirecturi = (router.query.redirecturi ?? window.location.origin + "/") as string;
  const toast = useToast({
    position: "bottom",
  });

  const {
    register,
    handleSubmit,
    setError: setFormError,
    formState: { errors },
  } = useForm<FormFields>({
    mode: "onChange",
    resolver: zodResolver(userLoginSchema),
  });

  const onUnauthenticatedError = () => {
    toast({
      title: "ログインに失敗しました",
      status: "error",
    });
    const err = { type: "custom", message: "ユーザ名またはパスワードが間違っています" } as const;
    setFormError("username", err);
    setFormError("password", err);
  };

  const { isLoading, mutate } = useLogin(onUnauthenticatedError);

  const onSubmit = handleSubmit((values) => {
    console.log("onSubmit(): values=", values);
    mutate(values, {
      onSuccess: (data) => {
        console.log("[LoginComponent] onSuccess");
        toast({
          title: `${data.user!.username} にログインしました`,
          status: "success",
        });
        router.push(redirecturi);
      },
    });
  });

  const formId = "login-form";

  return (
    <Container maxW="48rem" py="7rem" px="2rem">
      <Card p="2rem">
        <CardHeader textAlign="center">
          <Heading as="h1" size="lg" mb={4}>ログイン</Heading>
          <p>
            ユーザ登録をしていない方はまず{" "}
            <Link href={"/register?redirecturi=" + encodeURIComponent(redirecturi)}>こちら</Link> で登録してください
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
          <Button type="submit" form={formId} size="lg" colorScheme="teal" isLoading={isLoading}>
            ログイン
          </Button>
        </CardFooter>
      </Card>
    </Container>
  );
};
