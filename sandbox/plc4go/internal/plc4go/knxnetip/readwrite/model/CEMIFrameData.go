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
    "encoding/base64"
    "encoding/xml"
    "errors"
    "io"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/spi"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
    "reflect"
)

// The data-structure of this message
type CEMIFrameData struct {
    SourceAddress IKNXAddress
    DestinationAddress []int8
    GroupAddress bool
    HopCount uint8
    DataLength uint8
    Tcpi ITPCI
    Counter uint8
    Apci IAPCI
    DataFirstByte int8
    Data []int8
    Crc uint8
    CEMIFrame
}

// The corresponding interface
type ICEMIFrameData interface {
    ICEMIFrame
    Serialize(io utils.WriteBuffer) error
}

// Accessors for discriminator values.
func (m CEMIFrameData) NotAckFrame() bool {
    return true
}

func (m CEMIFrameData) StandardFrame() bool {
    return true
}

func (m CEMIFrameData) Polling() bool {
    return false
}

func (m CEMIFrameData) initialize(repeated bool, priority ICEMIPriority, acknowledgeRequested bool, errorFlag bool) spi.Message {
    m.Repeated = repeated
    m.Priority = priority
    m.AcknowledgeRequested = acknowledgeRequested
    m.ErrorFlag = errorFlag
    return m
}

func NewCEMIFrameData(sourceAddress IKNXAddress, destinationAddress []int8, groupAddress bool, hopCount uint8, dataLength uint8, tcpi ITPCI, counter uint8, apci IAPCI, dataFirstByte int8, data []int8, crc uint8) CEMIFrameInitializer {
    return &CEMIFrameData{SourceAddress: sourceAddress, DestinationAddress: destinationAddress, GroupAddress: groupAddress, HopCount: hopCount, DataLength: dataLength, Tcpi: tcpi, Counter: counter, Apci: apci, DataFirstByte: dataFirstByte, Data: data, Crc: crc}
}

func CastICEMIFrameData(structType interface{}) ICEMIFrameData {
    castFunc := func(typ interface{}) ICEMIFrameData {
        if iCEMIFrameData, ok := typ.(ICEMIFrameData); ok {
            return iCEMIFrameData
        }
        return nil
    }
    return castFunc(structType)
}

func CastCEMIFrameData(structType interface{}) CEMIFrameData {
    castFunc := func(typ interface{}) CEMIFrameData {
        if sCEMIFrameData, ok := typ.(CEMIFrameData); ok {
            return sCEMIFrameData
        }
        if sCEMIFrameData, ok := typ.(*CEMIFrameData); ok {
            return *sCEMIFrameData
        }
        return CEMIFrameData{}
    }
    return castFunc(structType)
}

func (m CEMIFrameData) LengthInBits() uint16 {
    var lengthInBits uint16 = m.CEMIFrame.LengthInBits()

    // Simple field (sourceAddress)
    lengthInBits += m.SourceAddress.LengthInBits()

    // Array field
    if len(m.DestinationAddress) > 0 {
        lengthInBits += 8 * uint16(len(m.DestinationAddress))
    }

    // Simple field (groupAddress)
    lengthInBits += 1

    // Simple field (hopCount)
    lengthInBits += 3

    // Simple field (dataLength)
    lengthInBits += 4

    // Enum Field (tcpi)
    lengthInBits += 2

    // Simple field (counter)
    lengthInBits += 4

    // Enum Field (apci)
    lengthInBits += 4

    // Simple field (dataFirstByte)
    lengthInBits += 6

    // Array field
    if len(m.Data) > 0 {
        lengthInBits += 8 * uint16(len(m.Data))
    }

    // Simple field (crc)
    lengthInBits += 8

    return lengthInBits
}

