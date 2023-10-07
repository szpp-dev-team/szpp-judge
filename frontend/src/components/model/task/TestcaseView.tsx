import { Box, BoxProps, Flex, Heading, Icon, Tooltip } from "@chakra-ui/react";
import { useCallback, useRef, useState } from "react";
import { IoCopyOutline } from "react-icons/io5";

export type TestcaseViewProps = {
  title: string;
  content: string;
  titleAs: "h1" | "h2" | "h3" | "h4" | "h5" | "h6";
} & Omit<BoxProps, "children">;

const BORDER_COLOR: BoxProps["borderColor"] = "gray.300";

export const TestcaseView = ({ title, titleAs, content, ...props }: TestcaseViewProps) => {
  const [copyActionState, setCopyActionState] = useState<"idle" | "copied">("idle");
  const copyActionTimeoutId = useRef<number | null>(null);

  const copyToClipboard = useCallback(() => {
    if (copyActionTimeoutId.current != null) {
      clearTimeout(copyActionTimeoutId.current);
      copyActionTimeoutId.current = null;
    }

    window.navigator.clipboard.writeText(content)
      .then(() => {
        setCopyActionState("copied");
        copyActionTimeoutId.current = window.setTimeout(() => setCopyActionState("idle"), 1000);
      }).catch((e) => {
        console.error("Failed to copy to clipboard:", e)
      });
  }, [content]);

  return (
    <Box as="section" rounded="md" border="1px" color="teal.900" borderColor={BORDER_COLOR} {...props}>
      <Flex
        justifyContent="space-between"
        roundedTop="inherit"
        alignItems="center"
        borderBottom="1px"
        borderColor={BORDER_COLOR}
        background="gray.100"
      >
        <Heading as={titleAs} fontSize="md" px={2} py={1}>{title}</Heading>
        <Tooltip hasArrow label="クリップボードにコピーしました" isOpen={copyActionState === "copied"}>
          <Flex
            as="button"
            justifyContent="center"
            alignItems="center"
            px={2}
            py={1}
            roundedTopRight="inherit"
            fontSize="xs"
            background="gray.50"
            borderLeft="1px"
            borderColor={BORDER_COLOR}
            disabled={copyActionState === "copied"}
            cursor={copyActionState === "copied" ? "default" : "pointer"}
            _hover={{ color: "cyan.600" }}
            onClick={copyToClipboard}
          >
            コピー
            <Icon as={IoCopyOutline} ml={1} fontSize={"xl"} />
          </Flex>
        </Tooltip>
      </Flex>
      <Box as="pre" overflow="auto" p={2}>
        <Box as="code">{content}</Box>
      </Box>
    </Box>
  );
};
