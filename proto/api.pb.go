// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/api.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/api.proto

It has these top-level messages:
	FilterRequest
	FilterResponse
	BookRequest
	BookResponse
	Book
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type FilterRequest struct {
	// these below are query params
	Page int32 `protobuf:"varint,1,opt,name=page" json:"page,omitempty"`
	Size int32 `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
}

func (m *FilterRequest) Reset()                    { *m = FilterRequest{} }
func (m *FilterRequest) String() string            { return proto1.CompactTextString(m) }
func (*FilterRequest) ProtoMessage()               {}
func (*FilterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FilterRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *FilterRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type FilterResponse struct {
	Book []*Book `protobuf:"bytes,1,rep,name=book" json:"book,omitempty"`
}

func (m *FilterResponse) Reset()                    { *m = FilterResponse{} }
func (m *FilterResponse) String() string            { return proto1.CompactTextString(m) }
func (*FilterResponse) ProtoMessage()               {}
func (*FilterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FilterResponse) GetBook() []*Book {
	if m != nil {
		return m.Book
	}
	return nil
}

type BookRequest struct {
	// Types that are valid to be assigned to BookReq:
	//	*BookRequest_Book
	//	*BookRequest_BookId
	BookReq isBookRequest_BookReq `protobuf_oneof:"book_req"`
}

func (m *BookRequest) Reset()                    { *m = BookRequest{} }
func (m *BookRequest) String() string            { return proto1.CompactTextString(m) }
func (*BookRequest) ProtoMessage()               {}
func (*BookRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isBookRequest_BookReq interface {
	isBookRequest_BookReq()
}

type BookRequest_Book struct {
	Book *Book `protobuf:"bytes,1,opt,name=book,oneof"`
}
type BookRequest_BookId struct {
	BookId int32 `protobuf:"varint,2,opt,name=book_id,json=bookId,oneof"`
}

func (*BookRequest_Book) isBookRequest_BookReq()   {}
func (*BookRequest_BookId) isBookRequest_BookReq() {}

func (m *BookRequest) GetBookReq() isBookRequest_BookReq {
	if m != nil {
		return m.BookReq
	}
	return nil
}

func (m *BookRequest) GetBook() *Book {
	if x, ok := m.GetBookReq().(*BookRequest_Book); ok {
		return x.Book
	}
	return nil
}

func (m *BookRequest) GetBookId() int32 {
	if x, ok := m.GetBookReq().(*BookRequest_BookId); ok {
		return x.BookId
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*BookRequest) XXX_OneofFuncs() (func(msg proto1.Message, b *proto1.Buffer) error, func(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error), func(msg proto1.Message) (n int), []interface{}) {
	return _BookRequest_OneofMarshaler, _BookRequest_OneofUnmarshaler, _BookRequest_OneofSizer, []interface{}{
		(*BookRequest_Book)(nil),
		(*BookRequest_BookId)(nil),
	}
}

func _BookRequest_OneofMarshaler(msg proto1.Message, b *proto1.Buffer) error {
	m := msg.(*BookRequest)
	// book_req
	switch x := m.BookReq.(type) {
	case *BookRequest_Book:
		b.EncodeVarint(1<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.Book); err != nil {
			return err
		}
	case *BookRequest_BookId:
		b.EncodeVarint(2<<3 | proto1.WireVarint)
		b.EncodeVarint(uint64(x.BookId))
	case nil:
	default:
		return fmt.Errorf("BookRequest.BookReq has unexpected type %T", x)
	}
	return nil
}

func _BookRequest_OneofUnmarshaler(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error) {
	m := msg.(*BookRequest)
	switch tag {
	case 1: // book_req.book
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(Book)
		err := b.DecodeMessage(msg)
		m.BookReq = &BookRequest_Book{msg}
		return true, err
	case 2: // book_req.book_id
		if wire != proto1.WireVarint {
			return true, proto1.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.BookReq = &BookRequest_BookId{int32(x)}
		return true, err
	default:
		return false, nil
	}
}

func _BookRequest_OneofSizer(msg proto1.Message) (n int) {
	m := msg.(*BookRequest)
	// book_req
	switch x := m.BookReq.(type) {
	case *BookRequest_Book:
		s := proto1.Size(x.Book)
		n += proto1.SizeVarint(1<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case *BookRequest_BookId:
		n += proto1.SizeVarint(2<<3 | proto1.WireVarint)
		n += proto1.SizeVarint(uint64(x.BookId))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type BookResponse struct {
	Book *Book `protobuf:"bytes,1,opt,name=book" json:"book,omitempty"`
}

func (m *BookResponse) Reset()                    { *m = BookResponse{} }
func (m *BookResponse) String() string            { return proto1.CompactTextString(m) }
func (*BookResponse) ProtoMessage()               {}
func (*BookResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BookResponse) GetBook() *Book {
	if m != nil {
		return m.Book
	}
	return nil
}

// domain
type Book struct {
	BookId int32  `protobuf:"varint,1,opt,name=book_id,json=bookId" json:"book_id,omitempty"`
	Isbn   string `protobuf:"bytes,2,opt,name=isbn" json:"isbn,omitempty"`
	Title  string `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Author string `protobuf:"bytes,4,opt,name=author" json:"author,omitempty"`
}

