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
type ModbusPDUReadWriteMultipleHoldingRegistersResponse struct {
	*ModbusPDU
	Value []byte
}

// The corresponding interface
type IModbusPDUReadWriteMultipleHoldingRegistersResponse interface {
	// GetValue returns Value
	GetValue() []byte
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
func (m *ModbusPDUReadWriteMultipleHoldingRegistersResponse) ErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUReadWriteMultipleHoldingRegistersResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUReadWriteMultipleHoldingRegistersResponse) FunctionFlag() uint8 {
	return 0x17
}

func (m *ModbusPDUReadWriteMultipleHoldingRegistersResponse) GetFunctionFlag() uint8 {
	return 0x17
}

func (m *ModbusPDUReadWriteMultipleHoldingRegistersResponse) Response() bool {
	return bool(true)
}

func (m *ModbusPDUReadWriteMultipleHoldingRegistersResponse) GetResponse() bool {
	return bool(true)
}

func (m *ModbusPDUReadWriteMultipleHoldingRegistersResponse) InitializeParent(parent *ModbusPDU) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *ModbusPDUReadWriteMultipleHoldingRegistersResponse) GetValue() []byte {
	return m.Value
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewModbusPDUReadWriteMultipleHoldingRegistersResponse factory function for ModbusPDUReadWriteMultipleHoldingRegistersResponse
func NewModbusPDUReadWriteMultipleHoldingRegistersResponse(value []byte) *ModbusPDU {
	child := &ModbusPDUReadWriteMultipleHoldingRegistersResponse{
		Value:     value,
		ModbusPDU: NewModbusPDU(),
	}
	child.Child = child
	return child.ModbusPDU
}

func CastModbusPDUReadWriteMultipleHoldingRegistersResponse(structType interface{}) *ModbusPDUReadWriteMultipleHoldingRegistersResponse {
	castFunc := func(typ interface{}) *ModbusPDUReadWriteMultipleHoldingRegistersResponse {
		if casted, ok := typ.(ModbusPDUReadWriteMultipleHoldingRegistersResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusPDUReadWriteMultipleHoldingRegistersResponse); ok {
			return casted
		}
		if casted, ok := typ.(ModbusPDU); ok {
			return CastModbusPDUReadWriteMultipleHoldingRegistersResponse(casted.Child)
		}
		if casted, ok := typ.(*ModbusPDU); ok {
			return CastModbusPDUReadWriteMultipleHoldingRegistersResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusPDUReadWriteMultipleHoldingRegistersResponse) GetTypeName() string {
	return "ModbusPDUReadWriteMultipleHoldingRegistersResponse"
}

func (m *ModbusPDUReadWriteMultipleHoldingRegistersResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ModbusPDUReadWriteMultipleHoldingRegistersResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 8 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *ModbusPDUReadWriteMultipleHoldingRegistersResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUReadWriteMultipleHoldingRegistersResponseParse(readBuffer utils.ReadBuffer, response bool) (*ModbusPDU, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUReadWriteMultipleHoldingRegistersResponse"); pullErr != nil {
		return nil, pullErr
	}

	// Implicit Field (byteCount) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	byteCount, _byteCountErr := readBuffer.ReadUint8("byteCount", 8)
	_ = byteCount
	if _byteCountErr != nil {
		return nil, errors.Wrap(_byteCountErr, "Error parsing 'byteCount' field")
	}
	// Byte Array field (value)
	numberOfBytesvalue := int(byteCount)
	value, _readArrayErr := readBuffer.ReadByteArray("value", numberOfBytesvalue)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'value' field")
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUReadWriteMultipleHoldingRegistersResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUReadWriteMultipleHoldingRegistersResponse{
		Value:     value,
		ModbusPDU: &ModbusPDU{},
	}
	_child.ModbusPDU.Child = _child
	return _child.ModbusPDU, nil
}

func (m *ModbusPDUReadWriteMultipleHoldingRegistersResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadWriteMultipleHoldingRegistersResponse"); pushErr != nil {
			return pushErr
		}

		// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		byteCount := uint8(uint8(len(m.GetValue())))
		_byteCountErr := writeBuffer.WriteUint8("byteCount", 8, (byteCount))
		if _byteCountErr != nil {
			return errors.Wrap(_byteCountErr, "Error serializing 'byteCount' field")
		}

		// Array Field (value)
		if m.Value != nil {
			// Byte Array field (value)
			_writeArrayErr := writeBuffer.WriteByteArray("value", m.Value)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'value' field")
			}
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadWriteMultipleHoldingRegistersResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUReadWriteMultipleHoldingRegistersResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
