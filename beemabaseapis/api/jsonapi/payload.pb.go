// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: protodef/api/jsonapi/payload.proto

package jsonapi

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

// A [resource identifier object](http://jsonapi.org/format/#document-resource-identifier-objects).
type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Unique id.
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protodef_api_jsonapi_payload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_protodef_api_jsonapi_payload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_protodef_api_jsonapi_payload_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Data) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Definition for resource identifier collection objects.
type DataCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Data []*Data `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *DataCollection) Reset() {
	*x = DataCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protodef_api_jsonapi_payload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataCollection) ProtoMessage() {}

func (x *DataCollection) ProtoReflect() protoreflect.Message {
	mi := &file_protodef_api_jsonapi_payload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataCollection.ProtoReflect.Descriptor instead.
func (*DataCollection) Descriptor() ([]byte, []int) {
	return file_protodef_api_jsonapi_payload_proto_rawDescGZIP(), []int{1}
}

func (x *DataCollection) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DataCollection) GetData() []*Data {
	if x != nil {
		return x.Data
	}
	return nil
}

// A container for http links.
type Links struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A http link. It points to the resource itself.
	Self string `protobuf:"bytes,1,opt,name=self,proto3" json:"self,omitempty"`
	// A http link. It points to a related resource.
	Related string `protobuf:"bytes,2,opt,name=related,proto3" json:"related,omitempty"`
}

func (x *Links) Reset() {
	*x = Links{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protodef_api_jsonapi_payload_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Links) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Links) ProtoMessage() {}

func (x *Links) ProtoReflect() protoreflect.Message {
	mi := &file_protodef_api_jsonapi_payload_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Links.ProtoReflect.Descriptor instead.
func (*Links) Descriptor() ([]byte, []int) {
	return file_protodef_api_jsonapi_payload_proto_rawDescGZIP(), []int{2}
}

func (x *Links) GetSelf() string {
	if x != nil {
		return x.Self
	}
	return ""
}

func (x *Links) GetRelated() string {
	if x != nil {
		return x.Related
	}
	return ""
}

// A container for pagination links.
type PaginationLinks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A http link to the resource itself.
	Self string `protobuf:"bytes,1,opt,name=self,proto3" json:"self,omitempty"`
	// A http link to the next page of data.
	Next string `protobuf:"bytes,2,opt,name=next,proto3" json:"next,omitempty"`
	// A http link to the previous page of data.
	Prev string `protobuf:"bytes,3,opt,name=prev,proto3" json:"prev,omitempty"`
	// A http link to the last page of data.
	Last string `protobuf:"bytes,4,opt,name=last,proto3" json:"last,omitempty"`
	// A http link to the first page of data.
	First string `protobuf:"bytes,5,opt,name=first,proto3" json:"first,omitempty"`
}

func (x *PaginationLinks) Reset() {
	*x = PaginationLinks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protodef_api_jsonapi_payload_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationLinks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationLinks) ProtoMessage() {}

func (x *PaginationLinks) ProtoReflect() protoreflect.Message {
	mi := &file_protodef_api_jsonapi_payload_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationLinks.ProtoReflect.Descriptor instead.
func (*PaginationLinks) Descriptor() ([]byte, []int) {
	return file_protodef_api_jsonapi_payload_proto_rawDescGZIP(), []int{3}
}

func (x *PaginationLinks) GetSelf() string {
	if x != nil {
		return x.Self
	}
	return ""
}

func (x *PaginationLinks) GetNext() string {
	if x != nil {
		return x.Next
	}
	return ""
}

func (x *PaginationLinks) GetPrev() string {
	if x != nil {
		return x.Prev
	}
	return ""
}

func (x *PaginationLinks) GetLast() string {
	if x != nil {
		return x.Last
	}
	return ""
}

func (x *PaginationLinks) GetFirst() string {
	if x != nil {
		return x.First
	}
	return ""
}

// A container for various pagination properties
type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Total number of entries, regardless of pages.
	Records int64 `protobuf:"varint,1,opt,name=records,proto3" json:"records,omitempty"`
	// Total number of pages.
	Total int64 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	// Number of entries per page.
	Size int64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	// Current page number.
	Number int64 `protobuf:"varint,4,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protodef_api_jsonapi_payload_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_protodef_api_jsonapi_payload_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_protodef_api_jsonapi_payload_proto_rawDescGZIP(), []int{4}
}