func (m CEMIFrameData) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func CEMIFrameDataParse(io *utils.ReadBuffer) (CEMIFrameInitializer, error) {

    // Simple Field (sourceAddress)
    _sourceAddressMessage, _err := KNXAddressParse(io)
    if _err != nil {
        return nil, errors.New("Error parsing simple field 'sourceAddress'. " + _err.Error())
    }
    var sourceAddress IKNXAddress
    sourceAddress, _sourceAddressOk := _sourceAddressMessage.(IKNXAddress)
    if !_sourceAddressOk {
        return nil, errors.New("Couldn't cast message of type " + reflect.TypeOf(_sourceAddressMessage).Name() + " to IKNXAddress")
    }

    // Array field (destinationAddress)
    // Count array
    destinationAddress := make([]int8, uint16(2))
    for curItem := uint16(0); curItem < uint16(uint16(2)); curItem++ {

        _item, _err := io.ReadInt8(8)
        if _err != nil {
            return nil, errors.New("Error parsing 'destinationAddress' field " + _err.Error())
        }
        destinationAddress[curItem] = _item
    }

    // Simple Field (groupAddress)
    groupAddress, _groupAddressErr := io.ReadBit()
    if _groupAddressErr != nil {
        return nil, errors.New("Error parsing 'groupAddress' field " + _groupAddressErr.Error())
    }

    // Simple Field (hopCount)
    hopCount, _hopCountErr := io.ReadUint8(3)
    if _hopCountErr != nil {
        return nil, errors.New("Error parsing 'hopCount' field " + _hopCountErr.Error())
    }

    // Simple Field (dataLength)
    dataLength, _dataLengthErr := io.ReadUint8(4)
    if _dataLengthErr != nil {
        return nil, errors.New("Error parsing 'dataLength' field " + _dataLengthErr.Error())
    }

    // Enum field (tcpi)
    tcpi, _tcpiErr := TPCIParse(io)
    if _tcpiErr != nil {
        return nil, errors.New("Error parsing 'tcpi' field " + _tcpiErr.Error())
    }

    // Simple Field (counter)
    counter, _counterErr := io.ReadUint8(4)
    if _counterErr != nil {
        return nil, errors.New("Error parsing 'counter' field " + _counterErr.Error())
    }

    // Enum field (apci)
    apci, _apciErr := APCIParse(io)
    if _apciErr != nil {
        return nil, errors.New("Error parsing 'apci' field " + _apciErr.Error())
    }

    // Simple Field (dataFirstByte)
    dataFirstByte, _dataFirstByteErr := io.ReadInt8(6)
    if _dataFirstByteErr != nil {
        return nil, errors.New("Error parsing 'dataFirstByte' field " + _dataFirstByteErr.Error())
    }

    // Array field (data)
    // Count array
    data := make([]int8, uint16(dataLength) - uint16(uint16(1)))
    for curItem := uint16(0); curItem < uint16(uint16(dataLength) - uint16(uint16(1))); curItem++ {

        _item, _err := io.ReadInt8(8)
        if _err != nil {
            return nil, errors.New("Error parsing 'data' field " + _err.Error())
        }
        data[curItem] = _item
    }

    // Simple Field (crc)
    crc, _crcErr := io.ReadUint8(8)
    if _crcErr != nil {
        return nil, errors.New("Error parsing 'crc' field " + _crcErr.Error())
    }

    // Create the instance
    return NewCEMIFrameData(sourceAddress, destinationAddress, groupAddress, hopCount, dataLength, tcpi, counter, apci, dataFirstByte, data, crc), nil
}

func (m CEMIFrameData) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (sourceAddress)
    sourceAddress := CastIKNXAddress(m.SourceAddress)
    _sourceAddressErr := sourceAddress.Serialize(io)
    if _sourceAddressErr != nil {
        return errors.New("Error serializing 'sourceAddress' field " + _sourceAddressErr.Error())
    }

    // Array Field (destinationAddress)
    if m.DestinationAddress != nil {
        for _, _element := range m.DestinationAddress {
            _elementErr := io.WriteInt8(8, _element)
            if _elementErr != nil {
                return errors.New("Error serializing 'destinationAddress' field " + _elementErr.Error())
            }
        }
    }

    // Simple Field (groupAddress)
    groupAddress := bool(m.GroupAddress)
    _groupAddressErr := io.WriteBit((bool) (groupAddress))
    if _groupAddressErr != nil {
        return errors.New("Error serializing 'groupAddress' field " + _groupAddressErr.Error())
    }

    // Simple Field (hopCount)
    hopCount := uint8(m.HopCount)
    _hopCountErr := io.WriteUint8(3, (hopCount))
    if _hopCountErr != nil {
        return errors.New("Error serializing 'hopCount' field " + _hopCountErr.Error())
    }

    // Simple Field (dataLength)
    dataLength := uint8(m.DataLength)
    _dataLengthErr := io.WriteUint8(4, (dataLength))
    if _dataLengthErr != nil {
        return errors.New("Error serializing 'dataLength' field " + _dataLengthErr.Error())
    }

    // Enum field (tcpi)
    tcpi := CastTPCI(m.Tcpi)
    _tcpiErr := tcpi.Serialize(io)
    if _tcpiErr != nil {
        return errors.New("Error serializing 'tcpi' field " + _tcpiErr.Error())
    }

    // Simple Field (counter)
    counter := uint8(m.Counter)
    _counterErr := io.WriteUint8(4, (counter))
    if _counterErr != nil {
        return errors.New("Error serializing 'counter' field " + _counterErr.Error())
    }

    // Enum field (apci)
    apci := CastAPCI(m.Apci)
    _apciErr := apci.Serialize(io)
    if _apciErr != nil {
        return errors.New("Error serializing 'apci' field " + _apciErr.Error())
    }

    // Simple Field (dataFirstByte)
    dataFirstByte := int8(m.DataFirstByte)
    _dataFirstByteErr := io.WriteInt8(6, (dataFirstByte))
    if _dataFirstByteErr != nil {
        return errors.New("Error serializing 'dataFirstByte' field " + _dataFirstByteErr.Error())
    }

    // Array Field (data)
    if m.Data != nil {
        for _, _element := range m.Data {
            _elementErr := io.WriteInt8(8, _element)
            if _elementErr != nil {
                return errors.New("Error serializing 'data' field " + _elementErr.Error())
            }
        }
    }

    // Simple Field (crc)
    crc := uint8(m.Crc)
    _crcErr := io.WriteUint8(8, (crc))
    if _crcErr != nil {
        return errors.New("Error serializing 'crc' field " + _crcErr.Error())
    }

        return nil
    }
    return CEMIFrameSerialize(io, m.CEMIFrame, CastICEMIFrame(m), ser)
}

