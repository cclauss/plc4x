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
type ApduDataExtIndividualAddressSerialNumberResponse struct {
	*ApduDataExt

	// Arguments.
	Length uint8
}

// The corresponding interface
type IApduDataExtIndividualAddressSerialNumberResponse interface {
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
func (m *ApduDataExtIndividualAddressSerialNumberResponse) ExtApciType() uint8 {
	return 0x1D
}

func (m *ApduDataExtIndividualAddressSerialNumberResponse) GetExtApciType() uint8 {
	return 0x1D
}

func (m *ApduDataExtIndividualAddressSerialNumberResponse) InitializeParent(parent *ApduDataExt) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewApduDataExtIndividualAddressSerialNumberResponse factory function for ApduDataExtIndividualAddressSerialNumberResponse
func NewApduDataExtIndividualAddressSerialNumberResponse(length uint8) *ApduDataExt {
	child := &ApduDataExtIndividualAddressSerialNumberResponse{
		ApduDataExt: NewApduDataExt(length),
	}
	child.Child = child
	return child.ApduDataExt
}

func CastApduDataExtIndividualAddressSerialNumberResponse(structType interface{}) *ApduDataExtIndividualAddressSerialNumberResponse {
	castFunc := func(typ interface{}) *ApduDataExtIndividualAddressSerialNumberResponse {
		if casted, ok := typ.(ApduDataExtIndividualAddressSerialNumberResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataExtIndividualAddressSerialNumberResponse); ok {
			return casted
		}
		if casted, ok := typ.(ApduDataExt); ok {
			return CastApduDataExtIndividualAddressSerialNumberResponse(casted.Child)
		}
		if casted, ok := typ.(*ApduDataExt); ok {
			return CastApduDataExtIndividualAddressSerialNumberResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataExtIndividualAddressSerialNumberResponse) GetTypeName() string {
	return "ApduDataExtIndividualAddressSerialNumberResponse"
}

func (m *ApduDataExtIndividualAddressSerialNumberResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataExtIndividualAddressSerialNumberResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataExtIndividualAddressSerialNumberResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtIndividualAddressSerialNumberResponseParse(readBuffer utils.ReadBuffer, length uint8) (*ApduDataExt, error) {
	if pullErr := readBuffer.PullContext("ApduDataExtIndividualAddressSerialNumberResponse"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("ApduDataExtIndividualAddressSerialNumberResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataExtIndividualAddressSerialNumberResponse{
		ApduDataExt: &ApduDataExt{},
	}
	_child.ApduDataExt.Child = _child
	return _child.ApduDataExt, nil
}

func (m *ApduDataExtIndividualAddressSerialNumberResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtIndividualAddressSerialNumberResponse"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("ApduDataExtIndividualAddressSerialNumberResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataExtIndividualAddressSerialNumberResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
