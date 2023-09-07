import { CreateUserRequest } from "@/src/gen/proto/backend/v1/messages_pb";
import { userRepo } from "@/src/repository/user";
import { userRegistrationSchema } from "@/src/zschema/user";
import { Button, Card, CardBody, CardFooter, CardHeader, Container, Heading, Input, useToast, Checkbox } from "@chakra-ui/react";
import { zodResolver } from "@hookform/resolvers/zod";
import { useRouter } from "next/router";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { InputOrganism } from "../../ui/InputOrganism";
import { Link } from "../../ui/Link";
import { PasswordInput } from "../../ui/PasswordInput";
import { useState } from "react";

type FormFields = z.infer<typeof userRegistrationSchema>;

export const Register = () => {
  const router = useRouter();
  const toast = useToast({
    position: "bottom",
  });

  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<FormFields>({
    mode: "onChange",
    resolver: zodResolver(userRegistrationSchema),
  });

  const onSubmit = handleSubmit(async (values) => {
    console.log("onSubmit(): values=", values);
    try {
      const resp = await userRepo.register(new CreateUserRequest(values));
      console.log("onSubmit(): resp=", resp);
      toast({
        title: `${resp.username} で登録しました`,
        status: "success",
      });
      router.push("/");
    } catch (e) {
      toast({
        title: "登録に失敗しました",
        status: "error",
      });
      console.error(typeof e, e);
    }
  });

  const formId = "register-form";

  const [policyAgreed, setPolicyAgreed] = useState(false);

  return (
    <Container maxW="48rem" py="7rem" px="2rem">
      <Card p="2rem">
        <CardHeader textAlign="center">
          <Heading as="h1" size="lg" mb={4}>ユーザ新規登録</Heading>
          <p>
            既に登録済みの方は <Link href="/login">ログインページ</Link> へ
          </p>
        </CardHeader>
        <CardBody textAlign="center">
          <form
            id={formId}
            onSubmit={onSubmit}
          >
            <InputOrganism
              schema={userRegistrationSchema}
              label="メールアドレス"
              name="email"
              errors={errors}
            >
              <Input {...register("email")} />
            </InputOrganism>
            <InputOrganism
              schema={userRegistrationSchema}
              label="ユーザ名"
              name="username"
              errors={errors}
            >
              <Input {...register("username")} />
            </InputOrganism>
            <InputOrganism
              schema={userRegistrationSchema}
              label="パスワード"
              name="password"
              errors={errors}
            >
              <PasswordInput {...register("password")} />
            </InputOrganism>
            <InputOrganism
              schema={userRegistrationSchema}
              label="パスワード(確認)"
              name="confPassword"
              errors={errors}
            >
              <PasswordInput {...register("confPassword")} />
            </InputOrganism>
            <Checkbox required isChecked={policyAgreed} onChange={() => setPolicyAgreed(!policyAgreed)}>
              <Link href="/tos">利用規約</Link> および <Link href="/privacy"> 個人情報の取り扱い</Link> に同意する
            </Checkbox>
          </form>
        </CardBody>
        <CardFooter flexDirection="column" alignItems="center" gap={2}>
        <Button type="submit" form={formId} size="lg" colorScheme={policyAgreed ? "teal" : "blackAlpha"} isLoading={isSubmitting}>
            登録
          </Button>
        </CardFooter>
      </Card>
    </Container>
  );
};
