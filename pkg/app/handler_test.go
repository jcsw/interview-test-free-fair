package app_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"interview-test-free-fair/pkg/app"
)

func TestSpec(t *testing.T) {

	fairHandler := &app.FairHandler{Service: &FairServiceMock{}}

	Convey("Find fair", t, func() {

		res := httptest.NewRecorder()
		handler := http.HandlerFunc(fairHandler.Find)

		Convey("When receive an existing code", func() {

			req, _ := http.NewRequest("GET", "/v1/fairies?fair_code=1020", nil)
			handler.ServeHTTP(res, req)

			Convey("Then return status code 200", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
			})
		})

		Convey("When receive a non-existent code", func() {

			req, _ := http.NewRequest("GET", "/v1/fairies?fair_code=1030", nil)
			handler.ServeHTTP(res, req)

			Convey("Then return status code 404", func() {
				So(res.Code, ShouldEqual, http.StatusNotFound)
			})
		})
	})

	Convey("Search fair", t, func() {

		res := httptest.NewRecorder()
		handler := http.HandlerFunc(fairHandler.Search)

		Convey("When receive an existing code", func() {

			req, _ := http.NewRequest("GET", "/v1/fairies/search?fair_name=TestFair", nil)
			handler.ServeHTTP(res, req)

			Convey("Then return status code 200", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
			})
		})

		Convey("When receive a non-existent code", func() {

			req, _ := http.NewRequest("GET", "/v1/fairies/search?fair_name=Other", nil)
			handler.ServeHTTP(res, req)

			Convey("Then return status code 200", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
			})
		})
	})

	Convey("Create new fair", t, func() {

		res := httptest.NewRecorder()
		handler := http.HandlerFunc(fairHandler.Create)

		Convey("When receive POST", func() {

			body := []byte(`{"fair_code":"1040", "fair_name":"TestNew"}`)
			req, _ := http.NewRequest("POST", "/v1/fairies", bytes.NewBuffer(body))
			handler.ServeHTTP(res, req)

			Convey("Then return status code 201", func() {
				So(res.Code, ShouldEqual, http.StatusCreated)
			})
		})
	})

	Convey("Update fair", t, func() {

		res := httptest.NewRecorder()
		handler := http.HandlerFunc(fairHandler.Update)

		Convey("When receive a PUT an existing fair", func() {

			body := []byte(`{"fair_code":"1020", "fair_name":"TestUpdate"}`)
			req, _ := http.NewRequest("PUT", "/v1/fairies?fair_code=1020", bytes.NewBuffer(body))
			handler.ServeHTTP(res, req)

			Convey("Then return status code 200", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
			})
		})

		Convey("When receive a PUT a non-existing fair", func() {

			body := []byte(`{"fair_code":"1060", "fair_name":"TestUpdate"}`)
			req, _ := http.NewRequest("PUT", "/v1/fairies?fair_code=1060", bytes.NewBuffer(body))
			handler.ServeHTTP(res, req)

			Convey("Then return status code 400", func() {
				So(res.Code, ShouldEqual, http.StatusBadRequest)
			})
		})
	})

	Convey("Delete fair", t, func() {

		res := httptest.NewRecorder()
		handler := http.HandlerFunc(fairHandler.Delete)

		Convey("When receive a DELETE an existing fair", func() {

			req, _ := http.NewRequest("DELETE", "/v1/fairies?fair_code=1020", nil)
			handler.ServeHTTP(res, req)

			Convey("Then return status code 200", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
			})
		})

		Convey("When receive a DELETE a non-existing fair", func() {

			req, _ := http.NewRequest("DELETE", "/v1/fairies?fair_code=1060", nil)
			handler.ServeHTTP(res, req)

			Convey("Then return status code 400", func() {
				So(res.Code, ShouldEqual, http.StatusBadRequest)
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
