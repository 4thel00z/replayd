package debug

import (
	"github.com/4thel00z/replayd/pkg/libreplay"
	"github.com/monzo/typhon"
)

func GetRoutesHandler(app libreplay.App) typhon.Service {
	return func(req typhon.Request) typhon.Response {

		response := req.Response(&GetRoutesResponse{
			Routes: app.Routes(),
		})

		response.StatusCode = 200
		return response
	}
}
