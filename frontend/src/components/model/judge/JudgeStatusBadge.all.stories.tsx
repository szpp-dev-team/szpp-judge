import { JudgeStatusValues } from "@/src/model/judge";
import { HStack, VStack } from "@chakra-ui/react";
import type { Meta, StoryObj } from "@storybook/react";
import type { FC } from "react";
import { JudgeStatusBadge } from "./JudgeStatusBadge";

const View: FC = () => (
  <VStack>
    {JudgeStatusValues.map((s) => (
      <HStack>
        <JudgeStatusBadge status={s} />
        <JudgeStatusBadge status={s} abbr />
        <JudgeStatusBadge status={s} progress={{ done: 100, total: 128 }} />
        <JudgeStatusBadge status={s} abbr progress={{ done: 100, total: 128 }} />
      </HStack>
    ))}
  </VStack>
);

const meta = {
  title: "szpp/JudgeStatusBadge-all",
  component: View,
  parameters: {
    layout: "centered",
  },
} satisfies Meta<typeof View>;

export default meta;
type Story = StoryObj<typeof meta>;

export const all: Story = {};
