import { AcceptedFileExt, acceptedFileExts, fileExt2langId, MAX_SOURCE_CODE_SIZE } from "@/src/config/submission";
import { type LangID, langIDs, langMetasBrief } from "@/src/gen/langs";
import { convertSb3ToCpp } from "@/src/util/scratch/converter";
import { Sb3ConvertError } from "@/src/util/scratch/errors";
import { ChevronDownIcon } from "@chakra-ui/icons";
import {
  Box,
  BoxProps,
  Button,
  ButtonProps,
  Flex,
  Heading,
  ListItem,
  Menu,
  MenuButton,
  MenuItem,
  MenuList,
  Text,
  UnorderedList,
  useToast,
} from "@chakra-ui/react";
import path from "path";
import { ChangeEventHandler, useCallback, useRef, useState } from "react";
import { Editor } from "../../ui/Editor";

const activeLangIds = langIDs.filter((id) => langMetasBrief[id].active);

const acceptedFileExtsCommnaJoined = acceptedFileExts.join(",");

export type SubmissionFormProps = {
  onSubmit?: (langId: LangID, sourceCode: string) => unknown;
  submitButtonProps?: ButtonProps;
} & Omit<BoxProps, "children" | "onSubmit">;

export const SubmissionForm = ({
  onSubmit,
  submitButtonProps,
}: SubmissionFormProps) => {
  const [sourceCode, setSourceCode] = useState("");
  const [langId, setLangId] = useState<LangID>("cpp/20/gcc");

  const fileInputRef = useRef<HTMLInputElement | null>(null);
  const [warnings, setWarnings] = useState<string[]>([]);

  const toast = useToast();

  const handleFileAttach = useCallback<ChangeEventHandler<HTMLInputElement>>((e) => {
    const files = e.target.files;
    if (files == null || files.length <= 0) return;

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
          setLangId("scratch/3/gcc");
          setSourceCode(cppSource);
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
      setSourceCode(s);
      const lang = fileExt2langId[path.extname(file.name) as AcceptedFileExt];
      if (lang) {
        setLangId(lang);
      }
    }).catch((e) => console.error(e));
  }, [toast]);

  return (
    <Box as="form">
      <Text my={2} fontSize="sm">ソースコード長の上限は {MAX_SOURCE_CODE_SIZE >> 10} KiB です。</Text>
      <Box rounded="md" border="1px" color="teal.900" borderColor="gray.300" overflow="hidden">
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
              type="button"
              fontSize="14px"
              px={3}
              py={2}
              rounded="md"
              textAlign="left"
              border="1px"
              borderColor="gray.300"
              bg="white"
              _hover={{ bg: "white" }}
              _expanded={{ bg: "white" }}
            >
              <Flex alignItems="center">
                <Text
                  as="span"
                  display="inline-block"
                  w="20ch"
                  whiteSpace="nowrap"
                  overflow="hidden"
                  textOverflow="ellipsis"
                >
                  {langMetasBrief[langId].name}
                </Text>
                <ChevronDownIcon ml={1} />
              </Flex>
            </MenuButton>
            <MenuList>
              {activeLangIds.map((id) => (
                <MenuItem key={id} onClick={() => setLangId(id)}>{langMetasBrief[id].name}</MenuItem>
              ))}
            </MenuList>
          </Menu>
          <Button
            fontSize="12px"
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
        <Editor
          lang={langId}
          value={sourceCode}
          onChange={setSourceCode}
          height="20rem"
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
      <Flex justifyContent="center" my={12}>
        <Button
          fontSize="xl"
          px={14}
          py={8}
          colorScheme="teal"
          onClick={onSubmit ? () => onSubmit(langId, sourceCode) : undefined}
          {...submitButtonProps}
        >
          提出
        </Button>
      </Flex>
    </Box>
  );
};
