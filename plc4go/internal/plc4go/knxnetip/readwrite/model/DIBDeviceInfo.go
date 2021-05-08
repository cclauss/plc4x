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
type DIBDeviceInfo struct {
	DescriptionType                uint8
	KnxMedium                      KnxMedium
	DeviceStatus                   *DeviceStatus
	KnxAddress                     *KnxAddress
	ProjectInstallationIdentifier  *ProjectInstallationIdentifier
	KnxNetIpDeviceSerialNumber     []int8
	KnxNetIpDeviceMulticastAddress *IPAddress
	KnxNetIpDeviceMacAddress       *MACAddress
	DeviceFriendlyName             []int8
}

// The corresponding interface
type IDIBDeviceInfo interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

func NewDIBDeviceInfo(descriptionType uint8, knxMedium KnxMedium, deviceStatus *DeviceStatus, knxAddress *KnxAddress, projectInstallationIdentifier *ProjectInstallationIdentifier, knxNetIpDeviceSerialNumber []int8, knxNetIpDeviceMulticastAddress *IPAddress, knxNetIpDeviceMacAddress *MACAddress, deviceFriendlyName []int8) *DIBDeviceInfo {
	return &DIBDeviceInfo{DescriptionType: descriptionType, KnxMedium: knxMedium, DeviceStatus: deviceStatus, KnxAddress: knxAddress, ProjectInstallationIdentifier: projectInstallationIdentifier, KnxNetIpDeviceSerialNumber: knxNetIpDeviceSerialNumber, KnxNetIpDeviceMulticastAddress: knxNetIpDeviceMulticastAddress, KnxNetIpDeviceMacAddress: knxNetIpDeviceMacAddress, DeviceFriendlyName: deviceFriendlyName}
}

func CastDIBDeviceInfo(structType interface{}) *DIBDeviceInfo {
	castFunc := func(typ interface{}) *DIBDeviceInfo {
		if casted, ok := typ.(DIBDeviceInfo); ok {
			return &casted
		}
		if casted, ok := typ.(*DIBDeviceInfo); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *DIBDeviceInfo) GetTypeName() string {
	return "DIBDeviceInfo"
}

func (m *DIBDeviceInfo) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *DIBDeviceInfo) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (structureLength)
	lengthInBits += 8

	// Simple field (descriptionType)
	lengthInBits += 8

	// Simple field (knxMedium)
	lengthInBits += 8

	// Simple field (deviceStatus)
	lengthInBits += m.DeviceStatus.LengthInBits()

	// Simple field (knxAddress)
	lengthInBits += m.KnxAddress.LengthInBits()

	// Simple field (projectInstallationIdentifier)
	lengthInBits += m.ProjectInstallationIdentifier.LengthInBits()

	// Array field
	if len(m.KnxNetIpDeviceSerialNumber) > 0 {
		lengthInBits += 8 * uint16(len(m.KnxNetIpDeviceSerialNumber))
	}

	// Simple field (knxNetIpDeviceMulticastAddress)
	lengthInBits += m.KnxNetIpDeviceMulticastAddress.LengthInBits()

	// Simple field (knxNetIpDeviceMacAddress)
	lengthInBits += m.KnxNetIpDeviceMacAddress.LengthInBits()

	// Array field
	if len(m.DeviceFriendlyName) > 0 {
		lengthInBits += 8 * uint16(len(m.DeviceFriendlyName))
	}

	return lengthInBits
}

