import { forwardRef, useBoolean } from "@chakra-ui/react";
import { ControllablePasswordInput, type ControllablePasswordInputProps } from "./ControllablePasswordInput";

export type PasswordInputProps = Omit<
  ControllablePasswordInputProps,
  "showPassword" | "onToggleButtonClicked"
>;

export const PasswordInput = forwardRef<PasswordInputProps, "input">(
  ({ ...props }, ref) => {
    const [showPassword, setShowPassword] = useBoolean(false);
    return (
      <ControllablePasswordInput
        ref={ref}
        showPassword={showPassword}
        onToggleButtonClick={setShowPassword.toggle}
        {...props}
      />
    );
  },
);
