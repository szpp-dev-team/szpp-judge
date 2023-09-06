import type { Meta, StoryObj } from "@storybook/react";
import { JudgeStatusBadge } from "./JudgeStatusBadge";

const meta = {
  title: "szpp/JudgeStatusBadge",
  component: JudgeStatusBadge,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
} satisfies Meta<typeof JudgeStatusBadge>;

export default meta;
type Story = StoryObj<typeof meta>;

export const WJ_Ja: Story = {
  args: {
    status: "WJ",
  },
};
export const IE_Ja: Story = {
  args: {
    status: "IE",
  },
};
export const AC_Ja: Story = {
  args: {
    status: "AC",
  },
};
export const CE_Ja: Story = {
  args: {
    status: "CE",
  },
};
export const WA_Ja: Story = {
  args: {
    status: "WA",
  },
};
export const WJ_InProgress_Ja: Story = {
  args: {
    status: "WJ",
    progress: {
      done: 100,
      total: 128,
    },
  },
};
export const TLE_InProgress_Ja: Story = {
  args: {
    status: "TLE",
    progress: {
      done: 100,
      total: 128,
    },
  },
};

export const WJ: Story = {
  args: {
    status: "WJ",
    abbr: true,
  },
};
export const IE: Story = {
  args: {
    status: "IE",
    abbr: true,
  },
};
export const AC: Story = {
  args: {
    status: "AC",
    abbr: true,
  },
};
export const CE: Story = {
  args: {
    status: "CE",
    abbr: true,
  },
};
export const WA: Story = {
  args: {
    status: "WA",
    abbr: true,
  },
};
export const WJ_InProgress: Story = {
  args: {
    status: "WJ",
    abbr: true,
    progress: {
      done: 100,
      total: 128,
    },
  },
};
export const TLE_InProgress: Story = {
  args: {
    status: "TLE",
    abbr: true,
    progress: {
      done: 100,
      total: 128,
    },
  },
};
