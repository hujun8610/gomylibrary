package util

import (
	"fmt"
	"testing"
)

func TestExecCmd(t *testing.T) {
	result, err := ExecCmd("ls -l /Users/hujun/")
	if err != nil {
		t.Errorf("result %s,expect %s", result, err)
	}
	fmt.Println(result)
}
