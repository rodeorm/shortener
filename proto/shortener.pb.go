// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0--rc3
// source: shortener.proto

package proto

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

// URL
type URL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalURL    string `protobuf:"bytes,1,opt,name=originalURL,proto3" json:"originalURL,omitempty"`        // Оригинальный урл
	Key            string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`                        // Ключ, использованный при сокращении
	UserKey        int32  `protobuf:"zigzag32,3,opt,name=userKey,proto3" json:"userKey,omitempty"`             // Пользователь, который сократил URL
	HasBeenShorted bool   `protobuf:"varint,4,opt,name=hasBeenShorted,proto3" json:"hasBeenShorted,omitempty"` // Признак, что сокращали ранее
	HasBeenDeleted bool   `protobuf:"varint,5,opt,name=hasBeenDeleted,proto3" json:"hasBeenDeleted,omitempty"` // Признал, что был удален
}

func (x *URL) Reset() {
	*x = URL{}
	mi := &file_shortener_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *URL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URL) ProtoMessage() {}

func (x *URL) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URL.ProtoReflect.Descriptor instead.
func (*URL) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{0}
}

func (x *URL) GetOriginalURL() string {
	if x != nil {
		return x.OriginalURL
	}
	return ""
}

func (x *URL) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *URL) GetUserKey() int32 {
	if x != nil {
		return x.UserKey
	}
	return 0
}

func (x *URL) GetHasBeenShorted() bool {
	if x != nil {
		return x.HasBeenShorted
	}
	return false
}

func (x *URL) GetHasBeenDeleted() bool {
	if x != nil {
		return x.HasBeenDeleted
	}
	return false
}

// User - пользователь сервиса
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key            int32  `protobuf:"zigzag32,1,opt,name=key,proto3" json:"key,omitempty"`                     // Уникальный идентификатор пользователя
	WasUnathorized bool   `protobuf:"varint,2,opt,name=wasUnathorized,proto3" json:"wasUnathorized,omitempty"` // Признак того, что пользователь был создан автоматически, после того как не получилось авторизовать его через куки
	Urls           []*URL `protobuf:"bytes,3,rep,name=urls,proto3" json:"urls,omitempty"`                      // Сокращенные пользователем URL
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_shortener_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetKey() int32 {
	if x != nil {
		return x.Key
	}
	return 0
}

func (x *User) GetWasUnathorized() bool {
	if x != nil {
		return x.WasUnathorized
	}
	return false
}

func (x *User) GetUrls() []*URL {
	if x != nil {
		return x.Urls
	}
	return nil
}

// Statistic - статистика по сокращенным URL и количеству пользователей в сервисе
type Statistic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls  int32 `protobuf:"zigzag32,1,opt,name=urls,proto3" json:"urls,omitempty"`   // Количество сокращённых URL в сервисе
	Users int32 `protobuf:"zigzag32,2,opt,name=users,proto3" json:"users,omitempty"` // Количество пользователей в сервисе
}

func (x *Statistic) Reset() {
	*x = Statistic{}
	mi := &file_shortener_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Statistic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Statistic) ProtoMessage() {}

func (x *Statistic) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Statistic.ProtoReflect.Descriptor instead.
func (*Statistic) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{2}
}

func (x *Statistic) GetUrls() int32 {
	if x != nil {
		return x.Urls
	}
	return 0
}

func (x *Statistic) GetUsers() int32 {
	if x != nil {
		return x.Users
	}
	return 0
}

// ShortenRequest запрос для Shorten
type ShortenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ShortenRequest) Reset() {
	*x = ShortenRequest{}
	mi := &file_shortener_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShortenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenRequest) ProtoMessage() {}

func (x *ShortenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[3]
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
	return file_shortener_proto_rawDescGZIP(), []int{3}
}

func (x *ShortenRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// ShortenResponse ответ для Shorten
type ShortenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ShortenResponse) Reset() {
	*x = ShortenResponse{}
	mi := &file_shortener_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShortenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenResponse) ProtoMessage() {}

func (x *ShortenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[4]
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
	return file_shortener_proto_rawDescGZIP(), []int{4}
}

func (x *ShortenResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// PingDBRequest запрос для PingDB
type PingDBRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingDBRequest) Reset() {
	*x = PingDBRequest{}
	mi := &file_shortener_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PingDBRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingDBRequest) ProtoMessage() {}

func (x *PingDBRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingDBRequest.ProtoReflect.Descriptor instead.
func (*PingDBRequest) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{5}
}

// PingDBResponse ответ для PingDB
type PingDBResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingDBResponse) Reset() {
	*x = PingDBResponse{}
	mi := &file_shortener_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PingDBResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingDBResponse) ProtoMessage() {}

func (x *PingDBResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingDBResponse.ProtoReflect.Descriptor instead.
func (*PingDBResponse) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{6}
}

// UserURLsRequest запрос для UserURLs
type UserURLsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserURLsRequest) Reset() {
	*x = UserURLsRequest{}
	mi := &file_shortener_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserURLsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserURLsRequest) ProtoMessage() {}

func (x *UserURLsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserURLsRequest.ProtoReflect.Descriptor instead.
func (*UserURLsRequest) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{7}
}

func (x *UserURLsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// UserURLsResponse ответ для UserURLs
type UserURLsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserURLsResponse) Reset() {
	*x = UserURLsResponse{}
	mi := &file_shortener_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserURLsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserURLsResponse) ProtoMessage() {}

func (x *UserURLsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[8]
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
	return file_shortener_proto_rawDescGZIP(), []int{8}
}

func (x *UserURLsResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

// DeleteURLsRequest запрос для DeleteURLs
type DeleteURLsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UrlsToDelete string `protobuf:"bytes,1,opt,name=urls_to_delete,json=urlsToDelete,proto3" json:"urls_to_delete,omitempty"`
}

func (x *DeleteURLsRequest) Reset() {
	*x = DeleteURLsRequest{}
	mi := &file_shortener_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteURLsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteURLsRequest) ProtoMessage() {}

func (x *DeleteURLsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteURLsRequest.ProtoReflect.Descriptor instead.
func (*DeleteURLsRequest) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteURLsRequest) GetUrlsToDelete() string {
	if x != nil {
		return x.UrlsToDelete
	}
	return ""
}

// DeleteURLsResponse ответ для DeleteURLs
type DeleteURLsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteURLsResponse) Reset() {
	*x = DeleteURLsResponse{}
	mi := &file_shortener_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteURLsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteURLsResponse) ProtoMessage() {}

func (x *DeleteURLsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteURLsResponse.ProtoReflect.Descriptor instead.
func (*DeleteURLsResponse) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{10}
}

// StatsRequest запрос для Stats
type StatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatsRequest) Reset() {
	*x = StatsRequest{}
	mi := &file_shortener_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsRequest) ProtoMessage() {}

func (x *StatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsRequest.ProtoReflect.Descriptor instead.
func (*StatsRequest) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{11}
}

// StatsRequest ответ для Stats
type StatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Statistic *Statistic `protobuf:"bytes,1,opt,name=statistic,proto3" json:"statistic,omitempty"`
}

func (x *StatsResponse) Reset() {
	*x = StatsResponse{}
	mi := &file_shortener_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsResponse) ProtoMessage() {}

func (x *StatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsResponse.ProtoReflect.Descriptor instead.
func (*StatsResponse) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{12}
}

func (x *StatsResponse) GetStatistic() *Statistic {
	if x != nil {
		return x.Statistic
	}
	return nil
}

// RootRequest запрос для Root
type RootRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *RootRequest) Reset() {
	*x = RootRequest{}
	mi := &file_shortener_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RootRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RootRequest) ProtoMessage() {}

func (x *RootRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RootRequest.ProtoReflect.Descriptor instead.
func (*RootRequest) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{13}
}

func (x *RootRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// RootResponse ответ для Root
type RootResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shorten string `protobuf:"bytes,1,opt,name=shorten,proto3" json:"shorten,omitempty"`
}

func (x *RootResponse) Reset() {
	*x = RootResponse{}
	mi := &file_shortener_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RootResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RootResponse) ProtoMessage() {}

func (x *RootResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RootResponse.ProtoReflect.Descriptor instead.
func (*RootResponse) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{14}
}

func (x *RootResponse) GetShorten() string {
	if x != nil {
		return x.Shorten
	}
	return ""
}

var File_shortener_proto protoreflect.FileDescriptor

var file_shortener_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x22, 0xa3, 0x01, 0x0a,
	0x03, 0x55, 0x52, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c,
	0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x11, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x4b,
	0x65, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x68, 0x61, 0x73, 0x42, 0x65, 0x65, 0x6e, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x68, 0x61, 0x73, 0x42,
	0x65, 0x65, 0x6e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x68, 0x61,
	0x73, 0x42, 0x65, 0x65, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0e, 0x68, 0x61, 0x73, 0x42, 0x65, 0x65, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x22, 0x64, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x0e,
	0x77, 0x61, 0x73, 0x55, 0x6e, 0x61, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x77, 0x61, 0x73, 0x55, 0x6e, 0x61, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x55,
	0x52, 0x4c, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x22, 0x35, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x11, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22,
	0x22, 0x0a, 0x0e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x22, 0x23, 0x0a, 0x0f, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x0f, 0x0a, 0x0d, 0x50, 0x69, 0x6e, 0x67,
	0x44, 0x42, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x50, 0x69, 0x6e,
	0x67, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x0a, 0x0f, 0x55,
	0x73, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x37, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x55,
	0x52, 0x4c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x22, 0x39, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x75, 0x72, 0x6c, 0x73, 0x5f, 0x74, 0x6f,
	0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75,
	0x72, 0x6c, 0x73, 0x54, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x43, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x22, 0x1f, 0x0a, 0x0b, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x28, 0x0a, 0x0c, 0x52, 0x6f, 0x6f, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x32, 0x99, 0x03, 0x0a, 0x0a, 0x55, 0x52, 0x4c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x40, 0x0a, 0x07, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x12, 0x19, 0x2e, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x50, 0x69, 0x6e, 0x67, 0x44, 0x42, 0x12, 0x18, 0x2e, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x44, 0x42, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x46, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x73,
	0x12, 0x1a, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x52, 0x4c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x73, 0x12, 0x1c, 0x2e, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x52,
	0x4c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x52, 0x4c, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x12, 0x17, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x16, 0x2e, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x2e, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a,
	0x0f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shortener_proto_rawDescOnce sync.Once
	file_shortener_proto_rawDescData = file_shortener_proto_rawDesc
)

func file_shortener_proto_rawDescGZIP() []byte {
	file_shortener_proto_rawDescOnce.Do(func() {
		file_shortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_shortener_proto_rawDescData)
	})
	return file_shortener_proto_rawDescData
}

var file_shortener_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_shortener_proto_goTypes = []any{
	(*URL)(nil),                // 0: shortener.URL
	(*User)(nil),               // 1: shortener.User
	(*Statistic)(nil),          // 2: shortener.Statistic
	(*ShortenRequest)(nil),     // 3: shortener.ShortenRequest
	(*ShortenResponse)(nil),    // 4: shortener.ShortenResponse
	(*PingDBRequest)(nil),      // 5: shortener.PingDBRequest
	(*PingDBResponse)(nil),     // 6: shortener.PingDBResponse
	(*UserURLsRequest)(nil),    // 7: shortener.UserURLsRequest
	(*UserURLsResponse)(nil),   // 8: shortener.UserURLsResponse
	(*DeleteURLsRequest)(nil),  // 9: shortener.DeleteURLsRequest
	(*DeleteURLsResponse)(nil), // 10: shortener.DeleteURLsResponse
	(*StatsRequest)(nil),       // 11: shortener.StatsRequest
	(*StatsResponse)(nil),      // 12: shortener.StatsResponse
	(*RootRequest)(nil),        // 13: shortener.RootRequest
	(*RootResponse)(nil),       // 14: shortener.RootResponse
}
var file_shortener_proto_depIdxs = []int32{
	0,  // 0: shortener.User.urls:type_name -> shortener.URL
	1,  // 1: shortener.UserURLsResponse.user:type_name -> shortener.User
	2,  // 2: shortener.StatsResponse.statistic:type_name -> shortener.Statistic
	3,  // 3: shortener.URLService.Shorten:input_type -> shortener.ShortenRequest
	5,  // 4: shortener.URLService.PingDB:input_type -> shortener.PingDBRequest
	7,  // 5: shortener.URLService.GetUserURLs:input_type -> shortener.UserURLsRequest
	9,  // 6: shortener.URLService.DeleteUserURLs:input_type -> shortener.DeleteURLsRequest
	11, // 7: shortener.URLService.Stats:input_type -> shortener.StatsRequest
	13, // 8: shortener.URLService.Root:input_type -> shortener.RootRequest
	4,  // 9: shortener.URLService.Shorten:output_type -> shortener.ShortenResponse
	6,  // 10: shortener.URLService.PingDB:output_type -> shortener.PingDBResponse
	8,  // 11: shortener.URLService.GetUserURLs:output_type -> shortener.UserURLsResponse
	10, // 12: shortener.URLService.DeleteUserURLs:output_type -> shortener.DeleteURLsResponse
	12, // 13: shortener.URLService.Stats:output_type -> shortener.StatsResponse
	14, // 14: shortener.URLService.Root:output_type -> shortener.RootResponse
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_shortener_proto_init() }
func file_shortener_proto_init() {
	if File_shortener_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shortener_proto_goTypes,
		DependencyIndexes: file_shortener_proto_depIdxs,
		MessageInfos:      file_shortener_proto_msgTypes,
	}.Build()
	File_shortener_proto = out.File
	file_shortener_proto_rawDesc = nil
	file_shortener_proto_goTypes = nil
	file_shortener_proto_depIdxs = nil
}
