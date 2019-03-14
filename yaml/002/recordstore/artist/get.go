package artist

import (
	"context"
	"github.com/graniticio/granitic/v2/ws"
)

type GetLogic struct{}

func (gl *GetLogic) Process(ctx context.Context, req *ws.Request, res *ws.Response) {

	a := new(Info)
	a.Name = "Hello, World!"

	res.Body = a
}

type Info struct {
	Name string
}
