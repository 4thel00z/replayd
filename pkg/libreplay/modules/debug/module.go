package debug

import "github.com/4thel00z/replayd/pkg/libreplay"

type Debug struct{}

var (
	Module = Debug{}
)

func (Y Debug) Version() string {
	return "v1"
}

func (Y Debug) Namespace() string {
	return "debug"
}

func (Y Debug) Routes() []libreplay.Route {
	return []libreplay.Route{
		// Add route definitions here
		{
			Path:        "routes",
			Method:      "GET",
			CurlExample: "curl http://<addr>/<version>/<namespace>/routes",
		},
	}
}
func (Y Debug) HandlerById(i int) libreplay.Service {
	switch i {
	// Add handlers for routes here
	case 0:
		return GetRoutesHandler
	}
	// This makes the server return a 404 by default
	return nil
}
