import { MAX_SOURCE_CODE_SIZE } from "@/src/config/submission";
import { type LangID, langIDs, langMetasBrief } from "@/src/gen/langs";
import { ChevronDownIcon } from "@chakra-ui/icons";
import {
  Box,
  BoxProps,
  Button,
  Flex,
  Icon,
  Menu,
  MenuButton,
  MenuItem,
  MenuList,
  Textarea,
  useToast,
} from "@chakra-ui/react";
import path from "path";
import { ChangeEventHandler, useRef } from "react";
import { IoPushSharp } from "react-icons/io5";

const activeLangIds = langIDs.filter((id) => langMetasBrief[id].active);

const acceptedFileExts = [
  ".c",
  ".cpp",
  ".cc",
  ".py",
  ".java",
  ".sb3",
] as const;

type AcceptedFileExt = (typeof acceptedFileExts)[number];

const acceptedFileExtsCommnaJoined = acceptedFileExts.join(",");

// NOTE* 将来、提出可能な言語が増えて C++17 や C++20 の clang 版などができた場合に対応・仕様の見直しが必要
const ext2lang = {
  ".c": "c/11/gcc",
  ".cc": "cpp/20/gcc",
  ".cpp": "cpp/20/gcc",
  ".py": "python/3.11/cpython",
  ".java": "java/21/openjdk",
  ".sb3": "scratch/3/gcc",
} satisfies Record<AcceptedFileExt, LangID>;

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

    file.text().then((s) => {
      onSourceCodeChange(s);
      const lang = ext2lang[path.extname(file.name) as AcceptedFileExt];
      if (lang) {
        onLangIdChange(lang);
      }
    }).catch((e) => console.error(e));
  };

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
  );
};
