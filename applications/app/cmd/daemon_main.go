package main

import (
	"github.com/sevlyar/go-daemon"
	"go-xb-song/applications/app/server"
	"log"
)

func main()  {

	cntxt := &daemon.Context{
		PidFileName: "sample.pid",
		PidFilePerm: 0644,
		LogFileName: "sample.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{"[go-daemon sample]"},
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()

	log.Print("- - - - - - - - - - - - - - -")
	log.Print("daemon started")

	srv := server.New()

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
