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
    "io"
)

// The data-structure of this message
type MFuncPropCon struct {
    Parent *CEMI
    IMFuncPropCon
}

// The corresponding interface
type IMFuncPropCon interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *MFuncPropCon) MessageCode() uint8 {
    return 0xFA
}


func (m *MFuncPropCon) InitializeParent(parent *CEMI) {
}

func NewMFuncPropCon() *CEMI {
    child := &MFuncPropCon{
        Parent: NewCEMI(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastMFuncPropCon(structType interface{}) *MFuncPropCon {
    castFunc := func(typ interface{}) *MFuncPropCon {
        if casted, ok := typ.(MFuncPropCon); ok {
            return &casted
        }
        if casted, ok := typ.(*MFuncPropCon); ok {
            return casted
        }
        if casted, ok := typ.(CEMI); ok {
            return CastMFuncPropCon(casted.Child)
        }
        if casted, ok := typ.(*CEMI); ok {
            return CastMFuncPropCon(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *MFuncPropCon) GetTypeName() string {
    return "MFuncPropCon"
}

func (m *MFuncPropCon) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    return lengthInBits
}

func (m *MFuncPropCon) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func MFuncPropConParse(io *utils.ReadBuffer) (*CEMI, error) {

    // Create a partially initialized instance
    _child := &MFuncPropCon{
        Parent: &CEMI{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *MFuncPropCon) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *MFuncPropCon) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
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

func (m *MFuncPropCon) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    return nil
}
