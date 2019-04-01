// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/extension_feed_item.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import common "google.golang.org/genproto/googleapis/ads/googleads/v1/common"
import enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// An extension feed item.
type ExtensionFeedItem struct {
	// The resource name of the extension feed item.
	// Extension feed item resource names have the form:
	//
	// `customers/{customer_id}/extensionFeedItems/{feed_item_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The extension type of the extension feed item.
	// This field is read-only.
	ExtensionType enums.ExtensionTypeEnum_ExtensionType `protobuf:"varint,13,opt,name=extension_type,json=extensionType,proto3,enum=google.ads.googleads.v1.enums.ExtensionTypeEnum_ExtensionType" json:"extension_type,omitempty"`
	// Start time in which this feed item is effective and can begin serving.
	// The format is "YYYY-MM-DD HH:MM:SS".
	// Examples: "2018-03-05 09:15:00" or "2018-02-01 14:34:30"
	StartDateTime *wrappers.StringValue `protobuf:"bytes,5,opt,name=start_date_time,json=startDateTime,proto3" json:"start_date_time,omitempty"`
	// End time in which this feed item is no longer effective and will stop
	// serving.
	// The format is "YYYY-MM-DD HH:MM:SS".
	// Examples: "2018-03-05 09:15:00" or "2018-02-01 14:34:30"
	EndDateTime *wrappers.StringValue `protobuf:"bytes,6,opt,name=end_date_time,json=endDateTime,proto3" json:"end_date_time,omitempty"`
	// Status of the feed item.
	// This field is read-only.
	Status enums.FeedItemStatusEnum_FeedItemStatus `protobuf:"varint,4,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.FeedItemStatusEnum_FeedItemStatus" json:"status,omitempty"`
	// Extension type.
	//
	// Types that are valid to be assigned to Extension:
	//	*ExtensionFeedItem_SitelinkFeedItem
	//	*ExtensionFeedItem_StructuredSnippetFeedItem
	//	*ExtensionFeedItem_AppFeedItem
	//	*ExtensionFeedItem_CallFeedItem
	//	*ExtensionFeedItem_CalloutFeedItem
	//	*ExtensionFeedItem_TextMessageFeedItem
	//	*ExtensionFeedItem_PriceFeedItem
	//	*ExtensionFeedItem_PromotionFeedItem
	Extension            isExtensionFeedItem_Extension `protobuf_oneof:"extension"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ExtensionFeedItem) Reset()         { *m = ExtensionFeedItem{} }
func (m *ExtensionFeedItem) String() string { return proto.CompactTextString(m) }
func (*ExtensionFeedItem) ProtoMessage()    {}
func (*ExtensionFeedItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_extension_feed_item_ecd9cb7f4df90d10, []int{0}
}
func (m *ExtensionFeedItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtensionFeedItem.Unmarshal(m, b)
}
func (m *ExtensionFeedItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtensionFeedItem.Marshal(b, m, deterministic)
}
func (dst *ExtensionFeedItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtensionFeedItem.Merge(dst, src)
}
func (m *ExtensionFeedItem) XXX_Size() int {
	return xxx_messageInfo_ExtensionFeedItem.Size(m)
}
func (m *ExtensionFeedItem) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtensionFeedItem.DiscardUnknown(m)
}

var xxx_messageInfo_ExtensionFeedItem proto.InternalMessageInfo

func (m *ExtensionFeedItem) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *ExtensionFeedItem) GetExtensionType() enums.ExtensionTypeEnum_ExtensionType {
	if m != nil {
		return m.ExtensionType
	}
	return enums.ExtensionTypeEnum_UNSPECIFIED
}

func (m *ExtensionFeedItem) GetStartDateTime() *wrappers.StringValue {
	if m != nil {
		return m.StartDateTime
	}
	return nil
}

func (m *ExtensionFeedItem) GetEndDateTime() *wrappers.StringValue {
	if m != nil {
		return m.EndDateTime
	}
	return nil
}

func (m *ExtensionFeedItem) GetStatus() enums.FeedItemStatusEnum_FeedItemStatus {
	if m != nil {
		return m.Status
	}
	return enums.FeedItemStatusEnum_UNSPECIFIED
}

type isExtensionFeedItem_Extension interface {
	isExtensionFeedItem_Extension()
}

type ExtensionFeedItem_SitelinkFeedItem struct {
	SitelinkFeedItem *common.SitelinkFeedItem `protobuf:"bytes,2,opt,name=sitelink_feed_item,json=sitelinkFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_StructuredSnippetFeedItem struct {
	StructuredSnippetFeedItem *common.StructuredSnippetFeedItem `protobuf:"bytes,3,opt,name=structured_snippet_feed_item,json=structuredSnippetFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_AppFeedItem struct {
	AppFeedItem *common.AppFeedItem `protobuf:"bytes,7,opt,name=app_feed_item,json=appFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_CallFeedItem struct {
	CallFeedItem *common.CallFeedItem `protobuf:"bytes,8,opt,name=call_feed_item,json=callFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_CalloutFeedItem struct {
	CalloutFeedItem *common.CalloutFeedItem `protobuf:"bytes,9,opt,name=callout_feed_item,json=calloutFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_TextMessageFeedItem struct {
	TextMessageFeedItem *common.TextMessageFeedItem `protobuf:"bytes,10,opt,name=text_message_feed_item,json=textMessageFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_PriceFeedItem struct {
	PriceFeedItem *common.PriceFeedItem `protobuf:"bytes,11,opt,name=price_feed_item,json=priceFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_PromotionFeedItem struct {
	PromotionFeedItem *common.PromotionFeedItem `protobuf:"bytes,12,opt,name=promotion_feed_item,json=promotionFeedItem,proto3,oneof"`
}

