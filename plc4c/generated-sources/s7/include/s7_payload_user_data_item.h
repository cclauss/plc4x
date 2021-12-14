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

#ifndef PLC4C_S7_READ_WRITE_S7_PAYLOAD_USER_DATA_ITEM_H_
#define PLC4C_S7_READ_WRITE_S7_PAYLOAD_USER_DATA_ITEM_H_

#include <stdbool.h>
#include <stdint.h>
#include <plc4c/driver_s7_static.h>
#include <plc4c/spi/read_buffer.h>
#include <plc4c/spi/write_buffer.h>
#include <plc4c/utils/list.h>
#include "data_transport_error_code.h"
#include "syntax_id_type.h"
#include "s7_payload_user_data_item.h"
#include "alarm_type.h"
#include "date_and_time.h"
#include "szl_id.h"
#include "data_transport_size.h"
#include "alarm_message_object_ack_type.h"
#include "alarm_message_ack_push_type.h"
#include "alarm_message_push_type.h"
#include "alarm_state_type.h"
#include "query_type.h"
#include "szl_data_tree_item.h"

// Code generated by code-generation. DO NOT EDIT.

#ifdef __cplusplus
extern "C" {
#endif


// Structure used to contain the discriminator values for discriminated types using this as a parent
struct plc4c_s7_read_write_s7_payload_user_data_item_discriminator {
  uint8_t cpuFunctionType;
  uint8_t cpuSubfunction;
  uint16_t dataLength;
};
typedef struct plc4c_s7_read_write_s7_payload_user_data_item_discriminator plc4c_s7_read_write_s7_payload_user_data_item_discriminator;

// Enum assigning each sub-type an individual id.
enum plc4c_s7_read_write_s7_payload_user_data_item_type {
  plc4c_s7_read_write_s7_payload_user_data_item_type_plc4c_s7_read_write_s7_payload_diagnostic_message = 0,
  plc4c_s7_read_write_s7_payload_user_data_item_type_plc4c_s7_read_write_s7_payload_alarm8 = 1,
  plc4c_s7_read_write_s7_payload_user_data_item_type_plc4c_s7_read_write_s7_payload_notify = 2,
  plc4c_s7_read_write_s7_payload_user_data_item_type_plc4c_s7_read_write_s7_payload_alarm_ack_ind = 3,
  plc4c_s7_read_write_s7_payload_user_data_item_type_plc4c_s7_read_write_s7_payload_alarm_sq = 4,
  plc4c_s7_read_write_s7_payload_user_data_item_type_plc4c_s7_read_write_s7_payload_alarm_s = 5,
  plc4c_s7_read_write_s7_payload_user_data_item_type_plc4c_s7_read_write_s7_payload_alarm_sc = 6,
  plc4c_s7_read_write_s7_payload_user_data_item_type_plc4c_s7_read_write_s7_payload_notify8 = 7,
  plc4c_s7_read_write_s7_payload_user_data_item_type_plc4c_s7_read_write_s7_payload_user_data_item_cpu_function_read_szl_request = 8,
  plc4c_s7_read_write_s7_payload_user_data_item_type_plc4c_s7_read_write_s7_payload_user_data_item_cpu_function_read_szl_response = 9,
  plc4c_s7_read_write_s7_payload_user_data_item_type_plc4c_s7_read_write_s7_payload_user_data_item_cpu_function_msg_subscription = 10,
  plc4c_s7_read_write_s7_payload_user_data_item_type_plc4c_s7_read_write_s7_payload_user_data_item_cpu_function_msg_subscription_response = 11,
  plc4c_s7_read_write_s7_payload_user_data_item_type_plc4c_s7_read_write_s7_payload_user_data_item_cpu_function_msg_subscription_sys_response = 12,
  plc4c_s7_read_write_s7_payload_user_data_item_type_plc4c_s7_read_write_s7_payload_user_data_item_cpu_function_msg_subscription_alarm_response = 13,
  plc4c_s7_read_write_s7_payload_user_data_item_type_plc4c_s7_read_write_s7_payload_user_data_item_cpu_function_alarm_ack = 14,
  plc4c_s7_read_write_s7_payload_user_data_item_type_plc4c_s7_read_write_s7_payload_user_data_item_cpu_function_alarm_ack_response = 15,
  plc4c_s7_read_write_s7_payload_user_data_item_type_plc4c_s7_read_write_s7_payload_user_data_item_cpu_function_alarm_query = 16,
  plc4c_s7_read_write_s7_payload_user_data_item_type_plc4c_s7_read_write_s7_payload_user_data_item_cpu_function_alarm_query_response = 17};
typedef enum plc4c_s7_read_write_s7_payload_user_data_item_type plc4c_s7_read_write_s7_payload_user_data_item_type;

// Function to get the discriminator values for a given type.
plc4c_s7_read_write_s7_payload_user_data_item_discriminator plc4c_s7_read_write_s7_payload_user_data_item_get_discriminator(plc4c_s7_read_write_s7_payload_user_data_item_type type);

// Constant values.
uint8_t PLC4C_S7_READ_WRITE_S7_PAYLOAD_USER_DATA_ITEM_CPU_FUNCTION_ALARM_QUERY_LENGTH();
uint8_t PLC4C_S7_READ_WRITE_S7_PAYLOAD_USER_DATA_ITEM_CPU_FUNCTION_ALARM_QUERY_NUMBER_MESSAGE_OBJ();
uint8_t PLC4C_S7_READ_WRITE_S7_PAYLOAD_USER_DATA_ITEM_CPU_FUNCTION_ALARM_QUERY_VARIABLE_SPEC();
uint8_t PLC4C_S7_READ_WRITE_S7_PAYLOAD_USER_DATA_ITEM_CPU_FUNCTION_ALARM_QUERY_FUNCTION_ID();
uint16_t PLC4C_S7_READ_WRITE_S7_PAYLOAD_USER_DATA_ITEM_CPU_FUNCTION_READ_SZL_RESPONSE_SZL_ITEM_LENGTH();
uint8_t PLC4C_S7_READ_WRITE_S7_PAYLOAD_USER_DATA_ITEM_CPU_FUNCTION_ALARM_QUERY_RESPONSE_FUNCTION_ID();
uint8_t PLC4C_S7_READ_WRITE_S7_PAYLOAD_USER_DATA_ITEM_CPU_FUNCTION_ALARM_QUERY_RESPONSE_NUMBER_MESSAGE_OBJ();

struct plc4c_s7_read_write_s7_payload_user_data_item {
  /* This is an abstract type so this property saves the type of this typed union */
  plc4c_s7_read_write_s7_payload_user_data_item_type _type;
  /* Properties */
  plc4c_s7_read_write_data_transport_error_code return_code;
  plc4c_s7_read_write_data_transport_size transport_size;
  union {
    struct { /* S7PayloadDiagnosticMessage */
      uint16_t s7_payload_diagnostic_message_event_id;
      uint8_t s7_payload_diagnostic_message_priority_class;
      uint8_t s7_payload_diagnostic_message_ob_number;
      uint16_t s7_payload_diagnostic_message_dat_id;
      uint16_t s7_payload_diagnostic_message_info1;
      uint32_t s7_payload_diagnostic_message_info2;
      plc4c_s7_read_write_date_and_time* s7_payload_diagnostic_message_time_stamp;
    };
    struct { /* S7PayloadAlarm8 */
      plc4c_s7_read_write_alarm_message_push_type* s7_payload_alarm8_alarm_message;
    };
    struct { /* S7PayloadNotify */
      plc4c_s7_read_write_alarm_message_push_type* s7_payload_notify_alarm_message;
    };
    struct { /* S7PayloadAlarmAckInd */
      plc4c_s7_read_write_alarm_message_ack_push_type* s7_payload_alarm_ack_ind_alarm_message;
    };
    struct { /* S7PayloadAlarmSQ */
      plc4c_s7_read_write_alarm_message_push_type* s7_payload_alarm_sq_alarm_message;
    };
    struct { /* S7PayloadAlarmS */
      plc4c_s7_read_write_alarm_message_push_type* s7_payload_alarm_s_alarm_message;
    };
    struct { /* S7PayloadAlarmSC */
      plc4c_s7_read_write_alarm_message_push_type* s7_payload_alarm_sc_alarm_message;
    };
    struct { /* S7PayloadNotify8 */
      plc4c_s7_read_write_alarm_message_push_type* s7_payload_notify8_alarm_message;
    };
    struct { /* S7PayloadUserDataItemCpuFunctionReadSzlRequest */
      plc4c_s7_read_write_szl_id* s7_payload_user_data_item_cpu_function_read_szl_request_szl_id;
      uint16_t s7_payload_user_data_item_cpu_function_read_szl_request_szl_index;
    };
    struct { /* S7PayloadUserDataItemCpuFunctionReadSzlResponse */
      plc4c_s7_read_write_szl_id* s7_payload_user_data_item_cpu_function_read_szl_response_szl_id;
      uint16_t s7_payload_user_data_item_cpu_function_read_szl_response_szl_index;
      plc4c_list* s7_payload_user_data_item_cpu_function_read_szl_response_items;
    };
    struct { /* S7PayloadUserDataItemCpuFunctionMsgSubscription */
      uint8_t s7_payload_user_data_item_cpu_function_msg_subscription_subscription;
      char* s7_payload_user_data_item_cpu_function_msg_subscription_magic_key;
      plc4c_s7_read_write_alarm_state_type* s7_payload_user_data_item_cpu_function_msg_subscription_alarmtype;
      uint8_t* s7_payload_user_data_item_cpu_function_msg_subscription_reserve;
    };
    struct { /* S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse */
      uint8_t s7_payload_user_data_item_cpu_function_msg_subscription_sys_response_result;
      uint8_t s7_payload_user_data_item_cpu_function_msg_subscription_sys_response_reserved01;
    };
    struct { /* S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse */
      uint8_t s7_payload_user_data_item_cpu_function_msg_subscription_alarm_response_result;
      uint8_t s7_payload_user_data_item_cpu_function_msg_subscription_alarm_response_reserved01;
      plc4c_s7_read_write_alarm_type s7_payload_user_data_item_cpu_function_msg_subscription_alarm_response_alarm_type;
      uint8_t s7_payload_user_data_item_cpu_function_msg_subscription_alarm_response_reserved02;
      uint8_t s7_payload_user_data_item_cpu_function_msg_subscription_alarm_response_reserved03;
    };
    struct { /* S7PayloadUserDataItemCpuFunctionAlarmAck */
      uint8_t s7_payload_user_data_item_cpu_function_alarm_ack_function_id;
      plc4c_list* s7_payload_user_data_item_cpu_function_alarm_ack_message_objects;
    };
    struct { /* S7PayloadUserDataItemCpuFunctionAlarmAckResponse */
      uint8_t s7_payload_user_data_item_cpu_function_alarm_ack_response_function_id;
      plc4c_list* s7_payload_user_data_item_cpu_function_alarm_ack_response_message_objects;
    };
    struct { /* S7PayloadUserDataItemCpuFunctionAlarmQuery */
      plc4c_s7_read_write_syntax_id_type s7_payload_user_data_item_cpu_function_alarm_query_syntax_id;
      plc4c_s7_read_write_query_type s7_payload_user_data_item_cpu_function_alarm_query_query_type;
      plc4c_s7_read_write_alarm_type s7_payload_user_data_item_cpu_function_alarm_query_alarm_type;
    };
    struct { /* S7PayloadUserDataItemCpuFunctionAlarmQueryResponse */
      plc4c_s7_read_write_data_transport_error_code s7_payload_user_data_item_cpu_function_alarm_query_response_pudicf_return_code;
      plc4c_s7_read_write_data_transport_size s7_payload_user_data_item_cpu_function_alarm_query_response_pudicftransport_size;
    };
  };
};
typedef struct plc4c_s7_read_write_s7_payload_user_data_item plc4c_s7_read_write_s7_payload_user_data_item;

// Create an empty NULL-struct
plc4c_s7_read_write_s7_payload_user_data_item plc4c_s7_read_write_s7_payload_user_data_item_null();

plc4c_return_code plc4c_s7_read_write_s7_payload_user_data_item_parse(plc4c_spi_read_buffer* readBuffer, uint8_t cpuFunctionType, uint8_t cpuSubfunction, plc4c_s7_read_write_s7_payload_user_data_item** message);

plc4c_return_code plc4c_s7_read_write_s7_payload_user_data_item_serialize(plc4c_spi_write_buffer* writeBuffer, plc4c_s7_read_write_s7_payload_user_data_item* message);

uint16_t plc4c_s7_read_write_s7_payload_user_data_item_length_in_bytes(plc4c_s7_read_write_s7_payload_user_data_item* message);

uint16_t plc4c_s7_read_write_s7_payload_user_data_item_length_in_bits(plc4c_s7_read_write_s7_payload_user_data_item* message);

#ifdef __cplusplus
}
#endif
#endif  // PLC4C_S7_READ_WRITE_S7_PAYLOAD_USER_DATA_ITEM_H_
