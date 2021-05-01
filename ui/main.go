package main

import (
	"flag"
	gohttp "net/http"
	"os"

	"github.com/gorilla/pat"
	"github.com/ian-kent/go-log/log"
	comcfg "github.com/ynori7/mailhog/config"
	"github.com/ynori7/mailhog/http"
	"github.com/ynori7/mailhog/ui/assets"
	"github.com/ynori7/mailhog/ui/config"
	"github.com/ynori7/mailhog/ui/web"
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
	// FIXME hacky
	web.APIHost = conf.APIHost
}

func main() {
	configure()

	// FIXME need to make API URL configurable

	if comconf.AuthFile != "" {
		http.AuthFile(comconf.AuthFile)
	}

	exitCh = make(chan int)
	cb := func(r gohttp.Handler) {
		web.CreateWeb(conf, r.(*pat.Router), assets.Asset)
	}
	go http.Listen(conf.UIBindAddr, assets.Asset, exitCh, cb)

	for {
		select {
		case <-exitCh:
			log.Printf("Received exit signal")
			os.Exit(0)
		}
	}
}
