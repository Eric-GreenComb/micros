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
	log.Fatalf("Usage: prepare_tpl_contact [-t filename] \n")
}

func main() {
	log.Println("Starting ...")
	defer log.Println("Shutdown complete!")

	var tpl_name = flag.String("t", "default", "the tpl file name")

	flag.Usage = usage
	flag.Parse()

	_f, _err := ioutil.ReadFile(*tpl_name + ".tpl")
	if _err != nil {
		fmt.Println(_err.Error())
		os.Exit(2)
	}

	_key := global.ETCD_KEY_TPL_CONTACT + *tpl_name
	_value := string(_f)
	etcd.Set(_key, _value)

	os.Exit(1)
}
