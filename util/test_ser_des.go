package util

import "fmt"

func UtilCheckDeserialize() {
	inputs := []string{
		"$-1\r\n",
		"*1\r\n$4\r\nping\r\n",
		"*2\r\n$4\r\necho\r\n$11\r\nhello world\r\n",
		"*2\r\n$3\r\nget\r\n$3\r\nkey\r\n",
		"+OK\r\n",
		"-Error message\r\n",
		"$0\r\n\r\n",
		"+hello world\r\n",
	}
	for _, in := range inputs {
		fmt.Println()
		var result []any
		var succ bool
		fmt.Println(in)
		result, succ = Deserialize([]byte(in), result)
		fmt.Println(succ)
		for _, r := range result {
			fmt.Printf("%v ", r)
		}
		fmt.Println()
		fmt.Println("===============================")
		fmt.Println()
	}
}


func UtilCheckSerialize() {
	inputs := []any {
		"hello world",
		1234,
		"12.34",
		[]any {"cristiano", 90, nil},
		nil,
	}
	for _, in := range inputs {
		fmt.Println()
		fmt.Printf("result : %s \n", Serialize(in))
		fmt.Println()
	}
}
