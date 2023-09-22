// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: backend/v1/task_resources.proto

package backendv1

import (
	v1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/judge/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Difficulty int32

const (
	Difficulty_DIFFICULTY_UNSPECIFIED Difficulty = 0
	Difficulty_BEGINNER               Difficulty = 1
	Difficulty_EASY                   Difficulty = 2
	Difficulty_MEDIUM                 Difficulty = 3
	Difficulty_HARD                   Difficulty = 4
	Difficulty_IMPOSSIBLE             Difficulty = 5
)

// Enum value maps for Difficulty.
var (
	Difficulty_name = map[int32]string{
		0: "DIFFICULTY_UNSPECIFIED",
		1: "BEGINNER",
		2: "EASY",
		3: "MEDIUM",
		4: "HARD",
		5: "IMPOSSIBLE",
	}
	Difficulty_value = map[string]int32{
		"DIFFICULTY_UNSPECIFIED": 0,
		"BEGINNER":               1,
		"EASY":                   2,
		"MEDIUM":                 3,
		"HARD":                   4,
		"IMPOSSIBLE":             5,
	}
)

func (x Difficulty) Enum() *Difficulty {
	p := new(Difficulty)
	*p = x
	return p
}

func (x Difficulty) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Difficulty) Descriptor() protoreflect.EnumDescriptor {
	return file_backend_v1_task_resources_proto_enumTypes[0].Descriptor()
}

func (Difficulty) Type() protoreflect.EnumType {
	return &file_backend_v1_task_resources_proto_enumTypes[0]
}

func (x Difficulty) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Difficulty.Descriptor instead.
func (Difficulty) EnumDescriptor() ([]byte, []int) {
	return file_backend_v1_task_resources_proto_rawDescGZIP(), []int{0}
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title           string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`                                               // 問題名
	Statement       string                 `protobuf:"bytes,3,opt,name=statement,proto3" json:"statement,omitempty"`                                       // 問題文
	ExecTimeLimit   int32                  `protobuf:"varint,4,opt,name=exec_time_limit,json=execTimeLimit,proto3" json:"exec_time_limit,omitempty"`       // 実行時間制限[ms]
	ExecMemoryLimit int32                  `protobuf:"varint,5,opt,name=exec_memory_limit,json=execMemoryLimit,proto3" json:"exec_memory_limit,omitempty"` // 実行メモリ制限[MB]
	JudgeType       *v1.JudgeType          `protobuf:"bytes,6,opt,name=judge_type,json=judgeType,proto3" json:"judge_type,omitempty"`                      // Judge の type(完全一致、誤差など)
	Difficulty      Difficulty             `protobuf:"varint,7,opt,name=difficulty,proto3,enum=backend.v1.Difficulty" json:"difficulty,omitempty"`
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3,oneof" json:"updated_at,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_v1_task_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_backend_v1_task_resources_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_backend_v1_task_resources_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Task) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Task) GetStatement() string {
	if x != nil {
		return x.Statement
	}
	return ""
}

func (x *Task) GetExecTimeLimit() int32 {
	if x != nil {
		return x.ExecTimeLimit
	}
	return 0
}

func (x *Task) GetExecMemoryLimit() int32 {
	if x != nil {
		return x.ExecMemoryLimit
	}
	return 0
}

func (x *Task) GetJudgeType() *v1.JudgeType {
	if x != nil {
		return x.JudgeType
	}
	return nil
}

func (x *Task) GetDifficulty() Difficulty {
	if x != nil {
		return x.Difficulty
	}
	return Difficulty_DIFFICULTY_UNSPECIFIED
}

func (x *Task) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Task) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type MutationTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              *int32        `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id,omitempty"`                                              // update の時のみ
	Title           string        `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`                                               // 問題名
	Statement       string        `protobuf:"bytes,3,opt,name=statement,proto3" json:"statement,omitempty"`                                       // 問題文
	ExecTimeLimit   int32         `protobuf:"varint,4,opt,name=exec_time_limit,json=execTimeLimit,proto3" json:"exec_time_limit,omitempty"`       // 実行時間制限[ms]
	ExecMemoryLimit int32         `protobuf:"varint,5,opt,name=exec_memory_limit,json=execMemoryLimit,proto3" json:"exec_memory_limit,omitempty"` // 実行メモリ制限[MB]
	JudgeType       *v1.JudgeType `protobuf:"bytes,6,opt,name=judge_type,json=judgeType,proto3" json:"judge_type,omitempty"`                      // Judge の type(完全一致、誤差など)
	Difficulty      Difficulty    `protobuf:"varint,7,opt,name=difficulty,proto3,enum=backend.v1.Difficulty" json:"difficulty,omitempty"`
	IsPublic        bool          `protobuf:"varint,8,opt,name=is_public,json=isPublic,proto3" json:"is_public,omitempty"`
}