func (x *Pagination) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *Pagination) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Pagination) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Pagination) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

// Top level meta container.
type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protodef_api_jsonapi_payload_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_protodef_api_jsonapi_payload_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_protodef_api_jsonapi_payload_proto_rawDescGZIP(), []int{5}
}

func (x *Meta) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_protodef_api_jsonapi_payload_proto protoreflect.FileDescriptor

var file_protodef_api_jsonapi_payload_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x64, 0x65, 0x66, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x64, 0x65, 0x66, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x22, 0x2a, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x50, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x64, 0x65,
	0x66, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x35, 0x0a, 0x05, 0x4c, 0x69, 0x6e, 0x6b,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x6c, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x65, 0x6c, 0x66, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x22,
	0x77, 0x0a, 0x0f, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e,
	0x6b, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x6c, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x73, 0x65, 0x6c, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x72,
	0x65, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x72, 0x65, 0x76, 0x12, 0x12,
	0x0a, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x22, 0x68, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x48, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x40, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x64, 0x65, 0x66, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6a, 0x73,
	0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x7a, 0x0a, 0x15,
	0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x65, 0x65, 0x6d, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6a, 0x73,
	0x6f, 0x6e, 0x61, 0x70, 0x69, 0x42, 0x0c, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x61, 0x6b, 0x65, 0x73, 0x68, 0x32, 0x30, 0x34, 0x35, 0x2f, 0x67, 0x6f, 0x2d,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x65, 0x65, 0x6d, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x61,
	0x70, 0x69, 0x3b, 0x6a, 0x73, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x08,
	0x44, 0x49, 0x43, 0x54, 0x59, 0x41, 0x50, 0x49, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protodef_api_jsonapi_payload_proto_rawDescOnce sync.Once
	file_protodef_api_jsonapi_payload_proto_rawDescData = file_protodef_api_jsonapi_payload_proto_rawDesc
)

func file_protodef_api_jsonapi_payload_proto_rawDescGZIP() []byte {
	file_protodef_api_jsonapi_payload_proto_rawDescOnce.Do(func() {
		file_protodef_api_jsonapi_payload_proto_rawDescData = protoimpl.X.CompressGZIP(file_protodef_api_jsonapi_payload_proto_rawDescData)
	})
	return file_protodef_api_jsonapi_payload_proto_rawDescData
}

var file_protodef_api_jsonapi_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protodef_api_jsonapi_payload_proto_goTypes = []interface{}{
	(*Data)(nil),            // 0: protodef.api.jsonapi.Data
	(*DataCollection)(nil),  // 1: protodef.api.jsonapi.DataCollection
	(*Links)(nil),           // 2: protodef.api.jsonapi.Links
	(*PaginationLinks)(nil), // 3: protodef.api.jsonapi.PaginationLinks
	(*Pagination)(nil),      // 4: protodef.api.jsonapi.Pagination
	(*Meta)(nil),            // 5: protodef.api.jsonapi.Meta
}
var file_protodef_api_jsonapi_payload_proto_depIdxs = []int32{
	0, // 0: protodef.api.jsonapi.DataCollection.data:type_name -> protodef.api.jsonapi.Data
	4, // 1: protodef.api.jsonapi.Meta.pagination:type_name -> protodef.api.jsonapi.Pagination
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protodef_api_jsonapi_payload_proto_init() }
func file_protodef_api_jsonapi_payload_proto_init() {
	if File_protodef_api_jsonapi_payload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protodef_api_jsonapi_payload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_protodef_api_jsonapi_payload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataCollection); i {
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
		file_protodef_api_jsonapi_payload_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Links); i {
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
		file_protodef_api_jsonapi_payload_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationLinks); i {
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
		file_protodef_api_jsonapi_payload_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
		file_protodef_api_jsonapi_payload_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
			RawDescriptor: file_protodef_api_jsonapi_payload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protodef_api_jsonapi_payload_proto_goTypes,
		DependencyIndexes: file_protodef_api_jsonapi_payload_proto_depIdxs,
		MessageInfos:      file_protodef_api_jsonapi_payload_proto_msgTypes,
	}.Build()
	File_protodef_api_jsonapi_payload_proto = out.File
	file_protodef_api_jsonapi_payload_proto_rawDesc = nil
	file_protodef_api_jsonapi_payload_proto_goTypes = nil
	file_protodef_api_jsonapi_payload_proto_depIdxs = nil
}
