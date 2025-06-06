// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: url.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Запрос на сокращение одного URL
type ShortenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShortenRequest) Reset() {
	*x = ShortenRequest{}
	mi := &file_url_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShortenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenRequest) ProtoMessage() {}

func (x *ShortenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_url_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenRequest.ProtoReflect.Descriptor instead.
func (*ShortenRequest) Descriptor() ([]byte, []int) {
	return file_url_proto_rawDescGZIP(), []int{0}
}

func (x *ShortenRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ShortenRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// Ответ на сокращение одного URL
type ShortenResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ShortUrl      string                 `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShortenResponse) Reset() {
	*x = ShortenResponse{}
	mi := &file_url_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShortenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenResponse) ProtoMessage() {}

func (x *ShortenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_url_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenResponse.ProtoReflect.Descriptor instead.
func (*ShortenResponse) Descriptor() ([]byte, []int) {
	return file_url_proto_rawDescGZIP(), []int{1}
}

func (x *ShortenResponse) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

// Запрос на получение оригинального URL
type URLRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ShortUrl      string                 `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *URLRequest) Reset() {
	*x = URLRequest{}
	mi := &file_url_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *URLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URLRequest) ProtoMessage() {}

func (x *URLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_url_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URLRequest.ProtoReflect.Descriptor instead.
func (*URLRequest) Descriptor() ([]byte, []int) {
	return file_url_proto_rawDescGZIP(), []int{2}
}

func (x *URLRequest) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

// Ответ с оригинальным URL и флагом удаления
type URLResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OriginalUrl   string                 `protobuf:"bytes,1,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
	IsDeleted     bool                   `protobuf:"varint,2,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *URLResponse) Reset() {
	*x = URLResponse{}
	mi := &file_url_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *URLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URLResponse) ProtoMessage() {}

func (x *URLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_url_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URLResponse.ProtoReflect.Descriptor instead.
func (*URLResponse) Descriptor() ([]byte, []int) {
	return file_url_proto_rawDescGZIP(), []int{3}
}

func (x *URLResponse) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

func (x *URLResponse) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

// Запрос на удаление URL пользователя
type DeleteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ShortUrls     []string               `protobuf:"bytes,1,rep,name=short_urls,json=shortUrls,proto3" json:"short_urls,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	mi := &file_url_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_url_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_url_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRequest) GetShortUrls() []string {
	if x != nil {
		return x.ShortUrls
	}
	return nil
}

func (x *DeleteRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// Ответ об успешности удаления
type DeleteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	mi := &file_url_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_url_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_url_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// Запрос на получение всех URL пользователя
type UserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	mi := &file_url_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_url_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_url_proto_rawDescGZIP(), []int{6}
}

func (x *UserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// Элемент из списка URL пользователя
type UserURLItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ShortUrl      string                 `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	OriginalUrl   string                 `protobuf:"bytes,2,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserURLItem) Reset() {
	*x = UserURLItem{}
	mi := &file_url_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserURLItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserURLItem) ProtoMessage() {}

func (x *UserURLItem) ProtoReflect() protoreflect.Message {
	mi := &file_url_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserURLItem.ProtoReflect.Descriptor instead.
func (*UserURLItem) Descriptor() ([]byte, []int) {
	return file_url_proto_rawDescGZIP(), []int{7}
}

func (x *UserURLItem) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

func (x *UserURLItem) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

// Ответ со списком URL
type UserURLsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Urls          []*UserURLItem         `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserURLsResponse) Reset() {
	*x = UserURLsResponse{}
	mi := &file_url_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserURLsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserURLsResponse) ProtoMessage() {}

func (x *UserURLsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_url_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserURLsResponse.ProtoReflect.Descriptor instead.
func (*UserURLsResponse) Descriptor() ([]byte, []int) {
	return file_url_proto_rawDescGZIP(), []int{8}
}

func (x *UserURLsResponse) GetUrls() []*UserURLItem {
	if x != nil {
		return x.Urls
	}
	return nil
}

// Запрос на batch-сокращение URL
type BatchRequestItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CorrelationId string                 `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	OriginalUrl   string                 `protobuf:"bytes,2,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BatchRequestItem) Reset() {
	*x = BatchRequestItem{}
	mi := &file_url_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchRequestItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchRequestItem) ProtoMessage() {}

func (x *BatchRequestItem) ProtoReflect() protoreflect.Message {
	mi := &file_url_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchRequestItem.ProtoReflect.Descriptor instead.
func (*BatchRequestItem) Descriptor() ([]byte, []int) {
	return file_url_proto_rawDescGZIP(), []int{9}
}

func (x *BatchRequestItem) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *BatchRequestItem) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type BatchRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Urls          []*BatchRequestItem    `protobuf:"bytes,2,rep,name=urls,proto3" json:"urls,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BatchRequest) Reset() {
	*x = BatchRequest{}
	mi := &file_url_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchRequest) ProtoMessage() {}

func (x *BatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_url_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchRequest.ProtoReflect.Descriptor instead.
func (*BatchRequest) Descriptor() ([]byte, []int) {
	return file_url_proto_rawDescGZIP(), []int{10}
}

func (x *BatchRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BatchRequest) GetUrls() []*BatchRequestItem {
	if x != nil {
		return x.Urls
	}
	return nil
}

// Ответ с batch-сокращёнными URL
type BatchResponseItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CorrelationId string                 `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	ShortUrl      string                 `protobuf:"bytes,2,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BatchResponseItem) Reset() {
	*x = BatchResponseItem{}
	mi := &file_url_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchResponseItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchResponseItem) ProtoMessage() {}

func (x *BatchResponseItem) ProtoReflect() protoreflect.Message {
	mi := &file_url_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchResponseItem.ProtoReflect.Descriptor instead.
func (*BatchResponseItem) Descriptor() ([]byte, []int) {
	return file_url_proto_rawDescGZIP(), []int{11}
}

func (x *BatchResponseItem) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *BatchResponseItem) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type BatchResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Urls          []*BatchResponseItem   `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BatchResponse) Reset() {
	*x = BatchResponse{}
	mi := &file_url_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchResponse) ProtoMessage() {}