func (x *MutationTask) Reset() {
	*x = MutationTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_v1_task_resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutationTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutationTask) ProtoMessage() {}

func (x *MutationTask) ProtoReflect() protoreflect.Message {
	mi := &file_backend_v1_task_resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutationTask.ProtoReflect.Descriptor instead.
func (*MutationTask) Descriptor() ([]byte, []int) {
	return file_backend_v1_task_resources_proto_rawDescGZIP(), []int{1}
}

func (x *MutationTask) GetId() int32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *MutationTask) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MutationTask) GetStatement() string {
	if x != nil {
		return x.Statement
	}
	return ""
}

func (x *MutationTask) GetExecTimeLimit() int32 {
	if x != nil {
		return x.ExecTimeLimit
	}
	return 0
}

func (x *MutationTask) GetExecMemoryLimit() int32 {
	if x != nil {
		return x.ExecMemoryLimit
	}
	return 0
}

func (x *MutationTask) GetJudgeType() *v1.JudgeType {
	if x != nil {
		return x.JudgeType
	}
	return nil
}

func (x *MutationTask) GetDifficulty() Difficulty {
	if x != nil {
		return x.Difficulty
	}
	return Difficulty_DIFFICULTY_UNSPECIFIED
}

func (x *MutationTask) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

// (slug, task_id) は複合 unique
type TestcaseSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Slug          string                 `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`                                // means name
	ScoreRatio    int32                  `protobuf:"varint,3,opt,name=score_ratio,json=scoreRatio,proto3" json:"score_ratio,omitempty"` // 得点比率(%)。TestcaseSet.score_ratio の総和が100になるように設定する。
	IsSample      bool                   `protobuf:"varint,4,opt,name=is_sample,json=isSample,proto3" json:"is_sample,omitempty"`
	TestcaseSlugs []string               `protobuf:"bytes,5,rep,name=testcase_slugs,json=testcaseSlugs,proto3" json:"testcase_slugs,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3,oneof" json:"updated_at,omitempty"`
	TaskId        int32                  `protobuf:"varint,8,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *TestcaseSet) Reset() {
	*x = TestcaseSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_v1_task_resources_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestcaseSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestcaseSet) ProtoMessage() {}

func (x *TestcaseSet) ProtoReflect() protoreflect.Message {
	mi := &file_backend_v1_task_resources_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestcaseSet.ProtoReflect.Descriptor instead.
func (*TestcaseSet) Descriptor() ([]byte, []int) {
	return file_backend_v1_task_resources_proto_rawDescGZIP(), []int{2}
}

func (x *TestcaseSet) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TestcaseSet) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *TestcaseSet) GetScoreRatio() int32 {
	if x != nil {
		return x.ScoreRatio
	}
	return 0
}

func (x *TestcaseSet) GetIsSample() bool {
	if x != nil {
		return x.IsSample
	}
	return false
}

func (x *TestcaseSet) GetTestcaseSlugs() []string {
	if x != nil {
		return x.TestcaseSlugs
	}
	return nil
}

func (x *TestcaseSet) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *TestcaseSet) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *TestcaseSet) GetTaskId() int32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

type MutationTestcaseSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slug          string   `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"` // means name
	ScoreRatio    int32    `protobuf:"varint,2,opt,name=score_ratio,json=scoreRatio,proto3" json:"score_ratio,omitempty"`
	IsSample      bool     `protobuf:"varint,3,opt,name=is_sample,json=isSample,proto3" json:"is_sample,omitempty"`
	TestcaseSlugs []string `protobuf:"bytes,4,rep,name=testcase_slugs,json=testcaseSlugs,proto3" json:"testcase_slugs,omitempty"`
}

func (x *MutationTestcaseSet) Reset() {
	*x = MutationTestcaseSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_v1_task_resources_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutationTestcaseSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutationTestcaseSet) ProtoMessage() {}

func (x *MutationTestcaseSet) ProtoReflect() protoreflect.Message {
	mi := &file_backend_v1_task_resources_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutationTestcaseSet.ProtoReflect.Descriptor instead.
func (*MutationTestcaseSet) Descriptor() ([]byte, []int) {
	return file_backend_v1_task_resources_proto_rawDescGZIP(), []int{3}
}

func (x *MutationTestcaseSet) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *MutationTestcaseSet) GetScoreRatio() int32 {
	if x != nil {
		return x.ScoreRatio
	}
	return 0
}

