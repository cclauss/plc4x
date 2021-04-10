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
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type ModbusPDUGetComEventCounterResponse struct {
	Status     uint16
	EventCount uint16
	Parent     *ModbusPDU
}

// The corresponding interface
type IModbusPDUGetComEventCounterResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ModbusPDUGetComEventCounterResponse) ErrorFlag() bool {
	return false
}

func (m *ModbusPDUGetComEventCounterResponse) FunctionFlag() uint8 {
	return 0x0B
}

func (m *ModbusPDUGetComEventCounterResponse) Response() bool {
	return true
}

func (m *ModbusPDUGetComEventCounterResponse) InitializeParent(parent *ModbusPDU) {
}

func NewModbusPDUGetComEventCounterResponse(status uint16, eventCount uint16) *ModbusPDU {
	child := &ModbusPDUGetComEventCounterResponse{
		Status:     status,
		EventCount: eventCount,
		Parent:     NewModbusPDU(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastModbusPDUGetComEventCounterResponse(structType interface{}) *ModbusPDUGetComEventCounterResponse {
	castFunc := func(typ interface{}) *ModbusPDUGetComEventCounterResponse {
		if casted, ok := typ.(ModbusPDUGetComEventCounterResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusPDUGetComEventCounterResponse); ok {
			return casted
		}
		if casted, ok := typ.(ModbusPDU); ok {
			return CastModbusPDUGetComEventCounterResponse(casted.Child)
		}
		if casted, ok := typ.(*ModbusPDU); ok {
			return CastModbusPDUGetComEventCounterResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusPDUGetComEventCounterResponse) GetTypeName() string {
	return "ModbusPDUGetComEventCounterResponse"
}

func (m *ModbusPDUGetComEventCounterResponse) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (status)
	lengthInBits += 16

	// Simple field (eventCount)
	lengthInBits += 16

	return lengthInBits
}

func (m *ModbusPDUGetComEventCounterResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUGetComEventCounterResponseParse(io *utils.ReadBuffer) (*ModbusPDU, error) {

	// Simple Field (status)
	status, _statusErr := io.ReadUint16(16)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field")
	}

	// Simple Field (eventCount)
	eventCount, _eventCountErr := io.ReadUint16(16)
	if _eventCountErr != nil {
		return nil, errors.Wrap(_eventCountErr, "Error parsing 'eventCount' field")
	}

	// Create a partially initialized instance
	_child := &ModbusPDUGetComEventCounterResponse{
		Status:     status,
		EventCount: eventCount,
		Parent:     &ModbusPDU{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ModbusPDUGetComEventCounterResponse) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Simple Field (status)
		status := uint16(m.Status)
		_statusErr := io.WriteUint16(16, (status))
		if _statusErr != nil {
			return errors.Wrap(_statusErr, "Error serializing 'status' field")
		}

		// Simple Field (eventCount)
		eventCount := uint16(m.EventCount)
		_eventCountErr := io.WriteUint16(16, (eventCount))
		if _eventCountErr != nil {
			return errors.Wrap(_eventCountErr, "Error serializing 'eventCount' field")
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *ModbusPDUGetComEventCounterResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "status":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Status = data
			case "eventCount":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.EventCount = data
			}
		}
		token, err = d.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

func (m *ModbusPDUGetComEventCounterResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.Status, xml.StartElement{Name: xml.Name{Local: "status"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.EventCount, xml.StartElement{Name: xml.Name{Local: "eventCount"}}); err != nil {
		return err
	}
	return nil
}

func (m ModbusPDUGetComEventCounterResponse) String() string {
	return string(m.Box("ModbusPDUGetComEventCounterResponse", utils.DefaultWidth*2))
}

func (m ModbusPDUGetComEventCounterResponse) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "ModbusPDUGetComEventCounterResponse"
	}
	boxes := make([]utils.AsciiBox, 0)
	boxes = append(boxes, utils.BoxAnything("Status", m.Status, width-2))
	boxes = append(boxes, utils.BoxAnything("EventCount", m.EventCount, width-2))
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