func (x *BatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_url_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchResponse.ProtoReflect.Descriptor instead.
func (*BatchResponse) Descriptor() ([]byte, []int) {
	return file_url_proto_rawDescGZIP(), []int{12}
}

func (x *BatchResponse) GetUrls() []*BatchResponseItem {
	if x != nil {
		return x.Urls
	}
	return nil
}

var File_url_proto protoreflect.FileDescriptor

const file_url_proto_rawDesc = "" +
	"\n" +
	"\turl.proto\x12\tshortener\";\n" +
	"\x0eShortenRequest\x12\x10\n" +
	"\x03url\x18\x01 \x01(\tR\x03url\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\".\n" +
	"\x0fShortenResponse\x12\x1b\n" +
	"\tshort_url\x18\x01 \x01(\tR\bshortUrl\")\n" +
	"\n" +
	"URLRequest\x12\x1b\n" +
	"\tshort_url\x18\x01 \x01(\tR\bshortUrl\"O\n" +
	"\vURLResponse\x12!\n" +
	"\foriginal_url\x18\x01 \x01(\tR\voriginalUrl\x12\x1d\n" +
	"\n" +
	"is_deleted\x18\x02 \x01(\bR\tisDeleted\"G\n" +
	"\rDeleteRequest\x12\x1d\n" +
	"\n" +
	"short_urls\x18\x01 \x03(\tR\tshortUrls\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\"*\n" +
	"\x0eDeleteResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"&\n" +
	"\vUserRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\"M\n" +
	"\vUserURLItem\x12\x1b\n" +
	"\tshort_url\x18\x01 \x01(\tR\bshortUrl\x12!\n" +
	"\foriginal_url\x18\x02 \x01(\tR\voriginalUrl\">\n" +
	"\x10UserURLsResponse\x12*\n" +
	"\x04urls\x18\x01 \x03(\v2\x16.shortener.UserURLItemR\x04urls\"\\\n" +
	"\x10BatchRequestItem\x12%\n" +
	"\x0ecorrelation_id\x18\x01 \x01(\tR\rcorrelationId\x12!\n" +
	"\foriginal_url\x18\x02 \x01(\tR\voriginalUrl\"X\n" +
	"\fBatchRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12/\n" +
	"\x04urls\x18\x02 \x03(\v2\x1b.shortener.BatchRequestItemR\x04urls\"W\n" +
	"\x11BatchResponseItem\x12%\n" +
	"\x0ecorrelation_id\x18\x01 \x01(\tR\rcorrelationId\x12\x1b\n" +
	"\tshort_url\x18\x02 \x01(\tR\bshortUrl\"A\n" +
	"\rBatchResponse\x120\n" +
	"\x04urls\x18\x01 \x03(\v2\x1c.shortener.BatchResponseItemR\x04urls2\xdd\x02\n" +
	"\n" +
	"URLService\x12C\n" +
	"\n" +
	"ShortenURL\x12\x19.shortener.ShortenRequest\x1a\x1a.shortener.ShortenResponse\x12?\n" +
	"\x0eGetOriginalURL\x12\x15.shortener.URLRequest\x1a\x16.shortener.URLResponse\x12E\n" +
	"\x0eDeleteUserURLs\x12\x18.shortener.DeleteRequest\x1a\x19.shortener.DeleteResponse\x12B\n" +
	"\vGetUserURLs\x12\x16.shortener.UserRequest\x1a\x1b.shortener.UserURLsResponse\x12>\n" +
	"\tSaveBatch\x12\x17.shortener.BatchRequest\x1a\x18.shortener.BatchResponseB6Z4github.com/Wrestler094/shortener/internal/grpc/pb;pbb\x06proto3"

var (
	file_url_proto_rawDescOnce sync.Once
	file_url_proto_rawDescData []byte
)

func file_url_proto_rawDescGZIP() []byte {
	file_url_proto_rawDescOnce.Do(func() {
		file_url_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_url_proto_rawDesc), len(file_url_proto_rawDesc)))
	})
	return file_url_proto_rawDescData
}

