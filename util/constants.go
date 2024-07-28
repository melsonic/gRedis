package util

import "fmt"

var CRLF string = "\r\n"
var CRLFByte = []byte(CRLF)
var CRLFLen int = len(CRLF)

var PlusByte byte = '+'
var MinusByte byte = '-'
var ColonByte byte = ':'
var DollerByte byte = '$'
var StarByte byte = '*'
var NullByte byte = 0

/// Command Names 
var PINGCommand string = "PING"
var PONGResponse string = "PONG"
var ECHOCommand string = "ECHO"
var SETCommand string = "SET"
var GETCommand string = "GET"
var EXISTSCommand string = "EXISTS"
var DELCommand string = "DEL"
var INCRCommand string = "INCR"
var DECRCommand string = "DECR"
var LPUSHCommand string = "LPUSH"
var RPUSHCommand string = "RPUSH"

/// Error messages 
var WRONG_ARG_NUM_ERROR string = "ERR wrong number of arguments for command" 
var	NO_ARG_ERROR string = "No argument passed"
var NOT_INTEGER_VALUE string = "value is not an integer or out of range"
var WRONGTYPE_VALUE_IN_KEY string = "WRONGTYPE Operation against a key holding the wrong kind of value"

/// Response messages 
var OK_RESPONSE string = "OK"

/// Empty byte 
var EMPTY_READ_BYTE []byte = make([]byte, 1024)

var PORT int = 6379
var ServerAddress string = fmt.Sprintf(":%d", PORT)

