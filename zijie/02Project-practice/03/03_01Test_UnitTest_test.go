package _3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_04Test(t *testing.T)  {
	output := "tom"
	output2 := "tom1"
	//断言比较
	assert.Equal(t, output, output2)
}