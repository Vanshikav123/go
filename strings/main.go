//cannot be changed=> read only slice of bytes

/*two types
type 1 strings in""
type 2 strings in ``*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	c := `my name is vanshika` //raw literals
	//escaping characters does not works in this type of strings
	fmt.Println(c)
	//f is a datatype in go called as rune used to return the bytes of characters
	teststr := "test"
	for index, char := range teststr {
		fmt.Println("indes:", index, "char:", char, "character:", string(char))
	}
	//string(char) type assertion

	fmt.Println("type assertion of string into bytes", []byte(teststr))
	//trim strings
	str1 := "@@some@strings+@=@"

	fmt.Println(strings.Trim(str1, "@"))
	fmt.Println(strings.TrimLeft(str1, "@"))
	fmt.Println(strings.TrimRight(str1, "@"))
	//similarily we have TrimSpace, TrimPrefix and TrimSuffix

	//split string

	str2 := "welcome , to channel , go guru ji"
	fmt.Println("before split:-", str2)

	strarr := strings.Split(str2, ",")
	fmt.Println("after split:-", strarr, len(strarr), strarr[0])

	strarr1 := strings.SplitAfter(str2, ",")
	fmt.Println("after splitAfter:-", strarr1, len(strarr1), strarr1[1])

	strarr2 := strings.SplitAfterN(str2, ",", 2)
	fmt.Println("after splitAfterN:-", strarr2, len(strarr2), strarr2[1])

	//similarily we have ToUpper,ToLower and ToTitle
	//ToTitle make changes in the copy of string

}
