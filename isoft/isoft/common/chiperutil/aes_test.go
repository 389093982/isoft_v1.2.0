package chiperutil

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func Test_Aes(t *testing.T) {
	key := []byte("0123456789abcdefffffffffffeeeeere")
	result, err := AesEncrypt([]byte("hello worrftfdfdld"), key)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	origData, err := AesDecrypt(result, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}
