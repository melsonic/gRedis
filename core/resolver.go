package core

import (
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
	return util.SerializeBulkString(value)
}

func Resolver(req []any) string {
	var response string = util.SerializeString("OK")
	if len(req) == 0 {
		return util.SerializeError(util.NO_ARG_ERROR)
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
	}
	return response
}
