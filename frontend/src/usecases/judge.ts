import {
  getJudgeProgress,
  getSubmissionDetail,
  listSubmissions,
  submit,
} from "@/src/gen/proto/backend/v1/judge_service-JudgeService_connectquery";
import {
  GetJudgeProgressRequest,
  GetSubmissionDetailRequest,
  ListSubmissionsRequest,
} from "@/src/gen/proto/backend/v1/judge_service_pb";
import { JudgeStatus } from "@/src/gen/proto/judge/v1/resources_pb";
import { PlainMessage } from "@bufbuild/protobuf";
import { useMutation, useQuery } from "@tanstack/react-query";
import { useState } from "react";
import { Duration } from "../util/time";

export const useGetSubmissionDetail = (input?: PlainMessage<GetSubmissionDetailRequest>) => {
  const { data, error, isLoading } = useQuery(getSubmissionDetail.useQuery(input));
  const submissionDetail = data?.submissionDetail;
  return { submissionDetail, error, isLoading };
};

export const useSubmit = () => {
  const { mutate, data, error, isLoading } = useMutation(submit.useMutation());
  return { mutate, data, error, isLoading };
};

export const useListSubmissions = (input?: PlainMessage<ListSubmissionsRequest>) => {
  return useQuery({
    ...listSubmissions.useQuery(input),
    staleTime: Duration.MINUTE,
  });
};

export const useGetJudgeProgress = (input?: PlainMessage<GetJudgeProgressRequest>) => {
  const [enabled, setEnabled] = useState(true);
  return useQuery({
    ...getJudgeProgress.useQuery(input),
    onSettled: (data) => {
      const status = data?.judgeProgress?.status;
      if (typeof status === "number" && status !== JudgeStatus.WJ) {
        setEnabled(false);
      }
    },
    staleTime: 2 * Duration.SECOND,
    cacheTime: 0,
    enabled,
  });
};
