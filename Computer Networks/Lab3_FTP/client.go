package main

import (
	"flag"
	"fmt"
	"github.com/jlaffaye/ftp"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var (
		mkdirFlag = flag.String("mkdir", "", "Create directory")
		rmdirFlag = flag.String("rmdir", "", "Remove directory")
		getFlag   = flag.String("get", "", "Download file")
		putFlag   = flag.String("put", "", "Upload file")
	)

	flag.Parse()

	c, err := ftp.Dial("localhost:1234")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Quit()

	err = c.Login("root", "root")
	if err != nil {
		log.Fatal(err)
	}

	if *mkdirFlag != "" {
		err := c.MakeDir(*mkdirFlag)
		if err != nil {
			log.Println("Failed to create dir", *mkdirFlag)
			return
		}
		log.Println("Successfully created remote dir", *mkdirFlag)
	} else if *rmdirFlag != "" {
		err := c.RemoveDirRecur(*rmdirFlag)
		if err != nil {
			log.Println("Failed to remove dir", *rmdirFlag)
			return
		}
		log.Println("Successfully removed remote dir", *rmdirFlag)
	} else if *getFlag != "" {
		r, err := c.Retr(*getFlag)
		if err != nil {
			log.Println("Failed to download file", *getFlag)
			return
		}
		buf, err := ioutil.ReadAll(r)
		if err != nil {
			log.Println("Failed to read file's contents")
			return
		}
		fmt.Printf("Reading file %s successful. File length: %d bytes", *getFlag, len(buf))
	} else if *putFlag != "" {
		f, err := os.Open(*putFlag)
		if err != nil{
			fmt.Println("Failed to read local file")
			return
		}
		err = c.Stor(*putFlag, f)
		if err != nil {
			log.Println("Failed to upload file", *putFlag)
			return
		}
		fmt.Println("Successfully uploaded file", *putFlag)
	} else {
		flag.Usage()
	}
}
