//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type AdsMultiRequestItemWrite struct {
	ItemIndexGroup  uint32
	ItemIndexOffset uint32
	ItemWriteLength uint32
	Parent          *AdsMultiRequestItem
}

// The corresponding interface
type IAdsMultiRequestItemWrite interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *AdsMultiRequestItemWrite) IndexGroup() uint32 {
	return 61569
}

func (m *AdsMultiRequestItemWrite) InitializeParent(parent *AdsMultiRequestItem) {
}

func NewAdsMultiRequestItemWrite(itemIndexGroup uint32, itemIndexOffset uint32, itemWriteLength uint32) *AdsMultiRequestItem {
	child := &AdsMultiRequestItemWrite{
		ItemIndexGroup:  itemIndexGroup,
		ItemIndexOffset: itemIndexOffset,
		ItemWriteLength: itemWriteLength,
		Parent:          NewAdsMultiRequestItem(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastAdsMultiRequestItemWrite(structType interface{}) *AdsMultiRequestItemWrite {
	castFunc := func(typ interface{}) *AdsMultiRequestItemWrite {
		if casted, ok := typ.(AdsMultiRequestItemWrite); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsMultiRequestItemWrite); ok {
			return casted
		}
		if casted, ok := typ.(AdsMultiRequestItem); ok {
			return CastAdsMultiRequestItemWrite(casted.Child)
		}
		if casted, ok := typ.(*AdsMultiRequestItem); ok {
			return CastAdsMultiRequestItemWrite(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsMultiRequestItemWrite) GetTypeName() string {
	return "AdsMultiRequestItemWrite"
}

func (m *AdsMultiRequestItemWrite) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AdsMultiRequestItemWrite) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (itemIndexGroup)
	lengthInBits += 32

	// Simple field (itemIndexOffset)
	lengthInBits += 32

	// Simple field (itemWriteLength)
	lengthInBits += 32

	return lengthInBits
}

func (m *AdsMultiRequestItemWrite) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AdsMultiRequestItemWriteParse(readBuffer utils.ReadBuffer) (*AdsMultiRequestItem, error) {
	if pullErr := readBuffer.PullContext("AdsMultiRequestItemWrite"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (itemIndexGroup)
	itemIndexGroup, _itemIndexGroupErr := readBuffer.ReadUint32("itemIndexGroup", 32)
	if _itemIndexGroupErr != nil {
		return nil, errors.Wrap(_itemIndexGroupErr, "Error parsing 'itemIndexGroup' field")
	}

	// Simple Field (itemIndexOffset)
	itemIndexOffset, _itemIndexOffsetErr := readBuffer.ReadUint32("itemIndexOffset", 32)
	if _itemIndexOffsetErr != nil {
		return nil, errors.Wrap(_itemIndexOffsetErr, "Error parsing 'itemIndexOffset' field")
	}

	// Simple Field (itemWriteLength)
	itemWriteLength, _itemWriteLengthErr := readBuffer.ReadUint32("itemWriteLength", 32)
	if _itemWriteLengthErr != nil {
		return nil, errors.Wrap(_itemWriteLengthErr, "Error parsing 'itemWriteLength' field")
	}

	if closeErr := readBuffer.CloseContext("AdsMultiRequestItemWrite"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &AdsMultiRequestItemWrite{
		ItemIndexGroup:  itemIndexGroup,
		ItemIndexOffset: itemIndexOffset,
		ItemWriteLength: itemWriteLength,
		Parent:          &AdsMultiRequestItem{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *AdsMultiRequestItemWrite) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsMultiRequestItemWrite"); pushErr != nil {
			return pushErr
		}

		// Simple Field (itemIndexGroup)
		itemIndexGroup := uint32(m.ItemIndexGroup)
		_itemIndexGroupErr := writeBuffer.WriteUint32("itemIndexGroup", 32, (itemIndexGroup))
		if _itemIndexGroupErr != nil {
			return errors.Wrap(_itemIndexGroupErr, "Error serializing 'itemIndexGroup' field")
		}

		// Simple Field (itemIndexOffset)
		itemIndexOffset := uint32(m.ItemIndexOffset)
		_itemIndexOffsetErr := writeBuffer.WriteUint32("itemIndexOffset", 32, (itemIndexOffset))
		if _itemIndexOffsetErr != nil {
			return errors.Wrap(_itemIndexOffsetErr, "Error serializing 'itemIndexOffset' field")
		}

		// Simple Field (itemWriteLength)
		itemWriteLength := uint32(m.ItemWriteLength)
		_itemWriteLengthErr := writeBuffer.WriteUint32("itemWriteLength", 32, (itemWriteLength))
		if _itemWriteLengthErr != nil {
			return errors.Wrap(_itemWriteLengthErr, "Error serializing 'itemWriteLength' field")
		}

		if popErr := writeBuffer.PopContext("AdsMultiRequestItemWrite"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *AdsMultiRequestItemWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
