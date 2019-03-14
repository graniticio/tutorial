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

	a := new(Info)
	a.Name = fmt.Sprintf("Hello, %s!", gl.EnvLabel)

	res.Body = a
}

type Info struct {
	Name string
}
