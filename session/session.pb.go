// Code generated by protoc-gen-gogo.
// source: session.proto
// DO NOT EDIT!

/*
	Package session is a generated protocol buffer package.

	It is generated from these files:
		session.proto

	It has these top-level messages:
		Session
		Archive
*/
package session

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/types"
import _ "github.com/gogo/protobuf/gogoproto"
import sync "github.com/havoc-io/mutagen/sync"
import url "github.com/havoc-io/mutagen/url"

import time "time"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Version int32

const (
	Version_Unknown  Version = 0
	Version_Version1 Version = 1
)

var Version_name = map[int32]string{
	0: "Unknown",
	1: "Version1",
}
var Version_value = map[string]int32{
	"Unknown":  0,
	"Version1": 1,
}

func (x Version) String() string {
	return proto.EnumName(Version_name, int32(x))
}
func (Version) EnumDescriptor() ([]byte, []int) { return fileDescriptorSession, []int{0} }

type Session struct {
	Identifier           string    `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Version              Version   `protobuf:"varint,2,opt,name=version,proto3,enum=session.Version" json:"version,omitempty"`
	CreationTime         time.Time `protobuf:"bytes,3,opt,name=creationTime,stdtime" json:"creationTime"`
	CreatingVersionMajor uint32    `protobuf:"varint,4,opt,name=creatingVersionMajor,proto3" json:"creatingVersionMajor,omitempty"`
	CreatingVersionMinor uint32    `protobuf:"varint,5,opt,name=creatingVersionMinor,proto3" json:"creatingVersionMinor,omitempty"`
	CreatingVersionPatch uint32    `protobuf:"varint,6,opt,name=creatingVersionPatch,proto3" json:"creatingVersionPatch,omitempty"`
	Alpha                *url.URL  `protobuf:"bytes,7,opt,name=alpha" json:"alpha,omitempty"`
	Beta                 *url.URL  `protobuf:"bytes,8,opt,name=beta" json:"beta,omitempty"`
	Ignores              []string  `protobuf:"bytes,9,rep,name=ignores" json:"ignores,omitempty"`
	Paused               bool      `protobuf:"varint,10,opt,name=paused,proto3" json:"paused,omitempty"`
}

func (m *Session) Reset()                    { *m = Session{} }
func (m *Session) String() string            { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()               {}
func (*Session) Descriptor() ([]byte, []int) { return fileDescriptorSession, []int{0} }

func (m *Session) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *Session) GetVersion() Version {
	if m != nil {
		return m.Version
	}
	return Version_Unknown
}

func (m *Session) GetCreationTime() time.Time {
	if m != nil {
		return m.CreationTime
	}
	return time.Time{}
}

func (m *Session) GetCreatingVersionMajor() uint32 {
	if m != nil {
		return m.CreatingVersionMajor
	}
	return 0
}

func (m *Session) GetCreatingVersionMinor() uint32 {
	if m != nil {
		return m.CreatingVersionMinor
	}
	return 0
}

func (m *Session) GetCreatingVersionPatch() uint32 {
	if m != nil {
		return m.CreatingVersionPatch
	}
	return 0
}

func (m *Session) GetAlpha() *url.URL {
	if m != nil {
		return m.Alpha
	}
	return nil
}

func (m *Session) GetBeta() *url.URL {
	if m != nil {
		return m.Beta
	}
	return nil
}

func (m *Session) GetIgnores() []string {
	if m != nil {
		return m.Ignores
	}
	return nil
}

func (m *Session) GetPaused() bool {
	if m != nil {
		return m.Paused
	}
	return false
}

// Archive is a wrapper around Entry that allows identification of non-existent
// directories. In-memory, a nil-Entry (that arrives without any error)
// represents an absence of content on the filesystem. Unfortunately, there is
// no way to represent that as an encoded message (an empty byte slice would
// successfully decode to an empty directory entry). By adding a level of
// indirection that allows for an unset root entry, we can encode Entry messages
// in a way that allows us to represent absence.
type Archive struct {
	Root *sync.Entry `protobuf:"bytes,1,opt,name=root" json:"root,omitempty"`
}

func (m *Archive) Reset()                    { *m = Archive{} }
func (m *Archive) String() string            { return proto.CompactTextString(m) }
func (*Archive) ProtoMessage()               {}
func (*Archive) Descriptor() ([]byte, []int) { return fileDescriptorSession, []int{1} }

func (m *Archive) GetRoot() *sync.Entry {
	if m != nil {
		return m.Root
	}
	return nil
}

func init() {
	proto.RegisterType((*Session)(nil), "session.Session")
	proto.RegisterType((*Archive)(nil), "session.Archive")
	proto.RegisterEnum("session.Version", Version_name, Version_value)
}
func (m *Session) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Session) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Identifier) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSession(dAtA, i, uint64(len(m.Identifier)))
		i += copy(dAtA[i:], m.Identifier)
	}
	if m.Version != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSession(dAtA, i, uint64(m.Version))
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintSession(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.CreationTime)))
	n1, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreationTime, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.CreatingVersionMajor != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintSession(dAtA, i, uint64(m.CreatingVersionMajor))
	}
	if m.CreatingVersionMinor != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintSession(dAtA, i, uint64(m.CreatingVersionMinor))
	}
	if m.CreatingVersionPatch != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintSession(dAtA, i, uint64(m.CreatingVersionPatch))
	}
	if m.Alpha != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintSession(dAtA, i, uint64(m.Alpha.Size()))
		n2, err := m.Alpha.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Beta != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintSession(dAtA, i, uint64(m.Beta.Size()))
		n3, err := m.Beta.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if len(m.Ignores) > 0 {
		for _, s := range m.Ignores {
			dAtA[i] = 0x4a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.Paused {
		dAtA[i] = 0x50
		i++
		if m.Paused {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *Archive) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Archive) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Root != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSession(dAtA, i, uint64(m.Root.Size()))
		n4, err := m.Root.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func encodeFixed64Session(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Session(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintSession(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Session) Size() (n int) {
	var l int
	_ = l
	l = len(m.Identifier)
	if l > 0 {
		n += 1 + l + sovSession(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovSession(uint64(m.Version))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreationTime)
	n += 1 + l + sovSession(uint64(l))
	if m.CreatingVersionMajor != 0 {
		n += 1 + sovSession(uint64(m.CreatingVersionMajor))
	}
	if m.CreatingVersionMinor != 0 {
		n += 1 + sovSession(uint64(m.CreatingVersionMinor))
	}
	if m.CreatingVersionPatch != 0 {
		n += 1 + sovSession(uint64(m.CreatingVersionPatch))
	}
	if m.Alpha != nil {
		l = m.Alpha.Size()
		n += 1 + l + sovSession(uint64(l))
	}
	if m.Beta != nil {
		l = m.Beta.Size()
		n += 1 + l + sovSession(uint64(l))
	}
	if len(m.Ignores) > 0 {
		for _, s := range m.Ignores {
			l = len(s)
			n += 1 + l + sovSession(uint64(l))
		}
	}
	if m.Paused {
		n += 2
	}
	return n
}

func (m *Archive) Size() (n int) {
	var l int
	_ = l
	if m.Root != nil {
		l = m.Root.Size()
		n += 1 + l + sovSession(uint64(l))
	}
	return n
}

func sovSession(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSession(x uint64) (n int) {
	return sovSession(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Session) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSession
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Session: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Session: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= (Version(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreationTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatingVersionMajor", wireType)
			}
			m.CreatingVersionMajor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatingVersionMajor |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatingVersionMinor", wireType)
			}
			m.CreatingVersionMinor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatingVersionMinor |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatingVersionPatch", wireType)
			}
			m.CreatingVersionPatch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatingVersionPatch |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alpha", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Alpha == nil {
				m.Alpha = &url.URL{}
			}
			if err := m.Alpha.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Beta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Beta == nil {
				m.Beta = &url.URL{}
			}
			if err := m.Beta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ignores", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ignores = append(m.Ignores, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Paused", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Paused = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipSession(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSession
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Archive) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSession
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Archive: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Archive: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Root", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Root == nil {
				m.Root = &sync.Entry{}
			}
			if err := m.Root.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSession(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSession
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSession(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSession
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSession
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSession
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSession
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSession
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSession(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSession = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSession   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("session.proto", fileDescriptorSession) }

var fileDescriptorSession = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0x86, 0x77, 0xdc, 0x6e, 0x93, 0x9e, 0x76, 0xa5, 0x0c, 0x22, 0x43, 0x91, 0x34, 0x2c, 0x5e,
	0xc4, 0xe2, 0x4e, 0xb0, 0x3e, 0x81, 0x0b, 0x82, 0x17, 0x0a, 0x32, 0xee, 0x7a, 0x3f, 0xcd, 0xce,
	0x26, 0xa3, 0xcd, 0x9c, 0x32, 0x99, 0x54, 0xf6, 0x2d, 0x7c, 0x0f, 0x5f, 0x64, 0x2f, 0x7d, 0x02,
	0x95, 0x3e, 0x89, 0x64, 0x92, 0xe0, 0x2a, 0xc1, 0x8b, 0x42, 0xff, 0xff, 0xff, 0x7e, 0x32, 0xe7,
	0xcc, 0xc0, 0x69, 0xa5, 0xaa, 0x4a, 0xa3, 0xe1, 0x3b, 0x8b, 0x0e, 0x69, 0xd0, 0xc9, 0xc5, 0x32,
	0x47, 0xcc, 0xb7, 0x2a, 0xf5, 0xf6, 0xa6, 0xbe, 0x49, 0x9d, 0x2e, 0x55, 0xe5, 0x64, 0xb9, 0x6b,
	0xc9, 0xc5, 0x79, 0xae, 0x5d, 0x51, 0x6f, 0x78, 0x86, 0x65, 0x9a, 0x63, 0x8e, 0x7f, 0xc8, 0x46,
	0x79, 0xe1, 0xff, 0x75, 0xf8, 0xf3, 0x7b, 0x78, 0x21, 0xf7, 0x98, 0x9d, 0x6b, 0x4c, 0xcb, 0xda,
	0xc9, 0x5c, 0x99, 0xb4, 0xba, 0x35, 0x59, 0xaa, 0x8c, 0xb3, 0xb7, 0x1d, 0xfd, 0xec, 0x7f, 0x74,
	0x6d, 0xb7, 0xcd, 0xaf, 0x45, 0xcf, 0xbe, 0x1d, 0x43, 0xf0, 0xa1, 0x3d, 0x34, 0x8d, 0x00, 0xf4,
	0xb5, 0x32, 0x4e, 0xdf, 0x68, 0x65, 0x19, 0x89, 0x49, 0x32, 0x11, 0xf7, 0x1c, 0xba, 0x82, 0x60,
	0xaf, 0x6c, 0x83, 0xb2, 0x07, 0x31, 0x49, 0x1e, 0xae, 0xe7, 0xbc, 0x1f, 0xff, 0x63, 0xeb, 0x8b,
	0x1e, 0xa0, 0x6f, 0x60, 0x96, 0x59, 0x25, 0x9d, 0x46, 0x73, 0xa9, 0x4b, 0xc5, 0x8e, 0x63, 0x92,
	0x4c, 0xd7, 0x0b, 0xde, 0xee, 0x85, 0xf7, 0xd3, 0xf2, 0xcb, 0x7e, 0x2f, 0x17, 0xe1, 0xdd, 0x8f,
	0xe5, 0xd1, 0xd7, 0x9f, 0x4b, 0x22, 0xfe, 0x6a, 0xd2, 0x35, 0x3c, 0x6a, 0xb5, 0xc9, 0xbb, 0xaf,
	0xbc, 0x93, 0x9f, 0xd0, 0xb2, 0x51, 0x4c, 0x92, 0x53, 0x31, 0x98, 0x0d, 0x75, 0xb4, 0x41, 0xcb,
	0x4e, 0x86, 0x3b, 0x4d, 0x36, 0xd0, 0x79, 0x2f, 0x5d, 0x56, 0xb0, 0xf1, 0x60, 0xc7, 0x67, 0x34,
	0x82, 0x13, 0xb9, 0xdd, 0x15, 0x92, 0x05, 0x7e, 0xbc, 0x90, 0x37, 0x8b, 0xbd, 0x12, 0x6f, 0x45,
	0x6b, 0xd3, 0x27, 0x30, 0xda, 0x28, 0x27, 0x59, 0xf8, 0x4f, 0xec, 0x5d, 0xca, 0x20, 0xd0, 0xb9,
	0x41, 0xab, 0x2a, 0x36, 0x89, 0x8f, 0x93, 0x89, 0xe8, 0x25, 0x7d, 0x0c, 0xe3, 0x9d, 0xac, 0x2b,
	0x75, 0xcd, 0x20, 0x26, 0x49, 0x28, 0x3a, 0x75, 0xb6, 0x82, 0xe0, 0x95, 0xcd, 0x0a, 0xbd, 0x57,
	0x74, 0x09, 0x23, 0x8b, 0xe8, 0xfc, 0x35, 0x4d, 0xd7, 0x53, 0xde, 0x3c, 0x02, 0xfe, 0xba, 0x79,
	0x04, 0xc2, 0x07, 0xab, 0xa7, 0x10, 0x74, 0x67, 0xa5, 0x53, 0x08, 0xae, 0xcc, 0x67, 0x83, 0x5f,
	0xcc, 0xfc, 0x88, 0xce, 0x20, 0xec, 0xfc, 0x17, 0x73, 0x72, 0x31, 0xbb, 0x3b, 0x44, 0xe4, 0xfb,
	0x21, 0x22, 0xbf, 0x0e, 0x11, 0xd9, 0x8c, 0xfd, 0xbd, 0xbc, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff,
	0xf3, 0xc9, 0x00, 0x84, 0xd7, 0x02, 0x00, 0x00,
}
