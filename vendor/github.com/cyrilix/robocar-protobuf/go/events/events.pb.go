// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events/events.proto

package events

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
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

type DriveMode int32

const (
	DriveMode_INVALID DriveMode = 0
	DriveMode_USER    DriveMode = 1
	DriveMode_PILOT   DriveMode = 2
)

var DriveMode_name = map[int32]string{
	0: "INVALID",
	1: "USER",
	2: "PILOT",
}

var DriveMode_value = map[string]int32{
	"INVALID": 0,
	"USER":    1,
	"PILOT":   2,
}

func (x DriveMode) String() string {
	return proto.EnumName(DriveMode_name, int32(x))
}

func (DriveMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{0}
}

type TypeObject int32

const (
	TypeObject_ANY  TypeObject = 0
	TypeObject_CAR  TypeObject = 1
	TypeObject_BUMP TypeObject = 2
	TypeObject_PLOT TypeObject = 3
)

var TypeObject_name = map[int32]string{
	0: "ANY",
	1: "CAR",
	2: "BUMP",
	3: "PLOT",
}

var TypeObject_value = map[string]int32{
	"ANY":  0,
	"CAR":  1,
	"BUMP": 2,
	"PLOT": 3,
}

func (x TypeObject) String() string {
	return proto.EnumName(TypeObject_name, int32(x))
}

func (TypeObject) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{1}
}

type FrameRef struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string               `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FrameRef) Reset()         { *m = FrameRef{} }
func (m *FrameRef) String() string { return proto.CompactTextString(m) }
func (*FrameRef) ProtoMessage()    {}
func (*FrameRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{0}
}

func (m *FrameRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrameRef.Unmarshal(m, b)
}
func (m *FrameRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrameRef.Marshal(b, m, deterministic)
}
func (m *FrameRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrameRef.Merge(m, src)
}
func (m *FrameRef) XXX_Size() int {
	return xxx_messageInfo_FrameRef.Size(m)
}
func (m *FrameRef) XXX_DiscardUnknown() {
	xxx_messageInfo_FrameRef.DiscardUnknown(m)
}

var xxx_messageInfo_FrameRef proto.InternalMessageInfo

func (m *FrameRef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FrameRef) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FrameRef) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type FrameMessage struct {
	Id                   *FrameRef `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Frame                []byte    `protobuf:"bytes,2,opt,name=frame,proto3" json:"frame,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *FrameMessage) Reset()         { *m = FrameMessage{} }
func (m *FrameMessage) String() string { return proto.CompactTextString(m) }
func (*FrameMessage) ProtoMessage()    {}
func (*FrameMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{1}
}

func (m *FrameMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrameMessage.Unmarshal(m, b)
}
func (m *FrameMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrameMessage.Marshal(b, m, deterministic)
}
func (m *FrameMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrameMessage.Merge(m, src)
}
func (m *FrameMessage) XXX_Size() int {
	return xxx_messageInfo_FrameMessage.Size(m)
}
func (m *FrameMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_FrameMessage.DiscardUnknown(m)
}

var xxx_messageInfo_FrameMessage proto.InternalMessageInfo

func (m *FrameMessage) GetId() *FrameRef {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *FrameMessage) GetFrame() []byte {
	if m != nil {
		return m.Frame
	}
	return nil
}

type SteeringMessage struct {
	Steering             float32   `protobuf:"fixed32,1,opt,name=steering,proto3" json:"steering,omitempty"`
	Confidence           float32   `protobuf:"fixed32,2,opt,name=confidence,proto3" json:"confidence,omitempty"`
	FrameRef             *FrameRef `protobuf:"bytes,3,opt,name=frame_ref,json=frameRef,proto3" json:"frame_ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SteeringMessage) Reset()         { *m = SteeringMessage{} }
func (m *SteeringMessage) String() string { return proto.CompactTextString(m) }
func (*SteeringMessage) ProtoMessage()    {}
func (*SteeringMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{2}
}

