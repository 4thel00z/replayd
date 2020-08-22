package replayd

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"io"
	"io/ioutil"
	"net/http"
)

func Serialize(req Request) (string, error) {
	buffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(req)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buffer.Bytes()), nil
}

// go binary decoder
func Deserialize(str string) (Request, error) {
	req := Request{}
	rawBytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return Request{}, err
	}
	buffer := bytes.Buffer{}
	buffer.Write(rawBytes)
	decoder := gob.NewDecoder(&buffer)
	err = decoder.Decode(&req)

	if err != nil {
		return Request{}, err
	}
	return req, nil
}

func SaveRequest(req http.Request, stringWriterCloser StringWriterCloser) error {
	defer func() {
		_ = stringWriterCloser.Close()
	}()

	body, err := ioutil.ReadAll(req.Body)

	defer func() {
		_ = req.Body.Close()
	}()

	if err != nil {
		return err
	}

	// We do this, so we can use SaveRequest in a middleware and relay the request
	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	err = req.ParseForm()
	if err != nil {
		return err
	}

	payload, err := Serialize(Request{
		HTTPVersion: req.Proto,
		Host:        req.Host,
		Method:      req.Method,
		Headers:     req.Header,
		Body:        body,
		Form:        req.Form,
	})

	if err != nil {
		return err
	}

	_, err = stringWriterCloser.WriteString(payload)

	if err != nil {
		return err
	}

	return nil
}



func RestoreRequest(reader io.Reader) (Request, error) {
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return Request{}, err
	}
	// If this doesn't work, try
	return Deserialize(string(body))
}

func Init() {
	gob.Register(Request{})
}
