// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: v1/api.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Interval struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Start         *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End           *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Interval) Reset() {
	*x = Interval{}
	mi := &file_v1_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Interval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interval) ProtoMessage() {}

func (x *Interval) ProtoReflect() protoreflect.Message {
	mi := &file_v1_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interval.ProtoReflect.Descriptor instead.
func (*Interval) Descriptor() ([]byte, []int) {
	return file_v1_api_proto_rawDescGZIP(), []int{0}
}

func (x *Interval) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *Interval) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

type Event struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the name is an index for the lookup table of event names
	Name uint32 `protobuf:"varint,1,opt,name=name,proto3" json:"name,omitempty"`
	// each element is an index for the lookup table of tag names
	Tags          []uint32             `protobuf:"varint,2,rep,packed,name=tags,proto3" json:"tags,omitempty"`
	Interval      *Interval            `protobuf:"bytes,3,opt,name=interval,proto3" json:"interval,omitempty"`
	Duration      *durationpb.Duration `protobuf:"bytes,4,opt,name=duration,proto3" json:"duration,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Event) Reset() {
	*x = Event{}
	mi := &file_v1_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_v1_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_v1_api_proto_rawDescGZIP(), []int{1}
}

func (x *Event) GetName() uint32 {
	if x != nil {
		return x.Name
	}
	return 0
}

func (x *Event) GetTags() []uint32 {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Event) GetInterval() *Interval {
	if x != nil {
		return x.Interval
	}
	return nil
}

func (x *Event) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

// Calendar
type CalendarRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CalendarRequest) Reset() {
	*x = CalendarRequest{}
	mi := &file_v1_api_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CalendarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalendarRequest) ProtoMessage() {}

func (x *CalendarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_api_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalendarRequest.ProtoReflect.Descriptor instead.
func (*CalendarRequest) Descriptor() ([]byte, []int) {
	return file_v1_api_proto_rawDescGZIP(), []int{2}
}

type CalendarResponse struct {
	state         protoimpl.MessageState     `protogen:"open.v1"`
	Sources       []*CalendarResponse_Source `protobuf:"bytes,1,rep,name=sources,proto3" json:"sources,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CalendarResponse) Reset() {
	*x = CalendarResponse{}
	mi := &file_v1_api_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CalendarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalendarResponse) ProtoMessage() {}

func (x *CalendarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_api_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalendarResponse.ProtoReflect.Descriptor instead.
func (*CalendarResponse) Descriptor() ([]byte, []int) {
	return file_v1_api_proto_rawDescGZIP(), []int{3}
}

func (x *CalendarResponse) GetSources() []*CalendarResponse_Source {
	if x != nil {
		return x.Sources
	}
	return nil
}

// Events
type EventsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Interval      *Interval              `protobuf:"bytes,1,opt,name=interval,proto3" json:"interval,omitempty"`
	Timezone      string                 `protobuf:"bytes,2,opt,name=timezone,proto3" json:"timezone,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EventsRequest) Reset() {
	*x = EventsRequest{}
	mi := &file_v1_api_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventsRequest) ProtoMessage() {}

func (x *EventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_api_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventsRequest.ProtoReflect.Descriptor instead.
func (*EventsRequest) Descriptor() ([]byte, []int) {
	return file_v1_api_proto_rawDescGZIP(), []int{4}
}

func (x *EventsRequest) GetInterval() *Interval {
	if x != nil {
		return x.Interval
	}
	return nil
}

func (x *EventsRequest) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

type EventsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EventNames    []string               `protobuf:"bytes,1,rep,name=event_names,json=eventNames,proto3" json:"event_names,omitempty"`
	Tags          []string               `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	Events        []*Event               `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EventsResponse) Reset() {
	*x = EventsResponse{}
	mi := &file_v1_api_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventsResponse) ProtoMessage() {}

func (x *EventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_api_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventsResponse.ProtoReflect.Descriptor instead.
func (*EventsResponse) Descriptor() ([]byte, []int) {
	return file_v1_api_proto_rawDescGZIP(), []int{5}
}

func (x *EventsResponse) GetEventNames() []string {
	if x != nil {
		return x.EventNames
	}
	return nil
}

func (x *EventsResponse) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *EventsResponse) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type CalendarResponse_Source struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	CalendarServer string                 `protobuf:"bytes,1,opt,name=calendar_server,json=calendarServer,proto3" json:"calendar_server,omitempty"`
	Names          []string               `protobuf:"bytes,2,rep,name=names,proto3" json:"names,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *CalendarResponse_Source) Reset() {
	*x = CalendarResponse_Source{}
	mi := &file_v1_api_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CalendarResponse_Source) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalendarResponse_Source) ProtoMessage() {}

