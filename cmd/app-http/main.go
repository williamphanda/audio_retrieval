package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	audiofilehandler "audioretrieval/internal/handler/http/audio"
	filerepo "audioretrieval/internal/repo/file"
	phraserepo "audioretrieval/internal/repo/phrase"
	userrepo "audioretrieval/internal/repo/user"
	userphraserepo "audioretrieval/internal/repo/userphrase"
	audiousecase "audioretrieval/internal/usecase/audio"

	_ "github.com/lib/pq"
)

func main() {
	/*
		RESOURCE
	*/

	// TODO: move credential to a secure file like Google Secret Manager or equivalent.
	databaseDSN := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", "postgres", "5432", "superuser", "supersecretpassword", "audio-retrieval")

	db, err := sql.Open("postgres", databaseDSN)
	if err != nil {
		log.Println("error opening postgres: ", err)
		return
	}
	log.Println("open postgres successful")

	err = db.Ping()
	if err != nil {
		log.Println("error ping postgres: ", err)
		return
	}
	log.Println("success ping postgres")

	/*
		REPOSITORY
	*/
	fileRepo := filerepo.New()
	log.Println("fileRepo create successful")

	userRepo := userrepo.New(db)
	log.Println("userRepo create successful")

	phraseRepo := phraserepo.New(db)
	log.Println("phraseRepo create successful")

	userPhraseRepo := userphraserepo.New(db)
	log.Println("userPhraseRepo create successful")

	/*
		USE CASE
	*/
	audioUseCase := audiousecase.New(fileRepo, userRepo, phraseRepo, userPhraseRepo)
	log.Println("audioUseCase create successful")

	/*
		HANDLER
	*/
	httpHandler := audiofilehandler.New(audioUseCase)
	log.Println("httpHandler create successful")

	/*
		SERVER
	*/
	router := newRoutes(httpHandler)
	log.Println("new routes set")

	// TODO move configurations to a different file
	// could be a cold config (.yaml, .ini) or a hot config (consul kv, etc.)
	srv := http.Server{
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 10000,
		Addr:           ":8000",
	}

	log.Println("application running on port 8000")
	err = srv.ListenAndServe()
	if err != nil {
		log.Println("Server failed: %v", err)
		return
	}
}
