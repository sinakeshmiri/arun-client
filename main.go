package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/sinakeshmiri/arun-client/packages/arunclient"
)

func main() {
	url := flag.String("url", "", "function name")
	fName := flag.String("name", "", "function name")
	fSrc := flag.String("source", "", "source code")
	act := flag.String("act", "", "create,get or run the function")
	flag.Parse()
	if *url == "" {
		log.Fatal("spcify the url")
	}
	if *fName == "" {
		log.Fatal("spcify the name")
	}
	if *act == "add" {
		if *fSrc == "" {
			log.Fatal("to add a function you must spcify a source file")
		}

		content, err := ioutil.ReadFile(*fSrc)
		if err != nil {
			log.Fatal(err)
		}
		src := string(content)
		err = arunclient.Add(src, *fName, *url)
		if err != nil {
			log.Fatal(err)
		}

	}
	if *act == "run" {
		res, err := arunclient.Run(*fName, *url)
		if err != nil {
			log.Fatal(err)
		}
		var t []byte
		res.Body.Read(t)
		
		fmt.Println(res.StatusCode,string(t))
	}
	if *act == "get" {
		res, err := arunclient.Get(*fName, *url)
		if err != nil {
			log.Fatal(err)
		}
		var t []byte
		res.Body.Read(t)
		
		fmt.Println(res.StatusCode,string(t))
	}
}
