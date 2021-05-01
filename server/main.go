package main

import (
	"flag"
	gohttp "net/http"
	"os"

	"github.com/ian-kent/go-log/log"
	comcfg "github.com/ynori7/mailhog/config"
	"github.com/ynori7/mailhog/http"
	"github.com/ynori7/mailhog/server/api"
	"github.com/ynori7/mailhog/server/config"
	"github.com/ynori7/mailhog/server/smtp"
	"github.com/ynori7/mailhog/ui/assets"
)

var (
	conf    *config.Config
	comconf *comcfg.Config
	exitCh  chan int
)

func configure() {
	comcfg.RegisterFlags()
	config.RegisterFlags()
	flag.Parse()
	conf = config.Configure()
	comconf = comcfg.Configure()
}

func main() {
	configure()

	if comconf.AuthFile != "" {
		http.AuthFile(comconf.AuthFile)
	}

	exitCh = make(chan int)
	cb := func(r gohttp.Handler) {
		api.CreateAPI(conf, r)
	}
	go http.Listen(conf.APIBindAddr, assets.Asset, exitCh, cb)
	go smtp.Listen(conf, exitCh)

	for {
		select {
		case <-exitCh:
			log.Printf("Received exit signal")
			os.Exit(0)
		}
	}
}
