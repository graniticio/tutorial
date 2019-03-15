package artist

import (
	"context"
	"github.com/graniticio/granitic/v2/logging"
	"github.com/graniticio/granitic/v2/rdbms"
	"github.com/graniticio/granitic/v2/types"
	"github.com/graniticio/granitic/v2/ws"
	"net/http"
)

type GetLogic struct {
	EnvLabel        string
	Log             logging.Logger
	DbClientManager rdbms.ClientManager
}

func (gl *GetLogic) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, q *ArtistQuery) {

	// Obtain an RdmsClient from the rdbms.RdbmsClientManager injected into this component
	dbc, _ := gl.DbClientManager.Client()

	// Create a new object to store the results of our database call
	result := new(Info)

	// Call the database and populate our object
	if found, err := dbc.SelectBindSingleQIDParams("ARTIST_BY_ID", result, q); found {
		// Make our result object the body of the HTTP response we'll send
		res.Body = result

	} else if err != nil {
		// Something went wrong when communicating with the database - return HTTP 500
		gl.Log.LogErrorf(err.Error())
		res.HTTPStatus = http.StatusInternalServerError

	} else {
		// No results were returned by the database call - return HTTP 404
		res.HTTPStatus = http.StatusNotFound
	}

}

type Info struct {
	Name string
}

type ArtistQuery struct {
	ID            int
	NormaliseName *types.NilableBool
}
