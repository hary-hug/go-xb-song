package main

import (
	"context"
	"go-xb-song/applications/app/server"
	"log"
	"os"
	"os/signal"
	"time"
)

func main()  {


	srv := server.New()

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<- quit

	log.Println("Shutdown Server ...")


	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {

		log.Fatal("Server Shutdown:", err)
	}

}

