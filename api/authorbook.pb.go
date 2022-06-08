// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.20.1
// source: proto/authorbook.proto

package api

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

type AuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Author string `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *AuthorRequest) Reset() {
	*x = AuthorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_authorbook_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorRequest) ProtoMessage() {}

func (x *AuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_authorbook_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorRequest.ProtoReflect.Descriptor instead.
func (*AuthorRequest) Descriptor() ([]byte, []int) {
	return file_proto_authorbook_proto_rawDescGZIP(), []int{0}
}

func (x *AuthorRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type BookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Book string `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *BookRequest) Reset() {
	*x = BookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_authorbook_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookRequest) ProtoMessage() {}

func (x *BookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_authorbook_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookRequest.ProtoReflect.Descriptor instead.
func (*BookRequest) Descriptor() ([]byte, []int) {
	return file_proto_authorbook_proto_rawDescGZIP(), []int{1}
}

func (x *BookRequest) GetBook() string {
	if x != nil {
		return x.Book
	}
	return ""
}

type BookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *BookResponse) Reset() {
	*x = BookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_authorbook_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookResponse) ProtoMessage() {}

func (x *BookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_authorbook_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookResponse.ProtoReflect.Descriptor instead.
func (*BookResponse) Descriptor() ([]byte, []int) {
	return file_proto_authorbook_proto_rawDescGZIP(), []int{2}
}

func (x *BookResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type BooksResponseList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*BookResponse `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *BooksResponseList) Reset() {
	*x = BooksResponseList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_authorbook_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BooksResponseList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BooksResponseList) ProtoMessage() {}

func (x *BooksResponseList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_authorbook_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BooksResponseList.ProtoReflect.Descriptor instead.
func (*BooksResponseList) Descriptor() ([]byte, []int) {
	return file_proto_authorbook_proto_rawDescGZIP(), []int{3}
}

func (x *BooksResponseList) GetBooks() []*BookResponse {
	if x != nil {
		return x.Books
	}
	return nil
}

type AuthorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AuthorResponse) Reset() {
	*x = AuthorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_authorbook_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorResponse) ProtoMessage() {}

func (x *AuthorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_authorbook_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorResponse.ProtoReflect.Descriptor instead.
func (*AuthorResponse) Descriptor() ([]byte, []int) {
	return file_proto_authorbook_proto_rawDescGZIP(), []int{4}
}

func (x *AuthorResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AuthorsResponseList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authors []*AuthorResponse `protobuf:"bytes,1,rep,name=authors,proto3" json:"authors,omitempty"`
}

func (x *AuthorsResponseList) Reset() {
	*x = AuthorsResponseList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_authorbook_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorsResponseList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorsResponseList) ProtoMessage() {}

func (x *AuthorsResponseList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_authorbook_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorsResponseList.ProtoReflect.Descriptor instead.
func (*AuthorsResponseList) Descriptor() ([]byte, []int) {
	return file_proto_authorbook_proto_rawDescGZIP(), []int{5}
}

func (x *AuthorsResponseList) GetAuthors() []*AuthorResponse {
	if x != nil {
		return x.Authors
	}
	return nil
}

var File_proto_authorbook_proto protoreflect.FileDescriptor

var file_proto_authorbook_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x62, 0x6f,
	0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x27, 0x0a,
	0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x21, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x22, 0x0a, 0x0c, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3c, 0x0a,
	0x11, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x27, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x24, 0x0a, 0x0e, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x44, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x32, 0x98, 0x01, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64,
	0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x41, 0x0a, 0x11,
	0x46, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12,
	0x41, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x42, 0x79,
	0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_authorbook_proto_rawDescOnce sync.Once
	file_proto_authorbook_proto_rawDescData = file_proto_authorbook_proto_rawDesc
)

func file_proto_authorbook_proto_rawDescGZIP() []byte {
	file_proto_authorbook_proto_rawDescOnce.Do(func() {
		file_proto_authorbook_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_authorbook_proto_rawDescData)
	})
	return file_proto_authorbook_proto_rawDescData
}

var file_proto_authorbook_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_authorbook_proto_goTypes = []interface{}{
	(*AuthorRequest)(nil),       // 0: api.AuthorRequest
	(*BookRequest)(nil),         // 1: api.BookRequest
	(*BookResponse)(nil),        // 2: api.BookResponse
	(*BooksResponseList)(nil),   // 3: api.BooksResponseList
	(*AuthorResponse)(nil),      // 4: api.AuthorResponse
	(*AuthorsResponseList)(nil), // 5: api.AuthorsResponseList
}
var file_proto_authorbook_proto_depIdxs = []int32{
	2, // 0: api.BooksResponseList.books:type_name -> api.BookResponse
	4, // 1: api.AuthorsResponseList.authors:type_name -> api.AuthorResponse
	0, // 2: api.FindBooksAuthors.FindBooksByAuthor:input_type -> api.AuthorRequest
	1, // 3: api.FindBooksAuthors.FindAuthorsByBook:input_type -> api.BookRequest
	3, // 4: api.FindBooksAuthors.FindBooksByAuthor:output_type -> api.BooksResponseList
	5, // 5: api.FindBooksAuthors.FindAuthorsByBook:output_type -> api.AuthorsResponseList
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_authorbook_proto_init() }
func file_proto_authorbook_proto_init() {
	if File_proto_authorbook_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_authorbook_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorRequest); i {
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
		file_proto_authorbook_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookRequest); i {
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
		file_proto_authorbook_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookResponse); i {
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
		file_proto_authorbook_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BooksResponseList); i {
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
		file_proto_authorbook_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorResponse); i {
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
		file_proto_authorbook_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorsResponseList); i {
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
			RawDescriptor: file_proto_authorbook_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_authorbook_proto_goTypes,
		DependencyIndexes: file_proto_authorbook_proto_depIdxs,
		MessageInfos:      file_proto_authorbook_proto_msgTypes,
	}.Build()
	File_proto_authorbook_proto = out.File
	file_proto_authorbook_proto_rawDesc = nil
	file_proto_authorbook_proto_goTypes = nil
	file_proto_authorbook_proto_depIdxs = nil
}