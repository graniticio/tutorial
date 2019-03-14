package artist

import (
	"context"
	"github.com/graniticio/granitic/v2/logging"
	"github.com/graniticio/granitic/v2/types"
	"github.com/graniticio/granitic/v2/ws"
	"strings"
)

type GetLogic struct {
	EnvLabel string
	Log      logging.Logger
}

func (gl *GetLogic) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, q *ArtistQuery) {

	gl.Log.LogTracef("Request for artist with ID %d", q.ID)

	name := "Some Artist"

	if q.NormaliseName.Bool() {
		name = strings.ToUpper(name)
	}

	res.Body = Info{
		Name: name,
	}

}

type Info struct {
	Name string
}

type ArtistQuery struct {
	ID            int
	NormaliseName *types.NilableBool
}
