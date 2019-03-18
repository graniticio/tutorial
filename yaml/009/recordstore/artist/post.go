package artist

import (
	"context"
	"github.com/graniticio/granitic/v2/logging"
	"github.com/graniticio/granitic/v2/rdbms"
	"github.com/graniticio/granitic/v2/types"
	"github.com/graniticio/granitic/v2/ws"
	"net/http"
)

type PostLogic struct {
	Log             logging.Logger
	DbClientManager rdbms.ClientManager
}

func (pl *PostLogic) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, s *Submission) {
	// Obtain a Client from the rdbms.ClientManager injected into this component
	dbc, _ := pl.DbClientManager.Client()

	defer dbc.Rollback()

	// Start a database transaction
	dbc.StartTransaction()

	// Declare a variable to capture the ID of the newly inserted artist
	var id int64

	// Execute the insert, storing the generated ID in our variable
	if err := dbc.InsertCaptureQIDParams("CREATE_ARTIST", &id, s); err != nil {
		// Something went wrong when communicating with the database - return HTTP 500
		pl.Log.LogErrorf(err.Error())
		res.HTTPStatus = http.StatusInternalServerError
	}

	// Insert a row for each related artist
	params := make(map[string]interface{})
	params["ArtistID"] = id

	for _, raID := range s.RelatedArtists {
		params["RelatedArtistID"] = raID

		if _, err := dbc.InsertQIDParams("RELATE_ARTIST", params); err != nil {
			// Something went wrong inserting the relationship
			pl.Log.LogErrorf(err.Error())
			res.HTTPStatus = http.StatusInternalServerError

			return
		}

	}

	// Commit the transaction
	dbc.CommitTransaction()

	// Use the new ID as the HTTP response, wrapped in a struct
	res.Body = CreatedResource{id}
}

type Submission struct {
	Name            *types.NilableString
	FirstYearActive *types.NilableInt64
	RelatedArtists  []int64
}

type CreatedResource struct {
	ID int64
}
