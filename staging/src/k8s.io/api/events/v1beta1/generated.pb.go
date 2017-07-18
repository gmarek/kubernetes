/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo.
// source: k8s.io/kubernetes/vendor/k8s.io/api/events/v1beta1/generated.proto
// DO NOT EDIT!

/*
	Package v1beta1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/vendor/k8s.io/api/events/v1beta1/generated.proto

	It has these top-level messages:
		Event
		EventAction
		EventSeries
		EventSource
*/
package v1beta1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import k8s_io_api_core_v1 "k8s.io/api/core/v1"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *Event) Reset()                    { *m = Event{} }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *EventAction) Reset()                    { *m = EventAction{} }
func (*EventAction) ProtoMessage()               {}
func (*EventAction) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *EventSeries) Reset()                    { *m = EventSeries{} }
func (*EventSeries) ProtoMessage()               {}
func (*EventSeries) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func (m *EventSource) Reset()                    { *m = EventSource{} }
func (*EventSource) ProtoMessage()               {}
func (*EventSource) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{3} }

func init() {
	proto.RegisterType((*Event)(nil), "k8s.io.api.events.v1beta1.Event")
	proto.RegisterType((*EventAction)(nil), "k8s.io.api.events.v1beta1.EventAction")
	proto.RegisterType((*EventSeries)(nil), "k8s.io.api.events.v1beta1.EventSeries")
	proto.RegisterType((*EventSource)(nil), "k8s.io.api.events.v1beta1.EventSource")
}
func (m *Event) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Event) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ObjectMeta.Size()))
	n1, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.EventTime.Size()))
	n2, err := m.EventTime.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if m.Series != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Series.Size()))
		n3, err := m.Series.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	dAtA[i] = 0x22
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Origin.Size()))
	n4, err := m.Origin.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	dAtA[i] = 0x2a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Action.Size()))
	n5, err := m.Action.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	if m.Object != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Object.Size()))
		n6, err := m.Object.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.SecondaryObject != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.SecondaryObject.Size()))
		n7, err := m.SecondaryObject.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	dAtA[i] = 0x42
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Severity)))
	i += copy(dAtA[i:], m.Severity)
	dAtA[i] = 0x4a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Message)))
	i += copy(dAtA[i:], m.Message)
	return i, nil
}

func (m *EventAction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventAction) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Action)))
	i += copy(dAtA[i:], m.Action)
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Reason)))
	i += copy(dAtA[i:], m.Reason)
	return i, nil
}

func (m *EventSeries) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSeries) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.UID)))
	i += copy(dAtA[i:], m.UID)
	dAtA[i] = 0x10
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Count))
	dAtA[i] = 0x1a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.LastObservedTime.Size()))
	n8, err := m.LastObservedTime.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n8
	dAtA[i] = 0x22
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.LastHeartbeat.Size()))
	n9, err := m.LastHeartbeat.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n9
	dAtA[i] = 0x28
	i++
	if m.FinishMarker {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	return i, nil
}

func (m *EventSource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSource) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Component)))
	i += copy(dAtA[i:], m.Component)
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.ID)))
	i += copy(dAtA[i:], m.ID)
	return i, nil
}

func encodeFixed64Generated(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Generated(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Event) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.EventTime.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if m.Series != nil {
		l = m.Series.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = m.Origin.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Action.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if m.Object != nil {
		l = m.Object.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.SecondaryObject != nil {
		l = m.SecondaryObject.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.Severity)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Message)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *EventAction) Size() (n int) {
	var l int
	_ = l
	l = len(m.Action)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Reason)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *EventSeries) Size() (n int) {
	var l int
	_ = l
	l = len(m.UID)
	n += 1 + l + sovGenerated(uint64(l))
	n += 1 + sovGenerated(uint64(m.Count))
	l = m.LastObservedTime.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.LastHeartbeat.Size()
	n += 1 + l + sovGenerated(uint64(l))
	n += 2
	return n
}

