import { FormControl, FormErrorMessage, FormHelperText, FormLabel } from "@chakra-ui/react";
import { ReactNode } from "react";
import { FieldErrors, FieldValues } from "react-hook-form";
import { z, ZodType } from "zod";

export type InputOrganismProps<T extends ZodType<any, any, any>, Fields extends FieldValues = z.infer<T>> = {
  label?: string;
  schema: T;
  name: keyof Fields;
  errors?: FieldErrors<Fields>;
  children: ReactNode;
};

export const InputOrganism = <T extends ZodType<any, any, any>>({
  label,
  schema,
  name,
  errors,
  children,
}: InputOrganismProps<T>) => {
  // @ts-ignore
  const desc: string | undefined = schema.shape?.[name]?.description;

  return (
    <FormControl isInvalid={errors?.[name] != null} mb={12} position="relative">
      {label && <FormLabel mb={desc ? 0 : 1}>{label}</FormLabel>}
      {desc && <FormHelperText mt={0} mb={2} fontSize="xs">{desc}</FormHelperText>}
      {children}
      {/*@ts-ignore*/}
      <FormErrorMessage>{errors?.[name]?.message}</FormErrorMessage>
    </FormControl>
  );
};