func (*ExtensionFeedItem_SitelinkFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_StructuredSnippetFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_AppFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_CallFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_CalloutFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_TextMessageFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_PriceFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_PromotionFeedItem) isExtensionFeedItem_Extension() {}

func (m *ExtensionFeedItem) GetExtension() isExtensionFeedItem_Extension {
	if m != nil {
		return m.Extension
	}
	return nil
}

func (m *ExtensionFeedItem) GetSitelinkFeedItem() *common.SitelinkFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_SitelinkFeedItem); ok {
		return x.SitelinkFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetStructuredSnippetFeedItem() *common.StructuredSnippetFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_StructuredSnippetFeedItem); ok {
		return x.StructuredSnippetFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetAppFeedItem() *common.AppFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_AppFeedItem); ok {
		return x.AppFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetCallFeedItem() *common.CallFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_CallFeedItem); ok {
		return x.CallFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetCalloutFeedItem() *common.CalloutFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_CalloutFeedItem); ok {
		return x.CalloutFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetTextMessageFeedItem() *common.TextMessageFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_TextMessageFeedItem); ok {
		return x.TextMessageFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetPriceFeedItem() *common.PriceFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_PriceFeedItem); ok {
		return x.PriceFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetPromotionFeedItem() *common.PromotionFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_PromotionFeedItem); ok {
		return x.PromotionFeedItem
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ExtensionFeedItem) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ExtensionFeedItem_OneofMarshaler, _ExtensionFeedItem_OneofUnmarshaler, _ExtensionFeedItem_OneofSizer, []interface{}{
		(*ExtensionFeedItem_SitelinkFeedItem)(nil),
		(*ExtensionFeedItem_StructuredSnippetFeedItem)(nil),
		(*ExtensionFeedItem_AppFeedItem)(nil),
		(*ExtensionFeedItem_CallFeedItem)(nil),
		(*ExtensionFeedItem_CalloutFeedItem)(nil),
		(*ExtensionFeedItem_TextMessageFeedItem)(nil),
		(*ExtensionFeedItem_PriceFeedItem)(nil),
		(*ExtensionFeedItem_PromotionFeedItem)(nil),
	}
}

func _ExtensionFeedItem_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ExtensionFeedItem)
	// extension
	switch x := m.Extension.(type) {
	case *ExtensionFeedItem_SitelinkFeedItem:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SitelinkFeedItem); err != nil {
			return err
		}
	case *ExtensionFeedItem_StructuredSnippetFeedItem:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StructuredSnippetFeedItem); err != nil {
			return err
		}
	case *ExtensionFeedItem_AppFeedItem:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AppFeedItem); err != nil {
			return err
		}
	case *ExtensionFeedItem_CallFeedItem:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CallFeedItem); err != nil {
			return err
		}
	case *ExtensionFeedItem_CalloutFeedItem:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CalloutFeedItem); err != nil {
			return err
		}
	case *ExtensionFeedItem_TextMessageFeedItem:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TextMessageFeedItem); err != nil {
			return err
		}
	case *ExtensionFeedItem_PriceFeedItem:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PriceFeedItem); err != nil {
			return err
		}
	case *ExtensionFeedItem_PromotionFeedItem:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PromotionFeedItem); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ExtensionFeedItem.Extension has unexpected type %T", x)
	}
	return nil
}