func (x *CalendarResponse_Source) ProtoReflect() protoreflect.Message {
	mi := &file_v1_api_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalendarResponse_Source.ProtoReflect.Descriptor instead.
func (*CalendarResponse_Source) Descriptor() ([]byte, []int) {
	return file_v1_api_proto_rawDescGZIP(), []int{3, 0}
}

func (x *CalendarResponse_Source) GetCalendarServer() string {
	if x != nil {
		return x.CalendarServer
	}
	return ""
}

func (x *CalendarResponse_Source) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

var File_v1_api_proto protoreflect.FileDescriptor

const file_v1_api_proto_rawDesc = "" +
	"\n" +
	"\fv1/api.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/duration.proto\"j\n" +
	"\bInterval\x120\n" +
	"\x05start\x18\x01 \x01(\v2\x1a.google.protobuf.TimestampR\x05start\x12,\n" +
	"\x03end\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\x03end\"\x8d\x01\n" +
	"\x05Event\x12\x12\n" +
	"\x04name\x18\x01 \x01(\rR\x04name\x12\x12\n" +
	"\x04tags\x18\x02 \x03(\rR\x04tags\x12%\n" +
	"\binterval\x18\x03 \x01(\v2\t.IntervalR\binterval\x125\n" +
	"\bduration\x18\x04 \x01(\v2\x19.google.protobuf.DurationR\bduration\"\x11\n" +
	"\x0fCalendarRequest\"\x8f\x01\n" +
	"\x10CalendarResponse\x122\n" +
	"\asources\x18\x01 \x03(\v2\x18.CalendarResponse.SourceR\asources\x1aG\n" +
	"\x06Source\x12'\n" +
	"\x0fcalendar_server\x18\x01 \x01(\tR\x0ecalendarServer\x12\x14\n" +
	"\x05names\x18\x02 \x03(\tR\x05names\"R\n" +
	"\rEventsRequest\x12%\n" +
	"\binterval\x18\x01 \x01(\v2\t.IntervalR\binterval\x12\x1a\n" +
	"\btimezone\x18\x02 \x01(\tR\btimezone\"e\n" +
	"\x0eEventsResponse\x12\x1f\n" +
	"\vevent_names\x18\x01 \x03(\tR\n" +
	"eventNames\x12\x12\n" +
	"\x04tags\x18\x02 \x03(\tR\x04tags\x12\x1e\n" +
	"\x06events\x18\x03 \x03(\v2\x06.EventR\x06events2m\n" +
	"\x0fCalendarService\x12/\n" +
	"\bCalendar\x12\x10.CalendarRequest\x1a\x11.CalendarResponse\x12)\n" +
	"\x06Events\x12\x0e.EventsRequest\x1a\x0f.EventsResponseB(B\bApiProtoP\x01Z\x1aschedule-statistics/api/v1b\x06proto3"

var (
	file_v1_api_proto_rawDescOnce sync.Once
	file_v1_api_proto_rawDescData []byte
)

func file_v1_api_proto_rawDescGZIP() []byte {
	file_v1_api_proto_rawDescOnce.Do(func() {
		file_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_v1_api_proto_rawDesc), len(file_v1_api_proto_rawDesc)))
	})
	return file_v1_api_proto_rawDescData
}

var file_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_v1_api_proto_goTypes = []any{
	(*Interval)(nil),                // 0: Interval
	(*Event)(nil),                   // 1: Event
	(*CalendarRequest)(nil),         // 2: CalendarRequest
	(*CalendarResponse)(nil),        // 3: CalendarResponse
	(*EventsRequest)(nil),           // 4: EventsRequest
	(*EventsResponse)(nil),          // 5: EventsResponse
	(*CalendarResponse_Source)(nil), // 6: CalendarResponse.Source
	(*timestamppb.Timestamp)(nil),   // 7: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),     // 8: google.protobuf.Duration
}
var file_v1_api_proto_depIdxs = []int32{
	7, // 0: Interval.start:type_name -> google.protobuf.Timestamp
	7, // 1: Interval.end:type_name -> google.protobuf.Timestamp
	0, // 2: Event.interval:type_name -> Interval
	8, // 3: Event.duration:type_name -> google.protobuf.Duration
	6, // 4: CalendarResponse.sources:type_name -> CalendarResponse.Source
	0, // 5: EventsRequest.interval:type_name -> Interval
	1, // 6: EventsResponse.events:type_name -> Event
	2, // 7: CalendarService.Calendar:input_type -> CalendarRequest
	4, // 8: CalendarService.Events:input_type -> EventsRequest
	3, // 9: CalendarService.Calendar:output_type -> CalendarResponse
	5, // 10: CalendarService.Events:output_type -> EventsResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_v1_api_proto_init() }
func file_v1_api_proto_init() {
	if File_v1_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_v1_api_proto_rawDesc), len(file_v1_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_api_proto_goTypes,
		DependencyIndexes: file_v1_api_proto_depIdxs,
		MessageInfos:      file_v1_api_proto_msgTypes,
	}.Build()
	File_v1_api_proto = out.File
	file_v1_api_proto_goTypes = nil
	file_v1_api_proto_depIdxs = nil
}
