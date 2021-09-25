package main

import (
	"flag"
	"log"

	"github.com/goftp/file-driver"
	"github.com/goftp/server"
)

func main() {
	var (
		root = flag.String("root", "D:\\Downloads", "Root directory to serve")
		user = flag.String("user", "root", "Username for login")
		pass = flag.String("pass", "root", "Password for login")
		port = flag.Int("port", 1234, "Port")
		host = flag.String("host", "localhost", "Host")
	)
	flag.Parse()
	if *root == "" {
		log.Fatalf("Please set a root to serve with -root")
	}

	factory := &filedriver.FileDriverFactory{
		RootPath: *root,
		Perm:     server.NewSimplePerm("user", "group"),
	}

	opts := &server.ServerOpts{
		Factory:  factory,
		Port:     *port,
		Hostname: *host,
		Auth:     &server.SimpleAuth{Name: *user, Password: *pass},
	}

	log.Printf("Starting ftp serv on %v:%v", opts.Hostname, opts.Port)
	log.Printf("Username %v, Password %v", *user, *pass)
	serv := server.NewServer(opts)
	err := serv.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting serv:", err)
	}
}