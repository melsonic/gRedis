package util

import (
	"bytes"
	"strconv"
)

func Deserialize(input []byte, result []any) ([]any, bool) {
	if len(input) == 0 {
		return result, true
	}
	var t byte = input[0]
	var succ bool
	input_original := input
	input = input[1:]
	switch t {
	case PlusByte:
		result, succ = DeserializeString(input, result)
	case MinusByte:
		result, succ = DeserializeError(input, result)
	case ColonByte:
		result, succ = DeserializeInt(input, result)
	case DollerByte:
		result, succ = DeserializeBulkString(input, result)
	case StarByte:
		result, succ = DeserializeArray(input, result)
	case NullByte:
		succ = true
	default:
		result, succ = DeserializeInlineCommand(input_original, result)
	}
	return result, succ
}

func DeserializeString(input []byte, result []any) ([]any, bool) {
	var subsIndex int = bytes.Index(input, CRLFByte)
	if subsIndex == -1 {
		return result, false
	}
	var output string = string(input[0:subsIndex])
	result = append(result, output)
	return Deserialize(input[subsIndex+CRLFLen:], result)
}

func DeserializeError(input []byte, result []any) ([]any, bool) {
	var subsIndex int = bytes.Index(input, CRLFByte)
	if subsIndex == -1 {
		return result, false
	}
	var output string = string(input[0:subsIndex])
	result = append(result, output)
	return Deserialize(input[subsIndex+CRLFLen:], result)
}

func DeserializeInt(input []byte, result []any) ([]any, bool) {
	var output int64
	var subsIndex int = bytes.Index(input, CRLFByte)
	if subsIndex == -1 {
		return result, false
	}
	var outputErr error
	output, outputErr = strconv.ParseInt(string(input[0:subsIndex]), 10, 64)
	if outputErr != nil {
		return result, false
	}
	result = append(result, output)
	return Deserialize(input[subsIndex+CRLFLen:], result)
}

func DeserializeBulkString(input []byte, result []any) ([]any, bool) {
	var subsIndex int = bytes.Index(input, CRLFByte)
	if subsIndex == -1 {
		return result, false
	}
	outputLen, outputLenErr := strconv.ParseInt(string(input[0:subsIndex]), 10, 64)
	if outputLenErr != nil {
		return result, false
	}
	input = input[subsIndex+CRLFLen:]
	var succ bool = true
	if outputLen == -1 {
		result = append(result, nil)
		result, succ = Deserialize(input, result)
	} else {
		result, succ = DeserializeString(input, result)
	}
	return result, succ
}

func DeserializeArray(input []byte, result []any) ([]any, bool) {
	var subsIndex int = bytes.Index(input, CRLFByte)
	if subsIndex == -1 {
		return result, false
	}
	outputLen, outputLenErr := strconv.ParseInt(string(input[0:subsIndex]), 10, 64)
	if outputLenErr != nil {
		return result, false
	}
	input = input[subsIndex+CRLFLen:]
	if outputLen == -1 {
		result = append(result, nil)
	}
	return Deserialize(input, result)
}

func DeserializeInlineCommand(input []byte, result []any) ([]any, bool) {
	var insideString bool = false
	var currBytes []byte
	for _, inputByte := range input {
		if inputByte == 0 {
			break
		}
		if inputByte == DoubleQuoteByte {
			insideString = !insideString
		} else if inputByte == SpaceByte && !insideString {
			var subsIndex int = bytes.Index(currBytes, CRLFByte)
			if subsIndex != -1 {
				currBytes = currBytes[:len(currBytes)-2]
			}
			result = append(result, string(currBytes))
			currBytes = nil
		} else {
			currBytes = append(currBytes, inputByte)
		}
	}
	var subsIndex int = bytes.Index(currBytes, CRLFByte)
	if subsIndex != -1 {
		currBytes = currBytes[:len(currBytes)-2]
	}
	result = append(result, string(currBytes))
	return result, true
}
