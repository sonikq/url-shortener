// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: internal/app/proto/shortener.proto

package url_shortener

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type ShortenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Url    string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ShortenRequest) Reset() {
	*x = ShortenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_shortener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenRequest) ProtoMessage() {}

func (x *ShortenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_shortener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_internal_app_proto_shortener_proto_rawDescGZIP(), []int{0}
}

func (x *ShortenRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ShortenRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ShortenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shorten string `protobuf:"bytes,1,opt,name=shorten,proto3" json:"shorten,omitempty"`
}

func (x *ShortenResponse) Reset() {
	*x = ShortenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_shortener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenResponse) ProtoMessage() {}

func (x *ShortenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_shortener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_internal_app_proto_shortener_proto_rawDescGZIP(), []int{1}
}

func (x *ShortenResponse) GetShorten() string {
	if x != nil {
		return x.Shorten
	}
	return ""
}

type ExpandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=shortUrl,proto3" json:"shortUrl,omitempty"`
}

func (x *ExpandRequest) Reset() {
	*x = ExpandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_shortener_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandRequest) ProtoMessage() {}

func (x *ExpandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_shortener_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandRequest.ProtoReflect.Descriptor instead.
func (*ExpandRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_shortener_proto_rawDescGZIP(), []int{2}
}

func (x *ExpandRequest) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type ExpandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ExpandResponse) Reset() {
	*x = ExpandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_shortener_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandResponse) ProtoMessage() {}

