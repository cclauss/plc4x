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
    "errors"
    "io"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/spi"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
)

// The data-structure of this message
type BACnetTagApplicationReal struct {
    Value float32
    BACnetTag
}

// The corresponding interface
type IBACnetTagApplicationReal interface {
    IBACnetTag
    Serialize(io utils.WriteBuffer) error
}

// Accessors for discriminator values.
func (m BACnetTagApplicationReal) ContextSpecificTag() uint8 {
    return 0
}

func (m BACnetTagApplicationReal) initialize(typeOrTagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8) spi.Message {
    m.TypeOrTagNumber = typeOrTagNumber
    m.LengthValueType = lengthValueType
    m.ExtTagNumber = extTagNumber
    m.ExtLength = extLength
    return m
}

func NewBACnetTagApplicationReal(value float32) BACnetTagInitializer {
    return &BACnetTagApplicationReal{Value: value}
}

func CastIBACnetTagApplicationReal(structType interface{}) IBACnetTagApplicationReal {
    castFunc := func(typ interface{}) IBACnetTagApplicationReal {
        if iBACnetTagApplicationReal, ok := typ.(IBACnetTagApplicationReal); ok {
            return iBACnetTagApplicationReal
        }
        return nil
    }
    return castFunc(structType)
}

func CastBACnetTagApplicationReal(structType interface{}) BACnetTagApplicationReal {
    castFunc := func(typ interface{}) BACnetTagApplicationReal {
        if sBACnetTagApplicationReal, ok := typ.(BACnetTagApplicationReal); ok {
            return sBACnetTagApplicationReal
        }
        if sBACnetTagApplicationReal, ok := typ.(*BACnetTagApplicationReal); ok {
            return *sBACnetTagApplicationReal
        }
        return BACnetTagApplicationReal{}
    }
    return castFunc(structType)
}

func (m BACnetTagApplicationReal) LengthInBits() uint16 {
    var lengthInBits uint16 = m.BACnetTag.LengthInBits()

    // Simple field (value)
    lengthInBits += 32

    return lengthInBits
}

func (m BACnetTagApplicationReal) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func BACnetTagApplicationRealParse(io *utils.ReadBuffer, lengthValueType uint8, extLength uint8) (BACnetTagInitializer, error) {

    // Simple Field (value)
    value, _valueErr := io.ReadFloat32(32)
    if _valueErr != nil {
        return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
    }

    // Create the instance
    return NewBACnetTagApplicationReal(value), nil
}

func (m BACnetTagApplicationReal) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (value)
    value := float32(m.Value)
    _valueErr := io.WriteFloat32(32, (value))
    if _valueErr != nil {
        return errors.New("Error serializing 'value' field " + _valueErr.Error())
    }

        return nil
    }
    return BACnetTagSerialize(io, m.BACnetTag, CastIBACnetTag(m), ser)
}

func (m *BACnetTagApplicationReal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    for {
        token, err := d.Token()
        if err != nil {
            if err == io.EOF {
                return nil
            }
            return err
        }
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "value":
                var data float32
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Value = data
            }
        }
    }
}

func (m BACnetTagApplicationReal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.bacnetip.readwrite.BACnetTagApplicationReal"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Value, xml.StartElement{Name: xml.Name{Local: "value"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}
