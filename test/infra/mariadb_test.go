package mongodb_test

import (
	testing "testing"

	. "github.com/smartystreets/goconvey/convey"

	"interview-test-free-fair/pkg/infra/mariadb"
	sys "interview-test-free-fair/pkg/infra/system"
)

func TestSpec(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping integration test")
	}

	Convey("Given a valid Mariadb URI", t, func() {

		sys.Properties = sys.AppProperties{
			Mariadb: "free_fair:free_fair_pdw@tcp(localhost:3306)/free_fair_adm",
		}

		Convey("When connect on Mariadb", func() {
			mariadb.Connect()
			defer mariadb.Disconnect()

			Convey("Then Mariadb session it's alive", func() {
				So(mariadb.IsAlive(), ShouldEqual, true)
			})
		})
	})
}
