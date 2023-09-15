import { useMutation } from "@tanstack/react-query";
import { login } from "../gen/proto/backend/v1/services-AuthService_connectquery";

export const useLogin = () => {
  const loginMutation = useMutation(login.useMutation());
  return { loginMutation };
};