func (x *MutationTestcaseSet) GetIsSample() bool {
	if x != nil {
		return x.IsSample
	}
	return false
}

func (x *MutationTestcaseSet) GetTestcaseSlugs() []string {
	if x != nil {
		return x.TestcaseSlugs
	}
	return nil
}

// (slug, task_id) は複合 unique
type Testcase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Slug        string                 `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	Description *string                `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Input       string                 `protobuf:"bytes,4,opt,name=input,proto3" json:"input,omitempty"`
	Output      string                 `protobuf:"bytes,5,opt,name=output,proto3" json:"output,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3,oneof" json:"updated_at,omitempty"`
	TaskId      int32                  `protobuf:"varint,8,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *Testcase) Reset() {
	*x = Testcase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_v1_task_resources_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Testcase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Testcase) ProtoMessage() {}

func (x *Testcase) ProtoReflect() protoreflect.Message {
	mi := &file_backend_v1_task_resources_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Testcase.ProtoReflect.Descriptor instead.
func (*Testcase) Descriptor() ([]byte, []int) {
	return file_backend_v1_task_resources_proto_rawDescGZIP(), []int{4}
}

func (x *Testcase) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Testcase) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *Testcase) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *Testcase) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *Testcase) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

func (x *Testcase) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Testcase) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Testcase) GetTaskId() int32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

type MutationTestcase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slug        string  `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	Description *string `protobuf:"bytes,2,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Input       string  `protobuf:"bytes,3,opt,name=input,proto3" json:"input,omitempty"`
	Output      string  `protobuf:"bytes,4,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *MutationTestcase) Reset() {
	*x = MutationTestcase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_v1_task_resources_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutationTestcase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutationTestcase) ProtoMessage() {}

func (x *MutationTestcase) ProtoReflect() protoreflect.Message {
	mi := &file_backend_v1_task_resources_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutationTestcase.ProtoReflect.Descriptor instead.
func (*MutationTestcase) Descriptor() ([]byte, []int) {
	return file_backend_v1_task_resources_proto_rawDescGZIP(), []int{5}
}

func (x *MutationTestcase) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *MutationTestcase) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *MutationTestcase) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *MutationTestcase) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

var File_backend_v1_task_resources_proto protoreflect.FileDescriptor

var file_backend_v1_task_resources_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73,
	0x6b, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x03, 0x0a, 0x04, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x65, 0x78, 0x65, 0x63, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x65, 0x78, 0x65, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x2a,
	0x0a, 0x11, 0x65, 0x78, 0x65, 0x63, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x65, 0x78, 0x65, 0x63, 0x4d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x32, 0x0a, 0x0a, 0x6a, 0x75,
	0x64, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x09, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x36,
	0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x16, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x52, 0x0a, 0x64, 0x69, 0x66, 0x66,
	0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x3e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x48, 0x00, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01,
	0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x22, 0xbb, 0x02, 0x0a, 0x0c, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52,
	0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x65, 0x78,
	0x65, 0x63, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x65, 0x78, 0x65, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x78, 0x65, 0x63, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x65,
	0x78, 0x65, 0x63, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x32,
	0x0a, 0x0a, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x75,
	0x64, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x52, 0x0a,
	0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x22, 0xb9,
	0x02, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x53, 0x65, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c,
	0x75, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x61,
	0x74, 0x69, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x6c, 0x75,
	0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61,
	0x73, 0x65, 0x53, 0x6c, 0x75, 0x67, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x3e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x48, 0x00, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x8e, 0x01, 0x0a, 0x13, 0x4d,
	0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x53,
	0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x73, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x53, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65,
	0x5f, 0x73, 0x6c, 0x75, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x65,
	0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x53, 0x6c, 0x75, 0x67, 0x73, 0x22, 0xb6, 0x02, 0x0a, 0x08,
	0x54, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x25, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x88, 0x01, 0x01, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3e, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x07,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74,
	0x61, 0x73, 0x6b, 0x49, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x22, 0x8b, 0x01, 0x0a, 0x10, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x25, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x88, 0x01, 0x01, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x2a, 0x66, 0x0a, 0x0a, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79,
	0x12, 0x1a, 0x0a, 0x16, 0x44, 0x49, 0x46, 0x46, 0x49, 0x43, 0x55, 0x4c, 0x54, 0x59, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x42, 0x45, 0x47, 0x49, 0x4e, 0x4e, 0x45, 0x52, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x41,
	0x53, 0x59, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d, 0x10, 0x03,
	0x12, 0x08, 0x0a, 0x04, 0x48, 0x41, 0x52, 0x44, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4d,
	0x50, 0x4f, 0x53, 0x53, 0x49, 0x42, 0x4c, 0x45, 0x10, 0x05, 0x42, 0xb4, 0x01, 0x0a, 0x0e, 0x63,
	0x6f, 0x6d, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x7a, 0x70, 0x70, 0x2d, 0x64, 0x65, 0x76, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x73, 0x7a,
	0x70, 0x70, 0x2d, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x76, 0x31,
	0x3b, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58,
	0xaa, 0x02, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a,
	0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x42, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_backend_v1_task_resources_proto_rawDescOnce sync.Once
	file_backend_v1_task_resources_proto_rawDescData = file_backend_v1_task_resources_proto_rawDesc
)

