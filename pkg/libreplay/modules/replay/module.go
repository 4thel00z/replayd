package replay

import "github.com/4thel00z/replayd/pkg/libreplay"

type Replay struct{}

var (
	Module = Replay{}
)

func (r Replay) Version() string {
	return ""
}

func (r Replay) Namespace() string {
	return ""
}

func (r Replay) Routes() []libreplay.Route {
	return []libreplay.Route{
		// Add route definitions here
		{
			Path:        "/",
			Method:      "*",
			CurlExample: "curl http://<addr>/<version>/<namespace>/routes",
		},
	}
}
func (r Replay) HandlerById(i int) libreplay.Service {
	switch i {
	// Add handlers for routes here
	case 0:
		return ReplayHandler
	}
	// This makes the server return a 404 by default
	return nil
}
