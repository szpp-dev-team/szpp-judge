// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: judge/v1/judge_type.proto

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

// 行ごとの一致比較 (最後の改行忘れは許容, 改行コードの違いも許容)
type JudgeTypeNormal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CaseInsensitive *bool `protobuf:"varint,1,opt,name=case_insensitive,json=caseInsensitive,proto3,oneof" json:"case_insensitive,omitempty"`
}

func (x *JudgeTypeNormal) Reset() {
	*x = JudgeTypeNormal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_judge_v1_judge_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JudgeTypeNormal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JudgeTypeNormal) ProtoMessage() {}

func (x *JudgeTypeNormal) ProtoReflect() protoreflect.Message {
	mi := &file_judge_v1_judge_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JudgeTypeNormal.ProtoReflect.Descriptor instead.
func (*JudgeTypeNormal) Descriptor() ([]byte, []int) {
	return file_judge_v1_judge_type_proto_rawDescGZIP(), []int{0}
}

func (x *JudgeTypeNormal) GetCaseInsensitive() bool {
	if x != nil && x.CaseInsensitive != nil {
		return *x.CaseInsensitive
	}
	return false
}

// 浮動小数点数向けの許容誤差付き比較
type JudgeTypeEPS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 許容誤差の桁数
	// 例えば ndigits=5 なら絶対誤差または相対誤差のいずれかが 1e-5 未満なら正解
	Ndigits uint32 `protobuf:"varint,1,opt,name=ndigits,proto3" json:"ndigits,omitempty"`
}

func (x *JudgeTypeEPS) Reset() {
	*x = JudgeTypeEPS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_judge_v1_judge_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JudgeTypeEPS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JudgeTypeEPS) ProtoMessage() {}

func (x *JudgeTypeEPS) ProtoReflect() protoreflect.Message {
	mi := &file_judge_v1_judge_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JudgeTypeEPS.ProtoReflect.Descriptor instead.
func (*JudgeTypeEPS) Descriptor() ([]byte, []int) {
	return file_judge_v1_judge_type_proto_rawDescGZIP(), []int{1}
}

func (x *JudgeTypeEPS) GetNdigits() uint32 {
	if x != nil {
		return x.Ndigits
	}
	return 0
}

// インタラクティブジャッジ
// 提出プログラムとジャッジプログラムとでリアルタイムに入出力をやりとりする
type JudgeTypeInteractive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JudgeCodePath string `protobuf:"bytes,1,opt,name=judge_code_path,json=judgeCodePath,proto3" json:"judge_code_path,omitempty"`
}

func (x *JudgeTypeInteractive) Reset() {
	*x = JudgeTypeInteractive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_judge_v1_judge_type_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JudgeTypeInteractive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JudgeTypeInteractive) ProtoMessage() {}

func (x *JudgeTypeInteractive) ProtoReflect() protoreflect.Message {
	mi := &file_judge_v1_judge_type_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JudgeTypeInteractive.ProtoReflect.Descriptor instead.
func (*JudgeTypeInteractive) Descriptor() ([]byte, []int) {
	return file_judge_v1_judge_type_proto_rawDescGZIP(), []int{2}
}

func (x *JudgeTypeInteractive) GetJudgeCodePath() string {
	if x != nil {
		return x.JudgeCodePath
	}
	return ""
}

// スペシャルジャッジ
// 正解が複数考えられる場合にはこのジャッジ形式を使う
type JudgeTypeCustom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JudgeCodePath string `protobuf:"bytes,1,opt,name=judge_code_path,json=judgeCodePath,proto3" json:"judge_code_path,omitempty"`
}

func (x *JudgeTypeCustom) Reset() {
	*x = JudgeTypeCustom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_judge_v1_judge_type_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JudgeTypeCustom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JudgeTypeCustom) ProtoMessage() {}

