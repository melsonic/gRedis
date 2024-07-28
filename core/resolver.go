package core

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/melsonic/gRedis/util"
)

// input contains only PING command arguments
func resolvePING(input []any) string {
	var n int = len(input)
	if n > 1 {
		return util.SerializeError(util.WRONG_ARG_NUM_ERROR)
	} else if n == 1 {
		// if argument present send bulk string
		return util.SerializeBulkString(input[0].(string))
	}
	// if no argument send simple string
	return util.SerializeString(util.PONGResponse)
}

// input contains only the ECHO command argument
func resolveECHO(input []any) string {
	var n int = len(input)
	if n > 1 {
		return util.SerializeError(util.WRONG_ARG_NUM_ERROR)
	} else if n == 0 {
		fmt.Println("echo : ", input)
		return util.SerializeError(util.NO_ARG_ERROR)
	}
	return util.SerializeBulkString(input[0].(string))
}

func resolveSET(input []any) string {
	var n int = len(input)
	if n != 2 {
		return util.SerializeError(util.WRONG_ARG_NUM_ERROR)
	}
	var key string = input[0].(string)
	var value string = input[1].(string)
	mu.Lock()
	core_data[key] = value
	mu.Unlock()
	return util.SerializeString(util.OK_RESPONSE)
}

func resolveGET(input []any) string {
	var n int = len(input)
	if n != 1 {
		return util.SerializeError("wrong")
	}
	var key string = input[0].(string)
	mu.Lock()
	value, valuePresent := core_data[key]
	mu.Unlock()
	if !valuePresent {
		return util.Serialize(nil)
	}
	return util.SerializeBulkString(value.(string))
}

func resolveEXISTS(input []any) string {
	var n int = len(input)
	if n < 1 {
		return util.SerializeError(util.WRONG_ARG_NUM_ERROR)
	}
	var count int = 0
	for _, key := range input {
		if _, keyPresent := core_data[key.(string)]; keyPresent {
			count = count + 1
		}
	}
	return util.SerializeNumber(count)
}

func resolveDEL(input []any) string {
	var n int = len(input)
	if n < 1 {
		return util.SerializeError(util.WRONG_ARG_NUM_ERROR)
	}
	var count int = 0
	for _, key := range input {
		if _, keyPresent := core_data[key.(string)]; keyPresent {
			delete(core_data, key.(string))
			count = count + 1
		}
	}
	return util.SerializeNumber(count)
}

func resolveINCR(input []any) string {
	var n int = len(input)
	if n != 1 {
		return util.SerializeError(util.WRONG_ARG_NUM_ERROR)
	}
	var key string = input[0].(string)
	val, keyPresent := core_data[key]
	var intVal int64 = 0
	var intValErr error
	if keyPresent {
		// val of type any
		intVal, intValErr = strconv.ParseInt(string(val.(string)), 10, 64)
		if intValErr != nil {
			// value is not of int64 type
			return util.SerializeError(util.NOT_INTEGER_VALUE)
		}
	}
	intVal = intVal + 1
	core_data[key] = strconv.FormatInt(intVal, 10)
	return util.Serialize(intVal)
}

func resolveDECR(input []any) string {
	var n int = len(input)
	if n != 1 {
		return util.SerializeError(util.WRONG_ARG_NUM_ERROR)
	}
	var key string = input[0].(string)
	val, valPresent := core_data[key]
	var intVal int64 = 0
	var intValErr error
	if valPresent {
		intVal, intValErr = strconv.ParseInt(val.(string), 10, 64)
		if intValErr != nil {
			return util.SerializeError(util.NOT_INTEGER_VALUE)
		}
	}
	intVal = intVal - 1
	core_data[key] = strconv.FormatInt(intVal, 10)
	return util.Serialize(intVal)
}

// lpush command
func resolveLPUSH(input []any) string {
	var n int = len(input)
	if n < 2 {
		return util.SerializeError(util.WRONG_ARG_NUM_ERROR)
	}
	var key string = input[0].(string)
	input = input[1:]
	temp_val, keyPresent := core_data[key]
	var value_array []string
	if keyPresent {
		var typeOfValue reflect.Kind = reflect.TypeOf(temp_val).Kind()
		if typeOfValue != reflect.Slice {
			return util.SerializeError(util.WRONGTYPE_VALUE_IN_KEY)
		}
		value_array = temp_val.([]string)
	} else {
		value_array = []string{}
	}
	var temp_arg_values []string
	for _, arg := range input {
		temp_arg_values = append([]string{arg.(string)}, temp_arg_values...)
	}
	value_array = append(temp_arg_values, value_array...)
	core_data[key] = value_array
	return util.Serialize(len(value_array))
}

func resolveRPUSH(input []any) string {
	var n int = len(input)
	if n < 2 {
		return util.SerializeError(util.WRONG_ARG_NUM_ERROR)
	}
	var key string = input[0].(string)
	input = input[1:]
	temp_val, keyPresent := core_data[key]
	var value_array []string
	if keyPresent {
		var typeOfValue reflect.Kind = reflect.TypeOf(temp_val).Kind()
		if typeOfValue != reflect.Slice {
			return util.SerializeError(util.WRONGTYPE_VALUE_IN_KEY)
		}
		value_array = temp_val.([]string)
	} else {
		value_array = []string{}
	}
	for _, arg := range input {
		value_array = append(value_array, arg.(string))
	}
	core_data[key] = value_array
	return util.Serialize(len(value_array))
}

func Resolver(req []any) string {
	var response string = util.SerializeError("NOT_OK")
	if len(req) == 0 {
		fmt.Println("req : ", req)
		// return util.SerializeError(util.NO_ARG_ERROR)
		return util.SerializeString("OK")
	}
	var t string
	t = strings.ToUpper(req[0].(string))
	req = req[1:]
	switch t {
	case util.PINGCommand:
		response = resolvePING(req)
	case util.ECHOCommand:
		response = resolveECHO(req)
	case util.SETCommand:
		response = resolveSET(req)
	case util.GETCommand:
		response = resolveGET(req)
	case util.EXISTSCommand:
		response = resolveEXISTS(req)
	case util.DELCommand:
		response = resolveDEL(req)
	case util.INCRCommand:
		response = resolveINCR(req)
	case util.DECRCommand:
		response = resolveDECR(req)
	case util.LPUSHCommand:
		response = resolveLPUSH(req)
	case util.RPUSHCommand:
		response = resolveRPUSH(req)
	}
	return response
}
