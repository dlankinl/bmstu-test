//go:build e2e

package e2e

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	testDB := SetupTestDatabase()
	defer testDB.TearDown()
	TestDbInstance = testDB.DbInstance
	err := SeedTestData(TestDbInstance)
	if err != nil {
		log.Fatalln(err)
	}
	os.Exit(m.Run())
}
