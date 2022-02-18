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

#include <stdio.h>
#include <string.h>
#include <time.h>
#include <plc4c/data.h>
#include <plc4c/utils/list.h>
#include <plc4c/spi/evaluation_helper.h>

#include "plc4x_value.h"

// Code generated by code-generation. DO NOT EDIT.

// Parse function.
plc4c_return_code plc4c_plc4x_read_write_plc4x_value_parse(plc4c_spi_read_buffer* readBuffer, plc4c_plc4x_read_write_plc4x_value_type valueType, plc4c_data** data_item) {
    uint16_t startPos = plc4c_spi_read_get_pos(readBuffer);
    uint16_t curPos;
    plc4c_return_code _res = OK;

        if(valueType == plc4c_plc4x_read_write_plc4x_value_type_BOOL) { /* BOOL */

                // Reserved Field (Compartmentalized so the "reserved" variable can't leak)
                {
                    uint8_t _reserved = 0;
                    _res = plc4c_spi_read_unsigned_byte(readBuffer, 7, (uint8_t*) &_reserved);
                    if(_res != OK) {
                        return _res;
                    }
                    if(_reserved != 0x00) {
                      printf("Expected constant value '%d' but got '%d' for reserved field.", 0x00, _reserved);
                    }
                }

                // Simple Field (value)
                bool value = false;
                _res = plc4c_spi_read_bit(readBuffer, (bool*) &value);
                if(_res != OK) {
                    return _res;
                }

                    // Hurz
                *data_item = plc4c_data_create_bool_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_BYTE) { /* BitString */

                // Simple Field (value)
                uint8_t value = 0;
                _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_uint8_t_bit_string_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_WORD) { /* BitString */

                // Simple Field (value)
                uint16_t value = 0;
                _res = plc4c_spi_read_unsigned_short(readBuffer, 16, (uint16_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_uint16_t_bit_string_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_DWORD) { /* BitString */

                // Simple Field (value)
                uint32_t value = 0;
                _res = plc4c_spi_read_unsigned_int(readBuffer, 32, (uint32_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_uint32_t_bit_string_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_USINT) { /* USINT */

                // Simple Field (value)
                uint8_t value = 0;
                _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                    // Hurz
                *data_item = plc4c_data_create_uint8_t_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_UINT) { /* UINT */

                // Simple Field (value)
                uint16_t value = 0;
                _res = plc4c_spi_read_unsigned_short(readBuffer, 16, (uint16_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                    // Hurz
                *data_item = plc4c_data_create_uint16_t_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_UDINT) { /* UDINT */

                // Simple Field (value)
                uint32_t value = 0;
                _res = plc4c_spi_read_unsigned_int(readBuffer, 32, (uint32_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                    // Hurz
                *data_item = plc4c_data_create_uint32_t_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_ULINT) { /* ULINT */

                // Simple Field (value)
                uint64_t value = 0;
                _res = plc4c_spi_read_unsigned_long(readBuffer, 64, (uint64_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                    // Hurz
                *data_item = plc4c_data_create_uint64_t_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_SINT) { /* SINT */

                // Simple Field (value)
                int8_t value = 0;
                _res = plc4c_spi_read_signed_byte(readBuffer, 8, (int8_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                    // Hurz
                *data_item = plc4c_data_create_int8_t_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_INT) { /* INT */

                // Simple Field (value)
                int16_t value = 0;
                _res = plc4c_spi_read_signed_short(readBuffer, 16, (int16_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                    // Hurz
                *data_item = plc4c_data_create_int16_t_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_DINT) { /* DINT */

                // Simple Field (value)
                int32_t value = 0;
                _res = plc4c_spi_read_signed_int(readBuffer, 32, (int32_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                    // Hurz
                *data_item = plc4c_data_create_int32_t_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_LINT) { /* LINT */

                // Simple Field (value)
                int64_t value = 0;
                _res = plc4c_spi_read_signed_long(readBuffer, 64, (int64_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                    // Hurz
                *data_item = plc4c_data_create_int64_t_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_REAL) { /* REAL */

                // Simple Field (value)
                float value = 0.0f;
                _res = plc4c_spi_read_float(readBuffer, 32, (float*) &value);
                if(_res != OK) {
                    return _res;
                }

                    // Hurz
                *data_item = plc4c_data_create_float_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_LREAL) { /* LREAL */

                // Simple Field (value)
                double value = 0.0f;
                _res = plc4c_spi_read_double(readBuffer, 64, (double*) &value);
                if(_res != OK) {
                    return _res;
                }

                    // Hurz
                *data_item = plc4c_data_create_double_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_CHAR) { /* STRING */

                // Simple Field (value)
                char* value = "";
                _res = plc4c_spi_read_string(readBuffer, 8, "UTF-8", (char**) &value);
                if(_res != OK) {
                    return _res;
                }

                    // Hurz
                *data_item = plc4c_data_create_char*_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_WCHAR) { /* STRING */

                // Simple Field (value)
                char* value = "";
                _res = plc4c_spi_read_string(readBuffer, 16, "UTF-16", (char**) &value);
                if(_res != OK) {
                    return _res;
                }

                    // Hurz
                *data_item = plc4c_data_create_char*_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_STRING) { /* STRING */
    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_WSTRING) { /* STRING */
    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_TIME) { /* TIME */
    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_TIME_OF_DAY) { /* TIME_OF_DAY */
    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_DATE) { /* DATE */
    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_DATE_AND_TIME) { /* DATE_AND_TIME */
    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_Struct) { /* Struct */
    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_List) { /* List */
    }

  return OK;
}

plc4c_return_code plc4c_plc4x_read_write_plc4x_value_serialize(plc4c_spi_write_buffer* writeBuffer, plc4c_data** data_item) {
  plc4c_return_code _res = OK;

  return OK;
}

uint16_t plc4c_plc4x_read_write_plc4x_value_length_in_bytes(plc4c_data* data_item) {
  return plc4c_plc4x_read_write_plc4x_value_length_in_bits(data_item) / 8;
}

uint16_t plc4c_plc4x_read_write_plc4x_value_length_in_bits(plc4c_data* data_item) {
  uint16_t lengthInBits = 0;

  return lengthInBits;
}

