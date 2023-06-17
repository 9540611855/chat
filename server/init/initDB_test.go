package init

import (
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestInitDB(t *testing.T) {
	c.Convey("TestInitDB should return nil", t, func() {
		c.So(InitDB(), c.ShouldBeNil)
	})

}
