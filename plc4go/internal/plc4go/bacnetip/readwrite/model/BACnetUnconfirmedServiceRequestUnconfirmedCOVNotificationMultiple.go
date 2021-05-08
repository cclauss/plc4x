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
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple struct {
	Parent *BACnetUnconfirmedServiceRequest
}

// The corresponding interface
type IBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) ServiceChoice() uint8 {
	return 0x0B
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) InitializeParent(parent *BACnetUnconfirmedServiceRequest) {
}

func NewBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple() *BACnetUnconfirmedServiceRequest {
	child := &BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple{
		Parent: NewBACnetUnconfirmedServiceRequest(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple(structType interface{}) *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple {
	castFunc := func(typ interface{}) *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple {
		if casted, ok := typ.(BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple); ok {
			return casted
		}
		if casted, ok := typ.(BACnetUnconfirmedServiceRequest); ok {
			return CastBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple(casted.Child)
		}
		if casted, ok := typ.(*BACnetUnconfirmedServiceRequest); ok {
			return CastBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple"
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleParse(readBuffer utils.ReadBuffer) (*BACnetUnconfirmedServiceRequest, error) {
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple{
		Parent: &BACnetUnconfirmedServiceRequest{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