func (x *ExpandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_shortener_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandResponse.ProtoReflect.Descriptor instead.
func (*ExpandResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_shortener_proto_rawDescGZIP(), []int{3}
}

func (x *ExpandResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type GetBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetBatchRequest) Reset() {
	*x = GetBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_shortener_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBatchRequest) ProtoMessage() {}

func (x *GetBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_shortener_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBatchRequest.ProtoReflect.Descriptor instead.
func (*GetBatchRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_shortener_proto_rawDescGZIP(), []int{4}
}

func (x *GetBatchRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetBatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rows []*UrlRow `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (x *GetBatchResponse) Reset() {
	*x = GetBatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_shortener_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBatchResponse) ProtoMessage() {}

func (x *GetBatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_shortener_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBatchResponse.ProtoReflect.Descriptor instead.
func (*GetBatchResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_shortener_proto_rawDescGZIP(), []int{5}
}

func (x *GetBatchResponse) GetRows() []*UrlRow {
	if x != nil {
		return x.Rows
	}
	return nil
}

type UrlRow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalURL string `protobuf:"bytes,1,opt,name=originalURL,proto3" json:"originalURL,omitempty"`
	ShortURL    string `protobuf:"bytes,2,opt,name=shortURL,proto3" json:"shortURL,omitempty"`
}

func (x *UrlRow) Reset() {
	*x = UrlRow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_shortener_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UrlRow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlRow) ProtoMessage() {}

func (x *UrlRow) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_shortener_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlRow.ProtoReflect.Descriptor instead.
func (*UrlRow) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_shortener_proto_rawDescGZIP(), []int{6}
}

func (x *UrlRow) GetOriginalURL() string {
	if x != nil {
		return x.OriginalURL
	}
	return ""
}

func (x *UrlRow) GetShortURL() string {
	if x != nil {
		return x.ShortURL
	}
	return ""
}

type GetStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls  int32 `protobuf:"varint,1,opt,name=urls,proto3" json:"urls,omitempty"`
	Users int32 `protobuf:"varint,2,opt,name=users,proto3" json:"users,omitempty"`
}

func (x *GetStatsResponse) Reset() {
	*x = GetStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_shortener_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatsResponse) ProtoMessage() {}

func (x *GetStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_shortener_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatsResponse.ProtoReflect.Descriptor instead.
func (*GetStatsResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_shortener_proto_rawDescGZIP(), []int{7}
}

func (x *GetStatsResponse) GetUrls() int32 {
	if x != nil {
		return x.Urls
	}
	return 0
}

func (x *GetStatsResponse) GetUsers() int32 {
	if x != nil {
		return x.Users
	}
	return 0
}

type CorrelatedOriginalURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationId string `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	OriginalUrl   string `protobuf:"bytes,2,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *CorrelatedOriginalURL) Reset() {
	*x = CorrelatedOriginalURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_shortener_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CorrelatedOriginalURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CorrelatedOriginalURL) ProtoMessage() {}

func (x *CorrelatedOriginalURL) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_shortener_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CorrelatedOriginalURL.ProtoReflect.Descriptor instead.
func (*CorrelatedOriginalURL) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_shortener_proto_rawDescGZIP(), []int{8}
}

func (x *CorrelatedOriginalURL) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *CorrelatedOriginalURL) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type ShortBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string                   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Original []*CorrelatedOriginalURL `protobuf:"bytes,2,rep,name=original,proto3" json:"original,omitempty"`
}

func (x *ShortBatchRequest) Reset() {
	*x = ShortBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_shortener_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortBatchRequest) ProtoMessage() {}

func (x *ShortBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_shortener_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortBatchRequest.ProtoReflect.Descriptor instead.
func (*ShortBatchRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_shortener_proto_rawDescGZIP(), []int{9}
}

func (x *ShortBatchRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ShortBatchRequest) GetOriginal() []*CorrelatedOriginalURL {
	if x != nil {
		return x.Original
	}
	return nil
}

type CorrelationShortURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationId string `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	ShortUrl      string `protobuf:"bytes,2,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *CorrelationShortURL) Reset() {
	*x = CorrelationShortURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_shortener_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CorrelationShortURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CorrelationShortURL) ProtoMessage() {}

func (x *CorrelationShortURL) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_shortener_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CorrelationShortURL.ProtoReflect.Descriptor instead.
func (*CorrelationShortURL) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_shortener_proto_rawDescGZIP(), []int{10}
}

func (x *CorrelationShortURL) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *CorrelationShortURL) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type ShortBatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Original []*CorrelationShortURL `protobuf:"bytes,1,rep,name=original,proto3" json:"original,omitempty"`
}

func (x *ShortBatchResponse) Reset() {
	*x = ShortBatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_shortener_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortBatchResponse) ProtoMessage() {}

func (x *ShortBatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_shortener_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortBatchResponse.ProtoReflect.Descriptor instead.
func (*ShortBatchResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_shortener_proto_rawDescGZIP(), []int{11}
}

func (x *ShortBatchResponse) GetOriginal() []*CorrelationShortURL {
	if x != nil {
		return x.Original
	}
	return nil
}

var File_internal_app_proto_shortener_proto protoreflect.FileDescriptor

var file_internal_app_proto_shortener_proto_rawDesc = []byte{
	0x0a, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x0e,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x2b, 0x0a, 0x0f, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x22, 0x2b, 0x0a, 0x0d, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55,
	0x72, 0x6c, 0x22, 0x22, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x29, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x39, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e,
	0x75, 0x72, 0x6c, 0x52, 0x6f, 0x77, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x22, 0x46, 0x0a, 0x06,
	0x75, 0x72, 0x6c, 0x52, 0x6f, 0x77, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x55, 0x52, 0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x55, 0x52, 0x4c, 0x22, 0x3c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x22, 0x61, 0x0a, 0x15, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64,
	0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x12, 0x25, 0x0a, 0x0e, 0x63,
	0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x55, 0x72, 0x6c, 0x22, 0x6a, 0x0a, 0x11, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x2e, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x52, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x22, 0x59, 0x0a, 0x13, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x50, 0x0a, 0x12,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3a, 0x0a, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x2e, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x55, 0x52, 0x4c, 0x52, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x32, 0x90,
	0x03, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x07,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x12, 0x19, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d,
	0x0a, 0x06, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x12, 0x18, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x45,
	0x78, 0x70, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1a, 0x2e, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x44, 0x0a, 0x05, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1c, 0x2e, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x3f, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6e, 0x69, 0x6b, 0x71, 0x2f, 0x75, 0x72, 0x6c, 0x2d, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_app_proto_shortener_proto_rawDescOnce sync.Once
	file_internal_app_proto_shortener_proto_rawDescData = file_internal_app_proto_shortener_proto_rawDesc
)

func file_internal_app_proto_shortener_proto_rawDescGZIP() []byte {
	file_internal_app_proto_shortener_proto_rawDescOnce.Do(func() {
		file_internal_app_proto_shortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_app_proto_shortener_proto_rawDescData)
	})
	return file_internal_app_proto_shortener_proto_rawDescData
}

var file_internal_app_proto_shortener_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_internal_app_proto_shortener_proto_goTypes = []any{
	(*ShortenRequest)(nil),        // 0: shortener.ShortenRequest
	(*ShortenResponse)(nil),       // 1: shortener.ShortenResponse
	(*ExpandRequest)(nil),         // 2: shortener.ExpandRequest
	(*ExpandResponse)(nil),        // 3: shortener.ExpandResponse
	(*GetBatchRequest)(nil),       // 4: shortener.GetBatchRequest
	(*GetBatchResponse)(nil),      // 5: shortener.GetBatchResponse
	(*UrlRow)(nil),                // 6: shortener.urlRow
	(*GetStatsResponse)(nil),      // 7: shortener.GetStatsResponse
	(*CorrelatedOriginalURL)(nil), // 8: shortener.CorrelatedOriginalURL
	(*ShortBatchRequest)(nil),     // 9: shortener.ShortBatchRequest
	(*CorrelationShortURL)(nil),   // 10: shortener.CorrelationShortURL
	(*ShortBatchResponse)(nil),    // 11: shortener.ShortBatchResponse
	(*empty.Empty)(nil),           // 12: google.protobuf.Empty
}
var file_internal_app_proto_shortener_proto_depIdxs = []int32{
	6,  // 0: shortener.GetBatchResponse.rows:type_name -> shortener.urlRow
	8,  // 1: shortener.ShortBatchRequest.original:type_name -> shortener.CorrelatedOriginalURL
	10, // 2: shortener.ShortBatchResponse.original:type_name -> shortener.CorrelationShortURL
	0,  // 3: shortener.Shortener.Shorten:input_type -> shortener.ShortenRequest
	2,  // 4: shortener.Shortener.Expand:input_type -> shortener.ExpandRequest
	4,  // 5: shortener.Shortener.GetBatch:input_type -> shortener.GetBatchRequest
	9,  // 6: shortener.Shortener.Batch:input_type -> shortener.ShortBatchRequest
	12, // 7: shortener.Shortener.Ping:input_type -> google.protobuf.Empty
	12, // 8: shortener.Shortener.GetStats:input_type -> google.protobuf.Empty
	1,  // 9: shortener.Shortener.Shorten:output_type -> shortener.ShortenResponse
	3,  // 10: shortener.Shortener.Expand:output_type -> shortener.ExpandResponse
	5,  // 11: shortener.Shortener.GetBatch:output_type -> shortener.GetBatchResponse
	11, // 12: shortener.Shortener.Batch:output_type -> shortener.ShortBatchResponse
	12, // 13: shortener.Shortener.Ping:output_type -> google.protobuf.Empty
	7,  // 14: shortener.Shortener.GetStats:output_type -> shortener.GetStatsResponse
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_internal_app_proto_shortener_proto_init() }
func file_internal_app_proto_shortener_proto_init() {
	if File_internal_app_proto_shortener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_app_proto_shortener_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ShortenRequest); i {
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
		file_internal_app_proto_shortener_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ShortenResponse); i {
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
		file_internal_app_proto_shortener_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ExpandRequest); i {
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
		file_internal_app_proto_shortener_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ExpandResponse); i {
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
		file_internal_app_proto_shortener_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetBatchRequest); i {
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
		file_internal_app_proto_shortener_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetBatchResponse); i {
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
		file_internal_app_proto_shortener_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UrlRow); i {
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
		file_internal_app_proto_shortener_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetStatsResponse); i {
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
		file_internal_app_proto_shortener_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*CorrelatedOriginalURL); i {
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
		file_internal_app_proto_shortener_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*ShortBatchRequest); i {
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
		file_internal_app_proto_shortener_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*CorrelationShortURL); i {
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
		file_internal_app_proto_shortener_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*ShortBatchResponse); i {
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
			RawDescriptor: file_internal_app_proto_shortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_app_proto_shortener_proto_goTypes,
		DependencyIndexes: file_internal_app_proto_shortener_proto_depIdxs,
		MessageInfos:      file_internal_app_proto_shortener_proto_msgTypes,
	}.Build()
	File_internal_app_proto_shortener_proto = out.File
	file_internal_app_proto_shortener_proto_rawDesc = nil
	file_internal_app_proto_shortener_proto_goTypes = nil
	file_internal_app_proto_shortener_proto_depIdxs = nil
}
