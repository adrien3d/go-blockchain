package store

import (
	"context"
	"fmt"
	"github.com/adrien3d/monarch/utils"
	"github.com/flimzy/kivik"
	_ "github.com/go-kivik/couchdb" // The CouchDB driver
)

type CouchDB struct {
	DB *kivik.DB
}

func GetCouch() CouchDB {
	client, err := kivik.New(context.Background(), "couch", "http://adrien3d:adchapwd@localhost:5984/")
	utils.CheckErr(err)

	dbExists, err := client.DBExists(context.TODO(), "blocks")
	utils.CheckErr(err)

	if dbExists == false {
		err = client.CreateDB(context.Background(), "blocks")
	}

	db, err := client.DB(context.Background(), "blocks")

	return CouchDB{db}
}

func (couch CouchDB) GetValue(key string, value interface{}) error {
	row, err := couch.DB.Get(context.TODO(), key)
	utils.CheckErr(err)
	fmt.Println(row)

	return err
}
func (couch CouchDB) SetValue(key string, value interface{}) error {
	rev, err := couch.DB.Put(context.TODO(), key, value)
	utils.CheckErr(err)
	fmt.Printf("%d inserted with revision %s\n", key, rev)

	return err
}
