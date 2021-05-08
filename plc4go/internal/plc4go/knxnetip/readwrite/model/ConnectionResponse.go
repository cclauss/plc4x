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
type ConnectionResponse struct {
	CommunicationChannelId      uint8
	Status                      Status
	HpaiDataEndpoint            *HPAIDataEndpoint
	ConnectionResponseDataBlock *ConnectionResponseDataBlock
	Parent                      *KnxNetIpMessage
}

// The corresponding interface
type IConnectionResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ConnectionResponse) MsgType() uint16 {
	return 0x0206
}

func (m *ConnectionResponse) InitializeParent(parent *KnxNetIpMessage) {
}

func NewConnectionResponse(communicationChannelId uint8, status Status, hpaiDataEndpoint *HPAIDataEndpoint, connectionResponseDataBlock *ConnectionResponseDataBlock) *KnxNetIpMessage {
	child := &ConnectionResponse{
		CommunicationChannelId:      communicationChannelId,
		Status:                      status,
		HpaiDataEndpoint:            hpaiDataEndpoint,
		ConnectionResponseDataBlock: connectionResponseDataBlock,
		Parent:                      NewKnxNetIpMessage(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastConnectionResponse(structType interface{}) *ConnectionResponse {
	castFunc := func(typ interface{}) *ConnectionResponse {
		if casted, ok := typ.(ConnectionResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*ConnectionResponse); ok {
			return casted
		}
		if casted, ok := typ.(KnxNetIpMessage); ok {
			return CastConnectionResponse(casted.Child)
		}
		if casted, ok := typ.(*KnxNetIpMessage); ok {
			return CastConnectionResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ConnectionResponse) GetTypeName() string {
	return "ConnectionResponse"
}

func (m *ConnectionResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ConnectionResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (communicationChannelId)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	// Optional Field (hpaiDataEndpoint)
	if m.HpaiDataEndpoint != nil {
		lengthInBits += (*m.HpaiDataEndpoint).LengthInBits()
	}

	// Optional Field (connectionResponseDataBlock)
	if m.ConnectionResponseDataBlock != nil {
		lengthInBits += (*m.ConnectionResponseDataBlock).LengthInBits()
	}

	return lengthInBits
}

func (m *ConnectionResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ConnectionResponseParse(readBuffer utils.ReadBuffer) (*KnxNetIpMessage, error) {
	if pullErr := readBuffer.PullContext("ConnectionResponse"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (communicationChannelId)
	communicationChannelId, _communicationChannelIdErr := readBuffer.ReadUint8("communicationChannelId", 8)
	if _communicationChannelIdErr != nil {
		return nil, errors.Wrap(_communicationChannelIdErr, "Error parsing 'communicationChannelId' field")
	}

	if pullErr := readBuffer.PullContext("status"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (status)
	status, _statusErr := StatusParse(readBuffer)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field")
	}
	if closeErr := readBuffer.CloseContext("status"); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (hpaiDataEndpoint) (Can be skipped, if a given expression evaluates to false)
	var hpaiDataEndpoint *HPAIDataEndpoint = nil
	if bool((status) == (Status_NO_ERROR)) {
		_val, _err := HPAIDataEndpointParse(readBuffer)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'hpaiDataEndpoint' field")
		}
		hpaiDataEndpoint = _val
	}

	// Optional Field (connectionResponseDataBlock) (Can be skipped, if a given expression evaluates to false)
	var connectionResponseDataBlock *ConnectionResponseDataBlock = nil
	if bool((status) == (Status_NO_ERROR)) {
		_val, _err := ConnectionResponseDataBlockParse(readBuffer)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'connectionResponseDataBlock' field")
		}
		connectionResponseDataBlock = _val
	}

	if closeErr := readBuffer.CloseContext("ConnectionResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ConnectionResponse{
		CommunicationChannelId:      communicationChannelId,
		Status:                      status,
		HpaiDataEndpoint:            hpaiDataEndpoint,
		ConnectionResponseDataBlock: connectionResponseDataBlock,
		Parent:                      &KnxNetIpMessage{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ConnectionResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConnectionResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (communicationChannelId)
		communicationChannelId := uint8(m.CommunicationChannelId)
		_communicationChannelIdErr := writeBuffer.WriteUint8("communicationChannelId", 8, (communicationChannelId))
		if _communicationChannelIdErr != nil {
			return errors.Wrap(_communicationChannelIdErr, "Error serializing 'communicationChannelId' field")
		}

		// Simple Field (status)
		if pushErr := writeBuffer.PushContext("status"); pushErr != nil {
			return pushErr
		}
		_statusErr := m.Status.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("status"); popErr != nil {
			return popErr
		}
		if _statusErr != nil {
			return errors.Wrap(_statusErr, "Error serializing 'status' field")
		}

		// Optional Field (hpaiDataEndpoint) (Can be skipped, if the value is null)
		var hpaiDataEndpoint *HPAIDataEndpoint = nil
		if m.HpaiDataEndpoint != nil {
			hpaiDataEndpoint = m.HpaiDataEndpoint
			_hpaiDataEndpointErr := hpaiDataEndpoint.Serialize(writeBuffer)
			if _hpaiDataEndpointErr != nil {
				return errors.Wrap(_hpaiDataEndpointErr, "Error serializing 'hpaiDataEndpoint' field")
			}
		}

		// Optional Field (connectionResponseDataBlock) (Can be skipped, if the value is null)
		var connectionResponseDataBlock *ConnectionResponseDataBlock = nil
		if m.ConnectionResponseDataBlock != nil {
			connectionResponseDataBlock = m.ConnectionResponseDataBlock
			_connectionResponseDataBlockErr := connectionResponseDataBlock.Serialize(writeBuffer)
			if _connectionResponseDataBlockErr != nil {
				return errors.Wrap(_connectionResponseDataBlockErr, "Error serializing 'connectionResponseDataBlock' field")
			}
		}

		if popErr := writeBuffer.PopContext("ConnectionResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *ConnectionResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
