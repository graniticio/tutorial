package artist

import (
	"context"
	"github.com/graniticio/granitic/v2/ws"
)

type GetLogic struct{}

func (gl *GetLogic) Process(ctx context.Context, req *ws.Request, res *ws.Response) {

	res.Body = Info{
		Name: "Hello, World!",
	}
}

type Info struct {
	Name string
}