func (x *JudgeTypeCustom) ProtoReflect() protoreflect.Message {
	mi := &file_judge_v1_judge_type_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JudgeTypeCustom.ProtoReflect.Descriptor instead.
func (*JudgeTypeCustom) Descriptor() ([]byte, []int) {
	return file_judge_v1_judge_type_proto_rawDescGZIP(), []int{3}
}

func (x *JudgeTypeCustom) GetJudgeCodePath() string {
	if x != nil {
		return x.JudgeCodePath
	}
	return ""
}

type JudgeType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to JudgeType:
	//
	//	*JudgeType_Normal
	//	*JudgeType_Eps
	//	*JudgeType_Interactive
	//	*JudgeType_Custom
	JudgeType isJudgeType_JudgeType `protobuf_oneof:"judge_type"`
}

func (x *JudgeType) Reset() {
	*x = JudgeType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_judge_v1_judge_type_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JudgeType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JudgeType) ProtoMessage() {}

func (x *JudgeType) ProtoReflect() protoreflect.Message {
	mi := &file_judge_v1_judge_type_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JudgeType.ProtoReflect.Descriptor instead.
func (*JudgeType) Descriptor() ([]byte, []int) {
	return file_judge_v1_judge_type_proto_rawDescGZIP(), []int{4}
}

func (m *JudgeType) GetJudgeType() isJudgeType_JudgeType {
	if m != nil {
		return m.JudgeType
	}
	return nil
}

func (x *JudgeType) GetNormal() *JudgeTypeNormal {
	if x, ok := x.GetJudgeType().(*JudgeType_Normal); ok {
		return x.Normal
	}
	return nil
}

func (x *JudgeType) GetEps() *JudgeTypeEPS {
	if x, ok := x.GetJudgeType().(*JudgeType_Eps); ok {
		return x.Eps
	}
	return nil
}

func (x *JudgeType) GetInteractive() *JudgeTypeInteractive {
	if x, ok := x.GetJudgeType().(*JudgeType_Interactive); ok {
		return x.Interactive
	}
	return nil
}

func (x *JudgeType) GetCustom() *JudgeTypeCustom {
	if x, ok := x.GetJudgeType().(*JudgeType_Custom); ok {
		return x.Custom
	}
	return nil
}

type isJudgeType_JudgeType interface {
	isJudgeType_JudgeType()
}

type JudgeType_Normal struct {
	Normal *JudgeTypeNormal `protobuf:"bytes,1,opt,name=normal,proto3,oneof"`
}

type JudgeType_Eps struct {
	Eps *JudgeTypeEPS `protobuf:"bytes,2,opt,name=eps,proto3,oneof"`
}

type JudgeType_Interactive struct {
	Interactive *JudgeTypeInteractive `protobuf:"bytes,3,opt,name=interactive,proto3,oneof"`
}

type JudgeType_Custom struct {
	Custom *JudgeTypeCustom `protobuf:"bytes,4,opt,name=custom,proto3,oneof"`
}

func (*JudgeType_Normal) isJudgeType_JudgeType() {}

func (*JudgeType_Eps) isJudgeType_JudgeType() {}

func (*JudgeType_Interactive) isJudgeType_JudgeType() {}

func (*JudgeType_Custom) isJudgeType_JudgeType() {}

var File_judge_v1_judge_type_proto protoreflect.FileDescriptor

var file_judge_v1_judge_type_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6a, 0x75, 0x64,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x56, 0x0a, 0x0f, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x12, 0x2e, 0x0a, 0x10, 0x63, 0x61, 0x73, 0x65,
	0x5f, 0x69, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x73, 0x65, 0x6e, 0x73,
	0x69, 0x74, 0x69, 0x76, 0x65, 0x88, 0x01, 0x01, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x63, 0x61, 0x73,
	0x65, 0x5f, 0x69, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x22, 0x28, 0x0a,
	0x0c, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x45, 0x50, 0x53, 0x12, 0x18, 0x0a,
	0x07, 0x6e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x6e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x73, 0x22, 0x3e, 0x0a, 0x14, 0x4a, 0x75, 0x64, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12,
	0x26, 0x0a, 0x0f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0x39, 0x0a, 0x0f, 0x4a, 0x75, 0x64, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x12, 0x26, 0x0a, 0x0f, 0x6a, 0x75,
	0x64, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x22, 0xf3, 0x01, 0x0a, 0x09, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x33, 0x0a, 0x06, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x75, 0x64, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x06, 0x6e,
	0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x12, 0x2a, 0x0a, 0x03, 0x65, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x75,
	0x64, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x45, 0x50, 0x53, 0x48, 0x00, 0x52, 0x03, 0x65, 0x70,
	0x73, 0x12, 0x42, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x48, 0x00, 0x52, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x42, 0x0c, 0x0a, 0x0a, 0x6a, 0x75,
	0x64, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0xa2, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d,
	0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x4a, 0x75, 0x64, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x7a, 0x70, 0x70, 0x2d, 0x64, 0x65, 0x76,
	0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x73, 0x7a, 0x70, 0x70, 0x2d, 0x6a, 0x75, 0x64, 0x67, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x6a, 0x75,
	0x64, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x4a, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x08, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x4a, 0x75, 0x64,
	0x67, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x09, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_judge_v1_judge_type_proto_rawDescOnce sync.Once
	file_judge_v1_judge_type_proto_rawDescData = file_judge_v1_judge_type_proto_rawDesc
)