func file_backend_v1_task_resources_proto_rawDescGZIP() []byte {
	file_backend_v1_task_resources_proto_rawDescOnce.Do(func() {
		file_backend_v1_task_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_backend_v1_task_resources_proto_rawDescData)
	})
	return file_backend_v1_task_resources_proto_rawDescData
}

var file_backend_v1_task_resources_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_backend_v1_task_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_backend_v1_task_resources_proto_goTypes = []interface{}{
	(Difficulty)(0),               // 0: backend.v1.Difficulty
	(*Task)(nil),                  // 1: backend.v1.Task
	(*MutationTask)(nil),          // 2: backend.v1.MutationTask
	(*TestcaseSet)(nil),           // 3: backend.v1.TestcaseSet
	(*MutationTestcaseSet)(nil),   // 4: backend.v1.MutationTestcaseSet
	(*Testcase)(nil),              // 5: backend.v1.Testcase
	(*MutationTestcase)(nil),      // 6: backend.v1.MutationTestcase
	(*v1.JudgeType)(nil),          // 7: judge.v1.JudgeType
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_backend_v1_task_resources_proto_depIdxs = []int32{
	7,  // 0: backend.v1.Task.judge_type:type_name -> judge.v1.JudgeType
	0,  // 1: backend.v1.Task.difficulty:type_name -> backend.v1.Difficulty
	8,  // 2: backend.v1.Task.created_at:type_name -> google.protobuf.Timestamp
	8,  // 3: backend.v1.Task.updated_at:type_name -> google.protobuf.Timestamp
	7,  // 4: backend.v1.MutationTask.judge_type:type_name -> judge.v1.JudgeType
	0,  // 5: backend.v1.MutationTask.difficulty:type_name -> backend.v1.Difficulty
	8,  // 6: backend.v1.TestcaseSet.created_at:type_name -> google.protobuf.Timestamp
	8,  // 7: backend.v1.TestcaseSet.updated_at:type_name -> google.protobuf.Timestamp
	8,  // 8: backend.v1.Testcase.created_at:type_name -> google.protobuf.Timestamp
	8,  // 9: backend.v1.Testcase.updated_at:type_name -> google.protobuf.Timestamp
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_backend_v1_task_resources_proto_init() }
func file_backend_v1_task_resources_proto_init() {
	if File_backend_v1_task_resources_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_backend_v1_task_resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_backend_v1_task_resources_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutationTask); i {
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
		file_backend_v1_task_resources_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestcaseSet); i {
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
		file_backend_v1_task_resources_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutationTestcaseSet); i {
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
		file_backend_v1_task_resources_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Testcase); i {
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
		file_backend_v1_task_resources_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutationTestcase); i {
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
	file_backend_v1_task_resources_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_backend_v1_task_resources_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_backend_v1_task_resources_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_backend_v1_task_resources_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_backend_v1_task_resources_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_backend_v1_task_resources_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_backend_v1_task_resources_proto_goTypes,
		DependencyIndexes: file_backend_v1_task_resources_proto_depIdxs,
		EnumInfos:         file_backend_v1_task_resources_proto_enumTypes,
		MessageInfos:      file_backend_v1_task_resources_proto_msgTypes,
	}.Build()
	File_backend_v1_task_resources_proto = out.File
	file_backend_v1_task_resources_proto_rawDesc = nil
	file_backend_v1_task_resources_proto_goTypes = nil
	file_backend_v1_task_resources_proto_depIdxs = nil
}
