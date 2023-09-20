// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: backend/v1/task_service.proto

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

type CreateTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task *MutationTask `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *CreateTaskRequest) Reset() {
	*x = CreateTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_v1_task_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequest) ProtoMessage() {}

func (x *CreateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_v1_task_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return file_backend_v1_task_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTaskRequest) GetTask() *MutationTask {
	if x != nil {
		return x.Task
	}
	return nil
}

type CreateTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task *Task `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *CreateTaskResponse) Reset() {
	*x = CreateTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_v1_task_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskResponse) ProtoMessage() {}

func (x *CreateTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_v1_task_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskResponse.ProtoReflect.Descriptor instead.
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return file_backend_v1_task_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTaskResponse) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type GetTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId int32 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *GetTaskRequest) Reset() {
	*x = GetTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_v1_task_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskRequest) ProtoMessage() {}

func (x *GetTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_v1_task_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskRequest.ProtoReflect.Descriptor instead.
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return file_backend_v1_task_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetTaskRequest) GetTaskId() int32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

type GetTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task *Task `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *GetTaskResponse) Reset() {
	*x = GetTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_v1_task_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskResponse) ProtoMessage() {}

func (x *GetTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_v1_task_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskResponse.ProtoReflect.Descriptor instead.
func (*GetTaskResponse) Descriptor() ([]byte, []int) {
	return file_backend_v1_task_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetTaskResponse) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type UpdateTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId int32         `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Task   *MutationTask `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *UpdateTaskRequest) Reset() {
	*x = UpdateTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_v1_task_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskRequest) ProtoMessage() {}

func (x *UpdateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_v1_task_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskRequest.ProtoReflect.Descriptor instead.
func (*UpdateTaskRequest) Descriptor() ([]byte, []int) {
	return file_backend_v1_task_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTaskRequest) GetTaskId() int32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *UpdateTaskRequest) GetTask() *MutationTask {
	if x != nil {
		return x.Task
	}
	return nil
}

type UpdateTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task *Task `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *UpdateTaskResponse) Reset() {
	*x = UpdateTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_v1_task_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskResponse) ProtoMessage() {}

func (x *UpdateTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_v1_task_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskResponse.ProtoReflect.Descriptor instead.
func (*UpdateTaskResponse) Descriptor() ([]byte, []int) {
	return file_backend_v1_task_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateTaskResponse) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type GetTestcaseSetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId int32 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *GetTestcaseSetsRequest) Reset() {
	*x = GetTestcaseSetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_v1_task_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTestcaseSetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTestcaseSetsRequest) ProtoMessage() {}

func (x *GetTestcaseSetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_v1_task_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTestcaseSetsRequest.ProtoReflect.Descriptor instead.
func (*GetTestcaseSetsRequest) Descriptor() ([]byte, []int) {
	return file_backend_v1_task_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetTestcaseSetsRequest) GetTaskId() int32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

type GetTestcaseSetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestcaseSets []*TestcaseSet `protobuf:"bytes,1,rep,name=testcase_sets,json=testcaseSets,proto3" json:"testcase_sets,omitempty"`
	Testcases    []*Testcase    `protobuf:"bytes,2,rep,name=testcases,proto3" json:"testcases,omitempty"`
}

func (x *GetTestcaseSetsResponse) Reset() {
	*x = GetTestcaseSetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_v1_task_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTestcaseSetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTestcaseSetsResponse) ProtoMessage() {}