func file_judge_v1_judge_type_proto_rawDescGZIP() []byte {
	file_judge_v1_judge_type_proto_rawDescOnce.Do(func() {
		file_judge_v1_judge_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_judge_v1_judge_type_proto_rawDescData)
	})
	return file_judge_v1_judge_type_proto_rawDescData
}

var file_judge_v1_judge_type_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_judge_v1_judge_type_proto_goTypes = []interface{}{
	(*JudgeTypeNormal)(nil),      // 0: judge.v1.JudgeTypeNormal
	(*JudgeTypeEPS)(nil),         // 1: judge.v1.JudgeTypeEPS
	(*JudgeTypeInteractive)(nil), // 2: judge.v1.JudgeTypeInteractive
	(*JudgeTypeCustom)(nil),      // 3: judge.v1.JudgeTypeCustom
	(*JudgeType)(nil),            // 4: judge.v1.JudgeType
}
var file_judge_v1_judge_type_proto_depIdxs = []int32{
	0, // 0: judge.v1.JudgeType.normal:type_name -> judge.v1.JudgeTypeNormal
	1, // 1: judge.v1.JudgeType.eps:type_name -> judge.v1.JudgeTypeEPS
	2, // 2: judge.v1.JudgeType.interactive:type_name -> judge.v1.JudgeTypeInteractive
	3, // 3: judge.v1.JudgeType.custom:type_name -> judge.v1.JudgeTypeCustom
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_judge_v1_judge_type_proto_init() }
func file_judge_v1_judge_type_proto_init() {
	if File_judge_v1_judge_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_judge_v1_judge_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JudgeTypeNormal); i {
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
		file_judge_v1_judge_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JudgeTypeEPS); i {
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
		file_judge_v1_judge_type_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JudgeTypeInteractive); i {
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
		file_judge_v1_judge_type_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JudgeTypeCustom); i {
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
		file_judge_v1_judge_type_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JudgeType); i {
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
	file_judge_v1_judge_type_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_judge_v1_judge_type_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*JudgeType_Normal)(nil),
		(*JudgeType_Eps)(nil),
		(*JudgeType_Interactive)(nil),
		(*JudgeType_Custom)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_judge_v1_judge_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_judge_v1_judge_type_proto_goTypes,
		DependencyIndexes: file_judge_v1_judge_type_proto_depIdxs,
		MessageInfos:      file_judge_v1_judge_type_proto_msgTypes,
	}.Build()
	File_judge_v1_judge_type_proto = out.File
	file_judge_v1_judge_type_proto_rawDesc = nil
	file_judge_v1_judge_type_proto_goTypes = nil
	file_judge_v1_judge_type_proto_depIdxs = nil
}
