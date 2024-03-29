// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: common_v2.proto

package v2

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

// ============================================================
//
//                            Enums
//
// ============================================================
type DevicePlatform int32

const (
	DevicePlatform_UNKNOWN DevicePlatform = 0
	DevicePlatform_iOS     DevicePlatform = 1
	DevicePlatform_ANDROID DevicePlatform = 2
	DevicePlatform_WEB     DevicePlatform = 3
)

// Enum value maps for DevicePlatform.
var (
	DevicePlatform_name = map[int32]string{
		0: "UNKNOWN",
		1: "iOS",
		2: "ANDROID",
		3: "WEB",
	}
	DevicePlatform_value = map[string]int32{
		"UNKNOWN": 0,
		"iOS":     1,
		"ANDROID": 2,
		"WEB":     3,
	}
)

func (x DevicePlatform) Enum() *DevicePlatform {
	p := new(DevicePlatform)
	*p = x
	return p
}

func (x DevicePlatform) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DevicePlatform) Descriptor() protoreflect.EnumDescriptor {
	return file_common_v2_proto_enumTypes[0].Descriptor()
}

func (DevicePlatform) Type() protoreflect.EnumType {
	return &file_common_v2_proto_enumTypes[0]
}

func (x DevicePlatform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DevicePlatform.Descriptor instead.
func (DevicePlatform) EnumDescriptor() ([]byte, []int) {
	return file_common_v2_proto_rawDescGZIP(), []int{0}
}

// A list of possible sensor types that can be passed to the config.
// All sensor types may not be available in the SDK. For SDK configuration,
// follow the SDK documentation for which sensors are available.
type SensorType int32

const (
	SensorType_UNKNOWN_SENSOR       SensorType = 0
	SensorType_LOCATION             SensorType = 1
	SensorType_ACCELEROMETER        SensorType = 2
	SensorType_LINEAR_ACCELEROMETER SensorType = 3
	SensorType_MAGNETOMETER         SensorType = 4
	SensorType_GYROSCOPE            SensorType = 5
	SensorType_BATTERY              SensorType = 6
	SensorType_ORIENTATION          SensorType = 7
	SensorType_POINTER              SensorType = 14
	// Deprecated: Do not use.
	SensorType_TEXT_CHANGE  SensorType = 15
	SensorType_KEY_PRESS    SensorType = 16
	SensorType_FOCUS_CHANGE SensorType = 17
	SensorType_VIEW_SCROLL  SensorType = 18
	SensorType_MOUSE_WHEEL  SensorType = 19
	SensorType_CLICK        SensorType = 20
	SensorType_INPUT_CHANGE SensorType = 21
	SensorType_FORM_SUBMIT  SensorType = 22
	SensorType_CONTEXT_MENU SensorType = 23
	SensorType_FRAME_RATE   SensorType = 24
)

// Enum value maps for SensorType.
var (
	SensorType_name = map[int32]string{
		0:  "UNKNOWN_SENSOR",
		1:  "LOCATION",
		2:  "ACCELEROMETER",
		3:  "LINEAR_ACCELEROMETER",
		4:  "MAGNETOMETER",
		5:  "GYROSCOPE",
		6:  "BATTERY",
		7:  "ORIENTATION",
		14: "POINTER",
		15: "TEXT_CHANGE",
		16: "KEY_PRESS",
		17: "FOCUS_CHANGE",
		18: "VIEW_SCROLL",
		19: "MOUSE_WHEEL",
		20: "CLICK",
		21: "INPUT_CHANGE",
		22: "FORM_SUBMIT",
		23: "CONTEXT_MENU",
		24: "FRAME_RATE",
	}
	SensorType_value = map[string]int32{
		"UNKNOWN_SENSOR":       0,
		"LOCATION":             1,
		"ACCELEROMETER":        2,
		"LINEAR_ACCELEROMETER": 3,
		"MAGNETOMETER":         4,
		"GYROSCOPE":            5,
		"BATTERY":              6,
		"ORIENTATION":          7,
		"POINTER":              14,
		"TEXT_CHANGE":          15,
		"KEY_PRESS":            16,
		"FOCUS_CHANGE":         17,
		"VIEW_SCROLL":          18,
		"MOUSE_WHEEL":          19,
		"CLICK":                20,
		"INPUT_CHANGE":         21,
		"FORM_SUBMIT":          22,
		"CONTEXT_MENU":         23,
		"FRAME_RATE":           24,
	}
)

func (x SensorType) Enum() *SensorType {
	p := new(SensorType)
	*p = x
	return p
}

func (x SensorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SensorType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_v2_proto_enumTypes[1].Descriptor()
}

func (SensorType) Type() protoreflect.EnumType {
	return &file_common_v2_proto_enumTypes[1]
}

func (x SensorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SensorType.Descriptor instead.
func (SensorType) EnumDescriptor() ([]byte, []int) {
	return file_common_v2_proto_rawDescGZIP(), []int{1}
}

// An empty message used as response for any call
// that does not return a body.
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_v2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_common_v2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_common_v2_proto_rawDescGZIP(), []int{0}
}

// Returned for a response that has failed in some manner.
// This is used for almost all errors including
// messages that return a 400, 401, 403, 404, and 500
type ErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Param   string `protobuf:"bytes,2,opt,name=param,proto3" json:"param,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ErrorResponse) Reset() {
	*x = ErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_v2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResponse) ProtoMessage() {}

func (x *ErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_v2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResponse.ProtoReflect.Descriptor instead.
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return file_common_v2_proto_rawDescGZIP(), []int{1}
}

func (x *ErrorResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ErrorResponse) GetParam() string {
	if x != nil {
		return x.Param
	}
	return ""
}

func (x *ErrorResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// The response to a GET "token/self" request
type TokenSelfResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId        string `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	ProjectId    string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Scopes       string `protobuf:"bytes,3,opt,name=scopes,proto3" json:"scopes,omitempty"`
	UserId       string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CredentialId string `protobuf:"bytes,5,opt,name=credential_id,json=credentialId,proto3" json:"credential_id,omitempty"`
	Type         int64  `protobuf:"varint,6,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *TokenSelfResponse) Reset() {
	*x = TokenSelfResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_v2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenSelfResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenSelfResponse) ProtoMessage() {}

func (x *TokenSelfResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_v2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenSelfResponse.ProtoReflect.Descriptor instead.
func (*TokenSelfResponse) Descriptor() ([]byte, []int) {
	return file_common_v2_proto_rawDescGZIP(), []int{2}
}

func (x *TokenSelfResponse) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *TokenSelfResponse) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *TokenSelfResponse) GetScopes() string {
	if x != nil {
		return x.Scopes
	}
	return ""
}

func (x *TokenSelfResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *TokenSelfResponse) GetCredentialId() string {
	if x != nil {
		return x.CredentialId
	}
	return ""
}

func (x *TokenSelfResponse) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

var File_common_v2_proto protoreflect.FileDescriptor

var file_common_v2_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x53, 0x0a, 0x0d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xb3, 0x01, 0x0a, 0x11, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x2a, 0x3c, 0x0a, 0x0e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x07, 0x0a, 0x03, 0x69, 0x4f, 0x53, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x4e, 0x44, 0x52,
	0x4f, 0x49, 0x44, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x57, 0x45, 0x42, 0x10, 0x03, 0x2a, 0xae,
	0x03, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x45, 0x4e, 0x53, 0x4f, 0x52, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12,
	0x11, 0x0a, 0x0d, 0x41, 0x43, 0x43, 0x45, 0x4c, 0x45, 0x52, 0x4f, 0x4d, 0x45, 0x54, 0x45, 0x52,
	0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x4c, 0x49, 0x4e, 0x45, 0x41, 0x52, 0x5f, 0x41, 0x43, 0x43,
	0x45, 0x4c, 0x45, 0x52, 0x4f, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c,
	0x4d, 0x41, 0x47, 0x4e, 0x45, 0x54, 0x4f, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x10, 0x04, 0x12, 0x0d,
	0x0a, 0x09, 0x47, 0x59, 0x52, 0x4f, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x10, 0x05, 0x12, 0x0b, 0x0a,
	0x07, 0x42, 0x41, 0x54, 0x54, 0x45, 0x52, 0x59, 0x10, 0x06, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x52,
	0x49, 0x45, 0x4e, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x07, 0x12, 0x0b, 0x0a, 0x07, 0x50,
	0x4f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x10, 0x0e, 0x12, 0x13, 0x0a, 0x0b, 0x54, 0x45, 0x58, 0x54,
	0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x0f, 0x1a, 0x02, 0x08, 0x01, 0x12, 0x0d, 0x0a,
	0x09, 0x4b, 0x45, 0x59, 0x5f, 0x50, 0x52, 0x45, 0x53, 0x53, 0x10, 0x10, 0x12, 0x10, 0x0a, 0x0c,
	0x46, 0x4f, 0x43, 0x55, 0x53, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x11, 0x12, 0x0f,
	0x0a, 0x0b, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x53, 0x43, 0x52, 0x4f, 0x4c, 0x4c, 0x10, 0x12, 0x12,
	0x0f, 0x0a, 0x0b, 0x4d, 0x4f, 0x55, 0x53, 0x45, 0x5f, 0x57, 0x48, 0x45, 0x45, 0x4c, 0x10, 0x13,
	0x12, 0x09, 0x0a, 0x05, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x10, 0x14, 0x12, 0x10, 0x0a, 0x0c, 0x49,
	0x4e, 0x50, 0x55, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x15, 0x12, 0x0f, 0x0a,
	0x0b, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x53, 0x55, 0x42, 0x4d, 0x49, 0x54, 0x10, 0x16, 0x12, 0x10,
	0x0a, 0x0c, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x58, 0x54, 0x5f, 0x4d, 0x45, 0x4e, 0x55, 0x10, 0x17,
	0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x52, 0x41, 0x4d, 0x45, 0x5f, 0x52, 0x41, 0x54, 0x45, 0x10, 0x18,
	0x22, 0x04, 0x08, 0x08, 0x10, 0x08, 0x22, 0x04, 0x08, 0x09, 0x10, 0x09, 0x22, 0x04, 0x08, 0x0a,
	0x10, 0x0a, 0x22, 0x04, 0x08, 0x0b, 0x10, 0x0b, 0x22, 0x04, 0x08, 0x0c, 0x10, 0x0c, 0x22, 0x04,
	0x08, 0x0d, 0x10, 0x0d, 0x2a, 0x0b, 0x54, 0x45, 0x4d, 0x50, 0x45, 0x52, 0x41, 0x54, 0x55, 0x52,
	0x45, 0x2a, 0x05, 0x4c, 0x49, 0x47, 0x48, 0x54, 0x2a, 0x08, 0x50, 0x52, 0x45, 0x53, 0x53, 0x55,
	0x52, 0x45, 0x2a, 0x08, 0x48, 0x55, 0x4d, 0x49, 0x44, 0x49, 0x54, 0x59, 0x2a, 0x05, 0x53, 0x54,
	0x45, 0x50, 0x53, 0x2a, 0x0a, 0x48, 0x45, 0x41, 0x52, 0x54, 0x5f, 0x52, 0x41, 0x54, 0x45, 0x42,
	0x48, 0x0a, 0x16, 0x69, 0x6f, 0x2e, 0x6d, 0x6f, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x65, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x32, 0x42, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5a, 0x20, 0x6d, 0x6f, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x73, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x32, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_common_v2_proto_rawDescOnce sync.Once
	file_common_v2_proto_rawDescData = file_common_v2_proto_rawDesc
)

func file_common_v2_proto_rawDescGZIP() []byte {
	file_common_v2_proto_rawDescOnce.Do(func() {
		file_common_v2_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_v2_proto_rawDescData)
	})
	return file_common_v2_proto_rawDescData
}

var file_common_v2_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_common_v2_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_common_v2_proto_goTypes = []interface{}{
	(DevicePlatform)(0),       // 0: v2.common.DevicePlatform
	(SensorType)(0),           // 1: v2.common.SensorType
	(*Empty)(nil),             // 2: v2.common.Empty
	(*ErrorResponse)(nil),     // 3: v2.common.ErrorResponse
	(*TokenSelfResponse)(nil), // 4: v2.common.TokenSelfResponse
}
var file_common_v2_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_v2_proto_init() }
func file_common_v2_proto_init() {
	if File_common_v2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_v2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_common_v2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorResponse); i {
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
		file_common_v2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenSelfResponse); i {
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
			RawDescriptor: file_common_v2_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_v2_proto_goTypes,
		DependencyIndexes: file_common_v2_proto_depIdxs,
		EnumInfos:         file_common_v2_proto_enumTypes,
		MessageInfos:      file_common_v2_proto_msgTypes,
	}.Build()
	File_common_v2_proto = out.File
	file_common_v2_proto_rawDesc = nil
	file_common_v2_proto_goTypes = nil
	file_common_v2_proto_depIdxs = nil
}
