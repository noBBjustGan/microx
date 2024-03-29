package orm

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetEngine(t *testing.T) {
	c := Config{
		DriverName:     "mysql",
		DataSourceName: "root:123456@tcp(172.16.13.4:3306)/mx_passport",
		MaxIdleConn:    3,
		MaxOpenConn:    10,
	}

	Convey("should get engine", t, func() {
		engine := GetEngine(c)
		So(engine, ShouldNotBeNil)
	})
}

func TestDataSource_Open(t *testing.T) {
	c := Config{
		DriverName:     "mysql",
		DataSourceName: "root:123456@tcp(172.16.13.4:3306)/mx_passport",
		MaxIdleConn:    3,
		MaxOpenConn:    10,
	}

	dataSource := new(dataSource)
	Convey("open datasource", t, func() {
		Convey("should report error when close database if it's not opened", func() {
			err := dataSource.Close()
			So(err, ShouldEqual, DatabaseIsNotOpenedError)
		})
		Convey("should open database", func() {
			err := dataSource.Open(c)
			So(err, ShouldBeNil)
		})
		Convey("should close database", func() {
			err := dataSource.Close()
			So(err, ShouldBeNil)
		})
	})
}