func _ExtensionFeedItem_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ExtensionFeedItem)
	switch tag {
	case 2: // extension.sitelink_feed_item
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.SitelinkFeedItem)
		err := b.DecodeMessage(msg)
		m.Extension = &ExtensionFeedItem_SitelinkFeedItem{msg}
		return true, err
	case 3: // extension.structured_snippet_feed_item
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.StructuredSnippetFeedItem)
		err := b.DecodeMessage(msg)
		m.Extension = &ExtensionFeedItem_StructuredSnippetFeedItem{msg}
		return true, err
	case 7: // extension.app_feed_item
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.AppFeedItem)
		err := b.DecodeMessage(msg)
		m.Extension = &ExtensionFeedItem_AppFeedItem{msg}
		return true, err
	case 8: // extension.call_feed_item
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.CallFeedItem)
		err := b.DecodeMessage(msg)
		m.Extension = &ExtensionFeedItem_CallFeedItem{msg}
		return true, err
	case 9: // extension.callout_feed_item
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.CalloutFeedItem)
		err := b.DecodeMessage(msg)
		m.Extension = &ExtensionFeedItem_CalloutFeedItem{msg}
		return true, err
	case 10: // extension.text_message_feed_item
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.TextMessageFeedItem)
		err := b.DecodeMessage(msg)
		m.Extension = &ExtensionFeedItem_TextMessageFeedItem{msg}
		return true, err
	case 11: // extension.price_feed_item
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.PriceFeedItem)
		err := b.DecodeMessage(msg)
		m.Extension = &ExtensionFeedItem_PriceFeedItem{msg}
		return true, err
	case 12: // extension.promotion_feed_item
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.PromotionFeedItem)
		err := b.DecodeMessage(msg)
		m.Extension = &ExtensionFeedItem_PromotionFeedItem{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ExtensionFeedItem_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ExtensionFeedItem)
	// extension
	switch x := m.Extension.(type) {
	case *ExtensionFeedItem_SitelinkFeedItem:
		s := proto.Size(x.SitelinkFeedItem)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ExtensionFeedItem_StructuredSnippetFeedItem:
		s := proto.Size(x.StructuredSnippetFeedItem)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ExtensionFeedItem_AppFeedItem:
		s := proto.Size(x.AppFeedItem)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ExtensionFeedItem_CallFeedItem:
		s := proto.Size(x.CallFeedItem)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ExtensionFeedItem_CalloutFeedItem:
		s := proto.Size(x.CalloutFeedItem)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ExtensionFeedItem_TextMessageFeedItem:
		s := proto.Size(x.TextMessageFeedItem)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ExtensionFeedItem_PriceFeedItem:
		s := proto.Size(x.PriceFeedItem)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ExtensionFeedItem_PromotionFeedItem:
		s := proto.Size(x.PromotionFeedItem)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ExtensionFeedItem)(nil), "google.ads.googleads.v1.resources.ExtensionFeedItem")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/extension_feed_item.proto", fileDescriptor_extension_feed_item_ecd9cb7f4df90d10)
}

