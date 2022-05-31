package common

import (
	uuid "github.com/satori/go.uuid"
	"path"
)

func GenerateFileName(fileName string)string  {
	return uuid.NewV4().String() + path.Ext(fileName)
}