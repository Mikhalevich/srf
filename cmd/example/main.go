package main

import (
	"fmt"
	"net/http"

	"github.com/Mikhalevich/srf/mux"
)

type TestRequest struct {
	StringField string `json:"string_field"`
	IntField    int    `json:"int_field"`
}

type TestResponse struct {
	StringField string `json:"rsp_string_field"`
	IntField    int    `json:"rsp_int_field"`
}

func main() {
	mux.Get("/", func(h TestRequest) (TestResponse, error) {
		return TestResponse{
			StringField: h.StringField,
			IntField:    h.IntField,
		}, nil
	})

	mux.Get("/empty", func(h mux.Empty) (mux.Empty, error) {
		return mux.Empty{}, nil
	})

	if err := http.ListenAndServe(":8080", mux.Mux()); err != nil {
		fmt.Printf("lister and server error: %v\n", err)
	}
}
