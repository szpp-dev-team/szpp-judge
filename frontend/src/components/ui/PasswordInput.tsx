import { forwardRef, useBoolean } from "@chakra-ui/react";
import { ControllablePasswordInput, ControllablePasswordInputProps } from "./ControllablePasswordInput";

export type PasswordInputProps = Omit<
  Omit<ControllablePasswordInputProps, "showPassword">,
  "onToggleButtonClicked"
>;

export const PasswordInput = forwardRef<PasswordInputProps, "input">(
  ({ ...props }, ref) => {
    const [passwordVisible, setPasswordVisible] = useBoolean(false);
    return (
      <ControllablePasswordInput
        ref={ref}
        showPassword={passwordVisible}
        onToggleButtonClick={setPasswordVisible.toggle}
        {...props}
      />
    );
  },
);
