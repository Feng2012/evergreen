// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/keyword_plan_keyword_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v0/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible errors from applying a keyword plan keyword.
type KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError int32

const (
	// Enum unspecified.
	KeywordPlanKeywordErrorEnum_UNSPECIFIED KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError = 0
	// The received error code is not known in this version.
	KeywordPlanKeywordErrorEnum_UNKNOWN KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError = 1
	// A keyword or negative keyword has invalid match type.
	KeywordPlanKeywordErrorEnum_INVALID_KEYWORD_MATCH_TYPE KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError = 2
	// A keyword or negative keyword with same text and match type already
	// exists.
	KeywordPlanKeywordErrorEnum_DUPLICATE_KEYWORD KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError = 3
	// Keyword or negative keyword text exceeds the allowed limit.
	KeywordPlanKeywordErrorEnum_KEYWORD_TEXT_TOO_LONG KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError = 4
	// Keyword or negative keyword text has invalid characters or symbols.
	KeywordPlanKeywordErrorEnum_KEYWORD_HAS_INVALID_CHARS KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError = 5
	// Keyword or negative keyword text has too many words.
	KeywordPlanKeywordErrorEnum_KEYWORD_HAS_TOO_MANY_WORDS KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError = 6
	// Keyword or negative keyword has invalid text.
	KeywordPlanKeywordErrorEnum_INVALID_KEYWORD_TEXT KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError = 7
)

var KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "INVALID_KEYWORD_MATCH_TYPE",
	3: "DUPLICATE_KEYWORD",
	4: "KEYWORD_TEXT_TOO_LONG",
	5: "KEYWORD_HAS_INVALID_CHARS",
	6: "KEYWORD_HAS_TOO_MANY_WORDS",
	7: "INVALID_KEYWORD_TEXT",
}
var KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError_value = map[string]int32{
	"UNSPECIFIED":                0,
	"UNKNOWN":                    1,
	"INVALID_KEYWORD_MATCH_TYPE": 2,
	"DUPLICATE_KEYWORD":          3,
	"KEYWORD_TEXT_TOO_LONG":      4,
	"KEYWORD_HAS_INVALID_CHARS":  5,
	"KEYWORD_HAS_TOO_MANY_WORDS": 6,
	"INVALID_KEYWORD_TEXT":       7,
}

func (x KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError) String() string {
	return proto.EnumName(KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError_name, int32(x))
}
func (KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_keyword_error_6c1c9e3b2c275f9e, []int{0, 0}
}

// Container for enum describing possible errors from applying a keyword or a
// negative keyword from a keyword plan.
type KeywordPlanKeywordErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeywordPlanKeywordErrorEnum) Reset()         { *m = KeywordPlanKeywordErrorEnum{} }
func (m *KeywordPlanKeywordErrorEnum) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanKeywordErrorEnum) ProtoMessage()    {}
func (*KeywordPlanKeywordErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_keyword_error_6c1c9e3b2c275f9e, []int{0}
}
func (m *KeywordPlanKeywordErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanKeywordErrorEnum.Unmarshal(m, b)
}
func (m *KeywordPlanKeywordErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanKeywordErrorEnum.Marshal(b, m, deterministic)
}
func (dst *KeywordPlanKeywordErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanKeywordErrorEnum.Merge(dst, src)
}
func (m *KeywordPlanKeywordErrorEnum) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanKeywordErrorEnum.Size(m)
}
func (m *KeywordPlanKeywordErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanKeywordErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanKeywordErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*KeywordPlanKeywordErrorEnum)(nil), "google.ads.googleads.v0.errors.KeywordPlanKeywordErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError", KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError_name, KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/keyword_plan_keyword_error.proto", fileDescriptor_keyword_plan_keyword_error_6c1c9e3b2c275f9e)
}

