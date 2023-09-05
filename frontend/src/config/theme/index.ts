import { extendTheme } from "@chakra-ui/react";
import { tableTheme } from "./Table";

const theme = extendTheme({
  components: {
    Table: tableTheme,
  },
});

export default theme;
