package hash

import (
	"fmt"
	"testing"
)

func TestNewHash(t *testing.T) {
	hash := NewHash(
		WithDriver(BCRYPT),
		WithConst(DEFAULTCOST),
	)
	password := "123456"
	bytes, err := hash.Make([]byte(password))
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(bytes))

	err = hash.Check(bytes, []byte(password))
	if err != nil {
		t.Error(err)
	}
	fmt.Println("密码正确")
}
