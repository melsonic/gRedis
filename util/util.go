package util 

func SetBytesToZero(buffer []byte) {
	var n int = len(buffer)
	for i := 0; i < n; i++ {
		buffer[i] = 0
	}
}
