// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: spec/v1.0-draft/statement.proto

package v1_0_draft

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Proto representation of the in-toto v1.0 Statement.
// https://github.com/in-toto/attestation/tree/main/spec/v1.0-draft
// Validation of all fields is left to the users of this proto.
type Statement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Expected to always be "https://in-toto.io/Statement/v1.0"
	Type          string               `protobuf:"bytes,1,opt,name=type,json=_type,proto3" json:"type,omitempty"`
	Subject       []*Statement_Subject `protobuf:"bytes,2,rep,name=subject,proto3" json:"subject,omitempty"`
	PredicateType string               `protobuf:"bytes,3,opt,name=predicateType,proto3" json:"predicateType,omitempty"`
	Predicate     *structpb.Struct     `protobuf:"bytes,4,opt,name=predicate,proto3" json:"predicate,omitempty"`
}

func (x *Statement) Reset() {
	*x = Statement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_v1_0_draft_statement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Statement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Statement) ProtoMessage() {}

func (x *Statement) ProtoReflect() protoreflect.Message {
	mi := &file_spec_v1_0_draft_statement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Statement.ProtoReflect.Descriptor instead.
func (*Statement) Descriptor() ([]byte, []int) {
	return file_spec_v1_0_draft_statement_proto_rawDescGZIP(), []int{0}
}

func (x *Statement) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Statement) GetSubject() []*Statement_Subject {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *Statement) GetPredicateType() string {
	if x != nil {
		return x.PredicateType
	}
	return ""
}

func (x *Statement) GetPredicate() *structpb.Struct {
	if x != nil {
		return x.Predicate
	}
	return nil
}

type Statement_Subject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Digest map[string]string `protobuf:"bytes,2,rep,name=digest,proto3" json:"digest,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Statement_Subject) Reset() {
	*x = Statement_Subject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_v1_0_draft_statement_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Statement_Subject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Statement_Subject) ProtoMessage() {}

func (x *Statement_Subject) ProtoReflect() protoreflect.Message {
	mi := &file_spec_v1_0_draft_statement_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Statement_Subject.ProtoReflect.Descriptor instead.
func (*Statement_Subject) Descriptor() ([]byte, []int) {
	return file_spec_v1_0_draft_statement_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Statement_Subject) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Statement_Subject) GetDigest() map[string]string {
	if x != nil {
		return x.Digest
	}
	return nil
}

var File_spec_v1_0_draft_statement_proto protoreflect.FileDescriptor

var file_spec_v1_0_draft_statement_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x76, 0x31, 0x2e, 0x30, 0x2d, 0x64, 0x72, 0x61, 0x66,
	0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x69, 0x6e, 0x74, 0x6f, 0x74, 0x6f,
	0x2e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x70, 0x65,
	0x63, 0x2e, 0x76, 0x31, 0x5f, 0x30, 0x5f, 0x64, 0x72, 0x61, 0x66, 0x74, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x03, 0x0a, 0x09, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x13, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x56, 0x0a,
	0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x69, 0x6e, 0x74, 0x6f, 0x74, 0x6f, 0x2e, 0x61,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e,
	0x76, 0x31, 0x5f, 0x30, 0x5f, 0x64, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x70,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x09, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x1a, 0xba, 0x01, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x60, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x48, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x69, 0x6e, 0x74, 0x6f,
	0x74, 0x6f, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73,
	0x70, 0x65, 0x63, 0x2e, 0x76, 0x31, 0x5f, 0x30, 0x5f, 0x64, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x64, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e,
	0x2d, 0x74, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x76, 0x31, 0x2e, 0x30, 0x2d, 0x64,
	0x72, 0x61, 0x66, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spec_v1_0_draft_statement_proto_rawDescOnce sync.Once
	file_spec_v1_0_draft_statement_proto_rawDescData = file_spec_v1_0_draft_statement_proto_rawDesc
)

func file_spec_v1_0_draft_statement_proto_rawDescGZIP() []byte {
	file_spec_v1_0_draft_statement_proto_rawDescOnce.Do(func() {
		file_spec_v1_0_draft_statement_proto_rawDescData = protoimpl.X.CompressGZIP(file_spec_v1_0_draft_statement_proto_rawDescData)
	})
	return file_spec_v1_0_draft_statement_proto_rawDescData
}

var file_spec_v1_0_draft_statement_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_spec_v1_0_draft_statement_proto_goTypes = []interface{}{
	(*Statement)(nil),         // 0: github.intoto.attestation.spec.v1_0_draft.Statement
	(*Statement_Subject)(nil), // 1: github.intoto.attestation.spec.v1_0_draft.Statement.Subject
	nil,                       // 2: github.intoto.attestation.spec.v1_0_draft.Statement.Subject.DigestEntry
	(*structpb.Struct)(nil),   // 3: google.protobuf.Struct
}
var file_spec_v1_0_draft_statement_proto_depIdxs = []int32{
	1, // 0: github.intoto.attestation.spec.v1_0_draft.Statement.subject:type_name -> github.intoto.attestation.spec.v1_0_draft.Statement.Subject
	3, // 1: github.intoto.attestation.spec.v1_0_draft.Statement.predicate:type_name -> google.protobuf.Struct
	2, // 2: github.intoto.attestation.spec.v1_0_draft.Statement.Subject.digest:type_name -> github.intoto.attestation.spec.v1_0_draft.Statement.Subject.DigestEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_spec_v1_0_draft_statement_proto_init() }
func file_spec_v1_0_draft_statement_proto_init() {
	if File_spec_v1_0_draft_statement_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spec_v1_0_draft_statement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Statement); i {
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
		file_spec_v1_0_draft_statement_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Statement_Subject); i {
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
			RawDescriptor: file_spec_v1_0_draft_statement_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spec_v1_0_draft_statement_proto_goTypes,
		DependencyIndexes: file_spec_v1_0_draft_statement_proto_depIdxs,
		MessageInfos:      file_spec_v1_0_draft_statement_proto_msgTypes,
	}.Build()
	File_spec_v1_0_draft_statement_proto = out.File
	file_spec_v1_0_draft_statement_proto_rawDesc = nil
	file_spec_v1_0_draft_statement_proto_goTypes = nil
	file_spec_v1_0_draft_statement_proto_depIdxs = nil
}
