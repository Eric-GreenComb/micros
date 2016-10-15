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
	log.Fatalf("Usage: prepare_tpl_web [-t filename] \n")
}

func main() {
	log.Println("Starting ...")
	defer log.Println("Shutdown complete!")

	var tplName = flag.String("t", "hello", "the tpl file name")

	flag.Usage = usage
	flag.Parse()

	_f, _err := ioutil.ReadFile(*tplName + ".tpl")
	if _err != nil {
		fmt.Println(_err.Error())
		os.Exit(2)
	}

	fmt.Println(string(_f))

	_key := global.ETCD_KEY_TPL_WEB + *tplName
	_value := string(_f)
	etcd.Set(_key, _value)

	os.Exit(1)
}
