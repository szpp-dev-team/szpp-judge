import { userLoginSchema } from "@/src/zschema/user";
import { Button, Card, CardBody, CardFooter, CardHeader, Container, Heading } from "@chakra-ui/react";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { ControlledInput } from "../../ui/ControlledInput";
import { Link } from "../../ui/Link";

type FormFields = z.infer<typeof userLoginSchema>;

export const Login = () => {
  const formId = "login-form";

  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<FormFields>({
    mode: "onBlur",
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
            <ControlledInput
              label="ユーザ名"
              error={errors.username}
              helpText={userLoginSchema.shape.username.description}
              {...register("username")}
              formControlProps={{ mb: 8 }}
            />
            <ControlledInput label="パスワード" error={errors.password} {...register("password")} />
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
