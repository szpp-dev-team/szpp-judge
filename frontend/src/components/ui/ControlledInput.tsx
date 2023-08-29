import {
  FormControl,
  FormControlProps,
  FormErrorMessage,
  FormErrorMessageProps,
  FormHelperText,
  FormLabel,
  FormLabelProps,
  forwardRef,
  Input,
  InputProps,
} from "@chakra-ui/react";
import { FieldError, FieldErrorsImpl, Ref } from "react-hook-form";

export type ControlledInputProps = {
  label: string;
  error?: FieldError;
  ref: Ref;
  helpText?: string;
  isRequired?: boolean;
  formControlProps?: Omit<FormControlProps, "isInvalid" | "isRequired">;
  formLabelProps?: FormLabelProps;
  formErrorMessageProps?: FormErrorMessageProps;
} & InputProps;

export const ControlledInput = forwardRef<ControlledInputProps, "input">(
  (
    {
      label,
      error,
      helpText,
      isRequired,
      formControlProps,
      formLabelProps,
      formErrorMessageProps,
      ...rest
    }: Omit<ControlledInputProps, "ref">,
    ref,
  ) => {
    const isInvalid = Boolean(error);
    return (
      <FormControl
        isInvalid={isInvalid}
        isRequired={isRequired}
        {...formControlProps}
      >
        <FormLabel {...formLabelProps}>{label}</FormLabel>
        <Input {...rest} ref={ref} />
        {isInvalid
          ? (
            <FormErrorMessage {...formErrorMessageProps}>
              {error?.message}
            </FormErrorMessage>
          )
          : <FormHelperText>{helpText}</FormHelperText>}
      </FormControl>
    );
  },
);