var fileDescriptor_keyword_plan_keyword_error_6c1c9e3b2c275f9e = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xcb, 0xce, 0x93, 0x40,
	0x14, 0x16, 0x7e, 0x6d, 0x93, 0xe9, 0xc2, 0x71, 0x62, 0xa3, 0xf5, 0xd2, 0x05, 0x0f, 0x30, 0x90,
	0xb8, 0x1b, 0x17, 0x66, 0x0a, 0x63, 0x4b, 0xda, 0x02, 0x29, 0x94, 0x5a, 0x43, 0x32, 0x41, 0x21,
	0xc4, 0x48, 0x99, 0x86, 0xd1, 0x1a, 0xb7, 0x3e, 0x8a, 0x4b, 0x1f, 0xc5, 0xc7, 0x70, 0xe9, 0x0b,
	0xb8, 0x35, 0xc3, 0x94, 0xc6, 0x98, 0xf4, 0x5f, 0xcd, 0x37, 0xe7, 0xbb, 0x9c, 0x93, 0x73, 0xc0,
	0xab, 0x4a, 0x88, 0xaa, 0x2e, 0xed, 0xbc, 0x90, 0xb6, 0x86, 0x0a, 0x9d, 0x1c, 0xbb, 0x6c, 0x5b,
	0xd1, 0x4a, 0xfb, 0x63, 0xf9, 0xf5, 0x8b, 0x68, 0x0b, 0x7e, 0xac, 0xf3, 0x86, 0xf7, 0x9f, 0x8e,
	0xc3, 0xc7, 0x56, 0x7c, 0x12, 0x68, 0xaa, 0x5d, 0x38, 0x2f, 0x24, 0xbe, 0x04, 0xe0, 0x93, 0x83,
	0x75, 0x80, 0xf5, 0xcd, 0x04, 0x4f, 0x97, 0xda, 0x17, 0xd5, 0x79, 0x73, 0x86, 0x4c, 0x91, 0xac,
	0xf9, 0x7c, 0xb0, 0x7e, 0x19, 0xe0, 0xd1, 0x15, 0x1e, 0xdd, 0x07, 0xa3, 0x6d, 0x10, 0x47, 0xcc,
	0xf5, 0x5f, 0xfb, 0xcc, 0x83, 0x77, 0xd0, 0x08, 0x0c, 0xb7, 0xc1, 0x32, 0x08, 0x77, 0x01, 0x34,
	0xd0, 0x14, 0x3c, 0xf1, 0x83, 0x94, 0xae, 0x7c, 0x8f, 0x2f, 0xd9, 0x7e, 0x17, 0x6e, 0x3c, 0xbe,
	0xa6, 0x89, 0xbb, 0xe0, 0xc9, 0x3e, 0x62, 0xd0, 0x44, 0x63, 0xf0, 0xc0, 0xdb, 0x46, 0x2b, 0xdf,
	0xa5, 0x09, 0xeb, 0x15, 0xf0, 0x06, 0x4d, 0xc0, 0xb8, 0x97, 0x27, 0xec, 0x4d, 0xc2, 0x93, 0x30,
	0xe4, 0xab, 0x30, 0x98, 0xc3, 0xbb, 0xe8, 0x39, 0x98, 0xf4, 0xd4, 0x82, 0xc6, 0xbc, 0x4f, 0x77,
	0x17, 0x74, 0x13, 0xc3, 0x7b, 0xaa, 0xe1, 0xbf, 0xb4, 0x32, 0xae, 0x69, 0xb0, 0xe7, 0xaa, 0x12,
	0xc3, 0x01, 0x7a, 0x0c, 0x1e, 0xfe, 0x3f, 0x90, 0xea, 0x00, 0x87, 0xb3, 0x3f, 0x06, 0xb0, 0xde,
	0x8b, 0x03, 0xbe, 0x7d, 0x57, 0xb3, 0x67, 0x57, 0x16, 0x11, 0xa9, 0x4d, 0x47, 0xc6, 0x5b, 0xef,
	0xec, 0xaf, 0x44, 0x9d, 0x37, 0x15, 0x16, 0x6d, 0x65, 0x57, 0x65, 0xd3, 0xdd, 0xa1, 0x3f, 0xde,
	0xf1, 0x83, 0xbc, 0x76, 0xcb, 0x97, 0xfa, 0xf9, 0x6e, 0xde, 0xcc, 0x29, 0xfd, 0x61, 0x4e, 0xe7,
	0x3a, 0x8c, 0x16, 0x12, 0x6b, 0xa8, 0x50, 0xea, 0xe0, 0xae, 0xa5, 0xfc, 0xd9, 0x0b, 0x32, 0x5a,
	0xc8, 0xec, 0x22, 0xc8, 0x52, 0x27, 0xd3, 0x82, 0xdf, 0xa6, 0xa5, 0xab, 0x84, 0xd0, 0x42, 0x12,
	0x72, 0x91, 0x10, 0x92, 0x3a, 0x84, 0x68, 0xd1, 0xbb, 0x41, 0x37, 0xdd, 0x8b, 0xbf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x76, 0x56, 0xec, 0x81, 0x68, 0x02, 0x00, 0x00,
}
