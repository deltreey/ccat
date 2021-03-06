// Code generated by protoc-gen-gogo.
// source: ann.proto
// DO NOT EDIT!

/*
	Package ann is a generated protocol buffer package.

	It is generated from these files:
		ann.proto

	It has these top-level messages:
		Ann
*/
package ann

import "encoding/json"

import proto "github.com/jingweno/ccat/Godeps/_workspace/src/github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto/gogo.pb"

import io "io"
import fmt "fmt"
import github_com_gogo_protobuf_proto "github.com/jingweno/ccat/Godeps/_workspace/src/github.com/gogo/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// An Ann is a source code annotation.
//
// Annotations are unique on (Repo, CommitID, UnitType, Unit, File,
// Start, End, Type).
type Ann struct {
	// Repo is the VCS repository in which this ann exists.
	Repo string `protobuf:"bytes,1,opt,name=repo" json:"Repo,omitempty"`
	// CommitID is the ID of the VCS commit that this ann exists
	// in. The CommitID is always a full commit ID (40 hexadecimal
	// characters for git and hg), never a branch or tag name.
	CommitID string `protobuf:"bytes,2,opt,name=commit_id" json:"CommitID,omitempty"`
	// UnitType is the source unit type that the annotation exists
	// on. It is either the source unit type during whose processing
	// the annotation was detected/created. Multiple annotations may
	// exist on the same file from different source unit types if a
	// file is contained in multiple source units.
	UnitType string `protobuf:"bytes,3,opt,name=unit_type" json:"UnitType,omitempty"`
	// Unit is the name of the source unit that this ann exists in.
	Unit string `protobuf:"bytes,4,opt,name=unit" json:"Unit,omitempty"`
	// File is the filename in which this Ann exists.
	File string `protobuf:"bytes,5,opt,name=file" json:"File,omitempty"`
	// Start is the byte offset of this ann's first byte in File.
	Start uint32 `protobuf:"varint,6,opt,name=start" json:"Start"`
	// End is the byte offset of this ann's last byte in File.
	End uint32 `protobuf:"varint,7,opt,name=end" json:"End"`
	// Type is the type of the annotation. See this package's type
	// constants for a list of possible types.
	Type string `protobuf:"bytes,8,opt,name=type" json:"Type"`
	// Data contains arbitrary JSON data that is specific to this
	// annotation type (e.g., the link URL for Link annotations).
	Data json.RawMessage `protobuf:"bytes,9,opt,name=data" json:"Data,omitempty"`
}

func (m *Ann) Reset()         { *m = Ann{} }
func (m *Ann) String() string { return proto.CompactTextString(m) }
func (*Ann) ProtoMessage()    {}

func init() {
}
func (m *Ann) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Repo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Repo = string(data[index:postIndex])
			index = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommitID = string(data[index:postIndex])
			index = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnitType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UnitType = string(data[index:postIndex])
			index = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Unit = string(data[index:postIndex])
			index = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field File", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.File = string(data[index:postIndex])
			index = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				m.Start |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				m.End |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(data[index:postIndex])
			index = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append([]byte{}, data[index:postIndex]...)
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := github_com_gogo_protobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			index += skippy
		}
	}
	return nil
}
func (m *Ann) Size() (n int) {
	var l int
	_ = l
	l = len(m.Repo)
	n += 1 + l + sovAnn(uint64(l))
	l = len(m.CommitID)
	n += 1 + l + sovAnn(uint64(l))
	l = len(m.UnitType)
	n += 1 + l + sovAnn(uint64(l))
	l = len(m.Unit)
	n += 1 + l + sovAnn(uint64(l))
	l = len(m.File)
	n += 1 + l + sovAnn(uint64(l))
	n += 1 + sovAnn(uint64(m.Start))
	n += 1 + sovAnn(uint64(m.End))
	l = len(m.Type)
	n += 1 + l + sovAnn(uint64(l))
	if m.Data != nil {
		l = len(m.Data)
		n += 1 + l + sovAnn(uint64(l))
	}
	return n
}

func sovAnn(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAnn(x uint64) (n int) {
	return sovAnn(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Ann) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Ann) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintAnn(data, i, uint64(len(m.Repo)))
	i += copy(data[i:], m.Repo)
	data[i] = 0x12
	i++
	i = encodeVarintAnn(data, i, uint64(len(m.CommitID)))
	i += copy(data[i:], m.CommitID)
	data[i] = 0x1a
	i++
	i = encodeVarintAnn(data, i, uint64(len(m.UnitType)))
	i += copy(data[i:], m.UnitType)
	data[i] = 0x22
	i++
	i = encodeVarintAnn(data, i, uint64(len(m.Unit)))
	i += copy(data[i:], m.Unit)
	data[i] = 0x2a
	i++
	i = encodeVarintAnn(data, i, uint64(len(m.File)))
	i += copy(data[i:], m.File)
	data[i] = 0x30
	i++
	i = encodeVarintAnn(data, i, uint64(m.Start))
	data[i] = 0x38
	i++
	i = encodeVarintAnn(data, i, uint64(m.End))
	data[i] = 0x42
	i++
	i = encodeVarintAnn(data, i, uint64(len(m.Type)))
	i += copy(data[i:], m.Type)
	if m.Data != nil {
		data[i] = 0x4a
		i++
		i = encodeVarintAnn(data, i, uint64(len(m.Data)))
		i += copy(data[i:], m.Data)
	}
	return i, nil
}

func encodeFixed64Ann(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Ann(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintAnn(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
