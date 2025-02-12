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
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type ModbusPDUReadInputRegistersRequest struct {
	*ModbusPDU
	StartingAddress uint16
	Quantity        uint16
}

// The corresponding interface
type IModbusPDUReadInputRegistersRequest interface {
	// GetStartingAddress returns StartingAddress
	GetStartingAddress() uint16
	// GetQuantity returns Quantity
	GetQuantity() uint16
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
func (m *ModbusPDUReadInputRegistersRequest) ErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUReadInputRegistersRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUReadInputRegistersRequest) FunctionFlag() uint8 {
	return 0x04
}

func (m *ModbusPDUReadInputRegistersRequest) GetFunctionFlag() uint8 {
	return 0x04
}

func (m *ModbusPDUReadInputRegistersRequest) Response() bool {
	return bool(false)
}

func (m *ModbusPDUReadInputRegistersRequest) GetResponse() bool {
	return bool(false)
}

func (m *ModbusPDUReadInputRegistersRequest) InitializeParent(parent *ModbusPDU) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *ModbusPDUReadInputRegistersRequest) GetStartingAddress() uint16 {
	return m.StartingAddress
}

func (m *ModbusPDUReadInputRegistersRequest) GetQuantity() uint16 {
	return m.Quantity
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewModbusPDUReadInputRegistersRequest factory function for ModbusPDUReadInputRegistersRequest
func NewModbusPDUReadInputRegistersRequest(startingAddress uint16, quantity uint16) *ModbusPDU {
	child := &ModbusPDUReadInputRegistersRequest{
		StartingAddress: startingAddress,
		Quantity:        quantity,
		ModbusPDU:       NewModbusPDU(),
	}
	child.Child = child
	return child.ModbusPDU
}

func CastModbusPDUReadInputRegistersRequest(structType interface{}) *ModbusPDUReadInputRegistersRequest {
	castFunc := func(typ interface{}) *ModbusPDUReadInputRegistersRequest {
		if casted, ok := typ.(ModbusPDUReadInputRegistersRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusPDUReadInputRegistersRequest); ok {
			return casted
		}
		if casted, ok := typ.(ModbusPDU); ok {
			return CastModbusPDUReadInputRegistersRequest(casted.Child)
		}
		if casted, ok := typ.(*ModbusPDU); ok {
			return CastModbusPDUReadInputRegistersRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusPDUReadInputRegistersRequest) GetTypeName() string {
	return "ModbusPDUReadInputRegistersRequest"
}

func (m *ModbusPDUReadInputRegistersRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ModbusPDUReadInputRegistersRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (startingAddress)
	lengthInBits += 16

	// Simple field (quantity)
	lengthInBits += 16

	return lengthInBits
}

func (m *ModbusPDUReadInputRegistersRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUReadInputRegistersRequestParse(readBuffer utils.ReadBuffer, response bool) (*ModbusPDU, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUReadInputRegistersRequest"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (startingAddress)
	_startingAddress, _startingAddressErr := readBuffer.ReadUint16("startingAddress", 16)
	if _startingAddressErr != nil {
		return nil, errors.Wrap(_startingAddressErr, "Error parsing 'startingAddress' field")
	}
	startingAddress := _startingAddress

	// Simple Field (quantity)
	_quantity, _quantityErr := readBuffer.ReadUint16("quantity", 16)
	if _quantityErr != nil {
		return nil, errors.Wrap(_quantityErr, "Error parsing 'quantity' field")
	}
	quantity := _quantity

	if closeErr := readBuffer.CloseContext("ModbusPDUReadInputRegistersRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUReadInputRegistersRequest{
		StartingAddress: startingAddress,
		Quantity:        quantity,
		ModbusPDU:       &ModbusPDU{},
	}
	_child.ModbusPDU.Child = _child
	return _child.ModbusPDU, nil
}

func (m *ModbusPDUReadInputRegistersRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadInputRegistersRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (startingAddress)
		startingAddress := uint16(m.StartingAddress)
		_startingAddressErr := writeBuffer.WriteUint16("startingAddress", 16, (startingAddress))
		if _startingAddressErr != nil {
			return errors.Wrap(_startingAddressErr, "Error serializing 'startingAddress' field")
		}

		// Simple Field (quantity)
		quantity := uint16(m.Quantity)
		_quantityErr := writeBuffer.WriteUint16("quantity", 16, (quantity))
		if _quantityErr != nil {
			return errors.Wrap(_quantityErr, "Error serializing 'quantity' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadInputRegistersRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUReadInputRegistersRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
