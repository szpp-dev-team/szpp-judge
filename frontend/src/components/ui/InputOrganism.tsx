import { FormControl, FormErrorMessage, FormHelperText, FormLabel } from "@chakra-ui/react";
import { ReactNode } from "react";
import { FieldErrors, FieldValues } from "react-hook-form";
import { z, ZodTypeAny } from "zod";

export type InputOrganismProps<T extends ZodTypeAny, Fields extends FieldValues = z.infer<T>> = {
  label?: string;
  schema: T;
  name: keyof Fields;
  errors?: FieldErrors<Fields>;
  children: ReactNode;
};

export const InputOrganism = <T extends ZodTypeAny>({
  label,
  schema,
  name,
  errors,
  children,
}: InputOrganismProps<T>) => {
  // @ts-expect-error: Property 'shape' does not exist というエラーになってしまうため
  const desc: string | undefined = schema.shape?.[name]?.description;

  return (
    <FormControl isInvalid={errors?.[name] != null} mb={12} position="relative">
      {label && <FormLabel mb={desc ? 0 : 1}>{label}</FormLabel>}
      {desc && <FormHelperText mt={0} mb={2} fontSize="xs">{desc}</FormHelperText>}
      {children}
      {/* @ts-expect-error: message を ReactNode に代入できないというエラーになってしまうため */}
      <FormErrorMessage position="absolute" bottom="-1.5lh">{errors?.[name]?.message}</FormErrorMessage>
    </FormControl>
  );
};
