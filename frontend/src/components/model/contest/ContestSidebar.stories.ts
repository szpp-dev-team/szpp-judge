import type { Meta, StoryObj } from "@storybook/react";
import { ContestSidebar, ContestSidebarProps } from "./ContestSidebar";

const meta = {
  title: "szpp/ContestSidebar",
  component: ContestSidebar,
  parameters: {
    layout: "fullscreen",
  },
  tags: ["autodocs"],
  argTypes: {},
} satisfies Meta<typeof ContestSidebar>;

export default meta;
type Story = StoryObj<typeof meta>;

const fewTasks: ContestSidebarProps["tasks"] = [
  { id: 1000, title: "すずっぴー君のおつかい" },
  { id: 1001, title: "すずっぴー" },
  { id: 1002, title: "長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル" },
];

const manyTasks: ContestSidebarProps["tasks"] = [
  { id: 1000, title: "すずっぴー君のおつかい", scoreStatus: "perfect" },
  { id: 1001, title: "すずっぴー", scoreStatus: "partial" },
  { id: 1002, title: "ぴっぴー", scoreStatus: "zero" },
  { id: 1003, title: "ぴー" },
  { id: 1004, title: "XOR Sum", scoreStatus: "perfect" },
  { id: 1005, title: "longlonglonglonglonglonglonglonglonglonglonglonglonglong" },
  { id: 1006, title: "長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル", scoreStatus: "zero" },
  { id: 1007, title: "はい", scoreStatus: "perfect" },
  { id: 1008, title: "うん" },
  { id: 1009, title: "んんん...？" },
  { id: 1010, title: "そ、そんな..." },
  { id: 1011, title: "何も分からん" },
  { id: 1012, title: "椅子を温める" },
];

export const BeforeStart: Story = {
  args: {
    startAt: new Date("2023-09-02 21:00"),
    endAt: new Date("2023-09-04 01:00"),
    now: new Date("2023-09-01 07:00:51"),
    slug: "sbc001",
    tasks: [],
  },
};
export const BeforeStart_in24H: Story = {
  args: {
    startAt: new Date("2023-09-02 21:00"),
    endAt: new Date("2023-09-04 01:00"),
    now: new Date("2023-09-02 19:45:03"),
    slug: "sbc001",
    tasks: [],
  },
};
export const ContestNow_fewTasks: Story = {
  args: {
    startAt: new Date("2023-09-02 21:00"),
    endAt: new Date("2023-09-04 01:00"),
    now: new Date("2023-09-03 22:59:59"),
    slug: "sbc001",
    tasks: fewTasks,
  },
};
export const ContestNow_manyTasks: Story = {
  args: {
    startAt: new Date("2023-09-02 21:00"),
    endAt: new Date("2023-09-15 01:00"),
    now: new Date("2023-09-02 21:03:06"),
    slug: "sbc001",
    tasks: manyTasks,
  },
};
export const AfterEnd = {
  args: {
    startAt: new Date("2023-09-02 21:00"),
    endAt: new Date("2023-09-15 01:00"),
    now: new Date("2023-09-15 01:00:01"),
    slug: "sbc001",
    tasks: manyTasks,
  },
};
