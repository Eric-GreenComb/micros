package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/banerwai/global"
	"github.com/banerwai/gommon/etcd"
)

func usage() {
	log.Fatalf("Usage: prepare_json_category [-j filename] \n")
}

func main() {
	log.Println("Starting ...")
	defer log.Println("Shutdown complete!")

	var json_name = flag.String("j", "category", "the json file name")

	flag.Usage = usage
	flag.Parse()

	_f, _err := ioutil.ReadFile(*json_name + ".json")
	if _err != nil {
		fmt.Println(_err.Error())
		os.Exit(2)
	}

	fmt.Println(string(_f))

	_key := global.ETCD_KEY_JSON_CATEGORY + *json_name
	_value := string(_f)
	etcd.Set(_key, _value)

	os.Exit(1)
}
