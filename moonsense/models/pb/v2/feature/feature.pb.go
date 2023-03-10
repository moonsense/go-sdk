// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: feature.proto

package feature

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

type BytesList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value [][]byte `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *BytesList) Reset() {
	*x = BytesList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feature_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BytesList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesList) ProtoMessage() {}

func (x *BytesList) ProtoReflect() protoreflect.Message {
	mi := &file_feature_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BytesList.ProtoReflect.Descriptor instead.
func (*BytesList) Descriptor() ([]byte, []int) {
	return file_feature_proto_rawDescGZIP(), []int{0}
}

func (x *BytesList) GetValue() [][]byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type DoubleList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []float64 `protobuf:"fixed64,1,rep,packed,name=value,proto3" json:"value,omitempty"`
}

func (x *DoubleList) Reset() {
	*x = DoubleList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feature_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoubleList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoubleList) ProtoMessage() {}

func (x *DoubleList) ProtoReflect() protoreflect.Message {
	mi := &file_feature_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoubleList.ProtoReflect.Descriptor instead.
func (*DoubleList) Descriptor() ([]byte, []int) {
	return file_feature_proto_rawDescGZIP(), []int{1}
}

func (x *DoubleList) GetValue() []float64 {
	if x != nil {
		return x.Value
	}
	return nil
}

type DoubleMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value map[string]float64 `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (x *DoubleMap) Reset() {
	*x = DoubleMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feature_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoubleMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoubleMap) ProtoMessage() {}

func (x *DoubleMap) ProtoReflect() protoreflect.Message {
	mi := &file_feature_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoubleMap.ProtoReflect.Descriptor instead.
func (*DoubleMap) Descriptor() ([]byte, []int) {
	return file_feature_proto_rawDescGZIP(), []int{2}
}

func (x *DoubleMap) GetValue() map[string]float64 {
	if x != nil {
		return x.Value
	}
	return nil
}

type Int64List struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []int64 `protobuf:"varint,1,rep,packed,name=value,proto3" json:"value,omitempty"`
}

func (x *Int64List) Reset() {
	*x = Int64List{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feature_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int64List) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int64List) ProtoMessage() {}

