package db

import (
	"github.com/graniticio/granitic/v2/logging"
	"github.com/graniticio/granitic/v2/rdbms"
)

type ArtistExistsChecker struct {
	DbClientManager rdbms.ClientManager
	Log             logging.Logger
}

func (aec *ArtistExistsChecker) ValidInt64(id int64) (bool, error) {

	dbc, _ := aec.DbClientManager.Client()

	var count int64

	if _, err := dbc.SelectBindSingleQIDParam("CHECK_ARTIST", "ID", id, &count); err != nil {
		return false, err
	} else {
		return count > 0, nil
	}
}