func (m *CEMIFrameData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            case "sourceAddress":
                var data *KNXAddress
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.SourceAddress = CastIKNXAddress(data)
            case "destinationAddress":
                var _encoded string
                if err := d.DecodeElement(&_encoded, &tok); err != nil {
                    return err
                }
                _decoded := make([]byte, base64.StdEncoding.DecodedLen(len(_encoded)))
                _len, err := base64.StdEncoding.Decode(_decoded, []byte(_encoded))
                if err != nil {
                    return err
                }
                m.DestinationAddress = utils.ByteToInt8(_decoded[0:_len])
            case "groupAddress":
                var data bool
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.GroupAddress = data
            case "hopCount":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.HopCount = data
            case "dataLength":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.DataLength = data
            case "tcpi":
                var data *TPCI
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Tcpi = data
            case "counter":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Counter = data
            case "apci":
                var data *APCI
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Apci = data
            case "dataFirstByte":
                var data int8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.DataFirstByte = data
            case "data":
                var _encoded string
                if err := d.DecodeElement(&_encoded, &tok); err != nil {
                    return err
                }
                _decoded := make([]byte, base64.StdEncoding.DecodedLen(len(_encoded)))
                _len, err := base64.StdEncoding.Decode(_decoded, []byte(_encoded))
                if err != nil {
                    return err
                }
                m.Data = utils.ByteToInt8(_decoded[0:_len])
            case "crc":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Crc = data
            }
        }
    }
}

func (m CEMIFrameData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.knxnetip.readwrite.CEMIFrameData"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.SourceAddress, xml.StartElement{Name: xml.Name{Local: "sourceAddress"}}); err != nil {
        return err
    }
    _encodedDestinationAddress := make([]byte, base64.StdEncoding.EncodedLen(len(m.DestinationAddress)))
    base64.StdEncoding.Encode(_encodedDestinationAddress, utils.Int8ToByte(m.DestinationAddress))
    if err := e.EncodeElement(_encodedDestinationAddress, xml.StartElement{Name: xml.Name{Local: "destinationAddress"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.GroupAddress, xml.StartElement{Name: xml.Name{Local: "groupAddress"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.HopCount, xml.StartElement{Name: xml.Name{Local: "hopCount"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.DataLength, xml.StartElement{Name: xml.Name{Local: "dataLength"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Tcpi, xml.StartElement{Name: xml.Name{Local: "tcpi"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Counter, xml.StartElement{Name: xml.Name{Local: "counter"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Apci, xml.StartElement{Name: xml.Name{Local: "apci"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.DataFirstByte, xml.StartElement{Name: xml.Name{Local: "dataFirstByte"}}); err != nil {
        return err
    }
    _encodedData := make([]byte, base64.StdEncoding.EncodedLen(len(m.Data)))
    base64.StdEncoding.Encode(_encodedData, utils.Int8ToByte(m.Data))
    if err := e.EncodeElement(_encodedData, xml.StartElement{Name: xml.Name{Local: "data"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Crc, xml.StartElement{Name: xml.Name{Local: "crc"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}
