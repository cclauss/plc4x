/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetConfirmedServiceACKGetEventInformation struct {
	*BACnetConfirmedServiceACK
}

// The corresponding interface
type IBACnetConfirmedServiceACKGetEventInformation interface {
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetConfirmedServiceACKGetEventInformation) ServiceChoice() uint8 {
	return 0x1D
}

func (m *BACnetConfirmedServiceACKGetEventInformation) GetServiceChoice() uint8 {
	return 0x1D
}

func (m *BACnetConfirmedServiceACKGetEventInformation) InitializeParent(parent *BACnetConfirmedServiceACK) {
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceACKGetEventInformation factory function for BACnetConfirmedServiceACKGetEventInformation
func NewBACnetConfirmedServiceACKGetEventInformation() *BACnetConfirmedServiceACK {
	child := &BACnetConfirmedServiceACKGetEventInformation{
		BACnetConfirmedServiceACK: NewBACnetConfirmedServiceACK(),
	}
	child.Child = child
	return child.BACnetConfirmedServiceACK
}

func CastBACnetConfirmedServiceACKGetEventInformation(structType interface{}) *BACnetConfirmedServiceACKGetEventInformation {
	castFunc := func(typ interface{}) *BACnetConfirmedServiceACKGetEventInformation {
		if casted, ok := typ.(BACnetConfirmedServiceACKGetEventInformation); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetConfirmedServiceACKGetEventInformation); ok {
			return casted
		}
		if casted, ok := typ.(BACnetConfirmedServiceACK); ok {
			return CastBACnetConfirmedServiceACKGetEventInformation(casted.Child)
		}
		if casted, ok := typ.(*BACnetConfirmedServiceACK); ok {
			return CastBACnetConfirmedServiceACKGetEventInformation(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetConfirmedServiceACKGetEventInformation) GetTypeName() string {
	return "BACnetConfirmedServiceACKGetEventInformation"
}

func (m *BACnetConfirmedServiceACKGetEventInformation) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceACKGetEventInformation) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *BACnetConfirmedServiceACKGetEventInformation) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceACKGetEventInformationParse(readBuffer utils.ReadBuffer) (*BACnetConfirmedServiceACK, error) {
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceACKGetEventInformation"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceACKGetEventInformation"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceACKGetEventInformation{
		BACnetConfirmedServiceACK: &BACnetConfirmedServiceACK{},
	}
	_child.BACnetConfirmedServiceACK.Child = _child
	return _child.BACnetConfirmedServiceACK, nil
}

func (m *BACnetConfirmedServiceACKGetEventInformation) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceACKGetEventInformation"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceACKGetEventInformation"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceACKGetEventInformation) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
