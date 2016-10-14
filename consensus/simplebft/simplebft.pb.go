// Code generated by protoc-gen-go.
// source: simplebft.proto
// DO NOT EDIT!

/*
Package simplebft is a generated protocol buffer package.

It is generated from these files:
	simplebft.proto

It has these top-level messages:
	Config
	Msg
	Request
	SeqView
	BatchHeader
	Batch
	Preprepare
	Subject
	ViewChange
	Signed
	NewView
	Checkpoint
*/
package simplebft

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Config struct {
	N                  uint64 `protobuf:"varint,1,opt,name=n" json:"n,omitempty"`
	F                  uint64 `protobuf:"varint,2,opt,name=f" json:"f,omitempty"`
	BatchDurationNsec  uint64 `protobuf:"varint,3,opt,name=batch_duration_nsec" json:"batch_duration_nsec,omitempty"`
	BatchSizeBytes     uint64 `protobuf:"varint,4,opt,name=batch_size_bytes" json:"batch_size_bytes,omitempty"`
	RequestTimeoutNsec uint64 `protobuf:"varint,5,opt,name=request_timeout_nsec" json:"request_timeout_nsec,omitempty"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}

type Msg struct {
	// Types that are valid to be assigned to Type:
	//	*Msg_Request
	//	*Msg_Preprepare
	//	*Msg_Prepare
	//	*Msg_Commit
	//	*Msg_ViewChange
	//	*Msg_NewView
	//	*Msg_Checkpoint
	//	*Msg_Hello
	Type isMsg_Type `protobuf_oneof:"type"`
}

func (m *Msg) Reset()         { *m = Msg{} }
func (m *Msg) String() string { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()    {}

type isMsg_Type interface {
	isMsg_Type()
}

type Msg_Request struct {
	Request *Request `protobuf:"bytes,1,opt,name=request,oneof"`
}
type Msg_Preprepare struct {
	Preprepare *Preprepare `protobuf:"bytes,2,opt,name=preprepare,oneof"`
}
type Msg_Prepare struct {
	Prepare *Subject `protobuf:"bytes,3,opt,name=prepare,oneof"`
}
type Msg_Commit struct {
	Commit *Subject `protobuf:"bytes,4,opt,name=commit,oneof"`
}
type Msg_ViewChange struct {
	ViewChange *Signed `protobuf:"bytes,5,opt,name=view_change,oneof"`
}
type Msg_NewView struct {
	NewView *NewView `protobuf:"bytes,6,opt,name=new_view,oneof"`
}
type Msg_Checkpoint struct {
	Checkpoint *Checkpoint `protobuf:"bytes,7,opt,name=checkpoint,oneof"`
}
type Msg_Hello struct {
	Hello *Batch `protobuf:"bytes,8,opt,name=hello,oneof"`
}

func (*Msg_Request) isMsg_Type()    {}
func (*Msg_Preprepare) isMsg_Type() {}
func (*Msg_Prepare) isMsg_Type()    {}
func (*Msg_Commit) isMsg_Type()     {}
func (*Msg_ViewChange) isMsg_Type() {}
func (*Msg_NewView) isMsg_Type()    {}
func (*Msg_Checkpoint) isMsg_Type() {}
func (*Msg_Hello) isMsg_Type()      {}

func (m *Msg) GetType() isMsg_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Msg) GetRequest() *Request {
	if x, ok := m.GetType().(*Msg_Request); ok {
		return x.Request
	}
	return nil
}

func (m *Msg) GetPreprepare() *Preprepare {
	if x, ok := m.GetType().(*Msg_Preprepare); ok {
		return x.Preprepare
	}
	return nil
}

func (m *Msg) GetPrepare() *Subject {
	if x, ok := m.GetType().(*Msg_Prepare); ok {
		return x.Prepare
	}
	return nil
}

func (m *Msg) GetCommit() *Subject {
	if x, ok := m.GetType().(*Msg_Commit); ok {
		return x.Commit
	}
	return nil
}

func (m *Msg) GetViewChange() *Signed {
	if x, ok := m.GetType().(*Msg_ViewChange); ok {
		return x.ViewChange
	}
	return nil
}

func (m *Msg) GetNewView() *NewView {
	if x, ok := m.GetType().(*Msg_NewView); ok {
		return x.NewView
	}
	return nil
}

func (m *Msg) GetCheckpoint() *Checkpoint {
	if x, ok := m.GetType().(*Msg_Checkpoint); ok {
		return x.Checkpoint
	}
	return nil
}

func (m *Msg) GetHello() *Batch {
	if x, ok := m.GetType().(*Msg_Hello); ok {
		return x.Hello
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Msg) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), []interface{}) {
	return _Msg_OneofMarshaler, _Msg_OneofUnmarshaler, []interface{}{
		(*Msg_Request)(nil),
		(*Msg_Preprepare)(nil),
		(*Msg_Prepare)(nil),
		(*Msg_Commit)(nil),
		(*Msg_ViewChange)(nil),
		(*Msg_NewView)(nil),
		(*Msg_Checkpoint)(nil),
		(*Msg_Hello)(nil),
	}
}

func _Msg_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Msg)
	// type
	switch x := m.Type.(type) {
	case *Msg_Request:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Request); err != nil {
			return err
		}
	case *Msg_Preprepare:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Preprepare); err != nil {
			return err
		}
	case *Msg_Prepare:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Prepare); err != nil {
			return err
		}
	case *Msg_Commit:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Commit); err != nil {
			return err
		}
	case *Msg_ViewChange:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ViewChange); err != nil {
			return err
		}
	case *Msg_NewView:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NewView); err != nil {
			return err
		}
	case *Msg_Checkpoint:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Checkpoint); err != nil {
			return err
		}
	case *Msg_Hello:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Hello); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Msg.Type has unexpected type %T", x)
	}
	return nil
}

func _Msg_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Msg)
	switch tag {
	case 1: // type.request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Request)
		err := b.DecodeMessage(msg)
		m.Type = &Msg_Request{msg}
		return true, err
	case 2: // type.preprepare
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Preprepare)
		err := b.DecodeMessage(msg)
		m.Type = &Msg_Preprepare{msg}
		return true, err
	case 3: // type.prepare
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Subject)
		err := b.DecodeMessage(msg)
		m.Type = &Msg_Prepare{msg}
		return true, err
	case 4: // type.commit
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Subject)
		err := b.DecodeMessage(msg)
		m.Type = &Msg_Commit{msg}
		return true, err
	case 5: // type.view_change
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Signed)
		err := b.DecodeMessage(msg)
		m.Type = &Msg_ViewChange{msg}
		return true, err
	case 6: // type.new_view
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NewView)
		err := b.DecodeMessage(msg)
		m.Type = &Msg_NewView{msg}
		return true, err
	case 7: // type.checkpoint
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Checkpoint)
		err := b.DecodeMessage(msg)
		m.Type = &Msg_Checkpoint{msg}
		return true, err
	case 8: // type.hello
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Batch)
		err := b.DecodeMessage(msg)
		m.Type = &Msg_Hello{msg}
		return true, err
	default:
		return false, nil
	}
}

type Request struct {
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

type SeqView struct {
	View uint64 `protobuf:"varint,1,opt,name=view" json:"view,omitempty"`
	Seq  uint64 `protobuf:"varint,2,opt,name=seq" json:"seq,omitempty"`
}

func (m *SeqView) Reset()         { *m = SeqView{} }
func (m *SeqView) String() string { return proto.CompactTextString(m) }
func (*SeqView) ProtoMessage()    {}

type BatchHeader struct {
	Seq      uint64 `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	PrevHash []byte `protobuf:"bytes,2,opt,name=prev_hash,proto3" json:"prev_hash,omitempty"`
	DataHash []byte `protobuf:"bytes,3,opt,name=data_hash,proto3" json:"data_hash,omitempty"`
}

