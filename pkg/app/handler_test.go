package app_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"interview-test-free-fair/pkg/app"
)

func TestSpec(t *testing.T) {

	fairHandler := &app.FairHandler{Service: &FairServiceMock{}}

	Convey("Given a code of a fair", t, func() {

		res := httptest.NewRecorder()
		handler := http.HandlerFunc(fairHandler.Find)

		Convey("When find an existing code", func() {

			req, _ := http.NewRequest("GET", "/v1/fairies?fair_code=1020", nil)
			handler.ServeHTTP(res, req)

			Convey("Then return status code 200", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
			})
		})

		Convey("When find a non-existent code", func() {

			req, _ := http.NewRequest("GET", "/v1/fairies?fair_code=1030", nil)
			handler.ServeHTTP(res, req)

			Convey("Then return status code 404", func() {
				So(res.Code, ShouldEqual, http.StatusNotFound)
			})
		})
	})

	Convey("Given a name of a fair", t, func() {

		res := httptest.NewRecorder()
		handler := http.HandlerFunc(fairHandler.Search)

		Convey("When search an existing code", func() {

			req, _ := http.NewRequest("GET", "/v1/fairies/search?fair_name=TestFair", nil)
			handler.ServeHTTP(res, req)

			Convey("Then return status code 200", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
			})
		})

		Convey("When search a non-existent code", func() {

			req, _ := http.NewRequest("GET", "/v1/fairies/search?fair_name=Other", nil)
			handler.ServeHTTP(res, req)

			Convey("Then return status code 200", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
			})
		})
	})

}

type FairServiceMock struct{}

func (f *FairServiceMock) Fair(code string) (*app.Fair, error) {
	if code == "1020" {
		return &app.Fair{FairCode: "1020", FairName: "TestFair"}, nil
	}
	return nil, nil
}

func (f *FairServiceMock) Fairs(searchParameter app.FairSearchParameter) ([]app.Fair, error) {
	if searchParameter.FairName == "TestFair" {
		return []app.Fair{{FairCode: "1020", FairName: "TestFair"}}, nil
	}
	return nil, nil
}

func (f *FairServiceMock) CreateFair(fair *app.Fair) error {
	return nil
}

func (f *FairServiceMock) ImportFair(data []string) error {
	return nil
}

func (f *FairServiceMock) UpdateFair(fair *app.Fair) error {
	return nil
}

func (f *FairServiceMock) DeleteFair(code string) error {
	return nil
}
