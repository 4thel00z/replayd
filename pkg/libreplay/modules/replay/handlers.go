package replay

import (
	"fmt"
	"github.com/4thel00z/replayd"
	"github.com/4thel00z/replayd/pkg/libreplay"
	"github.com/monzo/typhon"
)

func ReplayHandler(app libreplay.App) typhon.Service {
	return func(req typhon.Request) typhon.Response {

		request, err := replayd.ConvertRequest(req.Request)

		if err != nil {
			response := req.Response("replayd.ConvertRequest is broken")
			response.StatusCode = 503
			return response
		}

		serialized, err := replayd.Serialize(request)
		if err != nil {
			response := req.Response("replayd.Serialize is broken")
			response.StatusCode = 503
			return response
		}

		path, err := app.GenerateUniquePath()
		if err != nil {
			response := req.Response("app.GenerateUniquePath is broken")
			response.StatusCode = 503
			return response
		}

		fmt.Println("Generated path:", path)

		file, err := libreplay.OpenWritableFile(path)

		if err != nil {
			response := req.Response("libreplay.OpenWritableFile is broken")
			response.StatusCode = 503
			return response
		}

		defer func() {
			_ = file.Close()
		}()

		n, err := file.WriteString(serialized)

		if err != nil {
			response := req.Response("file.write is broken")
			response.StatusCode = 503
			return response
		}

		response := req.Response(fmt.Sprintf("written %d bytes to: %s", n, path))
		response.StatusCode = 200
		return response
	}
}
