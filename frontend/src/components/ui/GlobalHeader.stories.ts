import type { Meta, StoryObj } from "@storybook/react";
import { GlobalHeader } from "./GlobalHeader";

const meta = {
  title: "szpp/GlobalHeader",
  component: GlobalHeader,
  parameters: {
    layout: "fullscreen",
  },
} satisfies Meta<typeof GlobalHeader>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {},
};
export const Contest: Story = {
  args: {
    contestSlug: "sbc001",
    contestTitle: "SZPP プログラミング講座 in 浜松 2023年4月1日",
  },
};
