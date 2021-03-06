// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/datalabeling/v1beta1/instruction.proto

package datalabeling

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Instruction of how to perform the labeling task for human operators.
// Currently two types of instruction are supported - CSV file and PDF.
// One of the two types instruction must be provided.
// CSV file is only supported for image classification task. Instructions for
// other task should be provided as PDF.
// For image classification, CSV and PDF can be provided at the same time.
type Instruction struct {
	// Output only. Instruction resource name, format:
	// projects/{project_id}/instructions/{instruction_id}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The display name of the instruction. Maximum of 64 characters.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Optional. User-provided description of the instruction.
	// The description can be up to 10000 characters long.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Output only. Creation time of instruction.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Last update time of instruction.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Required. The data type of this instruction.
	DataType DataType `protobuf:"varint,6,opt,name=data_type,json=dataType,proto3,enum=google.cloud.datalabeling.v1beta1.DataType" json:"data_type,omitempty"`
	// One of CSV or PDF instruction is required.
	// Instruction from a CSV file, such as for classification task.
	// The CSV file should have exact two columns, in the following format:
	//
	// * The first column is labeled data, such as an image reference, text.
	// * The second column is comma separated labels associated with data.
	CsvInstruction *CsvInstruction `protobuf:"bytes,7,opt,name=csv_instruction,json=csvInstruction,proto3" json:"csv_instruction,omitempty"`
	// One of CSV or PDF instruction is required.
	// Instruction from a PDF document. The PDF should be in a Cloud Storage
	// bucket.
	PdfInstruction *PdfInstruction `protobuf:"bytes,9,opt,name=pdf_instruction,json=pdfInstruction,proto3" json:"pdf_instruction,omitempty"`
	// Output only. The names of any related resources that are blocking changes
	// to the instruction.
	BlockingResources    []string `protobuf:"bytes,10,rep,name=blocking_resources,json=blockingResources,proto3" json:"blocking_resources,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Instruction) Reset()         { *m = Instruction{} }
func (m *Instruction) String() string { return proto.CompactTextString(m) }
func (*Instruction) ProtoMessage()    {}
func (*Instruction) Descriptor() ([]byte, []int) {
	return fileDescriptor_683081924dd940d7, []int{0}
}

func (m *Instruction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Instruction.Unmarshal(m, b)
}
func (m *Instruction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Instruction.Marshal(b, m, deterministic)
}
func (m *Instruction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Instruction.Merge(m, src)
}
func (m *Instruction) XXX_Size() int {
	return xxx_messageInfo_Instruction.Size(m)
}
func (m *Instruction) XXX_DiscardUnknown() {
	xxx_messageInfo_Instruction.DiscardUnknown(m)
}

var xxx_messageInfo_Instruction proto.InternalMessageInfo

func (m *Instruction) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Instruction) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Instruction) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Instruction) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Instruction) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *Instruction) GetDataType() DataType {
	if m != nil {
		return m.DataType
	}
	return DataType_DATA_TYPE_UNSPECIFIED
}

func (m *Instruction) GetCsvInstruction() *CsvInstruction {
	if m != nil {
		return m.CsvInstruction
	}
	return nil
}

func (m *Instruction) GetPdfInstruction() *PdfInstruction {
	if m != nil {
		return m.PdfInstruction
	}
	return nil
}

func (m *Instruction) GetBlockingResources() []string {
	if m != nil {
		return m.BlockingResources
	}
	return nil
}

// Instruction from a CSV file.
type CsvInstruction struct {
	// CSV file for the instruction. Only gcs path is allowed.
	GcsFileUri           string   `protobuf:"bytes,1,opt,name=gcs_file_uri,json=gcsFileUri,proto3" json:"gcs_file_uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CsvInstruction) Reset()         { *m = CsvInstruction{} }
func (m *CsvInstruction) String() string { return proto.CompactTextString(m) }
func (*CsvInstruction) ProtoMessage()    {}
func (*CsvInstruction) Descriptor() ([]byte, []int) {
	return fileDescriptor_683081924dd940d7, []int{1}
}

func (m *CsvInstruction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CsvInstruction.Unmarshal(m, b)
}
func (m *CsvInstruction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CsvInstruction.Marshal(b, m, deterministic)
}
func (m *CsvInstruction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CsvInstruction.Merge(m, src)
}
func (m *CsvInstruction) XXX_Size() int {
	return xxx_messageInfo_CsvInstruction.Size(m)
}
func (m *CsvInstruction) XXX_DiscardUnknown() {
	xxx_messageInfo_CsvInstruction.DiscardUnknown(m)
}

var xxx_messageInfo_CsvInstruction proto.InternalMessageInfo

func (m *CsvInstruction) GetGcsFileUri() string {
	if m != nil {
		return m.GcsFileUri
	}
	return ""
}

// Instruction from a PDF file.
type PdfInstruction struct {
	// PDF file for the instruction. Only gcs path is allowed.
	GcsFileUri           string   `protobuf:"bytes,1,opt,name=gcs_file_uri,json=gcsFileUri,proto3" json:"gcs_file_uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PdfInstruction) Reset()         { *m = PdfInstruction{} }
func (m *PdfInstruction) String() string { return proto.CompactTextString(m) }
func (*PdfInstruction) ProtoMessage()    {}
func (*PdfInstruction) Descriptor() ([]byte, []int) {
	return fileDescriptor_683081924dd940d7, []int{2}
}

func (m *PdfInstruction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PdfInstruction.Unmarshal(m, b)
}
func (m *PdfInstruction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PdfInstruction.Marshal(b, m, deterministic)
}
func (m *PdfInstruction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PdfInstruction.Merge(m, src)
}
func (m *PdfInstruction) XXX_Size() int {
	return xxx_messageInfo_PdfInstruction.Size(m)
}
func (m *PdfInstruction) XXX_DiscardUnknown() {
	xxx_messageInfo_PdfInstruction.DiscardUnknown(m)
}

var xxx_messageInfo_PdfInstruction proto.InternalMessageInfo

func (m *PdfInstruction) GetGcsFileUri() string {
	if m != nil {
		return m.GcsFileUri
	}
	return ""
}

func init() {
	proto.RegisterType((*Instruction)(nil), "google.cloud.datalabeling.v1beta1.Instruction")
	proto.RegisterType((*CsvInstruction)(nil), "google.cloud.datalabeling.v1beta1.CsvInstruction")
	proto.RegisterType((*PdfInstruction)(nil), "google.cloud.datalabeling.v1beta1.PdfInstruction")
}

func init() {
	proto.RegisterFile("google/cloud/datalabeling/v1beta1/instruction.proto", fileDescriptor_683081924dd940d7)
}

var fileDescriptor_683081924dd940d7 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x8b, 0x13, 0x31,
	0x14, 0xc6, 0x19, 0xb7, 0xae, 0x36, 0xb3, 0x54, 0xcc, 0x69, 0x28, 0x82, 0xb3, 0x0b, 0x42, 0x41,
	0xcc, 0xd0, 0xee, 0x71, 0x6f, 0x2a, 0xa2, 0x07, 0x65, 0x19, 0xd6, 0xcb, 0x5e, 0x86, 0x4c, 0xe6,
	0x35, 0x04, 0xd3, 0x24, 0x24, 0x99, 0x62, 0xff, 0x0c, 0xff, 0x63, 0x49, 0x26, 0x75, 0x3b, 0x17,
	0xa7, 0xb7, 0xe4, 0xbd, 0xef, 0xf7, 0xf5, 0xeb, 0x7b, 0x13, 0x74, 0xcb, 0xb5, 0xe6, 0x12, 0x2a,
	0x26, 0x75, 0xdf, 0x55, 0x1d, 0xf5, 0x54, 0xd2, 0x16, 0xa4, 0x50, 0xbc, 0xda, 0xaf, 0x5b, 0xf0,
	0x74, 0x5d, 0x09, 0xe5, 0xbc, 0xed, 0x99, 0x17, 0x5a, 0x11, 0x63, 0xb5, 0xd7, 0xf8, 0x7a, 0x80,
	0x48, 0x84, 0xc8, 0x29, 0x44, 0x12, 0xb4, 0x7c, 0x93, 0x7c, 0xa9, 0x11, 0x15, 0x55, 0x4a, 0x7b,
	0x1a, 0x78, 0x37, 0x18, 0x2c, 0xab, 0xe9, 0x5f, 0x0d, 0x45, 0x07, 0x3e, 0x01, 0x6f, 0x13, 0x10,
	0x6f, 0x6d, 0xbf, 0xad, 0xbc, 0xd8, 0x81, 0xf3, 0x74, 0x67, 0x06, 0xc1, 0xcd, 0x9f, 0x19, 0xca,
	0xbf, 0x3d, 0x05, 0xc5, 0x18, 0xcd, 0x14, 0xdd, 0x41, 0x91, 0x95, 0xd9, 0x6a, 0x5e, 0xc7, 0x33,
	0xbe, 0x46, 0x57, 0x9d, 0x70, 0x46, 0xd2, 0x43, 0x13, 0x7b, 0xcf, 0x62, 0x2f, 0x4f, 0xb5, 0x1f,
	0x41, 0x52, 0xa2, 0xbc, 0x03, 0xc7, 0xac, 0x30, 0xc1, 0xa5, 0xb8, 0x48, 0x8a, 0xa7, 0x12, 0xbe,
	0x43, 0x39, 0xb3, 0x40, 0x3d, 0x34, 0x21, 0x42, 0x31, 0x2b, 0xb3, 0x55, 0xbe, 0x59, 0x92, 0x34,
	0x91, 0x63, 0x3e, 0xf2, 0x70, 0xcc, 0x57, 0xa3, 0x41, 0x1e, 0x0a, 0x01, 0xee, 0x4d, 0xf7, 0x0f,
	0x7e, 0x3e, 0x0d, 0x0f, 0xf2, 0x08, 0x7f, 0x45, 0xf3, 0x30, 0x94, 0xc6, 0x1f, 0x0c, 0x14, 0x97,
	0x65, 0xb6, 0x5a, 0x6c, 0xde, 0x93, 0xc9, 0x4d, 0x90, 0xcf, 0xd4, 0xd3, 0x87, 0x83, 0x81, 0xfa,
	0x65, 0x97, 0x4e, 0xf8, 0x11, 0xbd, 0x62, 0x6e, 0xdf, 0x9c, 0x2c, 0xb6, 0x78, 0x11, 0xa3, 0xac,
	0xcf, 0xf0, 0xfb, 0xe4, 0xf6, 0x27, 0x83, 0xae, 0x17, 0x6c, 0x74, 0x0f, 0xde, 0xa6, 0xdb, 0x8e,
	0xbc, 0xe7, 0x67, 0x7b, 0xdf, 0x77, 0xdb, 0x91, 0xb7, 0x19, 0xdd, 0xf1, 0x07, 0x84, 0x5b, 0xa9,
	0xd9, 0x2f, 0xa1, 0x78, 0x63, 0xc1, 0xe9, 0xde, 0x32, 0x70, 0x05, 0x2a, 0x2f, 0x56, 0xf3, 0xfa,
	0xf5, 0xb1, 0x53, 0x1f, 0x1b, 0x37, 0x1b, 0xb4, 0x18, 0x87, 0xc5, 0x25, 0xba, 0xe2, 0xcc, 0x35,
	0x5b, 0x21, 0xa1, 0xe9, 0xad, 0x48, 0x5f, 0x07, 0xe2, 0xcc, 0x7d, 0x11, 0x12, 0x7e, 0x5a, 0x11,
	0x98, 0x71, 0x88, 0x69, 0xe6, 0xe3, 0x6f, 0xf4, 0x8e, 0xe9, 0xdd, 0xf4, 0xdf, 0xbb, 0xcf, 0x1e,
	0xbf, 0x27, 0x11, 0xd7, 0x92, 0x2a, 0x4e, 0xb4, 0xe5, 0x15, 0x07, 0x15, 0x17, 0x9f, 0xde, 0x04,
	0x35, 0xc2, 0xfd, 0xe7, 0x5d, 0xdc, 0x9d, 0x16, 0xdb, 0xcb, 0x48, 0xde, 0xfe, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0xb3, 0x4c, 0xf6, 0x74, 0xc6, 0x03, 0x00, 0x00,
}
