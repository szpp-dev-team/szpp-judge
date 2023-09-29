import { SubmissionDetail } from "@/src/gen/proto/backend/v1/judge_resources_pb";
import { JudgeStatus } from "@/src/gen/proto/judge/v1/resources_pb";
import type { Meta, StoryObj } from "@storybook/react";
import { SubmissionInfo } from "./SubmissionInfo";

const meta = {
  title: "szpp/SubmissionInfo",
  component: SubmissionInfo,
  tags: ["autodocs"],
} satisfies Meta<typeof SubmissionInfo>;

export default meta;

type Story = StoryObj<typeof meta>;

const submissionDetail1 = new SubmissionDetail({
  username: "user",
  taskTitle: "問題1",
  score: 100,
  langId: "C++",
  status: JudgeStatus.AC,
  execTimeMs: 1,
  execMemoryKib: 3000,
  submittedAt: { seconds: BigInt(1696141800) }, // 2023-10-01 15:30:00
  // testcaseResults etc...
});

export const JudgeCompleted: Story = {
  args: {
    submissionDetail: submissionDetail1,
  },
};
