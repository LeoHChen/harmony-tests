package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"

	"github.com/SebastianJ/harmony-tf/config"
	"github.com/SebastianJ/harmony-tf/logger"
	"github.com/SebastianJ/harmony-tf/testcases"

	"github.com/urfave/cli"
)

func main() {
	// Force usage of Go's own DNS implementation
	os.Setenv("GODEBUG", "netdns=go")

	app := cli.NewApp()
	app.Name = "Harmony tx tests"
	app.Version = fmt.Sprintf("%s/%s-%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	app.Usage = "Runs a set of Harmony tx tests"
	app.Authors = []cli.Author{{Name: "Sebastian Johnsson", Email: ""}}
	app.Flags = config.CLIFlags()

	app.Action = func(context *cli.Context) error {
		return run(context)
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println()
		logger.ErrorLog(err.Error(), true)
		os.Exit(1)
	}
}

func run(context *cli.Context) error {
	pprofPort := context.GlobalInt(config.CLIPprofPort.Name)
	if pprofPort > 0 {
		go func() {
			fmt.Println(http.ListenAndServe(fmt.Sprintf("localhost:%d", pprofPort), nil))
		}()
	}

	basePath, err := filepath.Abs(context.GlobalString(config.CLIPath.Name))
	if err != nil {
		return err
	}

	if err := config.Configure(basePath, context); err != nil {
		return err
	}

	if err := testcases.Execute(); err != nil {
		return err
	}

	// When we run using pprof we don't want to immediately exit,
	// We want to manually exit so that we have time to capture heap dumps etc. after the test suite has executed
	if pprofPort > 0 {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs
		os.Exit(1)
	}

	return nil
}
