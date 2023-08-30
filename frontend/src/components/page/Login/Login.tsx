import { userLoginSchema } from "@/src/zschema/user";
import { Button, Card, CardBody, CardFooter, CardHeader, Container, Heading, Input } from "@chakra-ui/react";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { InputOrganism } from "../../ui/InputOrganism";
import { Link } from "../../ui/Link";
import { PasswordInput } from "../../ui/PasswordInput";

type FormFields = z.infer<typeof userLoginSchema>;

export const Login = () => {
  const formId = "login-form";

  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<FormFields>({
    mode: "onChange",
    resolver: zodResolver(userLoginSchema),
  });

  const onSubmit = handleSubmit((values) => {
    console.log("submit:", values);
  });

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
        <CardFooter justifyContent="center">
          <Button type="submit" form={formId} size="lg" colorScheme="teal" isLoading={isSubmitting}>
            ログイン
          </Button>
        </CardFooter>
      </Card>
    </Container>
  );
};
