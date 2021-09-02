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

#ifndef PLC4C_S7_READ_WRITE_ALARM_MESSAGE_OBJECT_ACK_TYPE_H_
#define PLC4C_S7_READ_WRITE_ALARM_MESSAGE_OBJECT_ACK_TYPE_H_

#include <stdbool.h>
#include <stdint.h>
#include <plc4c/driver_s7_static.h>
#include <plc4c/spi/read_buffer.h>
#include <plc4c/spi/write_buffer.h>
#include <plc4c/utils/list.h>
#include "syntax_id_type.h"
#include "state.h"

// Code generated by code-generation. DO NOT EDIT.

#ifdef __cplusplus
extern "C" {
#endif


// Constant values.
uint8_t PLC4C_S7_READ_WRITE_ALARM_MESSAGE_OBJECT_ACK_TYPE_VARIABLE_SPEC();
uint8_t PLC4C_S7_READ_WRITE_ALARM_MESSAGE_OBJECT_ACK_TYPE_LENGTH();

struct plc4c_s7_read_write_alarm_message_object_ack_type {
  /* Properties */
  uint8_t variable_spec;
  uint8_t length;
  plc4c_s7_read_write_syntax_id_type syntax_id;
  uint8_t number_of_values;
  uint32_t event_id;
  plc4c_s7_read_write_state* ack_state_going;
  plc4c_s7_read_write_state* ack_state_coming;
};
typedef struct plc4c_s7_read_write_alarm_message_object_ack_type plc4c_s7_read_write_alarm_message_object_ack_type;

// Create an empty NULL-struct
plc4c_s7_read_write_alarm_message_object_ack_type plc4c_s7_read_write_alarm_message_object_ack_type_null();

plc4c_return_code plc4c_s7_read_write_alarm_message_object_ack_type_parse(plc4c_spi_read_buffer* readBuffer, plc4c_s7_read_write_alarm_message_object_ack_type** message);

plc4c_return_code plc4c_s7_read_write_alarm_message_object_ack_type_serialize(plc4c_spi_write_buffer* writeBuffer, plc4c_s7_read_write_alarm_message_object_ack_type* message);

uint16_t plc4c_s7_read_write_alarm_message_object_ack_type_length_in_bytes(plc4c_s7_read_write_alarm_message_object_ack_type* message);

uint16_t plc4c_s7_read_write_alarm_message_object_ack_type_length_in_bits(plc4c_s7_read_write_alarm_message_object_ack_type* message);

#ifdef __cplusplus
}
#endif
#endif  // PLC4C_S7_READ_WRITE_ALARM_MESSAGE_OBJECT_ACK_TYPE_H_