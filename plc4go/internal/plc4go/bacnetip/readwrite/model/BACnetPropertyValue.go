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
type BACnetPropertyValue struct {
	Child IBACnetPropertyValueChild
}

// The corresponding interface
type IBACnetPropertyValue interface {
	Identifier() BACnetPropertyIdentifier
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IBACnetPropertyValueParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetPropertyValue, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetPropertyValueChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BACnetPropertyValue)
	GetTypeName() string
	IBACnetPropertyValue
}

func NewBACnetPropertyValue() *BACnetPropertyValue {
	return &BACnetPropertyValue{}
}

func CastBACnetPropertyValue(structType interface{}) *BACnetPropertyValue {
	castFunc := func(typ interface{}) *BACnetPropertyValue {
		if casted, ok := typ.(BACnetPropertyValue); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetPropertyValue); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetPropertyValue) GetTypeName() string {
	return "BACnetPropertyValue"
}

func (m *BACnetPropertyValue) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetPropertyValue) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *BACnetPropertyValue) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *BACnetPropertyValue) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetPropertyValueParse(readBuffer utils.ReadBuffer, identifier BACnetPropertyIdentifier, actualLength uint32) (*BACnetPropertyValue, error) {
	if pullErr := readBuffer.PullContext("BACnetPropertyValue"); pullErr != nil {
		return nil, pullErr
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *BACnetPropertyValue
	var typeSwitchError error
	switch {
	case identifier == BACnetPropertyIdentifier_OBJECT_TYPE: // BACnetPropertyValueObjectType
		_parent, typeSwitchError = BACnetPropertyValueObjectTypeParse(readBuffer, identifier, actualLength)
	case identifier == BACnetPropertyIdentifier_PRIORITY_ARRAY: // BACnetPropertyValuePriorityValue
		_parent, typeSwitchError = BACnetPropertyValuePriorityValueParse(readBuffer, identifier, actualLength)
	case identifier == BACnetPropertyIdentifier_PRESENT_VALUE: // BACnetPropertyValuePresentValue
		_parent, typeSwitchError = BACnetPropertyValuePresentValueParse(readBuffer, identifier, actualLength)
	case identifier == BACnetPropertyIdentifier_RELINQUISH_DEFAULT: // BACnetPropertyValueRelinquishDefault
		_parent, typeSwitchError = BACnetPropertyValueRelinquishDefaultParse(readBuffer, identifier, actualLength)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyValue"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *BACnetPropertyValue) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BACnetPropertyValue) SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetPropertyValue, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("BACnetPropertyValue"); pushErr != nil {
		return pushErr
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetPropertyValue"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetPropertyValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}