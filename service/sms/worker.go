package main

import (
	"log"
	"net"
	"os"

	"github.com/mikespook/gearman-go/worker"
	"github.com/mikespook/golib/signal"

	"github.com/banerwai/gather/common/gearman"
	smsservice "github.com/banerwai/micros/command/sms/service"
)

func SendSms(job worker.Job) ([]byte, error) {
	log.Printf("SendSms: Data=[%s]\n", job.Data())
	var _smsservice smsservice.SmsService
	_err := _smsservice.LPOP4Redis(string(job.Data()))
	return []byte("OK"), _err
}

func main() {
	log.Println("Starting ...")
	defer log.Println("Shutdown complete!")
	w := worker.New(worker.Unlimited)
	defer w.Close()
	w.ErrorHandler = func(e error) {
		log.Println(e)
		if opErr, ok := e.(*net.OpError); ok {
			if !opErr.Temporary() {
				proc, err := os.FindProcess(os.Getpid())
				if err != nil {
					log.Println(err)
				}
				if err := proc.Signal(os.Interrupt); err != nil {
					log.Println(err)
				}
			}
		}
	}
	w.JobHandler = func(job worker.Job) error {
		log.Printf("Data=%s\n", job.Data())
		return nil
	}
	w.AddServer("tcp4", gearman.GearmanAddr)
	w.AddFunc("SendSms", SendSms, worker.Unlimited)
	if err := w.Ready(); err != nil {
		log.Fatal(err)
		return
	}
	go w.Work()
	signal.Bind(os.Interrupt, func() uint { return signal.BreakExit })
	signal.Wait()
}
