package artist

import (
	"context"
	"fmt"
	"github.com/graniticio/granitic/v2/ws"
)

type GetLogic struct {
	EnvLabel string
}

func (gl *GetLogic) Process(ctx context.Context, req *ws.Request, res *ws.Response) {

	an := fmt.Sprintf("Hello, %s!", gl.EnvLabel)

	res.Body = Info{
		Name: an,
	}
}

type Info struct {
	Name string
}
