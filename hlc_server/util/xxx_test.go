package util

import (
	"encoding/hex"
	"fmt"
	"github.com/EducationEKT/EKT/crypto"
	"testing"
)

const salt_code = "_salt_code_"

func TestPrice(t *testing.T) {
	//if success, err := regexp.MatchString("^((13[0-9])|(14[1]|[4-9])|(15([0-3]|[5-9]))|(16[2])|(16[5-7])|(17[0-3])|(17[5-8])|(18[0-9])|(19[1|8|9]))\\d{8}$", "19926407850"); !success || err != nil {
	//	fmt.Println(err, success)
	//}

	fmt.Println(hex.EncodeToString(crypto.Sha3_256([]byte("9984754" + salt_code))))
}
