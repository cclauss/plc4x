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
type KnxAddress struct {
	MainGroup   uint8
	MiddleGroup uint8
	SubGroup    uint8
}

// The corresponding interface
type IKnxAddress interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

func NewKnxAddress(mainGroup uint8, middleGroup uint8, subGroup uint8) *KnxAddress {
	return &KnxAddress{MainGroup: mainGroup, MiddleGroup: middleGroup, SubGroup: subGroup}
}

func CastKnxAddress(structType interface{}) *KnxAddress {
	castFunc := func(typ interface{}) *KnxAddress {
		if casted, ok := typ.(KnxAddress); ok {
			return &casted
		}
		if casted, ok := typ.(*KnxAddress); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *KnxAddress) GetTypeName() string {
	return "KnxAddress"
}

func (m *KnxAddress) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *KnxAddress) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (mainGroup)
	lengthInBits += 4

	// Simple field (middleGroup)
	lengthInBits += 4

	// Simple field (subGroup)
	lengthInBits += 8

	return lengthInBits
}

func (m *KnxAddress) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func KnxAddressParse(readBuffer utils.ReadBuffer) (*KnxAddress, error) {
	if pullErr := readBuffer.PullContext("KnxAddress"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (mainGroup)
	mainGroup, _mainGroupErr := readBuffer.ReadUint8("mainGroup", 4)
	if _mainGroupErr != nil {
		return nil, errors.Wrap(_mainGroupErr, "Error parsing 'mainGroup' field")
	}

	// Simple Field (middleGroup)
	middleGroup, _middleGroupErr := readBuffer.ReadUint8("middleGroup", 4)
	if _middleGroupErr != nil {
		return nil, errors.Wrap(_middleGroupErr, "Error parsing 'middleGroup' field")
	}

	// Simple Field (subGroup)
	subGroup, _subGroupErr := readBuffer.ReadUint8("subGroup", 8)
	if _subGroupErr != nil {
		return nil, errors.Wrap(_subGroupErr, "Error parsing 'subGroup' field")
	}

	if closeErr := readBuffer.CloseContext("KnxAddress"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewKnxAddress(mainGroup, middleGroup, subGroup), nil
}

func (m *KnxAddress) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("KnxAddress"); pushErr != nil {
		return pushErr
	}

	// Simple Field (mainGroup)
	mainGroup := uint8(m.MainGroup)
	_mainGroupErr := writeBuffer.WriteUint8("mainGroup", 4, (mainGroup))
	if _mainGroupErr != nil {
		return errors.Wrap(_mainGroupErr, "Error serializing 'mainGroup' field")
	}

	// Simple Field (middleGroup)
	middleGroup := uint8(m.MiddleGroup)
	_middleGroupErr := writeBuffer.WriteUint8("middleGroup", 4, (middleGroup))
	if _middleGroupErr != nil {
		return errors.Wrap(_middleGroupErr, "Error serializing 'middleGroup' field")
	}

	// Simple Field (subGroup)
	subGroup := uint8(m.SubGroup)
	_subGroupErr := writeBuffer.WriteUint8("subGroup", 8, (subGroup))
	if _subGroupErr != nil {
		return errors.Wrap(_subGroupErr, "Error serializing 'subGroup' field")
	}

	if popErr := writeBuffer.PopContext("KnxAddress"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *KnxAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
