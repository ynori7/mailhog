package main

import (
	"flag"
	"fmt"
	gohttp "net/http"
	_ "net/http/pprof"
	"os"

	"github.com/gorilla/pat"
	"github.com/ian-kent/envconf"
	"github.com/ian-kent/go-log/log"
	cfgcom "github.com/ynori7/mailhog/config"
	"github.com/ynori7/mailhog/http"
	"github.com/ynori7/mailhog/sendmail/cmd"
	"github.com/ynori7/mailhog/server/api"
	cfgapi "github.com/ynori7/mailhog/server/config"
	"github.com/ynori7/mailhog/server/smtp"
	"github.com/ynori7/mailhog/ui/assets"
	cfgui "github.com/ynori7/mailhog/ui/config"
	"github.com/ynori7/mailhog/ui/web"
	"golang.org/x/crypto/bcrypt"
)

var (
	apiconf          *cfgapi.Config
	uiconf           *cfgui.Config
	comconf          *cfgcom.Config
	profilingEnabled bool

	exitCh  chan int
	version string
)

func configure() {
	cfgcom.RegisterFlags()
	cfgapi.RegisterFlags()
	cfgui.RegisterFlags()
	flag.BoolVar(&profilingEnabled, "profiling-enabled", envconf.FromEnvP("PROFILING_ENABLED", false).(bool), "When set, profiling via pprof is enabled")
	flag.Parse()
	apiconf = cfgapi.Configure()
	uiconf = cfgui.Configure()
	comconf = cfgcom.Configure()

	apiconf.WebPath = comconf.WebPath
	uiconf.WebPath = comconf.WebPath
}

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "-version" || os.Args[1] == "--version") {
		fmt.Println("MailHog version: " + version)
		os.Exit(0)
	}

	if len(os.Args) > 1 && os.Args[1] == "sendmail" {
		args := os.Args
		os.Args = []string{args[0]}
		if len(args) > 2 {
			os.Args = append(os.Args, args[2:]...)
		}
		cmd.Go()
		return
	}

	if len(os.Args) > 1 && os.Args[1] == "bcrypt" {
		var pw string
		if len(os.Args) > 2 {
			pw = os.Args[2]
		} else {
			// TODO: read from stdin
		}
		b, err := bcrypt.GenerateFromPassword([]byte(pw), 4)
		if err != nil {
			log.Fatalf("error bcrypting password: %s", err)
			os.Exit(1)
		}
		fmt.Println(string(b))
		os.Exit(0)
	}

	configure()

	if comconf.AuthFile != "" {
		http.AuthFile(comconf.AuthFile)
	}

	exitCh = make(chan int)
	if uiconf.UIBindAddr == apiconf.APIBindAddr {
		cb := func(r gohttp.Handler) {
			web.CreateWeb(uiconf, r.(*pat.Router), assets.Asset)
			api.CreateAPI(apiconf, r.(*pat.Router))
		}
		go http.Listen(uiconf.UIBindAddr, assets.Asset, exitCh, cb)
	} else {
		cb1 := func(r gohttp.Handler) {
			api.CreateAPI(apiconf, r.(*pat.Router))
		}
		cb2 := func(r gohttp.Handler) {
			web.CreateWeb(uiconf, r.(*pat.Router), assets.Asset)
		}
		go http.Listen(apiconf.APIBindAddr, assets.Asset, exitCh, cb1)
		go http.Listen(uiconf.UIBindAddr, assets.Asset, exitCh, cb2)
	}

	if profilingEnabled {
		go func() {
			log.Println("Profiler on http://localhost:8080/debug/pprof")
			gohttp.ListenAndServe("localhost:8080", nil)
		}()
	}

	go smtp.Listen(apiconf, exitCh)

	for {
		select {
		case <-exitCh:
			log.Printf("Received exit signal")
			os.Exit(0)
		}
	}
}