func (m *Book) Reset()                    { *m = Book{} }
func (m *Book) String() string            { return proto1.CompactTextString(m) }
func (*Book) ProtoMessage()               {}
func (*Book) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Book) GetBookId() int32 {
	if m != nil {
		return m.BookId
	}
	return 0
}

func (m *Book) GetIsbn() string {
	if m != nil {
		return m.Isbn
	}
	return ""
}

func (m *Book) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Book) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func init() {
	proto1.RegisterType((*FilterRequest)(nil), "proto.FilterRequest")
	proto1.RegisterType((*FilterResponse)(nil), "proto.FilterResponse")
	proto1.RegisterType((*BookRequest)(nil), "proto.BookRequest")
	proto1.RegisterType((*BookResponse)(nil), "proto.BookResponse")
	proto1.RegisterType((*Book)(nil), "proto.Book")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BookService service

type BookServiceClient interface {
	Add(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*BookResponse, error)
	Get(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*BookResponse, error)
	Update(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*BookResponse, error)
	Delete(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*BookResponse, error)
	Filter(ctx context.Context, in *FilterRequest, opts ...grpc.CallOption) (*FilterResponse, error)
}

type bookServiceClient struct {
	cc *grpc.ClientConn
}

func NewBookServiceClient(cc *grpc.ClientConn) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) Add(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*BookResponse, error) {
	out := new(BookResponse)
	err := grpc.Invoke(ctx, "/proto.BookService/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Get(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*BookResponse, error) {
	out := new(BookResponse)
	err := grpc.Invoke(ctx, "/proto.BookService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Update(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*BookResponse, error) {
	out := new(BookResponse)
	err := grpc.Invoke(ctx, "/proto.BookService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Delete(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*BookResponse, error) {
	out := new(BookResponse)
	err := grpc.Invoke(ctx, "/proto.BookService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Filter(ctx context.Context, in *FilterRequest, opts ...grpc.CallOption) (*FilterResponse, error) {
	out := new(FilterResponse)
	err := grpc.Invoke(ctx, "/proto.BookService/Filter", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BookService service

type BookServiceServer interface {
	Add(context.Context, *BookRequest) (*BookResponse, error)
	Get(context.Context, *BookRequest) (*BookResponse, error)
	Update(context.Context, *BookRequest) (*BookResponse, error)
	Delete(context.Context, *BookRequest) (*BookResponse, error)
	Filter(context.Context, *FilterRequest) (*FilterResponse, error)
}

func RegisterBookServiceServer(s *grpc.Server, srv BookServiceServer) {
	s.RegisterService(&_BookService_serviceDesc, srv)
}

func _BookService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BookService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Add(ctx, req.(*BookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BookService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Get(ctx, req.(*BookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BookService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Update(ctx, req.(*BookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BookService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Delete(ctx, req.(*BookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Filter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Filter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BookService/Filter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Filter(ctx, req.(*FilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BookService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _BookService_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _BookService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _BookService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BookService_Delete_Handler,
		},
		{
			MethodName: "Filter",
			Handler:    _BookService_Filter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api.proto",
}

func init() { proto1.RegisterFile("proto/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcd, 0x4e, 0xdb, 0x40,
	0x10, 0xc7, 0xe3, 0xd8, 0x71, 0xd3, 0x71, 0x9b, 0xb6, 0xd3, 0xb4, 0x75, 0x2d, 0x24, 0x82, 0x4f,
	0x51, 0x0e, 0xb1, 0x08, 0x07, 0x24, 0x6e, 0x89, 0xf8, 0x08, 0xe2, 0x66, 0x84, 0x38, 0x22, 0x07,
	0x8f, 0xc2, 0x2a, 0x96, 0xd7, 0xb1, 0x37, 0x1c, 0x40, 0x5c, 0x78, 0x05, 0x1e, 0x8d, 0x57, 0xe0,
	0x2d, 0xb8, 0xa0, 0xdd, 0x35, 0x1f, 0xf9, 0x38, 0x00, 0xa7, 0x9d, 0xf9, 0xef, 0xcc, 0x6f, 0xe6,
	0xbf, 0x5a, 0xf8, 0x91, 0xe5, 0x5c, 0xf0, 0x20, 0xca, 0x58, 0x57, 0x45, 0x58, 0x53, 0x87, 0xb7,
	0x36, 0xe6, 0x7c, 0x9c, 0x90, 0xbc, 0x08, 0xa2, 0x34, 0xe5, 0x22, 0x12, 0x8c, 0xa7, 0x85, 0x2e,
	0xf2, 0xb7, 0xe1, 0xfb, 0x3e, 0x4b, 0x04, 0xe5, 0x21, 0x4d, 0x67, 0x54, 0x08, 0x44, 0xb0, 0xb2,
	0x68, 0x4c, 0xae, 0xd1, 0x32, 0xda, 0xb5, 0x50, 0xc5, 0x52, 0x2b, 0xd8, 0x15, 0xb9, 0x55, 0xad,
	0xc9, 0xd8, 0xdf, 0x84, 0xc6, 0x73, 0x63, 0x91, 0xf1, 0xb4, 0x20, 0x5c, 0x07, 0x6b, 0xc4, 0xf9,
	0xc4, 0x35, 0x5a, 0x66, 0xdb, 0xe9, 0x39, 0x7a, 0x40, 0x77, 0xc0, 0xf9, 0x24, 0x54, 0x17, 0xfe,
	0x29, 0x38, 0x2a, 0x2b, 0x27, 0x6d, 0xbc, 0xd4, 0x1b, 0x0b, 0xf5, 0xc3, 0x8a, 0xee, 0xc0, 0xff,
	0xf0, 0x45, 0x9e, 0x67, 0x2c, 0xd6, 0xb3, 0x87, 0x95, 0xd0, 0x96, 0xc2, 0x61, 0x3c, 0x00, 0xa8,
	0xab, 0xab, 0x9c, 0xa6, 0x7e, 0x00, 0xdf, 0x34, 0x78, 0x69, 0x13, 0x63, 0xf5, 0x26, 0x11, 0x58,
	0x32, 0xc3, 0x7f, 0xaf, 0x7c, 0xed, 0xb7, 0xa4, 0x4b, 0xc7, 0xac, 0x18, 0xa5, 0x6a, 0xea, 0xd7,
	0x50, 0xc5, 0xd8, 0x84, 0x9a, 0x60, 0x22, 0x21, 0xd7, 0x54, 0xa2, 0x4e, 0xf0, 0x2f, 0xd8, 0xd1,
	0x4c, 0x5c, 0xf0, 0xdc, 0xb5, 0x94, 0x5c, 0x66, 0xbd, 0xc7, 0xaa, 0x76, 0x7b, 0x4c, 0xf9, 0x25,
	0x3b, 0x27, 0xec, 0x83, 0xd9, 0x8f, 0x63, 0xc4, 0xb7, 0xcb, 0xe8, 0x87, 0xf0, 0x7e, 0xcf, 0x69,
	0xda, 0x83, 0xff, 0xeb, 0xf6, 0xfe, 0xe1, 0xae, 0xea, 0xf8, 0x76, 0x20, 0x57, 0x2a, 0x76, 0x8c,
	0x0e, 0x0e, 0xc1, 0x3c, 0x20, 0xf1, 0x7e, 0x84, 0xab, 0x10, 0x88, 0x3f, 0x35, 0x22, 0xb8, 0x2e,
	0xbd, 0xde, 0xe0, 0x11, 0xd8, 0x27, 0x59, 0x1c, 0x09, 0xfa, 0x30, 0xcc, 0x5b, 0x09, 0xdb, 0xa5,
	0x84, 0x3e, 0x01, 0xeb, 0x2c, 0xc3, 0xf6, 0xc0, 0xd6, 0xdf, 0x0a, 0x9b, 0x65, 0xe3, 0xdc, 0xf7,
	0xf4, 0xfe, 0x2c, 0xa8, 0x25, 0xb0, 0xa1, 0x80, 0x75, 0x2c, 0x5f, 0x6b, 0x64, 0xab, 0xaa, 0xad,
	0xa7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x24, 0xf0, 0xd1, 0x15, 0x03, 0x00, 0x00,
}
