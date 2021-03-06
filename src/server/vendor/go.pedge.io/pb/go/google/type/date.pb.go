// Code generated by protoc-gen-go.
// source: google/type/date.proto
// DO NOT EDIT!

package google_type

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Represents a whole calendar date, e.g. date of birth. The time of day and
// time zone are either specified elsewhere or are not significant. The date
// is relative to the Proleptic Gregorian Calendar. The day may be 0 to
// represent a year and month where the day is not significant, e.g. credit card
// expiration date. The year may be 0 to represent a month and day independent
// of year, e.g. anniversary date. Related types are [google.type.TimeOfDay][google.type.TimeOfDay]
// and [google.protobuf.Timestamp][google.protobuf.Timestamp].
type Date struct {
	// Year of date. Must be from 1 to 9,999, or 0 if specifying a date without
	// a year.
	Year int32 `protobuf:"varint,1,opt,name=year" json:"year,omitempty"`
	// Month of year of date. Must be from 1 to 12.
	Month int32 `protobuf:"varint,2,opt,name=month" json:"month,omitempty"`
	// Day of month. Must be from 1 to 31 and valid for the year and month, or 0
	// if specifying a year/month where the day is not sigificant.
	Day int32 `protobuf:"varint,3,opt,name=day" json:"day,omitempty"`
}

func (m *Date) Reset()                    { *m = Date{} }
func (m *Date) String() string            { return proto.CompactTextString(m) }
func (*Date) ProtoMessage()               {}
func (*Date) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func init() {
	proto.RegisterType((*Date)(nil), "google.type.Date")
}

var fileDescriptor1 = []byte{
	// 126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x4f, 0x49, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x86, 0x88, 0xeb, 0x81, 0xc4, 0x95, 0x0c, 0xb8, 0x58, 0x5c, 0x80, 0x52,
	0x42, 0x3c, 0x5c, 0x2c, 0x95, 0xa9, 0x89, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x42, 0xbc,
	0x5c, 0xac, 0xb9, 0xf9, 0x79, 0x25, 0x19, 0x12, 0x4c, 0x60, 0x2e, 0x37, 0x17, 0x73, 0x4a, 0x62,
	0xa5, 0x04, 0x33, 0x88, 0xe3, 0xa4, 0xc8, 0xc5, 0x9f, 0x9c, 0x9f, 0xab, 0x87, 0x64, 0x88, 0x13,
	0x27, 0xc8, 0x88, 0x00, 0x90, 0xe1, 0x01, 0x8c, 0x0b, 0x18, 0x19, 0x93, 0xd8, 0xc0, 0x16, 0x19,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x3f, 0x56, 0x0c, 0x82, 0x00, 0x00, 0x00,
}
