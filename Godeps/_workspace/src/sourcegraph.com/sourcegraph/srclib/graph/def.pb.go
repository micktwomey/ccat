// Code generated by protoc-gen-gogo.
// source: def.proto
// DO NOT EDIT!

/*
	Package graph is a generated protocol buffer package.

	It is generated from these files:
		def.proto
		doc.proto
		output.proto
		ref.proto

	It has these top-level messages:
		DefKey
		Def
		DefDoc
*/
package graph

import "encoding/json"

import proto "github.com/jingweno/ccat/Godeps/_workspace/src/github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto/gogo.pb"

import io "io"
import fmt "fmt"
import github_com_gogo_protobuf_proto "github.com/jingweno/ccat/Godeps/_workspace/src/github.com/gogo/protobuf/proto"

import strings "strings"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// START DefKey OMIT
// DefKey specifies a definition, either concretely or abstractly. A concrete
// definition key has a non-empty CommitID and refers to a definition defined in a
// specific commit. An abstract definition key has an empty CommitID and is
// considered to refer to definitions from any number of commits (so long as the
// Repo, UnitType, Unit, and Path match).
//
// You can think of CommitID as the time dimension. With an empty CommitID, you
// are referring to a definition that may or may not exist at various times. With a
// non-empty CommitID, you are referring to a specific definition of a definition at
// the time specified by the CommitID.
type DefKey struct {
	// Repo is the VCS repository that defines this definition.
	Repo string `protobuf:"bytes,1,opt,name=repo" json:"Repo,omitempty"`
	// CommitID is the ID of the VCS commit that this definition was defined in. The
	// CommitID is always a full commit ID (40 hexadecimal characters for git
	// and hg), never a branch or tag name.
	CommitID string `protobuf:"bytes,2,opt,name=commit_id" json:"CommitID,omitempty"`
	// UnitType is the type name of the source unit (obtained from unit.Type(u))
	// that this definition was defined in.
	UnitType string `protobuf:"bytes,3,opt,name=unit_type" json:"UnitType,omitempty"`
	// Unit is the name of the source unit (obtained from u.Name()) that this
	// definition was defined in.
	Unit string `protobuf:"bytes,4,opt,name=unit" json:"Unit,omitempty"`
	// Path is a unique identifier for the def, relative to the source unit.
	// It should remain stable across commits as long as the def is the
	// "same" def. Its Elasticsearch mapping is defined separately (because
	// it is a multi_field, which the struct tag can't currently represent).
	//
	// Path encodes no structural semantics. Its only meaning is to be a stable
	// unique identifier within a given source unit. In many languages, it is
	// convenient to use the namespace hierarchy (with some modifications) as
	// the Path, but this may not always be the case. I.e., don't rely on Path
	// to find parents or children or any other structural propreties of the
	// def hierarchy). See Def.TreePath instead.
	Path string `protobuf:"bytes,5,opt,name=path" json:"Path"`
}

// END DefKey OMIT

func (m *DefKey) Reset()         { *m = DefKey{} }
func (m *DefKey) String() string { return proto.CompactTextString(m) }
func (*DefKey) ProtoMessage()    {}

// START Def OMIT
// Def is a definition in code.
type Def struct {
	// DefKey is the natural unique key for a def. It is stable
	// (subsequent runs of a grapher will emit the same defs with the same
	// DefKeys).
	DefKey `protobuf:"bytes,1,req,name=key,embedded=key" json:""`
	// Name of the definition. This need not be unique.
	Name string `protobuf:"bytes,2,opt,name=name" json:"Name"`
	// Kind is the kind of thing this definition is. This is
	// language-specific. Possible values include "type", "func",
	// "var", etc.
	Kind     string `protobuf:"bytes,3,opt,name=kind" json:"Kind,omitempty"`
	File     string `protobuf:"bytes,4,opt,name=file" json:"File"`
	DefStart uint32 `protobuf:"varint,5,opt,name=start" json:"DefStart"`
	DefEnd   uint32 `protobuf:"varint,6,opt,name=end" json:"DefEnd"`
	// Exported is whether this def is part of a source unit's
	// public API. For example, in Java a "public" field is
	// Exported.
	Exported bool `protobuf:"varint,7,opt,name=exported" json:"Exported,omitempty"`
	// Local is whether this def is local to a function or some
	// other inner scope. Local defs do *not* have module,
	// package, or file scope. For example, in Java a function's
	// args are Local, but fields with "private" scope are not
	// Local.
	Local bool `protobuf:"varint,8,opt,name=local" json:"Local,omitempty"`
	// Test is whether this def is defined in test code (as opposed to main
	// code). For example, definitions in Go *_test.go files have Test = true.
	Test bool `protobuf:"varint,9,opt,name=test" json:"Test,omitempty"`
	// Data contains additional language- and toolchain-specific information
	// about the def. Data is used to construct function signatures,
	// import/require statements, language-specific type descriptions, etc.
	//
	// To use json.RawMessage:
	// optional bytes data = 10 [(gogoproto.nullable) = false, (gogoproto.customtype) = "encoding/json.RawMessage", (gogoproto.jsontag) = "Data,omitempty"];
	Data json.RawMessage `protobuf:"bytes,10,opt,name=data" json:"Data,omitempty"`
	// Docs are docstrings for this Def. This field is not set in the
	// Defs produced by graphers; they should emit docs in the
	// separate Docs field on the graph.Output struct.
	Docs []DefDoc `protobuf:"bytes,11,rep,name=docs" json:"Docs,omitempty"`
	// TreePath is a structurally significant path descriptor for a def. For
	// many languages, it may be identical or similar to DefKey.Path.
	// However, it has the following constraints, which allow it to define a
	// def tree.
	//
	// A tree-path is a chain of '/'-delimited components. A component is either a
	// def name or a ghost component.
	// - A def name satifies the regex [^/-][^/]*
	// - A ghost component satisfies the regex -[^/]*
	// Any prefix of a tree-path that terminates in a def name must be a valid
	// tree-path for some def.
	// The following regex captures the children of a tree-path X: X(/-[^/]*)*(/[^/-][^/]*)
	TreePath string `protobuf:"bytes,17,opt,name=tree_path" json:"TreePath,omitempty"`
}

