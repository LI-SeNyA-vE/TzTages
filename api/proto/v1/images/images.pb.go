// Для запуска
//protoc \
//    --proto_path=. \
//    --go_out=. \
//    --go-grpc_out=. \
//    api/proto/v1/images/images.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: api/proto/v1/images/images.proto

package images

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

// Метаданные изображения (без содержимого файла)
type ImageMetadata struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ImageId       string                 `protobuf:"bytes,1,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`        // Уникальный идентификатор изображения
	FileName      string                 `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`     // Имя файла
	SizeBytes     int64                  `protobuf:"varint,3,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"` // Размер изображения в байтах
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImageMetadata) Reset() {
	*x = ImageMetadata{}
	mi := &file_api_proto_v1_images_images_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageMetadata) ProtoMessage() {}

func (x *ImageMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_images_images_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageMetadata.ProtoReflect.Descriptor instead.
func (*ImageMetadata) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_images_images_proto_rawDescGZIP(), []int{0}
}

func (x *ImageMetadata) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

func (x *ImageMetadata) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *ImageMetadata) GetSizeBytes() int64 {
	if x != nil {
		return x.SizeBytes
	}
	return 0
}

// Запрос на загрузку изображения
type UploadImageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileName      string                 `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"` // Имя файла, заданное клиентом
	Data          []byte                 `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`                         // Бинарные данные изображения
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadImageRequest) Reset() {
	*x = UploadImageRequest{}
	mi := &file_api_proto_v1_images_images_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageRequest) ProtoMessage() {}

func (x *UploadImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_images_images_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageRequest.ProtoReflect.Descriptor instead.
func (*UploadImageRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_images_images_proto_rawDescGZIP(), []int{1}
}

func (x *UploadImageRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *UploadImageRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// Ответ после загрузки изображения
type UploadImageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ImageMetadata *ImageMetadata         `protobuf:"bytes,1,opt,name=image_metadata,json=imageMetadata,proto3" json:"image_metadata,omitempty"` // Уникальный идентификатор изображения
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadImageResponse) Reset() {
	*x = UploadImageResponse{}
	mi := &file_api_proto_v1_images_images_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageResponse) ProtoMessage() {}

func (x *UploadImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_images_images_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageResponse.ProtoReflect.Descriptor instead.
func (*UploadImageResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_images_images_proto_rawDescGZIP(), []int{2}
}

func (x *UploadImageResponse) GetImageMetadata() *ImageMetadata {
	if x != nil {
		return x.ImageMetadata
	}
	return nil
}

// Запрос на скачивание изображения
type DownloadImageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ImageId       string                 `protobuf:"bytes,1,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"` // Идентификатор изображения, которое нужно скачать
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadImageRequest) Reset() {
	*x = DownloadImageRequest{}
	mi := &file_api_proto_v1_images_images_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadImageRequest) ProtoMessage() {}

func (x *DownloadImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_images_images_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadImageRequest.ProtoReflect.Descriptor instead.
func (*DownloadImageRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_images_images_proto_rawDescGZIP(), []int{3}
}

func (x *DownloadImageRequest) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

// Ответ при скачивании изображения
type DownloadImageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ImageMetadata *ImageMetadata         `protobuf:"bytes,1,opt,name=image_metadata,json=imageMetadata,proto3" json:"image_metadata,omitempty"` // Имя скачиваемого файла
	Data          []byte                 `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`                                        // Бинарные данные изображения
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadImageResponse) Reset() {
	*x = DownloadImageResponse{}
	mi := &file_api_proto_v1_images_images_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadImageResponse) ProtoMessage() {}

func (x *DownloadImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_images_images_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadImageResponse.ProtoReflect.Descriptor instead.
func (*DownloadImageResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_images_images_proto_rawDescGZIP(), []int{4}
}

func (x *DownloadImageResponse) GetImageMetadata() *ImageMetadata {
	if x != nil {
		return x.ImageMetadata
	}
	return nil
}

func (x *DownloadImageResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// Ответ, содержащий список всех доступных изображений
type ListImagesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ImageMetadata []*ImageMetadata       `protobuf:"bytes,1,rep,name=image_metadata,json=imageMetadata,proto3" json:"image_metadata,omitempty"` // Список метаданных изображений
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListImagesResponse) Reset() {
	*x = ListImagesResponse{}
	mi := &file_api_proto_v1_images_images_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListImagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListImagesResponse) ProtoMessage() {}

func (x *ListImagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_images_images_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListImagesResponse.ProtoReflect.Descriptor instead.
func (*ListImagesResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_images_images_proto_rawDescGZIP(), []int{5}
}

func (x *ListImagesResponse) GetImageMetadata() []*ImageMetadata {
	if x != nil {
		return x.ImageMetadata
	}
	return nil
}

// Пустое сообщение, используется если запрос не требует параметров
type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_api_proto_v1_images_images_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_images_images_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_images_images_proto_rawDescGZIP(), []int{6}
}

var File_api_proto_v1_images_images_proto protoreflect.FileDescriptor

var file_api_proto_v1_images_images_proto_rawDesc = string([]byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x66, 0x0a,
	0x0d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x19,
	0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x69, 0x7a, 0x65,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x45, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x56, 0x0a, 0x13,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0d, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x31, 0x0a, 0x14, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x6c, 0x0a, 0x15, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3f, 0x0a, 0x0e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x0d, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x55, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0e, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0d, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xf5, 0x01, 0x0a, 0x0c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x10, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x0d, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x15, 0x5a,
	0x13, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_proto_v1_images_images_proto_rawDescOnce sync.Once
	file_api_proto_v1_images_images_proto_rawDescData []byte
)

func file_api_proto_v1_images_images_proto_rawDescGZIP() []byte {
	file_api_proto_v1_images_images_proto_rawDescOnce.Do(func() {
		file_api_proto_v1_images_images_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_proto_v1_images_images_proto_rawDesc), len(file_api_proto_v1_images_images_proto_rawDesc)))
	})
	return file_api_proto_v1_images_images_proto_rawDescData
}

