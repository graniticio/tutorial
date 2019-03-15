package artist

import (
	"context"
	"github.com/graniticio/granitic/v2/logging"
	"github.com/graniticio/granitic/v2/types"
	"github.com/graniticio/granitic/v2/ws"
)

type PostLogic struct {
	Log logging.Logger
}

func (pl *PostLogic) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, s *Submission) {
	pl.Log.LogInfof("New artist: '%s'", s.Name)

	res.Body = CreatedResource{0}
}

type Submission struct {
	Name            *types.NilableString
	FirstYearActive *types.NilableInt64
}

type CreatedResource struct {
	ID int
}
