// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: api/song/song.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SongRequestParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Genre string `protobuf:"bytes,1,opt,name=genre,proto3" json:"genre,omitempty"`
}

func (x *SongRequestParams) Reset() {
	*x = SongRequestParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_song_song_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SongRequestParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SongRequestParams) ProtoMessage() {}

func (x *SongRequestParams) ProtoReflect() protoreflect.Message {
	mi := &file_api_song_song_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SongRequestParams.ProtoReflect.Descriptor instead.
func (*SongRequestParams) Descriptor() ([]byte, []int) {
	return file_api_song_song_proto_rawDescGZIP(), []int{0}
}

func (x *SongRequestParams) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

type GetSongsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool    `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*Song `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetSongsResponse) Reset() {
	*x = GetSongsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_song_song_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSongsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSongsResponse) ProtoMessage() {}

func (x *GetSongsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_song_song_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSongsResponse.ProtoReflect.Descriptor instead.
func (*GetSongsResponse) Descriptor() ([]byte, []int) {
	return file_api_song_song_proto_rawDescGZIP(), []int{1}
}

func (x *GetSongsResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetSongsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetSongsResponse) GetData() []*Song {
	if x != nil {
		return x.Data
	}
	return nil
}

type Song struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title           string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	ArtistsNames    string   `protobuf:"bytes,3,opt,name=artists_names,json=artistsNames,proto3" json:"artists_names,omitempty"`
	Thumbnail       string   `protobuf:"bytes,4,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	ThumbnailMedium string   `protobuf:"bytes,5,opt,name=thumbnail_medium,json=thumbnailMedium,proto3" json:"thumbnail_medium,omitempty"`
	Lyric           string   `protobuf:"bytes,6,opt,name=lyric,proto3" json:"lyric,omitempty"`
	Duration        int32    `protobuf:"varint,7,opt,name=duration,proto3" json:"duration,omitempty"`
	Genres          []*Genre `protobuf:"bytes,8,rep,name=genres,proto3" json:"genres,omitempty"`
}

func (x *Song) Reset() {
	*x = Song{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_song_song_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Song) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Song) ProtoMessage() {}

func (x *Song) ProtoReflect() protoreflect.Message {
	mi := &file_api_song_song_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Song.ProtoReflect.Descriptor instead.
func (*Song) Descriptor() ([]byte, []int) {
	return file_api_song_song_proto_rawDescGZIP(), []int{2}
}

func (x *Song) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Song) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Song) GetArtistsNames() string {
	if x != nil {
		return x.ArtistsNames
	}
	return ""
}

func (x *Song) GetThumbnail() string {
	if x != nil {
		return x.Thumbnail
	}
	return ""
}

func (x *Song) GetThumbnailMedium() string {
	if x != nil {
		return x.ThumbnailMedium
	}
	return ""
}

func (x *Song) GetLyric() string {
	if x != nil {
		return x.Lyric
	}
	return ""
}

func (x *Song) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Song) GetGenres() []*Genre {
	if x != nil {
		return x.Genres
	}
	return nil
}

type Streaming struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Q128 string `protobuf:"bytes,1,opt,name=q128,proto3" json:"q128,omitempty"`
	Q320 string `protobuf:"bytes,2,opt,name=q320,proto3" json:"q320,omitempty"`
}

func (x *Streaming) Reset() {
	*x = Streaming{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_song_song_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Streaming) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Streaming) ProtoMessage() {}

func (x *Streaming) ProtoReflect() protoreflect.Message {
	mi := &file_api_song_song_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Streaming.ProtoReflect.Descriptor instead.
func (*Streaming) Descriptor() ([]byte, []int) {
	return file_api_song_song_proto_rawDescGZIP(), []int{3}
}

func (x *Streaming) GetQ128() string {
	if x != nil {
		return x.Q128
	}
	return ""
}

func (x *Streaming) GetQ320() string {
	if x != nil {
		return x.Q320
	}
	return ""
}

type Genre struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Alias string `protobuf:"bytes,3,opt,name=alias,proto3" json:"alias,omitempty"`
	Link  string `protobuf:"bytes,4,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *Genre) Reset() {
	*x = Genre{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_song_song_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Genre) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Genre) ProtoMessage() {}

func (x *Genre) ProtoReflect() protoreflect.Message {
	mi := &file_api_song_song_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Genre.ProtoReflect.Descriptor instead.
func (*Genre) Descriptor() ([]byte, []int) {
	return file_api_song_song_proto_rawDescGZIP(), []int{4}
}

func (x *Genre) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Genre) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Genre) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *Genre) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

var File_api_song_song_proto protoreflect.FileDescriptor