func (x *Int64List) ProtoReflect() protoreflect.Message {
	mi := &file_feature_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int64List.ProtoReflect.Descriptor instead.
func (*Int64List) Descriptor() ([]byte, []int) {
	return file_feature_proto_rawDescGZIP(), []int{3}
}

func (x *Int64List) GetValue() []int64 {
	if x != nil {
		return x.Value
	}
	return nil
}

type StringList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []string `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *StringList) Reset() {
	*x = StringList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feature_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringList) ProtoMessage() {}

func (x *StringList) ProtoReflect() protoreflect.Message {
	mi := &file_feature_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringList.ProtoReflect.Descriptor instead.
func (*StringList) Descriptor() ([]byte, []int) {
	return file_feature_proto_rawDescGZIP(), []int{4}
}

func (x *StringList) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}

type Feature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Kind:
	//	*Feature_BytesList
	//	*Feature_DoubleList
	//	*Feature_Int64List
	//	*Feature_StringList
	//	*Feature_DoubleMap
	Kind isFeature_Kind `protobuf_oneof:"kind"`
}

func (x *Feature) Reset() {
	*x = Feature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feature_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Feature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feature) ProtoMessage() {}

func (x *Feature) ProtoReflect() protoreflect.Message {
	mi := &file_feature_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feature.ProtoReflect.Descriptor instead.
func (*Feature) Descriptor() ([]byte, []int) {
	return file_feature_proto_rawDescGZIP(), []int{5}
}

func (m *Feature) GetKind() isFeature_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *Feature) GetBytesList() *BytesList {
	if x, ok := x.GetKind().(*Feature_BytesList); ok {
		return x.BytesList
	}
	return nil
}

func (x *Feature) GetDoubleList() *DoubleList {
	if x, ok := x.GetKind().(*Feature_DoubleList); ok {
		return x.DoubleList
	}
	return nil
}

func (x *Feature) GetInt64List() *Int64List {
	if x, ok := x.GetKind().(*Feature_Int64List); ok {
		return x.Int64List
	}
	return nil
}

func (x *Feature) GetStringList() *StringList {
	if x, ok := x.GetKind().(*Feature_StringList); ok {
		return x.StringList
	}
	return nil
}

func (x *Feature) GetDoubleMap() *DoubleMap {
	if x, ok := x.GetKind().(*Feature_DoubleMap); ok {
		return x.DoubleMap
	}
	return nil
}

type isFeature_Kind interface {
	isFeature_Kind()
}

type Feature_BytesList struct {
	BytesList *BytesList `protobuf:"bytes,1,opt,name=bytes_list,json=bytesList,proto3,oneof"`
}

type Feature_DoubleList struct {
	DoubleList *DoubleList `protobuf:"bytes,2,opt,name=double_list,json=doubleList,proto3,oneof"`
}

type Feature_Int64List struct {
	Int64List *Int64List `protobuf:"bytes,3,opt,name=int64_list,json=int64List,proto3,oneof"`
}

type Feature_StringList struct {
	StringList *StringList `protobuf:"bytes,4,opt,name=string_list,json=stringList,proto3,oneof"`
}

type Feature_DoubleMap struct {
	DoubleMap *DoubleMap `protobuf:"bytes,5,opt,name=double_map,json=doubleMap,proto3,oneof"`
}

func (*Feature_BytesList) isFeature_Kind() {}

func (*Feature_DoubleList) isFeature_Kind() {}

func (*Feature_Int64List) isFeature_Kind() {}

func (*Feature_StringList) isFeature_Kind() {}

func (*Feature_DoubleMap) isFeature_Kind() {}

var File_feature_proto protoreflect.FileDescriptor

var file_feature_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x76, 0x32, 0x2e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x21, 0x0a, 0x09, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x26,
	0x0a, 0x0a, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x01, 0x42, 0x02, 0x10, 0x01, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x7d, 0x0a, 0x09, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x4d, 0x61, 0x70, 0x12, 0x36, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x76, 0x32, 0x2e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e,
	0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x4d, 0x61, 0x70, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x38, 0x0a, 0x0a, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x25, 0x0a, 0x09, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x03, 0x42, 0x02, 0x10, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x22, 0x0a, 0x0a,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0xaf, 0x02, 0x0a, 0x07, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x36, 0x0a, 0x0a,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x76, 0x32, 0x2e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x09, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x76, 0x32, 0x2e, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x48, 0x00, 0x52, 0x0a, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x36, 0x0a, 0x0a, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x76, 0x32, 0x2e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x09, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x76,
	0x32, 0x2e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x36, 0x0a, 0x0a, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x70,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x76, 0x32, 0x2e, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x4d, 0x61, 0x70, 0x48, 0x00, 0x52,
	0x09, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x4d, 0x61, 0x70, 0x42, 0x06, 0x0a, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x42, 0x57, 0x0a, 0x1e, 0x69, 0x6f, 0x2e, 0x6d, 0x6f, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x73, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x42, 0x0d, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x5a, 0x26, 0x6d, 0x6f, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x69,
	0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x3b, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_feature_proto_rawDescOnce sync.Once
	file_feature_proto_rawDescData = file_feature_proto_rawDesc
)

func file_feature_proto_rawDescGZIP() []byte {
	file_feature_proto_rawDescOnce.Do(func() {
		file_feature_proto_rawDescData = protoimpl.X.CompressGZIP(file_feature_proto_rawDescData)
	})
	return file_feature_proto_rawDescData
}

var file_feature_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_feature_proto_goTypes = []interface{}{
	(*BytesList)(nil),  // 0: v2.feature.BytesList
	(*DoubleList)(nil), // 1: v2.feature.DoubleList
	(*DoubleMap)(nil),  // 2: v2.feature.DoubleMap
	(*Int64List)(nil),  // 3: v2.feature.Int64List
	(*StringList)(nil), // 4: v2.feature.StringList
	(*Feature)(nil),    // 5: v2.feature.Feature
	nil,                // 6: v2.feature.DoubleMap.ValueEntry
}
var file_feature_proto_depIdxs = []int32{
	6, // 0: v2.feature.DoubleMap.value:type_name -> v2.feature.DoubleMap.ValueEntry
	0, // 1: v2.feature.Feature.bytes_list:type_name -> v2.feature.BytesList
	1, // 2: v2.feature.Feature.double_list:type_name -> v2.feature.DoubleList
	3, // 3: v2.feature.Feature.int64_list:type_name -> v2.feature.Int64List
	4, // 4: v2.feature.Feature.string_list:type_name -> v2.feature.StringList
	2, // 5: v2.feature.Feature.double_map:type_name -> v2.feature.DoubleMap
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_feature_proto_init() }
func file_feature_proto_init() {
	if File_feature_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_feature_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BytesList); i {
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
		file_feature_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoubleList); i {
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
		file_feature_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoubleMap); i {
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
		file_feature_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Int64List); i {
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
		file_feature_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringList); i {
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
		file_feature_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Feature); i {
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
	file_feature_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*Feature_BytesList)(nil),
		(*Feature_DoubleList)(nil),
		(*Feature_Int64List)(nil),
		(*Feature_StringList)(nil),
		(*Feature_DoubleMap)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_feature_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_feature_proto_goTypes,
		DependencyIndexes: file_feature_proto_depIdxs,
		MessageInfos:      file_feature_proto_msgTypes,
	}.Build()
	File_feature_proto = out.File
	file_feature_proto_rawDesc = nil
	file_feature_proto_goTypes = nil
	file_feature_proto_depIdxs = nil
}