func (m *DIBDeviceInfo) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func DIBDeviceInfoParse(readBuffer utils.ReadBuffer) (*DIBDeviceInfo, error) {
	if pullErr := readBuffer.PullContext("DIBDeviceInfo"); pullErr != nil {
		return nil, pullErr
	}

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	structureLength, _structureLengthErr := readBuffer.ReadUint8("structureLength", 8)
	_ = structureLength
	if _structureLengthErr != nil {
		return nil, errors.Wrap(_structureLengthErr, "Error parsing 'structureLength' field")
	}

	// Simple Field (descriptionType)
	descriptionType, _descriptionTypeErr := readBuffer.ReadUint8("descriptionType", 8)
	if _descriptionTypeErr != nil {
		return nil, errors.Wrap(_descriptionTypeErr, "Error parsing 'descriptionType' field")
	}

	if pullErr := readBuffer.PullContext("knxMedium"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (knxMedium)
	knxMedium, _knxMediumErr := KnxMediumParse(readBuffer)
	if _knxMediumErr != nil {
		return nil, errors.Wrap(_knxMediumErr, "Error parsing 'knxMedium' field")
	}
	if closeErr := readBuffer.CloseContext("knxMedium"); closeErr != nil {
		return nil, closeErr
	}

	if pullErr := readBuffer.PullContext("deviceStatus"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (deviceStatus)
	deviceStatus, _deviceStatusErr := DeviceStatusParse(readBuffer)
	if _deviceStatusErr != nil {
		return nil, errors.Wrap(_deviceStatusErr, "Error parsing 'deviceStatus' field")
	}
	if closeErr := readBuffer.CloseContext("deviceStatus"); closeErr != nil {
		return nil, closeErr
	}

	if pullErr := readBuffer.PullContext("knxAddress"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (knxAddress)
	knxAddress, _knxAddressErr := KnxAddressParse(readBuffer)
	if _knxAddressErr != nil {
		return nil, errors.Wrap(_knxAddressErr, "Error parsing 'knxAddress' field")
	}
	if closeErr := readBuffer.CloseContext("knxAddress"); closeErr != nil {
		return nil, closeErr
	}

	if pullErr := readBuffer.PullContext("projectInstallationIdentifier"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (projectInstallationIdentifier)
	projectInstallationIdentifier, _projectInstallationIdentifierErr := ProjectInstallationIdentifierParse(readBuffer)
	if _projectInstallationIdentifierErr != nil {
		return nil, errors.Wrap(_projectInstallationIdentifierErr, "Error parsing 'projectInstallationIdentifier' field")
	}
	if closeErr := readBuffer.CloseContext("projectInstallationIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Array field (knxNetIpDeviceSerialNumber)
	if pullErr := readBuffer.PullContext("knxNetIpDeviceSerialNumber", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	knxNetIpDeviceSerialNumber := make([]int8, uint16(6))
	for curItem := uint16(0); curItem < uint16(uint16(6)); curItem++ {
		_item, _err := readBuffer.ReadInt8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'knxNetIpDeviceSerialNumber' field")
		}
		knxNetIpDeviceSerialNumber[curItem] = _item
	}
	if closeErr := readBuffer.CloseContext("knxNetIpDeviceSerialNumber", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if pullErr := readBuffer.PullContext("knxNetIpDeviceMulticastAddress"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (knxNetIpDeviceMulticastAddress)
	knxNetIpDeviceMulticastAddress, _knxNetIpDeviceMulticastAddressErr := IPAddressParse(readBuffer)
	if _knxNetIpDeviceMulticastAddressErr != nil {
		return nil, errors.Wrap(_knxNetIpDeviceMulticastAddressErr, "Error parsing 'knxNetIpDeviceMulticastAddress' field")
	}
	if closeErr := readBuffer.CloseContext("knxNetIpDeviceMulticastAddress"); closeErr != nil {
		return nil, closeErr
	}

	if pullErr := readBuffer.PullContext("knxNetIpDeviceMacAddress"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (knxNetIpDeviceMacAddress)
	knxNetIpDeviceMacAddress, _knxNetIpDeviceMacAddressErr := MACAddressParse(readBuffer)
	if _knxNetIpDeviceMacAddressErr != nil {
		return nil, errors.Wrap(_knxNetIpDeviceMacAddressErr, "Error parsing 'knxNetIpDeviceMacAddress' field")
	}
	if closeErr := readBuffer.CloseContext("knxNetIpDeviceMacAddress"); closeErr != nil {
		return nil, closeErr
	}

	// Array field (deviceFriendlyName)
	if pullErr := readBuffer.PullContext("deviceFriendlyName", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	deviceFriendlyName := make([]int8, uint16(30))
	for curItem := uint16(0); curItem < uint16(uint16(30)); curItem++ {
		_item, _err := readBuffer.ReadInt8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'deviceFriendlyName' field")
		}
		deviceFriendlyName[curItem] = _item
	}
	if closeErr := readBuffer.CloseContext("deviceFriendlyName", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("DIBDeviceInfo"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewDIBDeviceInfo(descriptionType, knxMedium, deviceStatus, knxAddress, projectInstallationIdentifier, knxNetIpDeviceSerialNumber, knxNetIpDeviceMulticastAddress, knxNetIpDeviceMacAddress, deviceFriendlyName), nil
}

func (m *DIBDeviceInfo) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("DIBDeviceInfo"); pushErr != nil {
		return pushErr
	}

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	structureLength := uint8(uint8(m.LengthInBytes()))
	_structureLengthErr := writeBuffer.WriteUint8("structureLength", 8, (structureLength))
	if _structureLengthErr != nil {
		return errors.Wrap(_structureLengthErr, "Error serializing 'structureLength' field")
	}

	// Simple Field (descriptionType)
	descriptionType := uint8(m.DescriptionType)
	_descriptionTypeErr := writeBuffer.WriteUint8("descriptionType", 8, (descriptionType))
	if _descriptionTypeErr != nil {
		return errors.Wrap(_descriptionTypeErr, "Error serializing 'descriptionType' field")
	}

	// Simple Field (knxMedium)
	if pushErr := writeBuffer.PushContext("knxMedium"); pushErr != nil {
		return pushErr
	}
	_knxMediumErr := m.KnxMedium.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("knxMedium"); popErr != nil {
		return popErr
	}
	if _knxMediumErr != nil {
		return errors.Wrap(_knxMediumErr, "Error serializing 'knxMedium' field")
	}

	// Simple Field (deviceStatus)
	if pushErr := writeBuffer.PushContext("deviceStatus"); pushErr != nil {
		return pushErr
	}
	_deviceStatusErr := m.DeviceStatus.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("deviceStatus"); popErr != nil {
		return popErr
	}
	if _deviceStatusErr != nil {
		return errors.Wrap(_deviceStatusErr, "Error serializing 'deviceStatus' field")
	}

	// Simple Field (knxAddress)
	if pushErr := writeBuffer.PushContext("knxAddress"); pushErr != nil {
		return pushErr
	}
	_knxAddressErr := m.KnxAddress.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("knxAddress"); popErr != nil {
		return popErr
	}
	if _knxAddressErr != nil {
		return errors.Wrap(_knxAddressErr, "Error serializing 'knxAddress' field")
	}

	// Simple Field (projectInstallationIdentifier)
	if pushErr := writeBuffer.PushContext("projectInstallationIdentifier"); pushErr != nil {
		return pushErr
	}
	_projectInstallationIdentifierErr := m.ProjectInstallationIdentifier.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("projectInstallationIdentifier"); popErr != nil {
		return popErr
	}
	if _projectInstallationIdentifierErr != nil {
		return errors.Wrap(_projectInstallationIdentifierErr, "Error serializing 'projectInstallationIdentifier' field")
	}

	// Array Field (knxNetIpDeviceSerialNumber)
	if m.KnxNetIpDeviceSerialNumber != nil {
		if pushErr := writeBuffer.PushContext("knxNetIpDeviceSerialNumber", utils.WithRenderAsList(true)); pushErr != nil {
			return pushErr
		}
		for _, _element := range m.KnxNetIpDeviceSerialNumber {
			_elementErr := writeBuffer.WriteInt8("", 8, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'knxNetIpDeviceSerialNumber' field")
			}
		}
		if popErr := writeBuffer.PopContext("knxNetIpDeviceSerialNumber", utils.WithRenderAsList(true)); popErr != nil {
			return popErr
		}
	}

	// Simple Field (knxNetIpDeviceMulticastAddress)
	if pushErr := writeBuffer.PushContext("knxNetIpDeviceMulticastAddress"); pushErr != nil {
		return pushErr
	}
	_knxNetIpDeviceMulticastAddressErr := m.KnxNetIpDeviceMulticastAddress.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("knxNetIpDeviceMulticastAddress"); popErr != nil {
		return popErr
	}
	if _knxNetIpDeviceMulticastAddressErr != nil {
		return errors.Wrap(_knxNetIpDeviceMulticastAddressErr, "Error serializing 'knxNetIpDeviceMulticastAddress' field")
	}

	// Simple Field (knxNetIpDeviceMacAddress)
	if pushErr := writeBuffer.PushContext("knxNetIpDeviceMacAddress"); pushErr != nil {
		return pushErr
	}
	_knxNetIpDeviceMacAddressErr := m.KnxNetIpDeviceMacAddress.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("knxNetIpDeviceMacAddress"); popErr != nil {
		return popErr
	}
	if _knxNetIpDeviceMacAddressErr != nil {
		return errors.Wrap(_knxNetIpDeviceMacAddressErr, "Error serializing 'knxNetIpDeviceMacAddress' field")
	}

	// Array Field (deviceFriendlyName)
	if m.DeviceFriendlyName != nil {
		if pushErr := writeBuffer.PushContext("deviceFriendlyName", utils.WithRenderAsList(true)); pushErr != nil {
			return pushErr
		}
		for _, _element := range m.DeviceFriendlyName {
			_elementErr := writeBuffer.WriteInt8("", 8, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'deviceFriendlyName' field")
			}
		}
		if popErr := writeBuffer.PopContext("deviceFriendlyName", utils.WithRenderAsList(true)); popErr != nil {
			return popErr
		}
	}

	if popErr := writeBuffer.PopContext("DIBDeviceInfo"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *DIBDeviceInfo) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
