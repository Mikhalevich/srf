package mux

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/Mikhalevich/srf/logger"
)

var (
	mux *chi.Mux
)

func init() {
	mux = chi.NewRouter()
}

type Empty struct {
}

type Handler[REQ any, RSP any] func(REQ) (RSP, error)

func makeHandler[REQ any, RSP any](h Handler[REQ, RSP]) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logger.FromContext(r.Context())

		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.WithError(err).Error("read body error")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.WithField("handler_request", string(b)).Info("request received")

		var t REQ
		if len(b) > 0 {
			if err := json.Unmarshal(b, &t); err != nil {
				log.WithError(err).Error("json decode error")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		rsp, err := h(t)
		if err != nil {
			log.WithError(err).Error("execute handler error")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		rspBytes, err := json.Marshal(rsp)
		if err != nil {
			log.WithError(err).Error("json marshal error")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.WithField("handler_response", string(rspBytes)).Info("send response")

		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		if _, err := w.Write(rspBytes); err != nil {
			log.WithError(err).Error("write bytes error")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func Mux() *chi.Mux {
	return mux
}

func Get[REQ any, RSP any](pattern string, h Handler[REQ, RSP]) {
	mux.Get(pattern, makeHandler(h))
}
