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

/// Error messages 
var WRONG_ARG_NUM_ERROR string = "wrong arguments passed" 
var	NO_ARG_ERROR string = "No argument passed"

/// Response messages 
var OK_RESPONSE string = "OK"

/// Empty byte 
var EMPTY_READ_BYTE []byte = make([]byte, 1024)

var PORT int = 6379
var ServerAddress string = fmt.Sprintf(":%d", PORT)

