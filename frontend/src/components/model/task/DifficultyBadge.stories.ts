import type { Meta, StoryObj } from "@storybook/react";
import { DifficultyBadge } from "./DifficultyBadge";

const meta = {
  title: "szpp/DifficultyBadge",
  component: DifficultyBadge,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
} satisfies Meta<typeof DifficultyBadge>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Beginner: Story = {
  args: {
    dif: "beginner",
  },
};
export const Easy: Story = {
  args: {
    dif: "easy",
  },
};
export const Medium: Story = {
  args: {
    dif: "medium",
  },
};
export const Hard: Story = {
  args: {
    dif: "hard",
  },
};
export const Impossible: Story = {
  args: {
    dif: "impossible",
  },
};
