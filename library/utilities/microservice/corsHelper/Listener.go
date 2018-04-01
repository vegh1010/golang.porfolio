package corsHelper

import (
	"os"
	"errors"
	"log"
	"time"
	"github.com/gorilla/handlers"
	"github.com/rs/cors"
	"net/http"
	"fmt"
	"github.com/vegh1010/golang.porfolio/library/utilities/routeHelper"
)

//standardize serve function for microservices
type Listener struct {
	Port             string
	AllowOrigins     []string
	AllowedMethods   []string
	AllowedHeaders   []string
	AllowCredentials bool
	Debug            bool
}

func (self *Listener) Init() (error) {
	self.Port = os.Getenv("PORT")
	if self.Port == "" {
		return errors.New("PORT Required")
	}
	self.AllowOrigins = []string{"127.0.0.1"}
	self.AllowedMethods = []string{"POST", "GET", "PUT", "PATCH", "OPTIONS", "DELETE"}
	self.AllowedHeaders = []string{"RequestID", "content-type", "Authorization"}
	self.AllowCredentials = true
	self.Debug = true

	return nil
}

func (self *Listener) AddHeader(header string) {
	self.AllowedHeaders = append(self.AllowedHeaders, header)
}

func (self *Listener) Serve(ServiceName string, Routes routeHelper.Routes, AllowOriginFunc func(origin string) bool) {
	COUNTER := os.Getenv("COUNTER")
	if COUNTER == "" {
		COUNTER = "1"
	}
	ServiceName += "_" + COUNTER
	log.SetOutput(os.Stdout)
	now := time.Now()
	fmt.Println("[SERVER] " + ServiceName + " starting at " + now.Format("2006-01-02 15:04:05") + " on port: " + self.Port)
	loggedRouter := handlers.LoggingHandler(os.Stdout, routeHelper.NewRouter(Routes))
	c := cors.New(cors.Options{
		AllowedOrigins:   self.AllowOrigins,
		AllowedMethods:   self.AllowedMethods,
		AllowedHeaders:   self.AllowedHeaders,
		AllowCredentials: self.AllowCredentials,
		Debug:            self.Debug,
	})
	handler := c.Handler(loggedRouter)
	log.Fatal(http.ListenAndServe(":"+self.Port, handler))
}
