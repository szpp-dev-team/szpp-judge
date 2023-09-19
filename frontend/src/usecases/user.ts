import { createUser } from "@/src/gen/proto/backend/v1/user_service-UserService_connectquery";
import { useMutation } from "@tanstack/react-query";

export const useRegister = () => {
  const { data, error, isLoading, mutate } = useMutation(createUser.useMutation());

  return { data, error, isLoading, mutate };
};
