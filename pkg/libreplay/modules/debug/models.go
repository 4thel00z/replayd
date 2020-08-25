package debug

import "github.com/4thel00z/replayd/pkg/libreplay"

type GenericResponse struct {
	Message interface{} `json:"message"`
	Error   *string     `json:"error,omitempty"`
}

type GetRoutesResponse struct {
	Routes []libreplay.Route `json:"routes"`
	Error  *string         `json:"error,omitempty"`
}
