package main

import (
	audiofilehandler "audioretrieval/internal/handler/http/audio"

	"github.com/go-chi/chi"
)

func setupAppRoutes(router chi.Router, audioFileHandler *audiofilehandler.Handler) {
	router.Route("/audio/user", func(r chi.Router) {
		r.Post("/{user_id}/phrase/{phrase_id}", audioFileHandler.StoreAudio)
		r.Get("/{user_id}/phrase/{phrase_id}/{audio_format}", audioFileHandler.GetAudio)
	})
}

func newRoutes(audioFileHandler *audiofilehandler.Handler) *chi.Mux {
	router := chi.NewRouter()

	router.Group(func(router chi.Router) {
		// TODO: add middleware for CORS, authentication, etc.
		// router.Use(middleware)

		setupAppRoutes(router, audioFileHandler)
	})

	return router
}
