package zipstore

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	var encodeString = Encode("你好世界")
	fmt.Println(encodeString)
	fmt.Println(Decode(encodeString))
}
