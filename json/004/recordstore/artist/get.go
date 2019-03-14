package artist

import (
	"context"
	"fmt"
	"github.com/graniticio/granitic/v2/logging"
	"github.com/graniticio/granitic/v2/ws"
)

type GetLogic struct {
	EnvLabel string
	Log      logging.Logger
}

func (gl *GetLogic) Process(ctx context.Context, req *ws.Request, res *ws.Response) {

	an := fmt.Sprintf("Hello, %s!", gl.EnvLabel)

	res.Body = Info{
		Name: an,
	}

	log := gl.Log

	log.LogInfof("Environment is set to '%s'", gl.EnvLabel)
	log.LogTracef("Request served")
}

type Info struct {
	Name string
}
