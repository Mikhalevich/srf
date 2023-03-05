package main

import (
	"net/http"

	"github.com/Mikhalevich/srf/logger"
	"github.com/Mikhalevich/srf/mux"
)

type TestRequest struct {
	StringField string `json:"string_field"`
	IntField    int    `json:"int_field"`
}

type TestResponse struct {
	RspStringField string `json:"rsp_string_field"`
	RspIntField    int    `json:"rsp_int_field"`
}

func main() {
	mux.Get("/value", func(h TestRequest) (TestResponse, error) {
		return TestResponse{
			RspStringField: h.StringField,
			RspIntField:    h.IntField,
		}, nil
	})

	mux.Get("/pointer", func(h *TestRequest) (*TestResponse, error) {
		return &TestResponse{
			RspStringField: h.StringField,
			RspIntField:    h.IntField,
		}, nil
	})

	mux.Get("/empty", func(h mux.Empty) (mux.Empty, error) {
		return mux.Empty{}, nil
	})

	log := logger.New()
	if err := http.ListenAndServe(":8080", mux.Mux()); err != nil {
		log.WithError(err).Error("lister and server error")
	}
}
