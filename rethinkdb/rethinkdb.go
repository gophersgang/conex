package rethinkdb

import (
	"testing"

	gorethink "gopkg.in/gorethink/gorethink.v3"

	"github.com/omeid/conex"
)

// Image to use for the box.
var Image = "rethinkdb"

func init() {
	conex.Require(func() string { return Image })
}

// Box returns a RethinkDB client connect to a RethinkDB
// container based on your provided tags.
func Box(t *testing.T, db string) (*gorethink.Session, func()) {
	c, done := conex.Box(t, Image)

	opts := gorethink.ConnectOpts{
		Address:  c.Address(),
		Database: db,
	}

	sesh, err := gorethink.Connect(opts)
	if err != nil {
		t.Fatal(err)
	}

	return sesh, done
}
