package util

import (
	"fmt"
	"reflect"
)

func Serialize(input any) string {
	var result string
	var inputValue reflect.Value = reflect.ValueOf(input)
	if(!inputValue.IsValid()) {
		result = SerializeNil()
		return result
	} 
	inputKind := inputValue.Kind()
	switch inputKind {
		case reflect.String:
			inputStr := input.(string)
			result = SerializeString(inputStr)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			result = SerializeNumber(input)
		case reflect.Slice, reflect.Array:
		  inputArray := input.([]any)
			result = SerializeArray(inputArray)
		default:
			result = SerializeError(input.(string))
	}
	return result
}

func SerializeNil() string {
	return "$-1\r\n"
}

func SerializeString(input string) string {
	return fmt.Sprintf("+%s\r\n", input)
}

func SerializeBulkString(input string) string {
	var n int = len(input)
	return fmt.Sprintf("$%d\r\n%s\r\n", n, input)
}

func SerializeNumber(input any) string {
	return fmt.Sprintf(":%d\r\n", input)
}

func SerializeArray(input []any) string {
	var size int = len(input)
	if size == 0 {
		return "*-1\r\n"
	}
	var output string = fmt.Sprintf("*%d\r\n", size)
	for _, item := range input {
		output += Serialize(item)
	}
	return output
}

func SerializeError(input string) string {
	return fmt.Sprintf("-%s\r\n", input)
}