func (m *EventSource) Size() (n int) {
	var l int
	_ = l
	l = len(m.Component)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.ID)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Event) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Event{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`EventTime:` + strings.Replace(strings.Replace(this.EventTime.String(), "MicroTime", "k8s_io_apimachinery_pkg_apis_meta_v1.MicroTime", 1), `&`, ``, 1) + `,`,
		`Series:` + strings.Replace(fmt.Sprintf("%v", this.Series), "EventSeries", "EventSeries", 1) + `,`,
		`Origin:` + strings.Replace(strings.Replace(this.Origin.String(), "EventSource", "EventSource", 1), `&`, ``, 1) + `,`,
		`Action:` + strings.Replace(strings.Replace(this.Action.String(), "EventAction", "EventAction", 1), `&`, ``, 1) + `,`,
		`Object:` + strings.Replace(fmt.Sprintf("%v", this.Object), "ObjectReference", "k8s_io_api_core_v1.ObjectReference", 1) + `,`,
		`SecondaryObject:` + strings.Replace(fmt.Sprintf("%v", this.SecondaryObject), "ObjectReference", "k8s_io_api_core_v1.ObjectReference", 1) + `,`,
		`Severity:` + fmt.Sprintf("%v", this.Severity) + `,`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`}`,
	}, "")
	return s
}
func (this *EventAction) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&EventAction{`,
		`Action:` + fmt.Sprintf("%v", this.Action) + `,`,
		`Reason:` + fmt.Sprintf("%v", this.Reason) + `,`,
		`}`,
	}, "")
	return s
}
func (this *EventSeries) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&EventSeries{`,
		`UID:` + fmt.Sprintf("%v", this.UID) + `,`,
		`Count:` + fmt.Sprintf("%v", this.Count) + `,`,
		`LastObservedTime:` + strings.Replace(strings.Replace(this.LastObservedTime.String(), "MicroTime", "k8s_io_apimachinery_pkg_apis_meta_v1.MicroTime", 1), `&`, ``, 1) + `,`,
		`LastHeartbeat:` + strings.Replace(strings.Replace(this.LastHeartbeat.String(), "MicroTime", "k8s_io_apimachinery_pkg_apis_meta_v1.MicroTime", 1), `&`, ``, 1) + `,`,
		`FinishMarker:` + fmt.Sprintf("%v", this.FinishMarker) + `,`,
		`}`,
	}, "")
	return s
}
func (this *EventSource) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&EventSource{`,
		`Component:` + fmt.Sprintf("%v", this.Component) + `,`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Event) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: Event: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Event: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EventTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Series", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Series == nil {
				m.Series = &EventSeries{}
			}
			if err := m.Series.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Origin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Origin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Action.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Object", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Object == nil {
				m.Object = &k8s_io_api_core_v1.ObjectReference{}
			}
			if err := m.Object.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecondaryObject", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SecondaryObject == nil {
				m.SecondaryObject = &k8s_io_api_core_v1.ObjectReference{}
			}
			if err := m.SecondaryObject.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Severity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Severity = EventSeverity(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *EventAction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: EventAction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventAction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Action = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *EventSeries) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: EventSeries: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSeries: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastObservedTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LastObservedTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastHeartbeat", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LastHeartbeat.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinishMarker", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
			m.FinishMarker = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *EventSource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: EventSource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Component", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Component = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
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
				next, err := skipGenerated(dAtA[start:])
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
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/api/events/v1beta1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 738 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x4f, 0x8f, 0xe3, 0x34,
	0x18, 0xc6, 0x9b, 0x96, 0xfe, 0x73, 0xb7, 0xcc, 0x62, 0x40, 0x0a, 0x95, 0x48, 0x97, 0xae, 0x34,
	0x5a, 0x0e, 0x9b, 0x30, 0x80, 0xd0, 0x5e, 0x38, 0x90, 0xdd, 0x05, 0x16, 0x6d, 0x19, 0xc9, 0x05,
	0x09, 0xad, 0x84, 0x34, 0x4e, 0xf2, 0x4e, 0x6a, 0x3a, 0xb1, 0x2b, 0xdb, 0xa9, 0x34, 0x37, 0x3e,
	0x02, 0x1f, 0x6b, 0x8e, 0x73, 0x9c, 0x53, 0xc5, 0x84, 0x2f, 0xc0, 0x99, 0x13, 0x8a, 0xe3, 0xfe,
	0x99, 0x29, 0x15, 0x85, 0x5b, 0xfc, 0xbc, 0xcf, 0xf3, 0x7b, 0xeb, 0xd7, 0x76, 0x51, 0x38, 0x7b,
	0xa6, 0x7c, 0x26, 0x82, 0x59, 0x1e, 0x81, 0xe4, 0xa0, 0x41, 0x05, 0x0b, 0xe0, 0x89, 0x90, 0x81,
	0x2d, 0xd0, 0x39, 0x0b, 0x60, 0x01, 0x5c, 0xab, 0x60, 0x71, 0x12, 0x81, 0xa6, 0x27, 0x41, 0x0a,
	0x1c, 0x24, 0xd5, 0x90, 0xf8, 0x73, 0x29, 0xb4, 0xc0, 0x1f, 0x54, 0x56, 0x9f, 0xce, 0x99, 0x5f,
	0x59, 0x7d, 0x6b, 0x1d, 0x3c, 0x4d, 0x99, 0x9e, 0xe6, 0x91, 0x1f, 0x8b, 0x2c, 0x48, 0x45, 0x2a,
	0x02, 0x93, 0x88, 0xf2, 0x73, 0xb3, 0x32, 0x0b, 0xf3, 0x55, 0x91, 0x06, 0xa3, 0xad, 0xa6, 0xb1,
	0x90, 0x10, 0x2c, 0x76, 0xba, 0x0d, 0x3e, 0xdf, 0x78, 0x32, 0x1a, 0x4f, 0x19, 0x07, 0x79, 0x19,
	0xcc, 0x67, 0x69, 0x29, 0xa8, 0x20, 0x03, 0x4d, 0xff, 0x29, 0x15, 0xec, 0x4b, 0xc9, 0x9c, 0x6b,
	0x96, 0xc1, 0x4e, 0xe0, 0x8b, 0x7f, 0x0b, 0xa8, 0x78, 0x0a, 0x19, 0xdd, 0xc9, 0x7d, 0xb6, 0x2f,
	0x97, 0x6b, 0x76, 0x11, 0x30, 0xae, 0x95, 0x96, 0xf7, 0x43, 0xa3, 0x9b, 0x26, 0x6a, 0xbe, 0x2c,
	0x27, 0x87, 0xcf, 0x50, 0xa7, 0xdc, 0x42, 0x42, 0x35, 0x75, 0x9d, 0x47, 0xce, 0x93, 0xde, 0xa7,
	0x9f, 0xf8, 0x9b, 0xf1, 0xae, 0x89, 0xfe, 0x7c, 0x96, 0x96, 0x82, 0xf2, 0x4b, 0xb7, 0xbf, 0x38,
	0xf1, 0x4f, 0xa3, 0x5f, 0x20, 0xd6, 0x63, 0xd0, 0x34, 0xc4, 0x57, 0xcb, 0x61, 0xad, 0x58, 0x0e,
	0xd1, 0x46, 0x23, 0x6b, 0x2a, 0x3e, 0x43, 0x5d, 0x73, 0x48, 0x3f, 0xb0, 0x0c, 0xdc, 0xba, 0x69,
	0x11, 0x1c, 0xd6, 0x62, 0xcc, 0x62, 0x29, 0xca, 0x58, 0xf8, 0x8e, 0xed, 0xd0, 0x7d, 0xb9, 0x22,
	0x91, 0x0d, 0x14, 0x7f, 0x87, 0x5a, 0x0a, 0x24, 0x03, 0xe5, 0x36, 0x0c, 0xfe, 0xd8, 0xdf, 0x7b,
	0x41, 0x7c, 0x03, 0x98, 0x18, 0x77, 0x88, 0x8a, 0xe5, 0xb0, 0x55, 0x7d, 0x13, 0x4b, 0xc0, 0x3f,
	0xa1, 0x9e, 0x49, 0x4c, 0x44, 0x2e, 0x63, 0x70, 0xdf, 0x3a, 0x10, 0x68, 0xdc, 0xe1, 0xdb, 0xf6,
	0x67, 0xb6, 0x4e, 0x25, 0x4b, 0x19, 0x27, 0xdb, 0x28, 0xfc, 0x3d, 0x6a, 0xd1, 0x58, 0x33, 0xc1,
	0xdd, 0xe6, 0x61, 0xd0, 0xaf, 0x8c, 0x7b, 0x03, 0xad, 0xd6, 0xc4, 0x52, 0xf0, 0x37, 0xa8, 0x25,
	0xcc, 0xbc, 0xdd, 0x96, 0xe1, 0x3d, 0xde, 0xe6, 0x95, 0x97, 0x79, 0x73, 0x4a, 0x04, 0xce, 0x41,
	0x02, 0x8f, 0xa1, 0xda, 0xb2, 0x15, 0x6d, 0x1c, 0x47, 0xe8, 0x48, 0x41, 0x2c, 0x78, 0x42, 0xe5,
	0x65, 0x55, 0x72, 0xdb, 0x87, 0x13, 0xdf, 0x2d, 0x96, 0xc3, 0xa3, 0xc9, 0xdd, 0x3c, 0xb9, 0x0f,
	0xc4, 0x5f, 0xa2, 0x8e, 0x82, 0x05, 0x48, 0xa6, 0x2f, 0xdd, 0xce, 0x23, 0xe7, 0x49, 0x37, 0xfc,
	0xc8, 0x6e, 0xab, 0x33, 0xb1, 0xfa, 0x5f, 0xcb, 0x61, 0xdf, 0x9e, 0x4e, 0x25, 0x90, 0x75, 0x04,
	0x7f, 0x8c, 0xda, 0x19, 0x28, 0x45, 0x53, 0x70, 0xbb, 0x26, 0x7d, 0x64, 0xd3, 0xed, 0x71, 0x25,
	0x93, 0x55, 0x7d, 0xf4, 0x33, 0xea, 0x6d, 0x4d, 0x0f, 0x1f, 0xaf, 0xa7, 0xee, 0x98, 0xe0, 0xbe,
	0x69, 0x1e, 0xa3, 0x96, 0x04, 0xaa, 0x04, 0x37, 0x57, 0x74, 0xcb, 0x47, 0x8c, 0x4a, 0x6c, 0x75,
	0xf4, 0x67, 0xdd, 0xf2, 0xab, 0x7b, 0x83, 0x3f, 0x44, 0x8d, 0x9c, 0x25, 0x16, 0xde, 0xb3, 0xa1,
	0xc6, 0x8f, 0xaf, 0x5e, 0x90, 0x52, 0xc7, 0x8f, 0x51, 0x33, 0x16, 0x39, 0xd7, 0x86, 0xda, 0x0c,
	0xfb, 0xd6, 0xd0, 0x7c, 0x5e, 0x8a, 0xa4, 0xaa, 0xe1, 0x1c, 0x3d, 0xbc, 0xa0, 0x4a, 0x9f, 0x46,
	0x0a, 0xe4, 0x02, 0x12, 0xf3, 0x50, 0x1a, 0xff, 0xef, 0xa1, 0xb8, 0xb6, 0xc1, 0xc3, 0xd7, 0xf7,
	0x80, 0x64, 0xa7, 0x05, 0xbe, 0x40, 0xfd, 0x52, 0xfb, 0x16, 0xa8, 0xd4, 0x11, 0x50, 0x6d, 0x2f,
	0xfb, 0x7f, 0xee, 0xf9, 0xbe, 0xed, 0xd9, 0x7f, 0xbd, 0x4d, 0x23, 0x77, 0xe1, 0xf8, 0x19, 0x7a,
	0x70, 0xce, 0x38, 0x53, 0xd3, 0x31, 0x95, 0x33, 0x90, 0xe6, 0x11, 0x74, 0xc2, 0xf7, 0x6c, 0xf6,
	0xc1, 0xd7, 0x5b, 0x35, 0x72, 0xc7, 0x39, 0x7a, 0xb3, 0x9a, 0x78, 0xf5, 0x8e, 0x02, 0xd4, 0x8d,
	0x45, 0x36, 0x17, 0x1c, 0xb8, 0xb6, 0x73, 0x5f, 0xff, 0x3d, 0x3c, 0x5f, 0x15, 0xc8, 0xc6, 0x83,
	0x07, 0xa8, 0xce, 0x12, 0x7b, 0xac, 0xc8, 0x3a, 0xeb, 0xaf, 0x5e, 0x90, 0x3a, 0x4b, 0xc2, 0xa7,
	0x57, 0xb7, 0x5e, 0xed, 0xfa, 0xd6, 0xab, 0xdd, 0xdc, 0x7a, 0xb5, 0x5f, 0x0b, 0xcf, 0xb9, 0x2a,
	0x3c, 0xe7, 0xba, 0xf0, 0x9c, 0x9b, 0xc2, 0x73, 0x7e, 0x2f, 0x3c, 0xe7, 0xb7, 0x3f, 0xbc, 0xda,
	0x9b, 0xb6, 0x7d, 0x97, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x07, 0xb5, 0xb2, 0xbe, 0x06,
	0x00, 0x00,
}