var file_api_song_song_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x6f, 0x6e, 0x67, 0x2f, 0x73, 0x6f, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x6f, 0x6e, 0x67, 0x22, 0x29, 0x0a, 0x11, 0x53,
	0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x22, 0x66, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x6e,
	0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x73,
	0x6f, 0x6e, 0x67, 0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xf1,
	0x01, 0x0a, 0x04, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c,
	0x12, 0x29, 0x0a, 0x10, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x6d, 0x65,
	0x64, 0x69, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x68, 0x75, 0x6d,
	0x62, 0x6e, 0x61, 0x69, 0x6c, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x79, 0x72, 0x69, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x79, 0x72, 0x69,
	0x63, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a,
	0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x73, 0x6f, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x72,
	0x65, 0x73, 0x22, 0x33, 0x0a, 0x09, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x71, 0x31, 0x32, 0x38, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x71,
	0x31, 0x32, 0x38, 0x12, 0x12, 0x0a, 0x04, 0x71, 0x33, 0x32, 0x30, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x71, 0x33, 0x32, 0x30, 0x22, 0x55, 0x0a, 0x05, 0x47, 0x65, 0x6e, 0x72, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69,
	0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x32, 0x4b,
	0x0a, 0x0c, 0x53, 0x6f, 0x6e, 0x67, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x6e, 0x67, 0x73, 0x12, 0x17, 0x2e, 0x73, 0x6f, 0x6e,
	0x67, 0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x1a, 0x16, 0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6f,
	0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x61, 0x69, 0x64, 0x7a,
	0x61, 0x69, 0x32, 0x38, 0x35, 0x2f, 0x64, 0x7a, 0x61, 0x69, 0x2d, 0x6d, 0x70, 0x33, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_song_song_proto_rawDescOnce sync.Once
	file_api_song_song_proto_rawDescData = file_api_song_song_proto_rawDesc
)

func file_api_song_song_proto_rawDescGZIP() []byte {
	file_api_song_song_proto_rawDescOnce.Do(func() {
		file_api_song_song_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_song_song_proto_rawDescData)
	})
	return file_api_song_song_proto_rawDescData
}

var file_api_song_song_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_song_song_proto_goTypes = []interface{}{
	(*SongRequestParams)(nil), // 0: song.SongRequestParams
	(*GetSongsResponse)(nil),  // 1: song.GetSongsResponse
	(*Song)(nil),              // 2: song.Song
	(*Streaming)(nil),         // 3: song.Streaming
	(*Genre)(nil),             // 4: song.Genre
}
var file_api_song_song_proto_depIdxs = []int32{
	2, // 0: song.GetSongsResponse.data:type_name -> song.Song
	4, // 1: song.Song.genres:type_name -> song.Genre
	0, // 2: song.SongsService.GetSongs:input_type -> song.SongRequestParams
	1, // 3: song.SongsService.GetSongs:output_type -> song.GetSongsResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_song_song_proto_init() }
func file_api_song_song_proto_init() {
	if File_api_song_song_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_song_song_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SongRequestParams); i {
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
		file_api_song_song_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSongsResponse); i {
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
		file_api_song_song_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Song); i {
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
		file_api_song_song_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Streaming); i {
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
		file_api_song_song_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Genre); i {
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
			RawDescriptor: file_api_song_song_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_song_song_proto_goTypes,
		DependencyIndexes: file_api_song_song_proto_depIdxs,
		MessageInfos:      file_api_song_song_proto_msgTypes,
	}.Build()
	File_api_song_song_proto = out.File
	file_api_song_song_proto_rawDesc = nil
	file_api_song_song_proto_goTypes = nil
	file_api_song_song_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SongsServiceClient is the client API for SongsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SongsServiceClient interface {
	GetSongs(ctx context.Context, in *SongRequestParams, opts ...grpc.CallOption) (*GetSongsResponse, error)
}

type songsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSongsServiceClient(cc grpc.ClientConnInterface) SongsServiceClient {
	return &songsServiceClient{cc}
}

func (c *songsServiceClient) GetSongs(ctx context.Context, in *SongRequestParams, opts ...grpc.CallOption) (*GetSongsResponse, error) {
	out := new(GetSongsResponse)
	err := c.cc.Invoke(ctx, "/song.SongsService/GetSongs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SongsServiceServer is the server API for SongsService service.
type SongsServiceServer interface {
	GetSongs(context.Context, *SongRequestParams) (*GetSongsResponse, error)
}

// UnimplementedSongsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSongsServiceServer struct {
}

func (*UnimplementedSongsServiceServer) GetSongs(context.Context, *SongRequestParams) (*GetSongsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSongs not implemented")
}

func RegisterSongsServiceServer(s *grpc.Server, srv SongsServiceServer) {
	s.RegisterService(&_SongsService_serviceDesc, srv)
}

func _SongsService_GetSongs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SongRequestParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SongsServiceServer).GetSongs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/song.SongsService/GetSongs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SongsServiceServer).GetSongs(ctx, req.(*SongRequestParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _SongsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "song.SongsService",
	HandlerType: (*SongsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSongs",
			Handler:    _SongsService_GetSongs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/song/song.proto",
}
