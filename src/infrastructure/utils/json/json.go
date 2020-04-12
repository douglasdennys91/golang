package json

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func ParserJSON(data io.ReadCloser) (interface{}, error) {
	var param interface{}
	parser, err := ioutil.ReadAll(data)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(parser, &param)
	return param, nil
}

func ParserByte(data io.ReadCloser) ([]byte, error) {
	parser, err := ioutil.ReadAll(data)
	if err != nil {
		return nil, err
	}
	return parser, nil
}