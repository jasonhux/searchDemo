package data

import (
	"encoding/json"
	"io/ioutil"
)

type Serializer interface {
	ReadFile(filePath string) ([]byte, error)
	Unmarshal(dataForSerialize []byte, v interface{}) error
}

type serializer struct{}

func NewSerializer() Serializer {
	return &serializer{}
}

func (s *serializer) ReadFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

func (s *serializer) Unmarshal(dataForSerialize []byte, v interface{}) error {
	return json.Unmarshal(dataForSerialize, v)
}
