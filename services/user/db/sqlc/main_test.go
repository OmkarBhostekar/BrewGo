package user

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/omkarbhostekar/brewgo/services/user/util"
)

var testQueries *Queries
var testStore Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	testDb, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	testStore = NewStore(testDb)
	if testDb == nil {
		log.Printf("testDb is nil")
	}
	testQueries = New(testDb)

	os.Exit(m.Run())
}