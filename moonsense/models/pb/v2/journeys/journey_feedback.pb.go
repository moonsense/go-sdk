// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: journey_feedback.proto

package v2

import (
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

// A method for tracking Fraud against a journey
type FraudFeedback struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique id for the feedback populated by the Moonsense servers.
	FeedbackId string `protobuf:"bytes,1,opt,name=feedback_id,json=feedbackId,proto3" json:"feedback_id,omitempty"`
	// A flag that indicates whether the journey contains fraud.
	IsFraud bool `protobuf:"varint,2,opt,name=is_fraud,json=isFraud,proto3" json:"is_fraud,omitempty"`
	// A timestamp that indicates when the fraud was reported.
	ReportedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=reported_at,json=reportedAt,proto3" json:"reported_at,omitempty"`
	// A reason the journey was indicated to be fraudulent.
	FraudReason string `protobuf:"bytes,4,opt,name=fraud_reason,json=fraudReason,proto3" json:"fraud_reason,omitempty"`
}

func (x *FraudFeedback) Reset() {
	*x = FraudFeedback{}
	if protoimpl.UnsafeEnabled {
		mi := &file_journey_feedback_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FraudFeedback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FraudFeedback) ProtoMessage() {}

func (x *FraudFeedback) ProtoReflect() protoreflect.Message {
	mi := &file_journey_feedback_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FraudFeedback.ProtoReflect.Descriptor instead.
func (*FraudFeedback) Descriptor() ([]byte, []int) {
	return file_journey_feedback_proto_rawDescGZIP(), []int{0}
}

func (x *FraudFeedback) GetFeedbackId() string {
	if x != nil {
		return x.FeedbackId
	}
	return ""
}

func (x *FraudFeedback) GetIsFraud() bool {
	if x != nil {
		return x.IsFraud
	}
	return false
}

func (x *FraudFeedback) GetReportedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ReportedAt
	}
	return nil
}

func (x *FraudFeedback) GetFraudReason() string {
	if x != nil {
		return x.FraudReason
	}
	return ""
}

// A method for tracking Chargebacks against a journey
type ChargebackFeedback struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique id for the feedback populated by the Moonsense servers.
	FeedbackId string `protobuf:"bytes,1,opt,name=feedback_id,json=feedbackId,proto3" json:"feedback_id,omitempty"`
	// A flag that indicates whether the journey contains a chargeback.
	IsChargeback bool `protobuf:"varint,2,opt,name=is_chargeback,json=isChargeback,proto3" json:"is_chargeback,omitempty"`
	// A timestamp that indicates when the chargeback was reported.
	ReportedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=reported_at,json=reportedAt,proto3" json:"reported_at,omitempty"`
	// A reason the journey was indicated to be fraudulent.
	ChargebackReason string `protobuf:"bytes,4,opt,name=chargeback_reason,json=chargebackReason,proto3" json:"chargeback_reason,omitempty"`
}

func (x *ChargebackFeedback) Reset() {
	*x = ChargebackFeedback{}
	if protoimpl.UnsafeEnabled {
		mi := &file_journey_feedback_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChargebackFeedback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChargebackFeedback) ProtoMessage() {}

func (x *ChargebackFeedback) ProtoReflect() protoreflect.Message {
	mi := &file_journey_feedback_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChargebackFeedback.ProtoReflect.Descriptor instead.
func (*ChargebackFeedback) Descriptor() ([]byte, []int) {
	return file_journey_feedback_proto_rawDescGZIP(), []int{1}
}

func (x *ChargebackFeedback) GetFeedbackId() string {
	if x != nil {
		return x.FeedbackId
	}
	return ""
}

func (x *ChargebackFeedback) GetIsChargeback() bool {
	if x != nil {
		return x.IsChargeback
	}
	return false
}

func (x *ChargebackFeedback) GetReportedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ReportedAt
	}
	return nil
}

func (x *ChargebackFeedback) GetChargebackReason() string {
	if x != nil {
		return x.ChargebackReason
	}
	return ""
}

// A method for tracking events against a journey
type JourneyFeedback struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The journey id.
	JourneyId string `protobuf:"bytes,1,opt,name=journey_id,json=journeyId,proto3" json:"journey_id,omitempty"`
	// The fraud feedback. This is set if the
	// the journey has been reported as fraudulent.
	FraudFeedback *FraudFeedback `protobuf:"bytes,3,opt,name=fraud_feedback,json=fraudFeedback,proto3" json:"fraud_feedback,omitempty"`
	// The chargeback feedback. This is set if the
	// the journey has been reported as a chargeback.
	ChargebackFeedback *ChargebackFeedback `protobuf:"bytes,4,opt,name=chargeback_feedback,json=chargebackFeedback,proto3" json:"chargeback_feedback,omitempty"`
}

