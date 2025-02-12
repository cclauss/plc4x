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
type BACnetContextTag struct {
	Header *BACnetTagHeader

	// Arguments.
	TagNumberArgument uint8
	Child             IBACnetContextTagChild
}

// The corresponding interface
type IBACnetContextTag interface {
	// DataType returns DataType
	DataType() BACnetDataType
	// GetHeader returns Header
	GetHeader() *BACnetTagHeader
	// GetTagNumber returns TagNumber
	GetTagNumber() uint8
	// GetActualLength returns ActualLength
	GetActualLength() uint32
	// GetIsNotOpeningOrClosingTag returns IsNotOpeningOrClosingTag
	GetIsNotOpeningOrClosingTag() bool
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IBACnetContextTagParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetContextTag, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetContextTagChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BACnetContextTag, header *BACnetTagHeader)
	GetTypeName() string
	IBACnetContextTag
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetContextTag) GetHeader() *BACnetTagHeader {
	return m.Header
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////
func (m *BACnetContextTag) GetTagNumber() uint8 {
	return m.GetHeader().GetTagNumber()
}

func (m *BACnetContextTag) GetActualLength() uint32 {
	return m.GetHeader().GetActualLength()
}

func (m *BACnetContextTag) GetIsNotOpeningOrClosingTag() bool {
	return bool(bool((m.GetHeader().GetLengthValueType()) != (6))) && bool(bool((m.GetHeader().GetLengthValueType()) != (7)))
}

// NewBACnetContextTag factory function for BACnetContextTag
func NewBACnetContextTag(header *BACnetTagHeader, tagNumberArgument uint8) *BACnetContextTag {
	return &BACnetContextTag{Header: header, TagNumberArgument: tagNumberArgument}
}

func CastBACnetContextTag(structType interface{}) *BACnetContextTag {
	castFunc := func(typ interface{}) *BACnetContextTag {
		if casted, ok := typ.(BACnetContextTag); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetContextTag); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetContextTag) GetTypeName() string {
	return "BACnetContextTag"
}

func (m *BACnetContextTag) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetContextTag) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *BACnetContextTag) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetContextTag) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetContextTagParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType) (*BACnetContextTag, error) {
	if pullErr := readBuffer.PullContext("BACnetContextTag"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, pullErr
	}
	_header, _headerErr := BACnetTagHeaderParse(readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field")
	}
	header := CastBACnetTagHeader(_header)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, closeErr
	}

	// Validation
	if !(bool((header.GetActualTagNumber()) == (tagNumberArgument))) {
		return nil, utils.ParseAssertError{"tagnumber doesn't match"}
	}

	// Validation
	if !(bool((header.GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS))) {
		return nil, utils.ParseAssertError{"should be a context tag"}
	}

	// Virtual field
	_tagNumber := header.GetTagNumber()
	tagNumber := uint8(_tagNumber)
	_ = tagNumber

	// Virtual field
	_actualLength := header.GetActualLength()
	actualLength := uint32(_actualLength)
	_ = actualLength

	// Virtual field
	_isNotOpeningOrClosingTag := bool(bool((header.GetLengthValueType()) != (6))) && bool(bool((header.GetLengthValueType()) != (7)))
	isNotOpeningOrClosingTag := bool(_isNotOpeningOrClosingTag)
	_ = isNotOpeningOrClosingTag

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *BACnetContextTag
	var typeSwitchError error
	switch {
	case dataType == BACnetDataType_BOOLEAN: // BACnetContextTagBoolean
		_parent, typeSwitchError = BACnetContextTagBooleanParse(readBuffer, tagNumberArgument, dataType, isNotOpeningOrClosingTag, header)
	case dataType == BACnetDataType_UNSIGNED_INTEGER: // BACnetContextTagUnsignedInteger
		_parent, typeSwitchError = BACnetContextTagUnsignedIntegerParse(readBuffer, tagNumberArgument, dataType, isNotOpeningOrClosingTag, header)
	case dataType == BACnetDataType_SIGNED_INTEGER: // BACnetContextTagSignedInteger
		_parent, typeSwitchError = BACnetContextTagSignedIntegerParse(readBuffer, tagNumberArgument, dataType, isNotOpeningOrClosingTag, header)
	case dataType == BACnetDataType_REAL: // BACnetContextTagReal
		_parent, typeSwitchError = BACnetContextTagRealParse(readBuffer, tagNumberArgument, dataType, isNotOpeningOrClosingTag)
	case dataType == BACnetDataType_DOUBLE: // BACnetContextTagDouble
		_parent, typeSwitchError = BACnetContextTagDoubleParse(readBuffer, tagNumberArgument, dataType, isNotOpeningOrClosingTag)
	case dataType == BACnetDataType_OCTET_STRING: // BACnetContextTagOctetString
		_parent, typeSwitchError = BACnetContextTagOctetStringParse(readBuffer, tagNumberArgument, dataType, isNotOpeningOrClosingTag, header)
	case dataType == BACnetDataType_CHARACTER_STRING: // BACnetContextTagCharacterString
		_parent, typeSwitchError = BACnetContextTagCharacterStringParse(readBuffer, tagNumberArgument, dataType, isNotOpeningOrClosingTag, header)
	case dataType == BACnetDataType_BIT_STRING: // BACnetContextTagBitString
		_parent, typeSwitchError = BACnetContextTagBitStringParse(readBuffer, tagNumberArgument, dataType, isNotOpeningOrClosingTag, header)
	case dataType == BACnetDataType_ENUMERATED: // BACnetContextTagEnumerated
		_parent, typeSwitchError = BACnetContextTagEnumeratedParse(readBuffer, tagNumberArgument, dataType, isNotOpeningOrClosingTag, header)
	case dataType == BACnetDataType_DATE: // BACnetContextTagDate
		_parent, typeSwitchError = BACnetContextTagDateParse(readBuffer, tagNumberArgument, dataType, isNotOpeningOrClosingTag)
	case dataType == BACnetDataType_TIME: // BACnetContextTagTime
		_parent, typeSwitchError = BACnetContextTagTimeParse(readBuffer, tagNumberArgument, dataType, isNotOpeningOrClosingTag)
	case dataType == BACnetDataType_BACNET_OBJECT_IDENTIFIER: // BACnetContextTagObjectIdentifier
		_parent, typeSwitchError = BACnetContextTagObjectIdentifierParse(readBuffer, tagNumberArgument, dataType, isNotOpeningOrClosingTag)
	case dataType == BACnetDataType_BACNET_PROPERTY_IDENTIFIER: // BACnetContextTagPropertyIdentifier
		_parent, typeSwitchError = BACnetContextTagPropertyIdentifierParse(readBuffer, tagNumberArgument, dataType, isNotOpeningOrClosingTag, actualLength)
	case dataType == BACnetDataType_EVENT_TYPE: // BACnetContextTagEventType
		_parent, typeSwitchError = BACnetContextTagEventTypeParse(readBuffer, tagNumberArgument, dataType, isNotOpeningOrClosingTag, actualLength)
	case dataType == BACnetDataType_EVENT_STATE: // BACnetContextTagEventState
		_parent, typeSwitchError = BACnetContextTagEventStateParse(readBuffer, tagNumberArgument, dataType, isNotOpeningOrClosingTag, actualLength)
	case dataType == BACnetDataType_NOTIFY_TYPE: // BACnetContextTagNotifyType
		_parent, typeSwitchError = BACnetContextTagNotifyTypeParse(readBuffer, tagNumberArgument, dataType, isNotOpeningOrClosingTag, actualLength)
	case dataType == BACnetDataType_BACNET_DEVICE_STATE: // BACnetContextTagDeviceState
		_parent, typeSwitchError = BACnetContextTagDeviceStateParse(readBuffer, tagNumberArgument, dataType, isNotOpeningOrClosingTag)
	case dataType == BACnetDataType_OPENING_TAG: // BACnetOpeningTag
		_parent, typeSwitchError = BACnetOpeningTagParse(readBuffer, tagNumberArgument, dataType, actualLength)
	case dataType == BACnetDataType_CLOSING_TAG: // BACnetClosingTag
		_parent, typeSwitchError = BACnetClosingTagParse(readBuffer, tagNumberArgument, dataType, actualLength)
	case true: // BACnetContextTagEmpty
		_parent, typeSwitchError = BACnetContextTagEmptyParse(readBuffer, tagNumberArgument, dataType)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("BACnetContextTag"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent, header)
	return _parent, nil
}

func (m *BACnetContextTag) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BACnetContextTag) SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetContextTag, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("BACnetContextTag"); pushErr != nil {
		return pushErr
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return pushErr
	}
	_headerErr := m.Header.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return popErr
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}
	// Virtual field
	if _tagNumberErr := writeBuffer.WriteVirtual("tagNumber", m.GetTagNumber()); _tagNumberErr != nil {
		return errors.Wrap(_tagNumberErr, "Error serializing 'tagNumber' field")
	}
	// Virtual field
	if _actualLengthErr := writeBuffer.WriteVirtual("actualLength", m.GetActualLength()); _actualLengthErr != nil {
		return errors.Wrap(_actualLengthErr, "Error serializing 'actualLength' field")
	}
	// Virtual field
	if _isNotOpeningOrClosingTagErr := writeBuffer.WriteVirtual("isNotOpeningOrClosingTag", m.GetIsNotOpeningOrClosingTag()); _isNotOpeningOrClosingTagErr != nil {
		return errors.Wrap(_isNotOpeningOrClosingTagErr, "Error serializing 'isNotOpeningOrClosingTag' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetContextTag"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetContextTag) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
