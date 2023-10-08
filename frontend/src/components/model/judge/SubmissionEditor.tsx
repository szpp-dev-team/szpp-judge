import { type LangID, langIDs, langMetasBrief } from "@/src/gen/langs";
import { ChevronDownIcon } from "@chakra-ui/icons";
import { Box, BoxProps, Button, Flex, Icon, Menu, MenuButton, MenuItem, MenuList, Textarea } from "@chakra-ui/react";
import { IoPushSharp } from "react-icons/io5";

const activeLangIds = langIDs.filter((id) => langMetasBrief[id].active);

export type SubmissionEditorProps = {
  sourceCode: string;
  langId: LangID;
  onLangIdChange: (id: LangID) => unknown;
  onSourceCodeChange: (code: string) => unknown;
} & Omit<BoxProps, "children">;

export const SubmissionEditor = ({
  sourceCode,
  langId,
  onLangIdChange,
  onSourceCodeChange,
  ...props
}: SubmissionEditorProps) => {
  return (
    <Box rounded="md" border="1px" color="teal.900" borderColor="gray.300" {...props}>
      <Flex
        justifyContent="space-between"
        roundedTop="inherit"
        alignItems="center"
        px={2}
        py={2}
        borderBottom="1px"
        borderColor="inherit"
        bg="gray.100"
      >
        <Menu autoSelect={false}>
          <MenuButton
            as={Button}
            textAlign="left"
            bg="white"
            _hover={{ bg: "white" }}
            _expanded={{ bg: "white" }}
            rightIcon={<ChevronDownIcon />}
          >
            {langMetasBrief[langId].name}
          </MenuButton>
          <MenuList>
            {activeLangIds.map((id) => (
              <MenuItem key={id} onClick={() => onLangIdChange(id)}>{langMetasBrief[id].name}</MenuItem>
            ))}
          </MenuList>
        </Menu>
        <Button
          leftIcon={<Icon as={IoPushSharp} />}
          colorScheme="orange"
        >
          ファイルを開く
        </Button>
      </Flex>
      <Textarea
        rows={15}
        value={sourceCode}
        onChange={(e) => onSourceCodeChange(e.target.value)}
      />
    </Box>
  );
};