func (x *GetTestcaseSetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_v1_task_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTestcaseSetsResponse.ProtoReflect.Descriptor instead.
func (*GetTestcaseSetsResponse) Descriptor() ([]byte, []int) {
	return file_backend_v1_task_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetTestcaseSetsResponse) GetTestcaseSets() []*TestcaseSet {
	if x != nil {
		return x.TestcaseSets
	}
	return nil
}

func (x *GetTestcaseSetsResponse) GetTestcases() []*Testcase {
	if x != nil {
		return x.Testcases
	}
	return nil
}

type SyncTestcaseSetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId       int32                  `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	TestcaseSets []*MutationTestcaseSet `protobuf:"bytes,2,rep,name=testcase_sets,json=testcaseSets,proto3" json:"testcase_sets,omitempty"`
	Testcases    []*MutationTestcase    `protobuf:"bytes,3,rep,name=testcases,proto3" json:"testcases,omitempty"`
}

func (x *SyncTestcaseSetsRequest) Reset() {
	*x = SyncTestcaseSetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_v1_task_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncTestcaseSetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncTestcaseSetsRequest) ProtoMessage() {}

func (x *SyncTestcaseSetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_v1_task_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncTestcaseSetsRequest.ProtoReflect.Descriptor instead.
func (*SyncTestcaseSetsRequest) Descriptor() ([]byte, []int) {
	return file_backend_v1_task_service_proto_rawDescGZIP(), []int{8}
}

func (x *SyncTestcaseSetsRequest) GetTaskId() int32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *SyncTestcaseSetsRequest) GetTestcaseSets() []*MutationTestcaseSet {
	if x != nil {
		return x.TestcaseSets
	}
	return nil
}

func (x *SyncTestcaseSetsRequest) GetTestcases() []*MutationTestcase {
	if x != nil {
		return x.Testcases
	}
	return nil
}

type SyncTestcaseSetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestcaseSets []*TestcaseSet `protobuf:"bytes,1,rep,name=testcase_sets,json=testcaseSets,proto3" json:"testcase_sets,omitempty"`
	Testcases    []*Testcase    `protobuf:"bytes,2,rep,name=testcases,proto3" json:"testcases,omitempty"`
}

func (x *SyncTestcaseSetsResponse) Reset() {
	*x = SyncTestcaseSetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_v1_task_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncTestcaseSetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncTestcaseSetsResponse) ProtoMessage() {}

func (x *SyncTestcaseSetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_v1_task_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncTestcaseSetsResponse.ProtoReflect.Descriptor instead.
func (*SyncTestcaseSetsResponse) Descriptor() ([]byte, []int) {
	return file_backend_v1_task_service_proto_rawDescGZIP(), []int{9}
}

func (x *SyncTestcaseSetsResponse) GetTestcaseSets() []*TestcaseSet {
	if x != nil {
		return x.TestcaseSets
	}
	return nil
}

func (x *SyncTestcaseSetsResponse) GetTestcases() []*Testcase {
	if x != nil {
		return x.Testcases
	}
	return nil
}

var File_backend_v1_task_service_proto protoreflect.FileDescriptor

var file_backend_v1_task_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73,
	0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x75, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22,
	0x3a, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x29, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x22, 0x37, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x61, 0x73,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22,
	0x5a, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x2c, 0x0a,
	0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x3a, 0x0a, 0x12, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x31, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x54, 0x65,
	0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x53, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x22, 0x8b, 0x01, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x53, 0x65, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61,
	0x73, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x63,
	0x61, 0x73, 0x65, 0x53, 0x65, 0x74, 0x52, 0x0c, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65,
	0x53, 0x65, 0x74, 0x73, 0x12, 0x32, 0x0a, 0x09, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x52, 0x09, 0x74,
	0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x73, 0x22, 0xb4, 0x01, 0x0a, 0x17, 0x53, 0x79, 0x6e,
	0x63, 0x54, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x53, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x44, 0x0a,
	0x0d, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x63, 0x61,
	0x73, 0x65, 0x53, 0x65, 0x74, 0x52, 0x0c, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x53,
	0x65, 0x74, 0x73, 0x12, 0x3a, 0x0a, 0x09, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x73, 0x74,
	0x63, 0x61, 0x73, 0x65, 0x52, 0x09, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x73, 0x22,
	0x8c, 0x01, 0x0a, 0x18, 0x53, 0x79, 0x6e, 0x63, 0x54, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65,
	0x53, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0d,
	0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x53, 0x65, 0x74, 0x52, 0x0c, 0x74, 0x65,
	0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x53, 0x65, 0x74, 0x73, 0x12, 0x32, 0x0a, 0x09, 0x74, 0x65,
	0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x63,
	0x61, 0x73, 0x65, 0x52, 0x09, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x73, 0x32, 0xa6,
	0x03, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1d, 0x2e, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1a, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4b, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1d, 0x2e,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x53, 0x65, 0x74, 0x73, 0x12,
	0x22, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x53, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x53, 0x65, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x10, 0x53, 0x79, 0x6e, 0x63,
	0x54, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x53, 0x65, 0x74, 0x73, 0x12, 0x23, 0x2e, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x54, 0x65,
	0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x53, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x79, 0x6e, 0x63, 0x54, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x53, 0x65, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xb2, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x54, 0x61, 0x73, 0x6b,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x7a, 0x70, 0x70, 0x2d,
	0x64, 0x65, 0x76, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x73, 0x7a, 0x70, 0x70, 0x2d, 0x6a, 0x75,
	0x64, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x42, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0b, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_backend_v1_task_service_proto_rawDescOnce sync.Once
	file_backend_v1_task_service_proto_rawDescData = file_backend_v1_task_service_proto_rawDesc
)

func file_backend_v1_task_service_proto_rawDescGZIP() []byte {
	file_backend_v1_task_service_proto_rawDescOnce.Do(func() {
		file_backend_v1_task_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_backend_v1_task_service_proto_rawDescData)
	})
	return file_backend_v1_task_service_proto_rawDescData
}

var file_backend_v1_task_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_backend_v1_task_service_proto_goTypes = []interface{}{
	(*CreateTaskRequest)(nil),        // 0: backend.v1.CreateTaskRequest
	(*CreateTaskResponse)(nil),       // 1: backend.v1.CreateTaskResponse
	(*GetTaskRequest)(nil),           // 2: backend.v1.GetTaskRequest
	(*GetTaskResponse)(nil),          // 3: backend.v1.GetTaskResponse
	(*UpdateTaskRequest)(nil),        // 4: backend.v1.UpdateTaskRequest
	(*UpdateTaskResponse)(nil),       // 5: backend.v1.UpdateTaskResponse
	(*GetTestcaseSetsRequest)(nil),   // 6: backend.v1.GetTestcaseSetsRequest
	(*GetTestcaseSetsResponse)(nil),  // 7: backend.v1.GetTestcaseSetsResponse
	(*SyncTestcaseSetsRequest)(nil),  // 8: backend.v1.SyncTestcaseSetsRequest
	(*SyncTestcaseSetsResponse)(nil), // 9: backend.v1.SyncTestcaseSetsResponse
	(*MutationTask)(nil),             // 10: backend.v1.MutationTask
	(*Task)(nil),                     // 11: backend.v1.Task
	(*TestcaseSet)(nil),              // 12: backend.v1.TestcaseSet
	(*Testcase)(nil),                 // 13: backend.v1.Testcase
	(*MutationTestcaseSet)(nil),      // 14: backend.v1.MutationTestcaseSet
	(*MutationTestcase)(nil),         // 15: backend.v1.MutationTestcase
}
var file_backend_v1_task_service_proto_depIdxs = []int32{
	10, // 0: backend.v1.CreateTaskRequest.task:type_name -> backend.v1.MutationTask
	11, // 1: backend.v1.CreateTaskResponse.task:type_name -> backend.v1.Task
	11, // 2: backend.v1.GetTaskResponse.task:type_name -> backend.v1.Task
	10, // 3: backend.v1.UpdateTaskRequest.task:type_name -> backend.v1.MutationTask
	11, // 4: backend.v1.UpdateTaskResponse.task:type_name -> backend.v1.Task
	12, // 5: backend.v1.GetTestcaseSetsResponse.testcase_sets:type_name -> backend.v1.TestcaseSet
	13, // 6: backend.v1.GetTestcaseSetsResponse.testcases:type_name -> backend.v1.Testcase
	14, // 7: backend.v1.SyncTestcaseSetsRequest.testcase_sets:type_name -> backend.v1.MutationTestcaseSet
	15, // 8: backend.v1.SyncTestcaseSetsRequest.testcases:type_name -> backend.v1.MutationTestcase
	12, // 9: backend.v1.SyncTestcaseSetsResponse.testcase_sets:type_name -> backend.v1.TestcaseSet
	13, // 10: backend.v1.SyncTestcaseSetsResponse.testcases:type_name -> backend.v1.Testcase
	0,  // 11: backend.v1.TaskService.CreateTask:input_type -> backend.v1.CreateTaskRequest
	2,  // 12: backend.v1.TaskService.GetTask:input_type -> backend.v1.GetTaskRequest
	4,  // 13: backend.v1.TaskService.UpdateTask:input_type -> backend.v1.UpdateTaskRequest
	6,  // 14: backend.v1.TaskService.GetTestcaseSets:input_type -> backend.v1.GetTestcaseSetsRequest
	8,  // 15: backend.v1.TaskService.SyncTestcaseSets:input_type -> backend.v1.SyncTestcaseSetsRequest
	1,  // 16: backend.v1.TaskService.CreateTask:output_type -> backend.v1.CreateTaskResponse
	3,  // 17: backend.v1.TaskService.GetTask:output_type -> backend.v1.GetTaskResponse
	5,  // 18: backend.v1.TaskService.UpdateTask:output_type -> backend.v1.UpdateTaskResponse
	7,  // 19: backend.v1.TaskService.GetTestcaseSets:output_type -> backend.v1.GetTestcaseSetsResponse
	9,  // 20: backend.v1.TaskService.SyncTestcaseSets:output_type -> backend.v1.SyncTestcaseSetsResponse
	16, // [16:21] is the sub-list for method output_type
	11, // [11:16] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_backend_v1_task_service_proto_init() }
func file_backend_v1_task_service_proto_init() {
	if File_backend_v1_task_service_proto != nil {
		return
	}
	file_backend_v1_task_resources_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_backend_v1_task_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaskRequest); i {
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
		file_backend_v1_task_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaskResponse); i {
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
		file_backend_v1_task_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTaskRequest); i {
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
		file_backend_v1_task_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTaskResponse); i {
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
		file_backend_v1_task_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTaskRequest); i {
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
		file_backend_v1_task_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTaskResponse); i {
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
		file_backend_v1_task_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTestcaseSetsRequest); i {
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
		file_backend_v1_task_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTestcaseSetsResponse); i {
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
		file_backend_v1_task_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncTestcaseSetsRequest); i {
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
		file_backend_v1_task_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncTestcaseSetsResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_backend_v1_task_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_backend_v1_task_service_proto_goTypes,
		DependencyIndexes: file_backend_v1_task_service_proto_depIdxs,
		MessageInfos:      file_backend_v1_task_service_proto_msgTypes,
	}.Build()
	File_backend_v1_task_service_proto = out.File
	file_backend_v1_task_service_proto_rawDesc = nil
	file_backend_v1_task_service_proto_goTypes = nil
	file_backend_v1_task_service_proto_depIdxs = nil
}