func (m *SteeringMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SteeringMessage.Unmarshal(m, b)
}
func (m *SteeringMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SteeringMessage.Marshal(b, m, deterministic)
}
func (m *SteeringMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SteeringMessage.Merge(m, src)
}
func (m *SteeringMessage) XXX_Size() int {
	return xxx_messageInfo_SteeringMessage.Size(m)
}
func (m *SteeringMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SteeringMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SteeringMessage proto.InternalMessageInfo

func (m *SteeringMessage) GetSteering() float32 {
	if m != nil {
		return m.Steering
	}
	return 0
}

func (m *SteeringMessage) GetConfidence() float32 {
	if m != nil {
		return m.Confidence
	}
	return 0
}

func (m *SteeringMessage) GetFrameRef() *FrameRef {
	if m != nil {
		return m.FrameRef
	}
	return nil
}

type ThrottleMessage struct {
	Throttle             float32   `protobuf:"fixed32,1,opt,name=throttle,proto3" json:"throttle,omitempty"`
	Confidence           float32   `protobuf:"fixed32,2,opt,name=confidence,proto3" json:"confidence,omitempty"`
	FrameRef             *FrameRef `protobuf:"bytes,3,opt,name=frame_ref,json=frameRef,proto3" json:"frame_ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ThrottleMessage) Reset()         { *m = ThrottleMessage{} }
func (m *ThrottleMessage) String() string { return proto.CompactTextString(m) }
func (*ThrottleMessage) ProtoMessage()    {}
func (*ThrottleMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{3}
}

func (m *ThrottleMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThrottleMessage.Unmarshal(m, b)
}
func (m *ThrottleMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThrottleMessage.Marshal(b, m, deterministic)
}
func (m *ThrottleMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThrottleMessage.Merge(m, src)
}
func (m *ThrottleMessage) XXX_Size() int {
	return xxx_messageInfo_ThrottleMessage.Size(m)
}
func (m *ThrottleMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ThrottleMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ThrottleMessage proto.InternalMessageInfo

func (m *ThrottleMessage) GetThrottle() float32 {
	if m != nil {
		return m.Throttle
	}
	return 0
}

func (m *ThrottleMessage) GetConfidence() float32 {
	if m != nil {
		return m.Confidence
	}
	return 0
}

func (m *ThrottleMessage) GetFrameRef() *FrameRef {
	if m != nil {
		return m.FrameRef
	}
	return nil
}

type DriveModeMessage struct {
	DriveMode            DriveMode `protobuf:"varint,1,opt,name=drive_mode,json=driveMode,proto3,enum=robocar.events.DriveMode" json:"drive_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DriveModeMessage) Reset()         { *m = DriveModeMessage{} }
func (m *DriveModeMessage) String() string { return proto.CompactTextString(m) }
func (*DriveModeMessage) ProtoMessage()    {}
func (*DriveModeMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{4}
}

func (m *DriveModeMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DriveModeMessage.Unmarshal(m, b)
}
func (m *DriveModeMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DriveModeMessage.Marshal(b, m, deterministic)
}
func (m *DriveModeMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DriveModeMessage.Merge(m, src)
}
func (m *DriveModeMessage) XXX_Size() int {
	return xxx_messageInfo_DriveModeMessage.Size(m)
}
func (m *DriveModeMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_DriveModeMessage.DiscardUnknown(m)
}

var xxx_messageInfo_DriveModeMessage proto.InternalMessageInfo

func (m *DriveModeMessage) GetDriveMode() DriveMode {
	if m != nil {
		return m.DriveMode
	}
	return DriveMode_INVALID
}

type ObjectsMessage struct {
	Objects              []*Object `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
	FrameRef             *FrameRef `protobuf:"bytes,2,opt,name=frame_ref,json=frameRef,proto3" json:"frame_ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ObjectsMessage) Reset()         { *m = ObjectsMessage{} }
func (m *ObjectsMessage) String() string { return proto.CompactTextString(m) }
func (*ObjectsMessage) ProtoMessage()    {}
func (*ObjectsMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{5}
}

func (m *ObjectsMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectsMessage.Unmarshal(m, b)
}
func (m *ObjectsMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectsMessage.Marshal(b, m, deterministic)
}
func (m *ObjectsMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectsMessage.Merge(m, src)
}
func (m *ObjectsMessage) XXX_Size() int {
	return xxx_messageInfo_ObjectsMessage.Size(m)
}
func (m *ObjectsMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectsMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectsMessage proto.InternalMessageInfo

func (m *ObjectsMessage) GetObjects() []*Object {
	if m != nil {
		return m.Objects
	}
	return nil
}

func (m *ObjectsMessage) GetFrameRef() *FrameRef {
	if m != nil {
		return m.FrameRef
	}
	return nil
}

// BoundingBox that contains an object
type Object struct {
	Type                 TypeObject `protobuf:"varint,1,opt,name=type,proto3,enum=robocar.events.TypeObject" json:"type,omitempty"`
	Left                 int32      `protobuf:"varint,2,opt,name=left,proto3" json:"left,omitempty"`
	Top                  int32      `protobuf:"varint,3,opt,name=top,proto3" json:"top,omitempty"`
	Right                int32      `protobuf:"varint,4,opt,name=right,proto3" json:"right,omitempty"`
	Bottom               int32      `protobuf:"varint,5,opt,name=bottom,proto3" json:"bottom,omitempty"`
	Confidence           float32    `protobuf:"fixed32,6,opt,name=confidence,proto3" json:"confidence,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Object) Reset()         { *m = Object{} }
func (m *Object) String() string { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()    {}
func (*Object) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{6}
}

func (m *Object) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Object.Unmarshal(m, b)
}
func (m *Object) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Object.Marshal(b, m, deterministic)
}
func (m *Object) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Object.Merge(m, src)
}
func (m *Object) XXX_Size() int {
	return xxx_messageInfo_Object.Size(m)
}
func (m *Object) XXX_DiscardUnknown() {
	xxx_messageInfo_Object.DiscardUnknown(m)
}

var xxx_messageInfo_Object proto.InternalMessageInfo

func (m *Object) GetType() TypeObject {
	if m != nil {
		return m.Type
	}
	return TypeObject_ANY
}

func (m *Object) GetLeft() int32 {
	if m != nil {
		return m.Left
	}
	return 0
}

func (m *Object) GetTop() int32 {
	if m != nil {
		return m.Top
	}
	return 0
}

func (m *Object) GetRight() int32 {
	if m != nil {
		return m.Right
	}
	return 0
}

func (m *Object) GetBottom() int32 {
	if m != nil {
		return m.Bottom
	}
	return 0
}

func (m *Object) GetConfidence() float32 {
	if m != nil {
		return m.Confidence
	}
	return 0
}

type SwitchRecordMessage struct {
	Enabled              bool     `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SwitchRecordMessage) Reset()         { *m = SwitchRecordMessage{} }
func (m *SwitchRecordMessage) String() string { return proto.CompactTextString(m) }
func (*SwitchRecordMessage) ProtoMessage()    {}
func (*SwitchRecordMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{7}
}

func (m *SwitchRecordMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SwitchRecordMessage.Unmarshal(m, b)
}
func (m *SwitchRecordMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SwitchRecordMessage.Marshal(b, m, deterministic)
}
func (m *SwitchRecordMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwitchRecordMessage.Merge(m, src)
}
func (m *SwitchRecordMessage) XXX_Size() int {
	return xxx_messageInfo_SwitchRecordMessage.Size(m)
}
func (m *SwitchRecordMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SwitchRecordMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SwitchRecordMessage proto.InternalMessageInfo

func (m *SwitchRecordMessage) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

// Road description
type RoadMessage struct {
	Contour              []*Point  `protobuf:"bytes,1,rep,name=contour,proto3" json:"contour,omitempty"`
	Ellipse              *Ellipse  `protobuf:"bytes,2,opt,name=ellipse,proto3" json:"ellipse,omitempty"`
	FrameRef             *FrameRef `protobuf:"bytes,3,opt,name=frame_ref,json=frameRef,proto3" json:"frame_ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RoadMessage) Reset()         { *m = RoadMessage{} }
func (m *RoadMessage) String() string { return proto.CompactTextString(m) }
func (*RoadMessage) ProtoMessage()    {}
func (*RoadMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{8}
}

func (m *RoadMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoadMessage.Unmarshal(m, b)
}
func (m *RoadMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoadMessage.Marshal(b, m, deterministic)
}
func (m *RoadMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoadMessage.Merge(m, src)
}
func (m *RoadMessage) XXX_Size() int {
	return xxx_messageInfo_RoadMessage.Size(m)
}
func (m *RoadMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RoadMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RoadMessage proto.InternalMessageInfo

func (m *RoadMessage) GetContour() []*Point {
	if m != nil {
		return m.Contour
	}
	return nil
}

func (m *RoadMessage) GetEllipse() *Ellipse {
	if m != nil {
		return m.Ellipse
	}
	return nil
}

func (m *RoadMessage) GetFrameRef() *FrameRef {
	if m != nil {
		return m.FrameRef
	}
	return nil
}

type Point struct {
	X                    int32    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{9}
}

func (m *Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Point.Unmarshal(m, b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Point.Marshal(b, m, deterministic)
}
func (m *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(m, src)
}
func (m *Point) XXX_Size() int {
	return xxx_messageInfo_Point.Size(m)
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Point) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

type Ellipse struct {
	Center               *Point   `protobuf:"bytes,1,opt,name=center,proto3" json:"center,omitempty"`
	Width                int32    `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Height               int32    `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Angle                float32  `protobuf:"fixed32,4,opt,name=angle,proto3" json:"angle,omitempty"`
	Confidence           float32  `protobuf:"fixed32,5,opt,name=confidence,proto3" json:"confidence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ellipse) Reset()         { *m = Ellipse{} }
func (m *Ellipse) String() string { return proto.CompactTextString(m) }
func (*Ellipse) ProtoMessage()    {}
func (*Ellipse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{10}
}

func (m *Ellipse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ellipse.Unmarshal(m, b)
}
func (m *Ellipse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ellipse.Marshal(b, m, deterministic)
}
func (m *Ellipse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ellipse.Merge(m, src)
}
func (m *Ellipse) XXX_Size() int {
	return xxx_messageInfo_Ellipse.Size(m)
}
func (m *Ellipse) XXX_DiscardUnknown() {
	xxx_messageInfo_Ellipse.DiscardUnknown(m)
}

var xxx_messageInfo_Ellipse proto.InternalMessageInfo

func (m *Ellipse) GetCenter() *Point {
	if m != nil {
		return m.Center
	}
	return nil
}

func (m *Ellipse) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Ellipse) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Ellipse) GetAngle() float32 {
	if m != nil {
		return m.Angle
	}
	return 0
}

func (m *Ellipse) GetConfidence() float32 {
	if m != nil {
		return m.Confidence
	}
	return 0
}

// Record message used to tensorflow learning
type RecordMessage struct {
	Frame                *FrameMessage    `protobuf:"bytes,1,opt,name=frame,proto3" json:"frame,omitempty"`
	Steering             *SteeringMessage `protobuf:"bytes,2,opt,name=steering,proto3" json:"steering,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *RecordMessage) Reset()         { *m = RecordMessage{} }
func (m *RecordMessage) String() string { return proto.CompactTextString(m) }
func (*RecordMessage) ProtoMessage()    {}
func (*RecordMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{11}
}

func (m *RecordMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordMessage.Unmarshal(m, b)
}
func (m *RecordMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordMessage.Marshal(b, m, deterministic)
}
func (m *RecordMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordMessage.Merge(m, src)
}
func (m *RecordMessage) XXX_Size() int {
	return xxx_messageInfo_RecordMessage.Size(m)
}
func (m *RecordMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RecordMessage proto.InternalMessageInfo

func (m *RecordMessage) GetFrame() *FrameMessage {
	if m != nil {
		return m.Frame
	}
	return nil
}

func (m *RecordMessage) GetSteering() *SteeringMessage {
	if m != nil {
		return m.Steering
	}
	return nil
}

func init() {
	proto.RegisterEnum("robocar.events.DriveMode", DriveMode_name, DriveMode_value)
	proto.RegisterEnum("robocar.events.TypeObject", TypeObject_name, TypeObject_value)
	proto.RegisterType((*FrameRef)(nil), "robocar.events.FrameRef")
	proto.RegisterType((*FrameMessage)(nil), "robocar.events.FrameMessage")
	proto.RegisterType((*SteeringMessage)(nil), "robocar.events.SteeringMessage")
	proto.RegisterType((*ThrottleMessage)(nil), "robocar.events.ThrottleMessage")
	proto.RegisterType((*DriveModeMessage)(nil), "robocar.events.DriveModeMessage")
	proto.RegisterType((*ObjectsMessage)(nil), "robocar.events.ObjectsMessage")
	proto.RegisterType((*Object)(nil), "robocar.events.Object")
	proto.RegisterType((*SwitchRecordMessage)(nil), "robocar.events.SwitchRecordMessage")
	proto.RegisterType((*RoadMessage)(nil), "robocar.events.RoadMessage")
	proto.RegisterType((*Point)(nil), "robocar.events.Point")
	proto.RegisterType((*Ellipse)(nil), "robocar.events.Ellipse")
	proto.RegisterType((*RecordMessage)(nil), "robocar.events.RecordMessage")
}

func init() { proto.RegisterFile("events/events.proto", fileDescriptor_8ec31f2d2a3db598) }

var fileDescriptor_8ec31f2d2a3db598 = []byte{
	// 689 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4b, 0x6f, 0xd3, 0x5a,
	0x10, 0xee, 0x71, 0xe2, 0x3c, 0x26, 0xbd, 0xb9, 0xd1, 0xe9, 0xbd, 0xbd, 0xbe, 0x11, 0xa2, 0x95,
	0xd9, 0x44, 0x95, 0xea, 0x40, 0x10, 0x12, 0x88, 0x55, 0x4a, 0x8b, 0x54, 0xa9, 0x8f, 0xe8, 0x34,
	0x45, 0x82, 0x4d, 0xe5, 0xd8, 0xe3, 0xc4, 0xc8, 0xf1, 0x89, 0xec, 0xd3, 0x47, 0x76, 0x2c, 0xf8,
	0x19, 0xfc, 0x00, 0x16, 0xfc, 0x48, 0xe4, 0xf3, 0x48, 0x5b, 0x43, 0x25, 0x84, 0xc4, 0x2a, 0xf3,
	0x4d, 0x66, 0xbe, 0xf9, 0x66, 0xe6, 0x78, 0x60, 0x03, 0xaf, 0x30, 0x15, 0x79, 0x5f, 0xfd, 0x78,
	0x8b, 0x8c, 0x0b, 0x4e, 0xdb, 0x19, 0x9f, 0xf0, 0xc0, 0xcf, 0x3c, 0xe5, 0xed, 0x6e, 0x4d, 0x39,
	0x9f, 0x26, 0xd8, 0x97, 0xff, 0x4e, 0x2e, 0xa3, 0xbe, 0x88, 0xe7, 0x98, 0x0b, 0x7f, 0xbe, 0x50,
	0x09, 0x6e, 0x0c, 0x8d, 0xb7, 0x99, 0x3f, 0x47, 0x86, 0x11, 0xa5, 0x50, 0x4d, 0xfd, 0x39, 0x3a,
	0x64, 0x9b, 0xf4, 0x9a, 0x4c, 0xda, 0xb4, 0x0d, 0x56, 0x1c, 0x3a, 0x96, 0xf4, 0x58, 0x71, 0x48,
	0x5f, 0x01, 0x04, 0x19, 0xfa, 0x02, 0xc3, 0x0b, 0x5f, 0x38, 0x95, 0x6d, 0xd2, 0x6b, 0x0d, 0xba,
	0x9e, 0xaa, 0xe2, 0x99, 0x2a, 0xde, 0xd8, 0x54, 0x61, 0x4d, 0x1d, 0x3d, 0x14, 0xee, 0x09, 0xac,
	0xcb, 0x52, 0xc7, 0x98, 0xe7, 0xfe, 0x14, 0x69, 0x4f, 0x52, 0x13, 0x49, 0xe1, 0x78, 0xf7, 0x85,
	0x7b, 0x46, 0x94, 0x2c, 0xfa, 0x0f, 0xd8, 0x51, 0x81, 0xa5, 0x8e, 0x75, 0xa6, 0x80, 0xfb, 0x99,
	0xc0, 0xdf, 0x67, 0x02, 0x31, 0x8b, 0xd3, 0xa9, 0xe1, 0xec, 0x42, 0x23, 0xd7, 0x2e, 0xc9, 0x6c,
	0xb1, 0x15, 0xa6, 0x8f, 0x01, 0x02, 0x9e, 0x46, 0x71, 0x88, 0x69, 0xa0, 0xa8, 0x2c, 0x76, 0xc7,
	0x43, 0x5f, 0x40, 0x53, 0x12, 0x5f, 0x64, 0x18, 0xe9, 0xce, 0x1e, 0x96, 0xd5, 0x88, 0xb4, 0x25,
	0x65, 0x8c, 0x67, 0x19, 0x17, 0x22, 0xc1, 0x3b, 0x32, 0x84, 0x76, 0x19, 0x19, 0x06, 0xff, 0x29,
	0x19, 0x47, 0xd0, 0xd9, 0xcf, 0xe2, 0x2b, 0x3c, 0xe6, 0xe1, 0x4a, 0xc6, 0x4b, 0x80, 0xb0, 0xf0,
	0x5d, 0xcc, 0x79, 0xa8, 0x84, 0xb4, 0x07, 0xff, 0x97, 0xb9, 0x56, 0x59, 0xac, 0x19, 0x1a, 0xd3,
	0x5d, 0x42, 0xfb, 0x74, 0xf2, 0x11, 0x03, 0x91, 0x1b, 0xae, 0xa7, 0x50, 0xe7, 0xca, 0xe3, 0x90,
	0xed, 0x4a, 0xaf, 0x35, 0xd8, 0x2c, 0x13, 0xa9, 0x04, 0x66, 0xc2, 0xee, 0x37, 0x62, 0xfd, 0x72,
	0x23, 0xdf, 0x08, 0xd4, 0x14, 0x15, 0xf5, 0xa0, 0x2a, 0x96, 0x0b, 0xa3, 0xbc, 0x5b, 0x4e, 0x1e,
	0x2f, 0x17, 0xa8, 0x8b, 0xca, 0xb8, 0xe2, 0x01, 0x27, 0x18, 0x09, 0x59, 0xcc, 0x66, 0xd2, 0xa6,
	0x1d, 0xa8, 0x08, 0xbe, 0x90, 0x83, 0xb4, 0x59, 0x61, 0x16, 0xaf, 0x29, 0x8b, 0xa7, 0x33, 0xe1,
	0x54, 0xa5, 0x4f, 0x01, 0xba, 0x09, 0xb5, 0x09, 0x17, 0x82, 0xcf, 0x1d, 0x5b, 0xba, 0x35, 0x2a,
	0xad, 0xab, 0x56, 0x5e, 0x97, 0xdb, 0x87, 0x8d, 0xb3, 0xeb, 0x58, 0x04, 0x33, 0x86, 0x01, 0xcf,
	0x42, 0x33, 0x2e, 0x07, 0xea, 0x98, 0xfa, 0x93, 0x04, 0xd5, 0x0b, 0x6f, 0x30, 0x03, 0xdd, 0xaf,
	0x04, 0x5a, 0x8c, 0xfb, 0xab, 0xc8, 0x3e, 0xd4, 0x03, 0x9e, 0x0a, 0x7e, 0x99, 0xe9, 0xc1, 0xfe,
	0x5b, 0xee, 0x73, 0xc4, 0xe3, 0x54, 0x30, 0x13, 0x45, 0x9f, 0x41, 0x1d, 0x93, 0x24, 0x5e, 0xe4,
	0xa8, 0xa7, 0xfa, 0x5f, 0x39, 0xe1, 0x40, 0xfd, 0xcd, 0x4c, 0xdc, 0xef, 0xbe, 0xa9, 0x27, 0x60,
	0xcb, 0xda, 0x74, 0x1d, 0xc8, 0x8d, 0xec, 0xc3, 0x66, 0xe4, 0xa6, 0x40, 0x4b, 0x3d, 0x63, 0xb2,
	0x74, 0xbf, 0x10, 0xa8, 0xeb, 0x82, 0x74, 0x17, 0x6a, 0x01, 0xa6, 0x02, 0x33, 0xfd, 0x59, 0x3f,
	0xd0, 0x8a, 0x0e, 0x2a, 0x36, 0x71, 0x1d, 0x87, 0x62, 0xa6, 0xc9, 0x14, 0x28, 0x36, 0x31, 0x43,
	0xb9, 0x20, 0xb5, 0x34, 0x8d, 0x8a, 0x68, 0x3f, 0x9d, 0x26, 0x28, 0xf7, 0x66, 0x31, 0x05, 0x4a,
	0xfb, 0xb1, 0x7f, 0xd8, 0xcf, 0x27, 0x02, 0x7f, 0xdd, 0x5f, 0xcd, 0xc0, 0x5c, 0x13, 0xa5, 0xf1,
	0xd1, 0x4f, 0x07, 0xa1, 0x83, 0xf5, 0xad, 0xa1, 0xaf, 0xef, 0xdc, 0x15, 0x35, 0xf4, 0xad, 0x72,
	0x5a, 0xe9, 0x14, 0xdd, 0x1e, 0x9e, 0x9d, 0x5d, 0x68, 0xae, 0x3e, 0x32, 0xda, 0x82, 0xfa, 0xe1,
	0xc9, 0xbb, 0xe1, 0xd1, 0xe1, 0x7e, 0x67, 0x8d, 0x36, 0xa0, 0x7a, 0x7e, 0x76, 0xc0, 0x3a, 0x84,
	0x36, 0xc1, 0x1e, 0x1d, 0x1e, 0x9d, 0x8e, 0x3b, 0xd6, 0xce, 0x00, 0xe0, 0xf6, 0x65, 0xd3, 0x3a,
	0x54, 0x86, 0x27, 0xef, 0x3b, 0x6b, 0x85, 0xf1, 0x66, 0x58, 0x84, 0x36, 0xa0, 0xba, 0x77, 0x7e,
	0x3c, 0xea, 0x58, 0x85, 0x35, 0x2a, 0x72, 0x2a, 0x7b, 0x8d, 0x0f, 0x35, 0x25, 0x63, 0x52, 0x93,
	0x47, 0xf8, 0xf9, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x97, 0x5e, 0xa7, 0x9f, 0x1f, 0x06, 0x00,
	0x00,
}
