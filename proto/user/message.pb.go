// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.1
// source: proto/user/message.proto

package user

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

type UserInfo_Gender int32

const (
	UserInfo_UNKNOWN UserInfo_Gender = 0
	UserInfo_MALE    UserInfo_Gender = 1
	UserInfo_FEMALE  UserInfo_Gender = 2
)

// Enum value maps for UserInfo_Gender.
var (
	UserInfo_Gender_name = map[int32]string{
		0: "UNKNOWN",
		1: "MALE",
		2: "FEMALE",
	}
	UserInfo_Gender_value = map[string]int32{
		"UNKNOWN": 0,
		"MALE":    1,
		"FEMALE":  2,
	}
)

func (x UserInfo_Gender) Enum() *UserInfo_Gender {
	p := new(UserInfo_Gender)
	*p = x
	return p
}

func (x UserInfo_Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserInfo_Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_user_message_proto_enumTypes[0].Descriptor()
}

func (UserInfo_Gender) Type() protoreflect.EnumType {
	return &file_proto_user_message_proto_enumTypes[0]
}

func (x UserInfo_Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserInfo_Gender.Descriptor instead.
func (UserInfo_Gender) EnumDescriptor() ([]byte, []int) {
	return file_proto_user_message_proto_rawDescGZIP(), []int{1, 0}
}

type UserID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserID) Reset() {
	*x = UserID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserID) ProtoMessage() {}

func (x *UserID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserID.ProtoReflect.Descriptor instead.
func (*UserID) Descriptor() ([]byte, []int) {
	return file_proto_user_message_proto_rawDescGZIP(), []int{0}
}

func (x *UserID) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age    int32           `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Gender UserInfo_Gender `protobuf:"varint,4,opt,name=gender,proto3,enum=user.service.v1.UserInfo_Gender" json:"gender,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_proto_user_message_proto_rawDescGZIP(), []int{1}
}

func (x *UserInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserInfo) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UserInfo) GetGender() UserInfo_Gender {
	if x != nil {
		return x.Gender
	}
	return UserInfo_UNKNOWN
}

var File_proto_user_message_proto protoreflect.FileDescriptor

var file_proto_user_message_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x18, 0x0a, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa7, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x22, 0x2b, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x4c,
	0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x42,
	0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x61,
	0x6f, 0x2d, 0x79, 0x69, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x64, 0x65, 0x6d,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_user_message_proto_rawDescOnce sync.Once
	file_proto_user_message_proto_rawDescData = file_proto_user_message_proto_rawDesc
)

func file_proto_user_message_proto_rawDescGZIP() []byte {
	file_proto_user_message_proto_rawDescOnce.Do(func() {
		file_proto_user_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_user_message_proto_rawDescData)
	})
	return file_proto_user_message_proto_rawDescData
}

var file_proto_user_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_user_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_user_message_proto_goTypes = []interface{}{
	(UserInfo_Gender)(0), // 0: user.service.v1.UserInfo.Gender
	(*UserID)(nil),       // 1: user.service.v1.UserID
	(*UserInfo)(nil),     // 2: user.service.v1.UserInfo
}
var file_proto_user_message_proto_depIdxs = []int32{
	0, // 0: user.service.v1.UserInfo.gender:type_name -> user.service.v1.UserInfo.Gender
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_user_message_proto_init() }
func file_proto_user_message_proto_init() {
	if File_proto_user_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_user_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserID); i {
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
		file_proto_user_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
			RawDescriptor: file_proto_user_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_user_message_proto_goTypes,
		DependencyIndexes: file_proto_user_message_proto_depIdxs,
		EnumInfos:         file_proto_user_message_proto_enumTypes,
		MessageInfos:      file_proto_user_message_proto_msgTypes,
	}.Build()
	File_proto_user_message_proto = out.File
	file_proto_user_message_proto_rawDesc = nil
	file_proto_user_message_proto_goTypes = nil
	file_proto_user_message_proto_depIdxs = nil
}
