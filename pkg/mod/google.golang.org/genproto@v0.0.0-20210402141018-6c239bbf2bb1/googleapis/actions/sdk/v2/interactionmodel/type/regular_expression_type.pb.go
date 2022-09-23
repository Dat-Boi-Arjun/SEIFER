// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: google/actions/sdk/v2/interactionmodel/type/regular_expression_type.proto

package _type

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Type that matches text by regular expressions.
// **This message is localizable.**
type RegularExpressionType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Named map of entities which each contain Regex strings.
	Entities map[string]*RegularExpressionType_Entity `protobuf:"bytes,1,rep,name=entities,proto3" json:"entities,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RegularExpressionType) Reset() {
	*x = RegularExpressionType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegularExpressionType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegularExpressionType) ProtoMessage() {}

func (x *RegularExpressionType) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegularExpressionType.ProtoReflect.Descriptor instead.
func (*RegularExpressionType) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_rawDescGZIP(), []int{0}
}

func (x *RegularExpressionType) GetEntities() map[string]*RegularExpressionType_Entity {
	if x != nil {
		return x.Entities
	}
	return nil
}

// Represents an entity object that contains the regular expression that is
// used for comparison.
type RegularExpressionType_Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. Elements that will be displayed on the canvas once an entity is
	// extracted from a query. Only relevant for canvas enabled apps.
	Display *EntityDisplay `protobuf:"bytes,1,opt,name=display,proto3" json:"display,omitempty"`
	// Required. Uses RE2 regex syntax (See
	// https://github.com/google/re2/wiki/Syntax for more details)
	RegularExpressions []string `protobuf:"bytes,2,rep,name=regular_expressions,json=regularExpressions,proto3" json:"regular_expressions,omitempty"`
}

func (x *RegularExpressionType_Entity) Reset() {
	*x = RegularExpressionType_Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegularExpressionType_Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegularExpressionType_Entity) ProtoMessage() {}

func (x *RegularExpressionType_Entity) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegularExpressionType_Entity.ProtoReflect.Descriptor instead.
func (*RegularExpressionType_Entity) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_rawDescGZIP(), []int{0, 0}
}

func (x *RegularExpressionType_Entity) GetDisplay() *EntityDisplay {
	if x != nil {
		return x.Display
	}
	return nil
}

func (x *RegularExpressionType_Entity) GetRegularExpressions() []string {
	if x != nil {
		return x.RegularExpressions
	}
	return nil
}

var File_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto protoreflect.FileDescriptor

var file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_rawDesc = []byte{
	0x0a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x72, 0x65,
	0x67, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e,
	0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x03, 0x0a, 0x15,
	0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x71, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x50, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x45, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x1a, 0x99, 0x01, 0x0a, 0x06, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x59, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x34,
	0x0a, 0x13, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x12, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x86, 0x01, 0x0a, 0x0d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x5f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x49, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x45, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0xa0, 0x01,
	0x0a, 0x2f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x42, 0x1a, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x4f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x3b, 0x74, 0x79, 0x70, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_rawDescOnce sync.Once
	file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_rawDescData = file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_rawDesc
)

func file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_rawDescGZIP() []byte {
	file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_rawDescOnce.Do(func() {
		file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_rawDescData)
	})
	return file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_rawDescData
}

var file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_goTypes = []interface{}{
	(*RegularExpressionType)(nil),        // 0: google.actions.sdk.v2.interactionmodel.type.RegularExpressionType
	(*RegularExpressionType_Entity)(nil), // 1: google.actions.sdk.v2.interactionmodel.type.RegularExpressionType.Entity
	nil,                                  // 2: google.actions.sdk.v2.interactionmodel.type.RegularExpressionType.EntitiesEntry
	(*EntityDisplay)(nil),                // 3: google.actions.sdk.v2.interactionmodel.type.EntityDisplay
}
var file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_depIdxs = []int32{
	2, // 0: google.actions.sdk.v2.interactionmodel.type.RegularExpressionType.entities:type_name -> google.actions.sdk.v2.interactionmodel.type.RegularExpressionType.EntitiesEntry
	3, // 1: google.actions.sdk.v2.interactionmodel.type.RegularExpressionType.Entity.display:type_name -> google.actions.sdk.v2.interactionmodel.type.EntityDisplay
	1, // 2: google.actions.sdk.v2.interactionmodel.type.RegularExpressionType.EntitiesEntry.value:type_name -> google.actions.sdk.v2.interactionmodel.type.RegularExpressionType.Entity
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_init() }
func file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_init() {
	if File_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto != nil {
		return
	}
	file_google_actions_sdk_v2_interactionmodel_type_entity_display_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegularExpressionType); i {
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
		file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegularExpressionType_Entity); i {
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
			RawDescriptor: file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_goTypes,
		DependencyIndexes: file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_depIdxs,
		MessageInfos:      file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_msgTypes,
	}.Build()
	File_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto = out.File
	file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_rawDesc = nil
	file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_goTypes = nil
	file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_depIdxs = nil
}
