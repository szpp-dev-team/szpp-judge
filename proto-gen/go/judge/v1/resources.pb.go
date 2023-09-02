// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: judge/v1/resources.proto

package judgev1

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

type JudgeStatus int32

const (
	JudgeStatus_JUDGE_STATUS_UNSPECIFIED JudgeStatus = 0
	JudgeStatus_AC                       JudgeStatus = 1
	JudgeStatus_CE                       JudgeStatus = 2
	JudgeStatus_IE                       JudgeStatus = 3
	JudgeStatus_MLE                      JudgeStatus = 4
	JudgeStatus_OLE                      JudgeStatus = 5
	JudgeStatus_TLE                      JudgeStatus = 6
	JudgeStatus_WA                       JudgeStatus = 7
)

// Enum value maps for JudgeStatus.
var (
	JudgeStatus_name = map[int32]string{
		0: "JUDGE_STATUS_UNSPECIFIED",
		1: "AC",
		2: "CE",
		3: "IE",
		4: "MLE",
		5: "OLE",
		6: "TLE",
		7: "WA",
	}
	JudgeStatus_value = map[string]int32{
		"JUDGE_STATUS_UNSPECIFIED": 0,
		"AC":                       1,
		"CE":                       2,
		"IE":                       3,
		"MLE":                      4,
		"OLE":                      5,
		"TLE":                      6,
		"WA":                       7,
	}
)

func (x JudgeStatus) Enum() *JudgeStatus {
	p := new(JudgeStatus)
	*p = x
	return p
}

func (x JudgeStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JudgeStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_judge_v1_resources_proto_enumTypes[0].Descriptor()
}

func (JudgeStatus) Type() protoreflect.EnumType {
	return &file_judge_v1_resources_proto_enumTypes[0]
}

func (x JudgeStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JudgeStatus.Descriptor instead.
func (JudgeStatus) EnumDescriptor() ([]byte, []int) {
	return file_judge_v1_resources_proto_rawDescGZIP(), []int{0}
}

type Testcase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InputPath    string `protobuf:"bytes,1,opt,name=input_path,json=inputPath,proto3" json:"input_path,omitempty"`
	ExpectedPath string `protobuf:"bytes,2,opt,name=expected_path,json=expectedPath,proto3" json:"expected_path,omitempty"`
}

func (x *Testcase) Reset() {
	*x = Testcase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_judge_v1_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Testcase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Testcase) ProtoMessage() {}

func (x *Testcase) ProtoReflect() protoreflect.Message {
	mi := &file_judge_v1_resources_proto_msgTypes[0]
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
	return file_judge_v1_resources_proto_rawDescGZIP(), []int{0}
}

func (x *Testcase) GetInputPath() string {
	if x != nil {
		return x.InputPath
	}
	return ""
}

func (x *Testcase) GetExpectedPath() string {
	if x != nil {
		return x.ExpectedPath
	}
	return ""
}

type ExecutionResultDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExitCode      *uint32 `protobuf:"varint,1,opt,name=exit_code,json=exitCode,proto3,oneof" json:"exit_code,omitempty"` // TLE 等で kill した場合は exit_code はナシ
	Stdin         string  `protobuf:"bytes,2,opt,name=stdin,proto3" json:"stdin,omitempty"`
	Stdout        string  `protobuf:"bytes,3,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr        string  `protobuf:"bytes,4,opt,name=stderr,proto3" json:"stderr,omitempty"`
	StdoutOmitted bool    `protobuf:"varint,5,opt,name=stdout_omitted,json=stdoutOmitted,proto3" json:"stdout_omitted,omitempty"` // stdout が巨大すぎて省略した場合は true
	StderrOmitted bool    `protobuf:"varint,6,opt,name=stderr_omitted,json=stderrOmitted,proto3" json:"stderr_omitted,omitempty"` // stderr が巨大すぎて省略した場合は true
}

func (x *ExecutionResultDetail) Reset() {
	*x = ExecutionResultDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_judge_v1_resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutionResultDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionResultDetail) ProtoMessage() {}

func (x *ExecutionResultDetail) ProtoReflect() protoreflect.Message {
	mi := &file_judge_v1_resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionResultDetail.ProtoReflect.Descriptor instead.
func (*ExecutionResultDetail) Descriptor() ([]byte, []int) {
	return file_judge_v1_resources_proto_rawDescGZIP(), []int{1}
}

func (x *ExecutionResultDetail) GetExitCode() uint32 {
	if x != nil && x.ExitCode != nil {
		return *x.ExitCode
	}
	return 0
}

func (x *ExecutionResultDetail) GetStdin() string {
	if x != nil {
		return x.Stdin
	}
	return ""
}

func (x *ExecutionResultDetail) GetStdout() string {
	if x != nil {
		return x.Stdout
	}
	return ""
}

func (x *ExecutionResultDetail) GetStderr() string {
	if x != nil {
		return x.Stderr
	}
	return ""
}

func (x *ExecutionResultDetail) GetStdoutOmitted() bool {
	if x != nil {
		return x.StdoutOmitted
	}
	return false
}

func (x *ExecutionResultDetail) GetStderrOmitted() bool {
	if x != nil {
		return x.StderrOmitted
	}
	return false
}

var File_judge_v1_resources_proto protoreflect.FileDescriptor

var file_judge_v1_resources_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6a, 0x75, 0x64, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x22, 0x4e, 0x0a, 0x08, 0x54, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x23, 0x0a, 0x0d, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x50, 0x61, 0x74, 0x68, 0x22, 0xdb, 0x01, 0x0a, 0x15, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x20,
	0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x48, 0x00, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x64, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x73, 0x74, 0x64, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74,
	0x5f, 0x6f, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x4f, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x12, 0x25, 0x0a,
	0x0e, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x5f, 0x6f, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x4f, 0x6d, 0x69,
	0x74, 0x74, 0x65, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x2a, 0x66, 0x0a, 0x0b, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1c, 0x0a, 0x18, 0x4a, 0x55, 0x44, 0x47, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x06, 0x0a, 0x02, 0x41, 0x43, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x43, 0x45, 0x10, 0x02, 0x12,
	0x06, 0x0a, 0x02, 0x49, 0x45, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x4c, 0x45, 0x10, 0x04,
	0x12, 0x07, 0x0a, 0x03, 0x4f, 0x4c, 0x45, 0x10, 0x05, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x4c, 0x45,
	0x10, 0x06, 0x12, 0x06, 0x0a, 0x02, 0x57, 0x41, 0x10, 0x07, 0x42, 0xa2, 0x01, 0x0a, 0x0c, 0x63,
	0x6f, 0x6d, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x7a, 0x70, 0x70, 0x2d, 0x64,
	0x65, 0x76, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x73, 0x7a, 0x70, 0x70, 0x2d, 0x6a, 0x75, 0x64,
	0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x4a, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x08, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x4a,
	0x75, 0x64, 0x67, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_judge_v1_resources_proto_rawDescOnce sync.Once
	file_judge_v1_resources_proto_rawDescData = file_judge_v1_resources_proto_rawDesc
)

func file_judge_v1_resources_proto_rawDescGZIP() []byte {
	file_judge_v1_resources_proto_rawDescOnce.Do(func() {
		file_judge_v1_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_judge_v1_resources_proto_rawDescData)
	})
	return file_judge_v1_resources_proto_rawDescData
}

var file_judge_v1_resources_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_judge_v1_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_judge_v1_resources_proto_goTypes = []interface{}{
	(JudgeStatus)(0),              // 0: judge.v1.JudgeStatus
	(*Testcase)(nil),              // 1: judge.v1.Testcase
	(*ExecutionResultDetail)(nil), // 2: judge.v1.ExecutionResultDetail
}
var file_judge_v1_resources_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_judge_v1_resources_proto_init() }
func file_judge_v1_resources_proto_init() {
	if File_judge_v1_resources_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_judge_v1_resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_judge_v1_resources_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutionResultDetail); i {
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
	file_judge_v1_resources_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_judge_v1_resources_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_judge_v1_resources_proto_goTypes,
		DependencyIndexes: file_judge_v1_resources_proto_depIdxs,
		EnumInfos:         file_judge_v1_resources_proto_enumTypes,
		MessageInfos:      file_judge_v1_resources_proto_msgTypes,
	}.Build()
	File_judge_v1_resources_proto = out.File
	file_judge_v1_resources_proto_rawDesc = nil
	file_judge_v1_resources_proto_goTypes = nil
	file_judge_v1_resources_proto_depIdxs = nil
}
