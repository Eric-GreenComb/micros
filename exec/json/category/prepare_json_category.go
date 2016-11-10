package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/banerwai/global/constant"
	"github.com/banerwai/gommon/etcd"
)

func usage() {
	log.Fatalf("Usage: prepare_json_category [-j filename] \n")
}

func main() {
	log.Println("Starting ...")
	defer log.Println("Shutdown complete!")

	var jsonName = flag.String("j", "category", "the json file name")

	flag.Usage = usage
	flag.Parse()

	_f, _err := ioutil.ReadFile(*jsonName + ".json")
	if _err != nil {
		fmt.Println(_err.Error())
		os.Exit(2)
	}

	fmt.Println(string(_f))

	_key := constant.EtcdKeyJSONCategory + *jsonName
	_value := string(_f)
	etcd.Set(_key, _value)

	os.Exit(1)
}
