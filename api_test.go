package replayd

import (
	"bytes"
	"fmt"
	"github.com/4thel00z/replayd/pkg/libreplay"
	"github.com/stretchr/testify/require"
	"net/http"
	"strings"
	"testing"
)

const jsonBytes = `{
	"payload":"test",
	"non_trivial_stuffz": 1337
}`

func TestSaveRequestAndRestore(t *testing.T) {
	originalRequest, err := http.NewRequest(http.MethodPost, "https://ransomware.host", bytes.NewReader([]byte(jsonBytes)))
	if err != nil {
		t.Error(err)
	}
	file := libreplay.MockFile("")
	err = SaveRequest(*originalRequest, &file)
	if err != nil {
		t.Error(err)
	}
	request, err := RestoreRequest(strings.NewReader(string(file)))
	if err != nil {
		t.Error(err)
	}

	require.Equal(t, strings.ToLower(http.MethodPost), strings.ToLower(request.Method))
	require.Equal(t, []byte(jsonBytes), request.Body)
	require.Equal(t, "ransomware.host", request.URL)

	fmt.Println(request)
}