var file_url_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_url_proto_goTypes = []any{
	(*ShortenRequest)(nil),    // 0: shortener.ShortenRequest
	(*ShortenResponse)(nil),   // 1: shortener.ShortenResponse
	(*URLRequest)(nil),        // 2: shortener.URLRequest
	(*URLResponse)(nil),       // 3: shortener.URLResponse
	(*DeleteRequest)(nil),     // 4: shortener.DeleteRequest
	(*DeleteResponse)(nil),    // 5: shortener.DeleteResponse
	(*UserRequest)(nil),       // 6: shortener.UserRequest
	(*UserURLItem)(nil),       // 7: shortener.UserURLItem
	(*UserURLsResponse)(nil),  // 8: shortener.UserURLsResponse
	(*BatchRequestItem)(nil),  // 9: shortener.BatchRequestItem
	(*BatchRequest)(nil),      // 10: shortener.BatchRequest
	(*BatchResponseItem)(nil), // 11: shortener.BatchResponseItem
	(*BatchResponse)(nil),     // 12: shortener.BatchResponse
}
var file_url_proto_depIdxs = []int32{
	7,  // 0: shortener.UserURLsResponse.urls:type_name -> shortener.UserURLItem
	9,  // 1: shortener.BatchRequest.urls:type_name -> shortener.BatchRequestItem
	11, // 2: shortener.BatchResponse.urls:type_name -> shortener.BatchResponseItem
	0,  // 3: shortener.URLService.ShortenURL:input_type -> shortener.ShortenRequest
	2,  // 4: shortener.URLService.GetOriginalURL:input_type -> shortener.URLRequest
	4,  // 5: shortener.URLService.DeleteUserURLs:input_type -> shortener.DeleteRequest
	6,  // 6: shortener.URLService.GetUserURLs:input_type -> shortener.UserRequest
	10, // 7: shortener.URLService.SaveBatch:input_type -> shortener.BatchRequest
	1,  // 8: shortener.URLService.ShortenURL:output_type -> shortener.ShortenResponse
	3,  // 9: shortener.URLService.GetOriginalURL:output_type -> shortener.URLResponse
	5,  // 10: shortener.URLService.DeleteUserURLs:output_type -> shortener.DeleteResponse
	8,  // 11: shortener.URLService.GetUserURLs:output_type -> shortener.UserURLsResponse
	12, // 12: shortener.URLService.SaveBatch:output_type -> shortener.BatchResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_url_proto_init() }
func file_url_proto_init() {
	if File_url_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_url_proto_rawDesc), len(file_url_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_url_proto_goTypes,
		DependencyIndexes: file_url_proto_depIdxs,
		MessageInfos:      file_url_proto_msgTypes,
	}.Build()
	File_url_proto = out.File
	file_url_proto_goTypes = nil
	file_url_proto_depIdxs = nil
}
