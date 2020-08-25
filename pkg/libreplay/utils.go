package libreplay

import (
	"fmt"
	"github.com/monzo/typhon"
	"os"
)

func OpenWritableFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0660)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func Default404Handler(app App) typhon.Service {
	return func(req typhon.Request) typhon.Response {
		// TODO: Change this body to a default 404 page
		response := req.Response(nil)
		response.StatusCode = 404
		return response
	}
}

func VerifyRequest(r *typhon.Request) error {
	if _, ok := r.MultipartForm.File["audio"]; !ok || len(r.MultipartForm.File["audio"]) < 1 {
		return fmt.Errorf("the field audio is absent")
	}
	return nil
}

func CleanStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

type MockFile string

func (m *MockFile) WriteString(s string) (n int, err error) {
	*m = MockFile(string(*m) + s)
	return len(s), nil
}

func (m MockFile) Close() error {
	return nil
}