func (x *JourneyFeedback) Reset() {
	*x = JourneyFeedback{}
	if protoimpl.UnsafeEnabled {
		mi := &file_journey_feedback_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JourneyFeedback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JourneyFeedback) ProtoMessage() {}

func (x *JourneyFeedback) ProtoReflect() protoreflect.Message {
	mi := &file_journey_feedback_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JourneyFeedback.ProtoReflect.Descriptor instead.
func (*JourneyFeedback) Descriptor() ([]byte, []int) {
	return file_journey_feedback_proto_rawDescGZIP(), []int{2}
}

func (x *JourneyFeedback) GetJourneyId() string {
	if x != nil {
		return x.JourneyId
	}
	return ""
}

func (x *JourneyFeedback) GetFraudFeedback() *FraudFeedback {
	if x != nil {
		return x.FraudFeedback
	}
	return nil
}

func (x *JourneyFeedback) GetChargebackFeedback() *ChargebackFeedback {
	if x != nil {
		return x.ChargebackFeedback
	}
	return nil
}

var File_journey_feedback_proto protoreflect.FileDescriptor

var file_journey_feedback_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61,
	0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x65,
	0x79, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x01, 0x0a, 0x0d,
	0x46, 0x72, 0x61, 0x75, 0x64, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x1f, 0x0a,
	0x0b, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x69, 0x73, 0x5f, 0x66, 0x72, 0x61, 0x75, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x69, 0x73, 0x46, 0x72, 0x61, 0x75, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72, 0x61, 0x75, 0x64, 0x5f,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x72,
	0x61, 0x75, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0xc4, 0x01, 0x0a, 0x12, 0x43, 0x68,
	0x61, 0x72, 0x67, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b,
	0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x49,
	0x64, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x62, 0x61,
	0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x43, 0x68, 0x61, 0x72,
	0x67, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x3b, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x62, 0x61, 0x63,
	0x6b, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x22, 0xcf, 0x01, 0x0a, 0x0f, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x46, 0x65, 0x65, 0x64,
	0x62, 0x61, 0x63, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x65,
	0x79, 0x49, 0x64, 0x12, 0x46, 0x0a, 0x0e, 0x66, 0x72, 0x61, 0x75, 0x64, 0x5f, 0x66, 0x65, 0x65,
	0x64, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6a, 0x6f,
	0x75, 0x72, 0x6e, 0x65, 0x79, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x46,
	0x72, 0x61, 0x75, 0x64, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x0d, 0x66, 0x72,
	0x61, 0x75, 0x64, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x55, 0x0a, 0x13, 0x63,
	0x68, 0x61, 0x72, 0x67, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61,
	0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6a, 0x6f, 0x75, 0x72, 0x6e,
	0x65, 0x79, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x43, 0x68, 0x61, 0x72,
	0x67, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x12,
	0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61,
	0x63, 0x6b, 0x42, 0x64, 0x0a, 0x27, 0x69, 0x6f, 0x2e, 0x6d, 0x6f, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x73, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x6a, 0x6f, 0x75,
	0x72, 0x6e, 0x65, 0x79, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x42, 0x15, 0x4a,
	0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x5a, 0x22, 0x6d, 0x6f, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x65, 0x2e,
	0x69, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x32, 0x2f, 0x6a, 0x6f, 0x75,
	0x72, 0x6e, 0x65, 0x79, 0x73, 0x3b, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_journey_feedback_proto_rawDescOnce sync.Once
	file_journey_feedback_proto_rawDescData = file_journey_feedback_proto_rawDesc
)

func file_journey_feedback_proto_rawDescGZIP() []byte {
	file_journey_feedback_proto_rawDescOnce.Do(func() {
		file_journey_feedback_proto_rawDescData = protoimpl.X.CompressGZIP(file_journey_feedback_proto_rawDescData)
	})
	return file_journey_feedback_proto_rawDescData
}

var file_journey_feedback_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_journey_feedback_proto_goTypes = []interface{}{
	(*FraudFeedback)(nil),         // 0: journey_feedback.FraudFeedback
	(*ChargebackFeedback)(nil),    // 1: journey_feedback.ChargebackFeedback
	(*JourneyFeedback)(nil),       // 2: journey_feedback.JourneyFeedback
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_journey_feedback_proto_depIdxs = []int32{
	3, // 0: journey_feedback.FraudFeedback.reported_at:type_name -> google.protobuf.Timestamp
	3, // 1: journey_feedback.ChargebackFeedback.reported_at:type_name -> google.protobuf.Timestamp
	0, // 2: journey_feedback.JourneyFeedback.fraud_feedback:type_name -> journey_feedback.FraudFeedback
	1, // 3: journey_feedback.JourneyFeedback.chargeback_feedback:type_name -> journey_feedback.ChargebackFeedback
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_journey_feedback_proto_init() }
func file_journey_feedback_proto_init() {
	if File_journey_feedback_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_journey_feedback_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FraudFeedback); i {
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
		file_journey_feedback_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChargebackFeedback); i {
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
		file_journey_feedback_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JourneyFeedback); i {
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
			RawDescriptor: file_journey_feedback_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_journey_feedback_proto_goTypes,
		DependencyIndexes: file_journey_feedback_proto_depIdxs,
		MessageInfos:      file_journey_feedback_proto_msgTypes,
	}.Build()
	File_journey_feedback_proto = out.File
	file_journey_feedback_proto_rawDesc = nil
	file_journey_feedback_proto_goTypes = nil
	file_journey_feedback_proto_depIdxs = nil
}