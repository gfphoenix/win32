// Copyright 2013 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package pdh

// PDH error codes, which can be returned by all Pdh* functions. Taken from mingw-w64 pdhmsg.h
const (
	PDH_CSTATUS_VALID_DATA                     = 0x00000000 // The returned data is valid.
	PDH_CSTATUS_NEW_DATA                       = 0x00000001 // The return data value is valid and different from the last sample.
	PDH_CSTATUS_NO_MACHINE                     = 0x800007D0 // Unable to connect to the specified computer, or the computer is offline.
	PDH_CSTATUS_NO_INSTANCE                    = 0x800007D1
	PDH_MORE_DATA                              = 0x800007D2 // The PdhGetFormattedCounterArray* function can return this if there's 'more data to be displayed'.
	PDH_CSTATUS_ITEM_NOT_VALIDATED             = 0x800007D3
	PDH_RETRY                                  = 0x800007D4
	PDH_NO_DATA                                = 0x800007D5 // The query does not currently contain any counters (for example, limited access)
	PDH_CALC_NEGATIVE_DENOMINATOR              = 0x800007D6
	PDH_CALC_NEGATIVE_TIMEBASE                 = 0x800007D7
	PDH_CALC_NEGATIVE_VALUE                    = 0x800007D8
	PDH_DIALOG_CANCELLED                       = 0x800007D9
	PDH_END_OF_LOG_FILE                        = 0x800007DA
	PDH_ASYNC_QUERY_TIMEOUT                    = 0x800007DB
	PDH_CANNOT_SET_DEFAULT_REALTIME_DATASOURCE = 0x800007DC
	PDH_CSTATUS_NO_OBJECT                      = 0xC0000BB8
	PDH_CSTATUS_NO_COUNTER                     = 0xC0000BB9 // The specified counter could not be found.
	PDH_CSTATUS_INVALID_DATA                   = 0xC0000BBA // The counter was successfully found, but the data returned is not valid.
	PDH_MEMORY_ALLOCATION_FAILURE              = 0xC0000BBB
	PDH_INVALID_HANDLE                         = 0xC0000BBC
	PDH_INVALID_ARGUMENT                       = 0xC0000BBD // Required argument is missing or incorrect.
	PDH_FUNCTION_NOT_FOUND                     = 0xC0000BBE
	PDH_CSTATUS_NO_COUNTERNAME                 = 0xC0000BBF
	PDH_CSTATUS_BAD_COUNTERNAME                = 0xC0000BC0 // Unable to parse the counter path. Check the format and syntax of the specified path.
	PDH_INVALID_BUFFER                         = 0xC0000BC1
	PDH_INSUFFICIENT_BUFFER                    = 0xC0000BC2
	PDH_CANNOT_CONNECT_MACHINE                 = 0xC0000BC3
	PDH_INVALID_PATH                           = 0xC0000BC4
	PDH_INVALID_INSTANCE                       = 0xC0000BC5
	PDH_INVALID_DATA                           = 0xC0000BC6 // specified counter does not contain valid data or a successful status code.
	PDH_NO_DIALOG_DATA                         = 0xC0000BC7
	PDH_CANNOT_READ_NAME_STRINGS               = 0xC0000BC8
	PDH_LOG_FILE_CREATE_ERROR                  = 0xC0000BC9
	PDH_LOG_FILE_OPEN_ERROR                    = 0xC0000BCA
	PDH_LOG_TYPE_NOT_FOUND                     = 0xC0000BCB
	PDH_NO_MORE_DATA                           = 0xC0000BCC
	PDH_ENTRY_NOT_IN_LOG_FILE                  = 0xC0000BCD
	PDH_DATA_SOURCE_IS_LOG_FILE                = 0xC0000BCE
	PDH_DATA_SOURCE_IS_REAL_TIME               = 0xC0000BCF
	PDH_UNABLE_READ_LOG_HEADER                 = 0xC0000BD0
	PDH_FILE_NOT_FOUND                         = 0xC0000BD1
	PDH_FILE_ALREADY_EXISTS                    = 0xC0000BD2
	PDH_NOT_IMPLEMENTED                        = 0xC0000BD3
	PDH_STRING_NOT_FOUND                       = 0xC0000BD4
	PDH_UNABLE_MAP_NAME_FILES                  = 0x80000BD5
	PDH_UNKNOWN_LOG_FORMAT                     = 0xC0000BD6
	PDH_UNKNOWN_LOGSVC_COMMAND                 = 0xC0000BD7
	PDH_LOGSVC_QUERY_NOT_FOUND                 = 0xC0000BD8
	PDH_LOGSVC_NOT_OPENED                      = 0xC0000BD9
	PDH_WBEM_ERROR                             = 0xC0000BDA
	PDH_ACCESS_DENIED                          = 0xC0000BDB
	PDH_LOG_FILE_TOO_SMALL                     = 0xC0000BDC
	PDH_INVALID_DATASOURCE                     = 0xC0000BDD
	PDH_INVALID_SQLDB                          = 0xC0000BDE
	PDH_NO_COUNTERS                            = 0xC0000BDF
	PDH_SQL_ALLOC_FAILED                       = 0xC0000BE0
	PDH_SQL_ALLOCCON_FAILED                    = 0xC0000BE1
	PDH_SQL_EXEC_DIRECT_FAILED                 = 0xC0000BE2
	PDH_SQL_FETCH_FAILED                       = 0xC0000BE3
	PDH_SQL_ROWCOUNT_FAILED                    = 0xC0000BE4
	PDH_SQL_MORE_RESULTS_FAILED                = 0xC0000BE5
	PDH_SQL_CONNECT_FAILED                     = 0xC0000BE6
	PDH_SQL_BIND_FAILED                        = 0xC0000BE7
	PDH_CANNOT_CONNECT_WMI_SERVER              = 0xC0000BE8
	PDH_PLA_COLLECTION_ALREADY_RUNNING         = 0xC0000BE9
	PDH_PLA_ERROR_SCHEDULE_OVERLAP             = 0xC0000BEA
	PDH_PLA_COLLECTION_NOT_FOUND               = 0xC0000BEB
	PDH_PLA_ERROR_SCHEDULE_ELAPSED             = 0xC0000BEC
	PDH_PLA_ERROR_NOSTART                      = 0xC0000BED
	PDH_PLA_ERROR_ALREADY_EXISTS               = 0xC0000BEE
	PDH_PLA_ERROR_TYPE_MISMATCH                = 0xC0000BEF
	PDH_PLA_ERROR_FILEPATH                     = 0xC0000BF0
	PDH_PLA_SERVICE_ERROR                      = 0xC0000BF1
	PDH_PLA_VALIDATION_ERROR                   = 0xC0000BF2
	PDH_PLA_VALIDATION_WARNING                 = 0x80000BF3
	PDH_PLA_ERROR_NAME_TOO_LONG                = 0xC0000BF4
	PDH_INVALID_SQL_LOG_FORMAT                 = 0xC0000BF5
	PDH_COUNTER_ALREADY_IN_QUERY               = 0xC0000BF6
	PDH_BINARY_LOG_CORRUPT                     = 0xC0000BF7
	PDH_LOG_SAMPLE_TOO_SMALL                   = 0xC0000BF8
	PDH_OS_LATER_VERSION                       = 0xC0000BF9
	PDH_OS_EARLIER_VERSION                     = 0xC0000BFA
	PDH_INCORRECT_APPEND_TIME                  = 0xC0000BFB
	PDH_UNMATCHED_APPEND_COUNTER               = 0xC0000BFC
	PDH_SQL_ALTER_DETAIL_FAILED                = 0xC0000BFD
	PDH_QUERY_PERF_DATA_TIMEOUT                = 0xC0000BFE
)

// Formatting options for GetFormattedCounterValue().
const (
	PDH_FMT_RAW          = 0x00000010
	PDH_FMT_ANSI         = 0x00000020
	PDH_FMT_UNICODE      = 0x00000040
	PDH_FMT_LONG         = 0x00000100 // Return data as a long int.
	PDH_FMT_DOUBLE       = 0x00000200 // Return data as a double precision floating point real.
	PDH_FMT_LARGE        = 0x00000400 // Return data as a 64 bit integer.
	PDH_FMT_NOSCALE      = 0x00001000 // can be OR-ed: Do not apply the counter's default scaling factor.
	PDH_FMT_1000         = 0x00002000 // can be OR-ed: multiply the actual value by 1,000.
	PDH_FMT_NODATA       = 0x00004000 // can be OR-ed: unknown what this is for, MSDN says nothing.
	PDH_FMT_NOCAP100     = 0x00008000 // can be OR-ed: do not cap values > 100.
	PERF_DETAIL_COSTLY   = 0x00010000
	PERF_DETAIL_STANDARD = 0x0000FFFF
)