var file_api_proto_v1_images_images_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_proto_v1_images_images_proto_goTypes = []any{
	(*ImageMetadata)(nil),         // 0: images.v1.ImageMetadata
	(*UploadImageRequest)(nil),    // 1: images.v1.UploadImageRequest
	(*UploadImageResponse)(nil),   // 2: images.v1.UploadImageResponse
	(*DownloadImageRequest)(nil),  // 3: images.v1.DownloadImageRequest
	(*DownloadImageResponse)(nil), // 4: images.v1.DownloadImageResponse
	(*ListImagesResponse)(nil),    // 5: images.v1.ListImagesResponse
	(*Empty)(nil),                 // 6: images.v1.Empty
}
var file_api_proto_v1_images_images_proto_depIdxs = []int32{
	0, // 0: images.v1.UploadImageResponse.image_metadata:type_name -> images.v1.ImageMetadata
	0, // 1: images.v1.DownloadImageResponse.image_metadata:type_name -> images.v1.ImageMetadata
	0, // 2: images.v1.ListImagesResponse.image_metadata:type_name -> images.v1.ImageMetadata
	1, // 3: images.v1.ImageService.UploadImage:input_type -> images.v1.UploadImageRequest
	6, // 4: images.v1.ImageService.ListImages:input_type -> images.v1.Empty
	3, // 5: images.v1.ImageService.DownloadImage:input_type -> images.v1.DownloadImageRequest
	2, // 6: images.v1.ImageService.UploadImage:output_type -> images.v1.UploadImageResponse
	5, // 7: images.v1.ImageService.ListImages:output_type -> images.v1.ListImagesResponse
	4, // 8: images.v1.ImageService.DownloadImage:output_type -> images.v1.DownloadImageResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_proto_v1_images_images_proto_init() }
func file_api_proto_v1_images_images_proto_init() {
	if File_api_proto_v1_images_images_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_proto_v1_images_images_proto_rawDesc), len(file_api_proto_v1_images_images_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_v1_images_images_proto_goTypes,
		DependencyIndexes: file_api_proto_v1_images_images_proto_depIdxs,
		MessageInfos:      file_api_proto_v1_images_images_proto_msgTypes,
	}.Build()
	File_api_proto_v1_images_images_proto = out.File
	file_api_proto_v1_images_images_proto_goTypes = nil
	file_api_proto_v1_images_images_proto_depIdxs = nil
}
