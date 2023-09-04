import { Table, Td, Th, Tr } from "@chakra-ui/react";
import type { Meta, StoryObj } from "@storybook/react";

const meta = {
  title: "szpp/chakra/Table",
  component: Table,
  parameters: {
    layout: "centered",
  },
  argTypes: {
    variant: {
      control: "select",
      options: ["simple", "striped", "unstyled", "bordered"],
    },
  },
} satisfies Meta<typeof Table>;

export default meta;
type Story = StoryObj<typeof meta>;

export const bordered: Story = {
  args: {
    variant: "bordered",
  },
  render: ({ ...props }) => (
    <Table {...props}>
      <Tr>
        <Th textAlign="center">#</Th>
        <Th textAlign="center">難易度</Th>
        <Th textAlign="left">タイトル</Th>
        <Th textAlign="center">自分の提出結果</Th>
      </Tr>
      {[
        [1, "Beginner", "すずっぴー君のおつかい", "AC"],
        [2, "Easy", "すずっぴー君のおつかい", "WA"],
        [3, "Easy", "すずっぴー君のおつかい", "TLE"],
        [4, "Hard", "すずっぴー君のおつかい", "13/57 ..."],
      ].map(([seq, d, title, s]) => (
        <Tr>
          <Td textAlign="center">{seq}</Td>
          <Td textAlign="center">{d}</Td>
          <Td textAlign="left">{title}</Td>
          <Td textAlign="center">{s}</Td>
        </Tr>
      ))}
    </Table>
  ),
};
