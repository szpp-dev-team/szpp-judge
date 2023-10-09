import { AcceptedFileExt, acceptedFileExts, fileExt2langId, MAX_SOURCE_CODE_SIZE } from "@/src/config/submission";
import { type LangID, langIDs, langMetasBrief } from "@/src/gen/langs";
import { convertSb3ToCpp } from "@/src/util/scratch/converter";
import { Sb3ConvertError } from "@/src/util/scratch/errors";
import { ChevronDownIcon } from "@chakra-ui/icons";
import {
  Box,
  BoxProps,
  Button,
  Flex,
  Heading,
  Icon,
  ListItem,
  Menu,
  MenuButton,
  MenuItem,
  MenuList,
  Textarea,
  UnorderedList,
  useToast,
} from "@chakra-ui/react";
import path from "path";
import { ChangeEventHandler, useRef, useState } from "react";
import { IoPushSharp } from "react-icons/io5";

const activeLangIds = langIDs.filter((id) => langMetasBrief[id].active);

const acceptedFileExtsCommnaJoined = acceptedFileExts.join(",");

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
  const fileInputRef = useRef<HTMLInputElement | null>(null);
  const [warnings, setWarnings] = useState<string[]>([]);

  const toast = useToast();

  const handleFileAttach: ChangeEventHandler<HTMLInputElement> = (e) => {
    const files = e.target.files;
    if (files == null || files.length <= 0) return;

    console.log("file attached:", files);

    const file = files[0]!;
    if (file.size > MAX_SOURCE_CODE_SIZE) {
      toast({
        title: `ソースコード長の上限 ${MAX_SOURCE_CODE_SIZE >> 10} KiB を超えています`,
        description: `${file.name} のサイズ: ${file.size / 1024} KiB`,
        status: "warning",
        duration: 8000,
      });
      return;
    }

    if (file.name.endsWith(".sb3")) {
      convertSb3ToCpp(file)
        .then(({ cppSource, warnings }) => {
          onLangIdChange("scratch/3/gcc");
          onSourceCodeChange(cppSource);
          setWarnings(warnings);
        })
        .catch((e) => {
          if (e instanceof Sb3ConvertError) {
            toast({
              title: e.message,
              description: e.detail,
              status: "error",
              duration: 1000 * 15,
            });
          } else {
            console.error("Unknown error while converting sb3 to cpp:", e);
          }
        });
      return;
    }

    file.text().then((s) => {
      onSourceCodeChange(s);
      const lang = fileExt2langId[path.extname(file.name) as AcceptedFileExt];
      if (lang) {
        onLangIdChange(lang);
      }
    }).catch((e) => console.error(e));
  };

  return (
    <Box>
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
            onClick={() => fileInputRef.current?.click()}
          >
            ファイルを開く
          </Button>
        </Flex>
        <input
          type="file"
          ref={fileInputRef}
          style={{ display: "none" }}
          accept={acceptedFileExtsCommnaJoined}
          onChange={handleFileAttach}
        />
        <Textarea
          rows={15}
          value={sourceCode}
          onChange={(e) => onSourceCodeChange(e.target.value)}
        />
      </Box>
      {warnings.length > 0 && (
        <Box as="section">
          <Heading as="h3" mt={4} mb={2} fontSize="lg">警告 (エラーではないので提出はできます)</Heading>
          <UnorderedList color="red.500">
            {warnings.map((msg) => <ListItem key={msg}>{msg}</ListItem>)}
          </UnorderedList>
        </Box>
      )}
    </Box>
  );
};