func (m *BatchHeader) Reset()         { *m = BatchHeader{} }
func (m *BatchHeader) String() string { return proto.CompactTextString(m) }
func (*BatchHeader) ProtoMessage()    {}

type Batch struct {
	Header     []byte            `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Payloads   [][]byte          `protobuf:"bytes,2,rep,name=payloads,proto3" json:"payloads,omitempty"`
	Signatures map[uint64][]byte `protobuf:"bytes,3,rep,name=signatures" json:"signatures,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Batch) Reset()         { *m = Batch{} }
func (m *Batch) String() string { return proto.CompactTextString(m) }
func (*Batch) ProtoMessage()    {}

func (m *Batch) GetSignatures() map[uint64][]byte {
	if m != nil {
		return m.Signatures
	}
	return nil
}

type Preprepare struct {
	Seq   *SeqView `protobuf:"bytes,1,opt,name=seq" json:"seq,omitempty"`
	Batch *Batch   `protobuf:"bytes,2,opt,name=batch" json:"batch,omitempty"`
}

func (m *Preprepare) Reset()         { *m = Preprepare{} }
func (m *Preprepare) String() string { return proto.CompactTextString(m) }
func (*Preprepare) ProtoMessage()    {}

func (m *Preprepare) GetSeq() *SeqView {
	if m != nil {
		return m.Seq
	}
	return nil
}

func (m *Preprepare) GetBatch() *Batch {
	if m != nil {
		return m.Batch
	}
	return nil
}

type Subject struct {
	Seq    *SeqView `protobuf:"bytes,1,opt,name=seq" json:"seq,omitempty"`
	Digest []byte   `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
}

func (m *Subject) Reset()         { *m = Subject{} }
func (m *Subject) String() string { return proto.CompactTextString(m) }
func (*Subject) ProtoMessage()    {}

func (m *Subject) GetSeq() *SeqView {
	if m != nil {
		return m.Seq
	}
	return nil
}

type ViewChange struct {
	View     uint64     `protobuf:"varint,1,opt,name=view" json:"view,omitempty"`
	Pset     []*Subject `protobuf:"bytes,2,rep,name=pset" json:"pset,omitempty"`
	Qset     []*Subject `protobuf:"bytes,3,rep,name=qset" json:"qset,omitempty"`
	Executed uint64     `protobuf:"varint,4,opt,name=executed" json:"executed,omitempty"`
}

func (m *ViewChange) Reset()         { *m = ViewChange{} }
func (m *ViewChange) String() string { return proto.CompactTextString(m) }
func (*ViewChange) ProtoMessage()    {}

func (m *ViewChange) GetPset() []*Subject {
	if m != nil {
		return m.Pset
	}
	return nil
}

func (m *ViewChange) GetQset() []*Subject {
	if m != nil {
		return m.Qset
	}
	return nil
}

type Signed struct {
	Data      []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Signed) Reset()         { *m = Signed{} }
func (m *Signed) String() string { return proto.CompactTextString(m) }
func (*Signed) ProtoMessage()    {}

type NewView struct {
	View  uint64             `protobuf:"varint,1,opt,name=view" json:"view,omitempty"`
	Vset  map[uint64]*Signed `protobuf:"bytes,2,rep,name=vset" json:"vset,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Xset  *Subject           `protobuf:"bytes,3,opt,name=xset" json:"xset,omitempty"`
	Batch *Batch             `protobuf:"bytes,4,opt,name=batch" json:"batch,omitempty"`
}

func (m *NewView) Reset()         { *m = NewView{} }
func (m *NewView) String() string { return proto.CompactTextString(m) }
func (*NewView) ProtoMessage()    {}

func (m *NewView) GetVset() map[uint64]*Signed {
	if m != nil {
		return m.Vset
	}
	return nil
}

func (m *NewView) GetXset() *Subject {
	if m != nil {
		return m.Xset
	}
	return nil
}

func (m *NewView) GetBatch() *Batch {
	if m != nil {
		return m.Batch
	}
	return nil
}

type Checkpoint struct {
	Seq       uint64 `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	Digest    []byte `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Checkpoint) Reset()         { *m = Checkpoint{} }
func (m *Checkpoint) String() string { return proto.CompactTextString(m) }
func (*Checkpoint) ProtoMessage()    {}
