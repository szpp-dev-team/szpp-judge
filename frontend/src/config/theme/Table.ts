/* 参考にしたソースコード：
 * https://github.com/chakra-ui/chakra-ui/blob/main/packages/components/theme/src/components/table.ts
 *
 *  ↑ このような theme のソースコードは、Chakra UI のドキュメントの各コンポーネントページ (例： https://chakra-ui.com/docs/components/table)
 *  の冒頭の「Theme source」ボタンをクリックすることで見つけることができる。
 */
import { tableAnatomy } from "@chakra-ui/anatomy";
import { createMultiStyleConfigHelpers, defineStyle } from "@chakra-ui/styled-system";

const { defineMultiStyleConfig, definePartsStyle } = createMultiStyleConfigHelpers(tableAnatomy.keys);

const numericStyles = defineStyle({
  "&[data-is-numeric=true]": {
    textAlign: "end",
  },
});

const variantBordered = definePartsStyle((props) => {
  const { colorScheme: c } = props;
  const borderColor = `${c}.300`;

  return {
    table: {
      borderTop: "1px",
      borderLeft: "1px",
      borderColor,
    },
    th: {
      fontSize: "sm",
      color: "gray.700",
      bg: "gray.50",
      borderBottom: "1px",
      borderRight: "1px",
      borderColor,
      ...numericStyles,
    },
    td: {
      borderBottom: "1px",
      borderRight: "1px",
      borderColor,
      ...numericStyles,
    },
    tfoot: {
      tr: {
        "&:last-of-type": {
          th: { borderBottomWidth: 0 },
        },
      },
    },
  };
});

export const tableTheme = defineMultiStyleConfig({
  variants: {
    bordered: variantBordered,
  },
});
