// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: wso2/discovery/subscription/application.proto

package subscription

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

// Application data model
type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid         string            `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name         string            `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	SubId        int32             `protobuf:"varint,4,opt,name=subId,proto3" json:"subId,omitempty"`
	SubName      string            `protobuf:"bytes,5,opt,name=subName,proto3" json:"subName,omitempty"`
	Policy       string            `protobuf:"bytes,6,opt,name=policy,proto3" json:"policy,omitempty"`
	TokenType    string            `protobuf:"bytes,7,opt,name=tokenType,proto3" json:"tokenType,omitempty"`
	GroupIds     []string          `protobuf:"bytes,8,rep,name=groupIds,proto3" json:"groupIds,omitempty"`
	Attributes   map[string]string `protobuf:"bytes,9,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	TenantId     int32             `protobuf:"varint,10,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	TenantDomain string            `protobuf:"bytes,11,opt,name=tenantDomain,proto3" json:"tenantDomain,omitempty"`
	Timestamp    int64             `protobuf:"varint,12,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_subscription_application_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_subscription_application_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_subscription_application_proto_rawDescGZIP(), []int{0}
}

func (x *Application) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Application) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Application) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Application) GetSubId() int32 {
	if x != nil {
		return x.SubId
	}
	return 0
}

func (x *Application) GetSubName() string {
	if x != nil {
		return x.SubName
	}
	return ""
}

func (x *Application) GetPolicy() string {
	if x != nil {
		return x.Policy
	}
	return ""
}

func (x *Application) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *Application) GetGroupIds() []string {
	if x != nil {
		return x.GroupIds
	}
	return nil
}

func (x *Application) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Application) GetTenantId() int32 {
	if x != nil {
		return x.TenantId
	}
	return 0
}

func (x *Application) GetTenantDomain() string {
	if x != nil {
		return x.TenantDomain
	}
	return ""
}

func (x *Application) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_wso2_discovery_subscription_application_proto protoreflect.FileDescriptor

var file_wso2_discovery_subscription_application_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1b, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xbe, 0x03, 0x0a,
	0x0b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x75, 0x62, 0x49, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x75, 0x62, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x62, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x73, 0x12, 0x58, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x77, 0x73, 0x6f,
	0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x1a, 0x3d,
	0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x95, 0x01,
	0x0a, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x65,
	0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x10, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x77, 0x73, 0x6f,
	0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wso2_discovery_subscription_application_proto_rawDescOnce sync.Once
	file_wso2_discovery_subscription_application_proto_rawDescData = file_wso2_discovery_subscription_application_proto_rawDesc
)

func file_wso2_discovery_subscription_application_proto_rawDescGZIP() []byte {
	file_wso2_discovery_subscription_application_proto_rawDescOnce.Do(func() {
		file_wso2_discovery_subscription_application_proto_rawDescData = protoimpl.X.CompressGZIP(file_wso2_discovery_subscription_application_proto_rawDescData)
	})
	return file_wso2_discovery_subscription_application_proto_rawDescData
}

var file_wso2_discovery_subscription_application_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_wso2_discovery_subscription_application_proto_goTypes = []interface{}{
	(*Application)(nil), // 0: wso2.discovery.subscription.Application
	nil,                 // 1: wso2.discovery.subscription.Application.AttributesEntry
}
var file_wso2_discovery_subscription_application_proto_depIdxs = []int32{
	1, // 0: wso2.discovery.subscription.Application.attributes:type_name -> wso2.discovery.subscription.Application.AttributesEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_wso2_discovery_subscription_application_proto_init() }
func file_wso2_discovery_subscription_application_proto_init() {
	if File_wso2_discovery_subscription_application_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wso2_discovery_subscription_application_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
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
			RawDescriptor: file_wso2_discovery_subscription_application_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wso2_discovery_subscription_application_proto_goTypes,
		DependencyIndexes: file_wso2_discovery_subscription_application_proto_depIdxs,
		MessageInfos:      file_wso2_discovery_subscription_application_proto_msgTypes,
	}.Build()
	File_wso2_discovery_subscription_application_proto = out.File
	file_wso2_discovery_subscription_application_proto_rawDesc = nil
	file_wso2_discovery_subscription_application_proto_goTypes = nil
	file_wso2_discovery_subscription_application_proto_depIdxs = nil
}