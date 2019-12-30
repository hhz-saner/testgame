package main

import (
	"fmt"
	"gameLog/config"
	"gameLog/routers"
	"net/http"
	"time"
)

func main(){
	router := routers.InitRouter()
	fmt.Println("http://"+config.Cfg.Http.Addr)

	s := &http.Server{
		Addr:           config.Cfg.Http.Addr,
		Handler:        router,
		ReadTimeout:    time.Duration(config.Cfg.Http.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.Cfg.Http.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err !=nil {
		panic(err.Error())
	}
}