var fileDescriptor_extension_feed_item_ecd9cb7f4df90d10 = []byte{
	// 672 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0xdf, 0x4e, 0x13, 0x4f,
	0x14, 0xc7, 0x7f, 0x2d, 0x3f, 0xd1, 0x4e, 0x29, 0xc8, 0x90, 0x90, 0x4a, 0x88, 0x01, 0x0d, 0x09,
	0x89, 0xba, 0x6b, 0xc1, 0x1b, 0x97, 0xc4, 0x50, 0x04, 0x51, 0x13, 0x0d, 0xb6, 0x4d, 0x35, 0xa6,
	0x66, 0x1d, 0x76, 0x0f, 0x9b, 0xd5, 0x9d, 0x3f, 0xd9, 0x99, 0x45, 0x48, 0xbc, 0xf3, 0x45, 0x8c,
	0x97, 0x3e, 0x8a, 0x8f, 0xe2, 0x53, 0x98, 0xce, 0x76, 0xb6, 0x53, 0x48, 0x59, 0xee, 0x66, 0xce,
	0x9c, 0xef, 0xe7, 0x7b, 0xce, 0x49, 0xf7, 0x14, 0xed, 0x44, 0x9c, 0x47, 0x09, 0xb8, 0x24, 0x94,
	0x6e, 0x7e, 0x1c, 0x9e, 0x4e, 0x5b, 0x6e, 0x0a, 0x92, 0x67, 0x69, 0x00, 0xd2, 0x85, 0x33, 0x05,
	0x4c, 0xc6, 0x9c, 0xf9, 0x27, 0x00, 0xa1, 0x1f, 0x2b, 0xa0, 0x8e, 0x48, 0xb9, 0xe2, 0x78, 0x3d,
	0x57, 0x38, 0x24, 0x94, 0x4e, 0x21, 0x76, 0x4e, 0x5b, 0x4e, 0x21, 0x5e, 0x71, 0xa7, 0xf1, 0x03,
	0x4e, 0x29, 0x67, 0x63, 0xb8, 0xcc, 0x99, 0x2b, 0x5b, 0xd3, 0x04, 0xc0, 0x32, 0x6a, 0x17, 0xa3,
	0xce, 0x05, 0x8c, 0x34, 0x4f, 0xae, 0xd6, 0x14, 0x65, 0xfb, 0x52, 0x11, 0x95, 0x19, 0xa7, 0xbb,
	0x23, 0x95, 0xbe, 0x1d, 0x67, 0x27, 0xee, 0xb7, 0x94, 0x08, 0x01, 0xa9, 0x79, 0x5f, 0x35, 0x54,
	0x11, 0xbb, 0x84, 0x31, 0xae, 0x88, 0x1a, 0xd7, 0x79, 0xef, 0x67, 0x0d, 0x2d, 0x1e, 0x98, 0x62,
	0x5e, 0x00, 0x84, 0xaf, 0x14, 0x50, 0x7c, 0x1f, 0x35, 0x4c, 0xef, 0x3e, 0x23, 0x14, 0x9a, 0x95,
	0xb5, 0xca, 0x66, 0xad, 0x33, 0x67, 0x82, 0x6f, 0x09, 0x05, 0x0c, 0x68, 0x7e, 0xb2, 0x8d, 0x66,
	0x63, 0xad, 0xb2, 0x39, 0xbf, 0xf5, 0xcc, 0x99, 0x36, 0x4f, 0xdd, 0x87, 0x53, 0xd8, 0xf5, 0xce,
	0x05, 0x1c, 0xb0, 0x8c, 0x4e, 0x46, 0x3a, 0x0d, 0xb0, 0xaf, 0x78, 0x1f, 0x2d, 0x48, 0x45, 0x52,
	0xe5, 0x87, 0x44, 0x81, 0xaf, 0x62, 0x0a, 0xcd, 0x1b, 0x6b, 0x95, 0xcd, 0xfa, 0xd6, 0xaa, 0xf1,
	0x31, 0x9d, 0x3b, 0x5d, 0x95, 0xc6, 0x2c, 0xea, 0x93, 0x24, 0x83, 0x4e, 0x43, 0x8b, 0xf6, 0x89,
	0x82, 0x5e, 0x4c, 0x01, 0xef, 0xa2, 0x06, 0xb0, 0xd0, 0x62, 0xcc, 0x5e, 0x83, 0x51, 0x07, 0x16,
	0x16, 0x84, 0x0f, 0x68, 0x36, 0x9f, 0x7b, 0xf3, 0x7f, 0xdd, 0xe6, 0x6e, 0x49, 0x9b, 0x66, 0x98,
	0x5d, 0x2d, 0xd2, 0x7d, 0x4e, 0x86, 0x3a, 0x23, 0x1e, 0xfe, 0x8c, 0xb0, 0x8c, 0x15, 0x24, 0x31,
	0xfb, 0x3a, 0xfe, 0x6d, 0x36, 0xab, 0xba, 0xc0, 0xc7, 0x53, 0x5d, 0xf2, 0x5f, 0x9e, 0xd3, 0x1d,
	0x29, 0x0d, 0xfb, 0xe5, 0x7f, 0x9d, 0xdb, 0xf2, 0x42, 0x0c, 0x7f, 0x47, 0xab, 0x52, 0xa5, 0x59,
	0xa0, 0xb2, 0x14, 0x42, 0x5f, 0xb2, 0x58, 0x08, 0x50, 0x96, 0xd7, 0x8c, 0xf6, 0x7a, 0x5a, 0xea,
	0x55, 0x30, 0xba, 0x39, 0xc2, 0x32, 0xbd, 0x23, 0xa7, 0x3d, 0xe2, 0x77, 0xa8, 0x41, 0x84, 0xb0,
	0xec, 0x6e, 0x6a, 0xbb, 0x07, 0x65, 0x76, 0x6d, 0x21, 0x2c, 0x83, 0x3a, 0x19, 0x5f, 0x71, 0x0f,
	0xcd, 0x07, 0x24, 0x49, 0x2c, 0xe6, 0x2d, 0xcd, 0x7c, 0x58, 0xc6, 0x7c, 0x4e, 0x92, 0xc4, 0x82,
	0xce, 0x05, 0xd6, 0x1d, 0x7f, 0x42, 0x8b, 0xc3, 0x3b, 0xcf, 0xec, 0xd9, 0xd4, 0x34, 0xd8, 0xbd,
	0x0e, 0x98, 0x67, 0xf6, 0x44, 0x16, 0x82, 0xc9, 0x10, 0xfe, 0x82, 0x96, 0x15, 0x9c, 0x29, 0x9f,
	0x82, 0x94, 0x24, 0x02, 0xcb, 0x03, 0x69, 0x8f, 0xed, 0x32, 0x8f, 0x1e, 0x9c, 0xa9, 0x37, 0xb9,
	0xd8, 0xf2, 0x59, 0x52, 0x97, 0xc3, 0xf8, 0x3d, 0x5a, 0x10, 0x69, 0x1c, 0xd8, 0x26, 0x75, 0x6d,
	0xf2, 0xa8, 0xcc, 0xe4, 0x68, 0x28, 0xb3, 0xf0, 0x0d, 0x61, 0x07, 0x70, 0x80, 0x96, 0x44, 0xca,
	0x29, 0x57, 0x13, 0x9b, 0xb4, 0x39, 0xa7, 0xe1, 0xad, 0x72, 0xf8, 0x48, 0x6a, 0x19, 0x2c, 0x8a,
	0x8b, 0xc1, 0xbd, 0x3a, 0xaa, 0x15, 0x4b, 0x60, 0xef, 0x47, 0x15, 0x6d, 0x04, 0x9c, 0x3a, 0xa5,
	0x5b, 0x7a, 0x6f, 0xf9, 0xd2, 0x26, 0x3b, 0x1a, 0x7e, 0xd6, 0x47, 0x95, 0x8f, 0xaf, 0x47, 0xe2,
	0x88, 0x27, 0x84, 0x45, 0x0e, 0x4f, 0x23, 0x37, 0x02, 0xa6, 0x3f, 0x7a, 0xb3, 0x6a, 0x45, 0x2c,
	0xaf, 0xf8, 0xfb, 0xd8, 0x29, 0x4e, 0xbf, 0xaa, 0x33, 0x87, 0xed, 0xf6, 0xef, 0xea, 0xfa, 0x61,
	0x8e, 0x6c, 0x87, 0xd2, 0xc9, 0x8f, 0xc3, 0x53, 0xbf, 0xe5, 0x74, 0x4c, 0xe6, 0x1f, 0x93, 0x33,
	0x68, 0x87, 0x72, 0x50, 0xe4, 0x0c, 0xfa, 0xad, 0x41, 0x91, 0xf3, 0xb7, 0xba, 0x91, 0x3f, 0x78,
	0x5e, 0x3b, 0x94, 0x9e, 0x57, 0x64, 0x79, 0x5e, 0xbf, 0xe5, 0x79, 0x45, 0xde, 0xf1, 0xac, 0x2e,
	0x76, 0xfb, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc4, 0xbd, 0x73, 0x03, 0xea, 0x06, 0x00, 0x00,
}
