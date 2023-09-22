// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: backend/v1/judge_service.proto

package backendv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SubmitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContestId  *int32 `protobuf:"varint,1,opt,name=contest_id,json=contestId,proto3,oneof" json:"contest_id,omitempty"`
	TaskId     int32  `protobuf:"varint,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	LangId     string `protobuf:"bytes,3,opt,name=lang_id,json=langId,proto3" json:"lang_id,omitempty"`
	SourceCode string `protobuf:"bytes,4,opt,name=source_code,json=sourceCode,proto3" json:"source_code,omitempty"`
}

func (x *SubmitRequest) Reset() {
	*x = SubmitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_v1_judge_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitRequest) ProtoMessage() {}

func (x *SubmitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_v1_judge_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitRequest.ProtoReflect.Descriptor instead.
func (*SubmitRequest) Descriptor() ([]byte, []int) {
	return file_backend_v1_judge_service_proto_rawDescGZIP(), []int{0}
}

func (x *SubmitRequest) GetContestId() int32 {
	if x != nil && x.ContestId != nil {
		return *x.ContestId
	}
	return 0
}

func (x *SubmitRequest) GetTaskId() int32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *SubmitRequest) GetLangId() string {
	if x != nil {
		return x.LangId
	}
	return ""
}

func (x *SubmitRequest) GetSourceCode() string {
	if x != nil {
		return x.SourceCode
	}
	return ""
}

type SubmitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubmissionId int32 `protobuf:"varint,1,opt,name=submission_id,json=submissionId,proto3" json:"submission_id,omitempty"`
}

func (x *SubmitResponse) Reset() {
	*x = SubmitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_v1_judge_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitResponse) ProtoMessage() {}

func (x *SubmitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_v1_judge_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitResponse.ProtoReflect.Descriptor instead.
func (*SubmitResponse) Descriptor() ([]byte, []int) {
	return file_backend_v1_judge_service_proto_rawDescGZIP(), []int{1}
}

func (x *SubmitResponse) GetSubmissionId() int32 {
	if x != nil {
		return x.SubmissionId
	}
	return 0
}

type GetSubmissionDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetSubmissionDetailRequest) Reset() {
	*x = GetSubmissionDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_v1_judge_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubmissionDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubmissionDetailRequest) ProtoMessage() {}

func (x *GetSubmissionDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_v1_judge_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubmissionDetailRequest.ProtoReflect.Descriptor instead.
func (*GetSubmissionDetailRequest) Descriptor() ([]byte, []int) {
	return file_backend_v1_judge_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetSubmissionDetailRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetSubmissionDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubmissionDetail *SubmissionDetail `protobuf:"bytes,1,opt,name=submission_detail,json=submissionDetail,proto3" json:"submission_detail,omitempty"`
}

func (x *GetSubmissionDetailResponse) Reset() {
	*x = GetSubmissionDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_v1_judge_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubmissionDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubmissionDetailResponse) ProtoMessage() {}

func (x *GetSubmissionDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_v1_judge_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubmissionDetailResponse.ProtoReflect.Descriptor instead.
func (*GetSubmissionDetailResponse) Descriptor() ([]byte, []int) {
	return file_backend_v1_judge_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetSubmissionDetailResponse) GetSubmissionDetail() *SubmissionDetail {
	if x != nil {
		return x.SubmissionDetail
	}
	return nil
}

type ListSubmissionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContestId *int32  `protobuf:"varint,1,opt,name=contest_id,json=contestId,proto3,oneof" json:"contest_id,omitempty"`
	Username  *string `protobuf:"bytes,2,opt,name=username,proto3,oneof" json:"username,omitempty"`
}

func (x *ListSubmissionsRequest) Reset() {
	*x = ListSubmissionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_v1_judge_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSubmissionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubmissionsRequest) ProtoMessage() {}

func (x *ListSubmissionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_v1_judge_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubmissionsRequest.ProtoReflect.Descriptor instead.
func (*ListSubmissionsRequest) Descriptor() ([]byte, []int) {
	return file_backend_v1_judge_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListSubmissionsRequest) GetContestId() int32 {
	if x != nil && x.ContestId != nil {
		return *x.ContestId
	}
	return 0
}

func (x *ListSubmissionsRequest) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

type ListSubmissionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Submissions []*SubmissionSummary `protobuf:"bytes,1,rep,name=submissions,proto3" json:"submissions,omitempty"`
}

func (x *ListSubmissionsResponse) Reset() {
	*x = ListSubmissionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_v1_judge_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSubmissionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubmissionsResponse) ProtoMessage() {}

func (x *ListSubmissionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_v1_judge_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubmissionsResponse.ProtoReflect.Descriptor instead.
func (*ListSubmissionsResponse) Descriptor() ([]byte, []int) {
	return file_backend_v1_judge_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListSubmissionsResponse) GetSubmissions() []*SubmissionSummary {
	if x != nil {
		return x.Submissions
	}
	return nil
}

type GetJudgeProgressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubmissionId int32 `protobuf:"varint,1,opt,name=submission_id,json=submissionId,proto3" json:"submission_id,omitempty"`
}

func (x *GetJudgeProgressRequest) Reset() {
	*x = GetJudgeProgressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_v1_judge_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJudgeProgressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJudgeProgressRequest) ProtoMessage() {}

func (x *GetJudgeProgressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_v1_judge_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJudgeProgressRequest.ProtoReflect.Descriptor instead.
func (*GetJudgeProgressRequest) Descriptor() ([]byte, []int) {
	return file_backend_v1_judge_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetJudgeProgressRequest) GetSubmissionId() int32 {
	if x != nil {
		return x.SubmissionId
	}
	return 0
}

type GetJudgeProgressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JudgeProgress *JudgeProgress `protobuf:"bytes,1,opt,name=judge_progress,json=judgeProgress,proto3" json:"judge_progress,omitempty"`
}

func (x *GetJudgeProgressResponse) Reset() {
	*x = GetJudgeProgressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_v1_judge_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJudgeProgressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJudgeProgressResponse) ProtoMessage() {}

func (x *GetJudgeProgressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_v1_judge_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJudgeProgressResponse.ProtoReflect.Descriptor instead.
func (*GetJudgeProgressResponse) Descriptor() ([]byte, []int) {
	return file_backend_v1_judge_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetJudgeProgressResponse) GetJudgeProgress() *JudgeProgress {
	if x != nil {
		return x.JudgeProgress
	}
	return nil
}

var File_backend_v1_judge_service_proto protoreflect.FileDescriptor

var file_backend_v1_judge_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x75, 0x64,
	0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x5f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95,
	0x01, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x22, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6c, 0x61, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x2c, 0x0a,
	0x1a, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x68, 0x0a, 0x1b, 0x47,
	0x65, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x11, 0x73, 0x75,
	0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x10, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x79, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x5a, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x73,
	0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52,
	0x0b, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x3e, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x5c, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0e, 0x6a, 0x75, 0x64, 0x67,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x75,
	0x64, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0d, 0x6a, 0x75, 0x64,
	0x67, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x32, 0xf4, 0x02, 0x0a, 0x0c, 0x4a,
	0x75, 0x64, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x12, 0x19, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x26, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x22, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x5f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x23, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30,
	0x01, 0x42, 0xb3, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x7a, 0x70, 0x70, 0x2d, 0x64, 0x65, 0x76, 0x2d, 0x74,
	0x65, 0x61, 0x6d, 0x2f, 0x73, 0x7a, 0x70, 0x70, 0x2d, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x42, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x16, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x42, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_backend_v1_judge_service_proto_rawDescOnce sync.Once
	file_backend_v1_judge_service_proto_rawDescData = file_backend_v1_judge_service_proto_rawDesc
)

func file_backend_v1_judge_service_proto_rawDescGZIP() []byte {
	file_backend_v1_judge_service_proto_rawDescOnce.Do(func() {
		file_backend_v1_judge_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_backend_v1_judge_service_proto_rawDescData)
	})
	return file_backend_v1_judge_service_proto_rawDescData
}

var file_backend_v1_judge_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_backend_v1_judge_service_proto_goTypes = []interface{}{
	(*SubmitRequest)(nil),               // 0: backend.v1.SubmitRequest
	(*SubmitResponse)(nil),              // 1: backend.v1.SubmitResponse
	(*GetSubmissionDetailRequest)(nil),  // 2: backend.v1.GetSubmissionDetailRequest
	(*GetSubmissionDetailResponse)(nil), // 3: backend.v1.GetSubmissionDetailResponse
	(*ListSubmissionsRequest)(nil),      // 4: backend.v1.ListSubmissionsRequest
	(*ListSubmissionsResponse)(nil),     // 5: backend.v1.ListSubmissionsResponse
	(*GetJudgeProgressRequest)(nil),     // 6: backend.v1.GetJudgeProgressRequest
	(*GetJudgeProgressResponse)(nil),    // 7: backend.v1.GetJudgeProgressResponse
	(*SubmissionDetail)(nil),            // 8: backend.v1.SubmissionDetail
	(*SubmissionSummary)(nil),           // 9: backend.v1.SubmissionSummary
	(*JudgeProgress)(nil),               // 10: backend.v1.JudgeProgress
}
var file_backend_v1_judge_service_proto_depIdxs = []int32{
	8,  // 0: backend.v1.GetSubmissionDetailResponse.submission_detail:type_name -> backend.v1.SubmissionDetail
	9,  // 1: backend.v1.ListSubmissionsResponse.submissions:type_name -> backend.v1.SubmissionSummary
	10, // 2: backend.v1.GetJudgeProgressResponse.judge_progress:type_name -> backend.v1.JudgeProgress
	0,  // 3: backend.v1.JudgeService.Submit:input_type -> backend.v1.SubmitRequest
	2,  // 4: backend.v1.JudgeService.GetSubmissionDetail:input_type -> backend.v1.GetSubmissionDetailRequest
	4,  // 5: backend.v1.JudgeService.ListSubmissions:input_type -> backend.v1.ListSubmissionsRequest
	6,  // 6: backend.v1.JudgeService.GetJudgeProgress:input_type -> backend.v1.GetJudgeProgressRequest
	1,  // 7: backend.v1.JudgeService.Submit:output_type -> backend.v1.SubmitResponse
	3,  // 8: backend.v1.JudgeService.GetSubmissionDetail:output_type -> backend.v1.GetSubmissionDetailResponse
	5,  // 9: backend.v1.JudgeService.ListSubmissions:output_type -> backend.v1.ListSubmissionsResponse
	7,  // 10: backend.v1.JudgeService.GetJudgeProgress:output_type -> backend.v1.GetJudgeProgressResponse
	7,  // [7:11] is the sub-list for method output_type
	3,  // [3:7] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_backend_v1_judge_service_proto_init() }
func file_backend_v1_judge_service_proto_init() {
	if File_backend_v1_judge_service_proto != nil {
		return
	}
	file_backend_v1_judge_resources_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_backend_v1_judge_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_backend_v1_judge_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_backend_v1_judge_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubmissionDetailRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_backend_v1_judge_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubmissionDetailResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_backend_v1_judge_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSubmissionsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_backend_v1_judge_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSubmissionsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_backend_v1_judge_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJudgeProgressRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_backend_v1_judge_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJudgeProgressResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_backend_v1_judge_service_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_backend_v1_judge_service_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_backend_v1_judge_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_backend_v1_judge_service_proto_goTypes,
		DependencyIndexes: file_backend_v1_judge_service_proto_depIdxs,
		MessageInfos:      file_backend_v1_judge_service_proto_msgTypes,
	}.Build()
	File_backend_v1_judge_service_proto = out.File
	file_backend_v1_judge_service_proto_rawDesc = nil
	file_backend_v1_judge_service_proto_goTypes = nil
	file_backend_v1_judge_service_proto_depIdxs = nil
}
