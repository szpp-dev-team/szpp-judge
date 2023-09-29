import { getSubmissionDetail } from "@/src/gen/proto/backend/v1/judge_service-JudgeService_connectquery";
import { PlainMessage } from "@bufbuild/protobuf";
import { useQuery } from "@tanstack/react-query";
import { GetSubmissionDetailRequest } from "../gen/proto/backend/v1/judge_service_pb";

export const useGetSubmissionDetail = (input?: PlainMessage<GetSubmissionDetailRequest>) => {
  const { data, error, isLoading } = useQuery(getSubmissionDetail.useQuery(input));
  const submissionDetail = data?.submissionDetail;
  return { submissionDetail, error, isLoading };
};
