package util

import (
	"fmt"
)

var (
	CRLF     string = "\r\n"
	CRLFByte        = []byte(CRLF)
	CRLFLen  int    = len(CRLF)
)

var (
	PlusByte        byte = '+'
	MinusByte       byte = '-'
	ColonByte       byte = ':'
	DollerByte      byte = '$'
	StarByte        byte = '*'
	NullByte        byte = 0
	SpaceByte       byte = ' '
	DoubleQuoteByte byte = '"'
)

// / Command Names
var (
	PINGCommand   string = "PING"
	PONGResponse  string = "PONG"
	ECHOCommand   string = "ECHO"
	SETCommand    string = "SET"
	GETCommand    string = "GET"
	EXISTSCommand string = "EXISTS"
	DELCommand    string = "DEL"
	INCRCommand   string = "INCR"
	DECRCommand   string = "DECR"
	LPUSHCommand  string = "LPUSH"
	RPUSHCommand  string = "RPUSH"
)

// / Error messages
var (
	WRONG_ARG_NUM_ERROR    string = "ERR wrong number of arguments for command"
	NO_ARG_ERROR           string = "No argument passed"
	NOT_INTEGER_VALUE      string = "value is not an integer or out of range"
	WRONGTYPE_VALUE_IN_KEY string = "WRONGTYPE Operation against a key holding the wrong kind of value"
)

// / Response messages
var OK_RESPONSE string = "OK"

var (
	PORT          int    = 6379
	ServerAddress string = fmt.Sprintf(":%d", PORT)
)
