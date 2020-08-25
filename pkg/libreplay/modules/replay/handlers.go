package replay

import (
	"github.com/4thel00z/replayd/pkg/libreplay"
	"github.com/monzo/typhon"
)

func ReplayHandler(app libreplay.App) typhon.Service {
	return func(req typhon.Request) typhon.Response {

		response := req.Response("not implemented")
		response.StatusCode = 200
		return response
	}
}
