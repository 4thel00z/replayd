package replayd

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"github.com/4thel00z/replayd/pkg/libreplay"
	"io"
	"io/ioutil"
	"net/http"
)

func Serialize(req libreplay.Request) (string, error) {
	buffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(req)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buffer.Bytes()), nil
}

// go binary decoder
func Deserialize(str string) (libreplay.Request, error) {
	req := libreplay.Request{}
	rawBytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return libreplay.Request{}, err
	}
	buffer := bytes.Buffer{}
	buffer.Write(rawBytes)
	decoder := gob.NewDecoder(&buffer)
	err = decoder.Decode(&req)

	if err != nil {
		return libreplay.Request{}, err
	}
	return req, nil
}

func SaveRequest(req http.Request, stringWriterCloser libreplay.StringWriterCloser) error {
	defer func() {
		_ = stringWriterCloser.Close()
	}()

	request, err := ToInternalRequest(req)
	if err != nil {
		return err
	}
	payload, err := Serialize(request)

	if err != nil {
		return err
	}

	_, err = stringWriterCloser.WriteString(payload)

	if err != nil {
		return err
	}

	return nil
}

func ToInternalRequest(req http.Request) (libreplay.Request, error) {
	body, err := ioutil.ReadAll(req.Body)

	defer func() {
		_ = req.Body.Close()
	}()

	if err != nil {
		return libreplay.Request{}, err
	}
	// We do this, so we can use SaveRequest in a middleware and relay the request
	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	err = req.ParseForm()
	if err != nil {
		return libreplay.Request{}, err
	}

	request := libreplay.Request{
		HTTPVersion: req.Proto,
		URL:         req.RequestURI	,
		Method:      req.Method,
		Headers:     req.Header,
		Body:        body,
		Form:        req.Form,
	}

	return request, nil
}

func ToHTTPRequest(req libreplay.Request) (*http.Request, error) {
	request, err := http.NewRequest(req.Method, req.URL, bytes.NewReader(req.Body))
	if err != nil {
		return nil, err
	}

	request.Header = req.Headers
	request.Form = req.Form
	request.Proto = req.HTTPVersion

	return request, nil
}

func RestoreRequest(reader io.Reader) (libreplay.Request, error) {
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return libreplay.Request{}, err
	}
	// If this doesn't work, try
	return Deserialize(string(body))
}

func Init() {
	gob.Register(libreplay.Request{})
}

func Invoke(request libreplay.Request) (*http.Response, error) {
	httpRequest, err := ToHTTPRequest(request)
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(httpRequest)
}
