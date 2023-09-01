import { ViewIcon, ViewOffIcon } from "@chakra-ui/icons";
import { Flex, forwardRef, Input, InputGroup, InputRightElement, Text } from "@chakra-ui/react";
import type { InputGroupProps, InputProps, TextProps } from "@chakra-ui/react";
import { MouseEventHandler, ReactNode } from "react";

export type ControllablePasswordInputProps = {
  showPassword?: boolean;
  onToggleButtonClick?: MouseEventHandler<HTMLButtonElement>;
  leftElem?: ReactNode;
  inputGroupProps?: InputGroupProps;
} & InputProps;

export const ControllablePasswordInput = forwardRef<ControllablePasswordInputProps, "input">(
  ({
    showPassword,
    onToggleButtonClick,
    leftElem,
    inputGroupProps,
    ...props
  }, ref) => {
    const btnTextProps: TextProps = {
      as: "span",
      fontSize: "sm",
    };
    return (
      <InputGroup {...inputGroupProps}>
        {leftElem}
        <Input
          type={showPassword ? "text" : "password"}
          fontFamily="monospace"
          ref={ref}
          {...props}
        />
        <InputRightElement w="6.5em" p="1px" color="gray.500" justifyContent="end">
          <Flex
            as="button"
            type="button"
            alignItems="center"
            gap={1}
            px={3}
            h="100%"
            bg="whitesmoke"
            borderTopRightRadius="var(--input-border-radius)"
            borderBottomRightRadius="var(--input-border-radius)"
            // @ts-expect-error MouseEventHandler<Button> を MouseEventHandler<Div> に代入できないため型エラー
            // Flex に as="button" を指定しているけど onClick 属性が許容する型は div のままになってそう
            onClick={onToggleButtonClick}
          >
            {showPassword
              ? (
                <>
                  <ViewIcon />
                  <Text {...btnTextProps}>隠す</Text>
                </>
              )
              : (
                <>
                  <ViewOffIcon />
                  <Text {...btnTextProps}>表示する</Text>
                </>
              )}
          </Flex>
        </InputRightElement>
      </InputGroup>
    );
  },
);