// END Def OMIT

func (m *Def) Reset()         { *m = Def{} }
func (m *Def) String() string { return proto.CompactTextString(m) }
func (*Def) ProtoMessage()    {}

// DefDoc is documentation on a Def.
type DefDoc struct {
	// Format is the the MIME-type that the documentation is stored
	// in. Valid formats include 'text/html', 'text/plain',
	// 'text/x-markdown', text/x-rst'.
	Format string `protobuf:"bytes,1,req,name=format" json:"Format"`
	// Data is the actual documentation text.
	Data string `protobuf:"bytes,2,opt,name=data" json:"Data"`
}

func (m *DefDoc) Reset()         { *m = DefDoc{} }
func (m *DefDoc) String() string { return proto.CompactTextString(m) }
func (*DefDoc) ProtoMessage()    {}

func init() {
}
func (m *DefKey) Unmarshal(data []byte) error {
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
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
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
			m.Path = string(data[index:postIndex])
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
func (m *Def) Unmarshal(data []byte) error {
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
				return fmt.Errorf("proto: wrong wireType = %d for field DefKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DefKey.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
			m.Name = string(data[index:postIndex])
			index = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
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
			m.Kind = string(data[index:postIndex])
			index = postIndex
		case 4:
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
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefStart", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				m.DefStart |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefEnd", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				m.DefEnd |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exported", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Exported = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Local", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Local = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Test", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Test = bool(v != 0)
		case 10:
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
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Docs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Docs = append(m.Docs, DefDoc{})
			m.Docs[len(m.Docs)-1].Unmarshal(data[index:postIndex])
			index = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TreePath", wireType)
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
			m.TreePath = string(data[index:postIndex])
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
func (m *DefDoc) Unmarshal(data []byte) error {
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
				return fmt.Errorf("proto: wrong wireType = %d for field Format", wireType)
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
			m.Format = string(data[index:postIndex])
			index = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
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
			m.Data = string(data[index:postIndex])
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
func (m *DefKey) Size() (n int) {
	var l int
	_ = l
	l = len(m.Repo)
	n += 1 + l + sovDef(uint64(l))
	l = len(m.CommitID)
	n += 1 + l + sovDef(uint64(l))
	l = len(m.UnitType)
	n += 1 + l + sovDef(uint64(l))
	l = len(m.Unit)
	n += 1 + l + sovDef(uint64(l))
	l = len(m.Path)
	n += 1 + l + sovDef(uint64(l))
	return n
}

func (m *Def) Size() (n int) {
	var l int
	_ = l
	l = m.DefKey.Size()
	n += 1 + l + sovDef(uint64(l))
	l = len(m.Name)
	n += 1 + l + sovDef(uint64(l))
	l = len(m.Kind)
	n += 1 + l + sovDef(uint64(l))
	l = len(m.File)
	n += 1 + l + sovDef(uint64(l))
	n += 1 + sovDef(uint64(m.DefStart))
	n += 1 + sovDef(uint64(m.DefEnd))
	n += 2
	n += 2
	n += 2
	if m.Data != nil {
		l = len(m.Data)
		n += 1 + l + sovDef(uint64(l))
	}
	if len(m.Docs) > 0 {
		for _, e := range m.Docs {
			l = e.Size()
			n += 1 + l + sovDef(uint64(l))
		}
	}
	l = len(m.TreePath)
	n += 2 + l + sovDef(uint64(l))
	return n
}

func (m *DefDoc) Size() (n int) {
	var l int
	_ = l
	l = len(m.Format)
	n += 1 + l + sovDef(uint64(l))
	l = len(m.Data)
	n += 1 + l + sovDef(uint64(l))
	return n
}

func sovDef(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDef(x uint64) (n int) {
	return sovDef(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DefKey) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *DefKey) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintDef(data, i, uint64(len(m.Repo)))
	i += copy(data[i:], m.Repo)
	data[i] = 0x12
	i++
	i = encodeVarintDef(data, i, uint64(len(m.CommitID)))
	i += copy(data[i:], m.CommitID)
	data[i] = 0x1a
	i++
	i = encodeVarintDef(data, i, uint64(len(m.UnitType)))
	i += copy(data[i:], m.UnitType)
	data[i] = 0x22
	i++
	i = encodeVarintDef(data, i, uint64(len(m.Unit)))
	i += copy(data[i:], m.Unit)
	data[i] = 0x2a
	i++
	i = encodeVarintDef(data, i, uint64(len(m.Path)))
	i += copy(data[i:], m.Path)
	return i, nil
}

func (m *Def) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Def) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintDef(data, i, uint64(m.DefKey.Size()))
	n1, err := m.DefKey.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	data[i] = 0x12
	i++
	i = encodeVarintDef(data, i, uint64(len(m.Name)))
	i += copy(data[i:], m.Name)
	data[i] = 0x1a
	i++
	i = encodeVarintDef(data, i, uint64(len(m.Kind)))
	i += copy(data[i:], m.Kind)
	data[i] = 0x22
	i++
	i = encodeVarintDef(data, i, uint64(len(m.File)))
	i += copy(data[i:], m.File)
	data[i] = 0x28
	i++
	i = encodeVarintDef(data, i, uint64(m.DefStart))
	data[i] = 0x30
	i++
	i = encodeVarintDef(data, i, uint64(m.DefEnd))
	data[i] = 0x38
	i++
	if m.Exported {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	data[i] = 0x40
	i++
	if m.Local {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	data[i] = 0x48
	i++
	if m.Test {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	if m.Data != nil {
		data[i] = 0x52
		i++
		i = encodeVarintDef(data, i, uint64(len(m.Data)))
		i += copy(data[i:], m.Data)
	}
	if len(m.Docs) > 0 {
		for _, msg := range m.Docs {
			data[i] = 0x5a
			i++
			i = encodeVarintDef(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	data[i] = 0x8a
	i++
	data[i] = 0x1
	i++
	i = encodeVarintDef(data, i, uint64(len(m.TreePath)))
	i += copy(data[i:], m.TreePath)
	return i, nil
}

func (m *DefDoc) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *DefDoc) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintDef(data, i, uint64(len(m.Format)))
	i += copy(data[i:], m.Format)
	data[i] = 0x12
	i++
	i = encodeVarintDef(data, i, uint64(len(m.Data)))
	i += copy(data[i:], m.Data)
	return i, nil
}

func encodeFixed64Def(data []byte, offset int, v uint64) int {
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
func encodeFixed32Def(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintDef(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (this *DefKey) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&graph.DefKey{` +
		`Repo:` + fmt.Sprintf("%#v", this.Repo),
		`CommitID:` + fmt.Sprintf("%#v", this.CommitID),
		`UnitType:` + fmt.Sprintf("%#v", this.UnitType),
		`Unit:` + fmt.Sprintf("%#v", this.Unit),
		`Path:` + fmt.Sprintf("%#v", this.Path) + `}`}, ", ")
	return s
}
func (this *Def) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&graph.Def{` +
		`DefKey:` + strings.Replace(this.DefKey.GoString(), `&`, ``, 1),
		`Name:` + fmt.Sprintf("%#v", this.Name),
		`Kind:` + fmt.Sprintf("%#v", this.Kind),
		`File:` + fmt.Sprintf("%#v", this.File),
		`DefStart:` + fmt.Sprintf("%#v", this.DefStart),
		`DefEnd:` + fmt.Sprintf("%#v", this.DefEnd),
		`Exported:` + fmt.Sprintf("%#v", this.Exported),
		`Local:` + fmt.Sprintf("%#v", this.Local),
		`Test:` + fmt.Sprintf("%#v", this.Test),
		`Data:` + fmt.Sprintf("%#v", this.Data),
		`Docs:` + strings.Replace(fmt.Sprintf("%#v", this.Docs), `&`, ``, 1),
		`TreePath:` + fmt.Sprintf("%#v", this.TreePath) + `}`}, ", ")
	return s
}
func (this *DefDoc) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&graph.DefDoc{` +
		`Format:` + fmt.Sprintf("%#v", this.Format),
		`Data:` + fmt.Sprintf("%#v", this.Data) + `}`}, ", ")
	return s
}
func valueToGoStringDef(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringDef(e map[int32]github_com_gogo_protobuf_proto.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "}"
	return s
}
