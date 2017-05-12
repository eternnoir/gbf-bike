package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/eternnoir/gbf-bike/bike"
	"github.com/eternnoir/gbf-bike/server"
	"github.com/urfave/cli"
	"os"
)

// InitApp function is use to create new cli.App.
func InitApp() *cli.App {
	app := cli.NewApp()
	app.Name = "GBF Bike"
	app.Usage = "GBF Battle grab server."
	app.Flags = cmdFlags
	app.Action = func(c *cli.Context) error {
		return Start(c)
	}
	return app
}

func Start(c *cli.Context) error {
	if Debug {
		log.Warn("Use debug mode")
		log.SetLevel(log.DebugLevel)
	}
	gb, err := bike.NewGbfBike(ConsumerKey, ConsumerSecret, AccessToken, AccessTokenSecret)
	if err != nil {
		return err
	}

	apiServer := server.NewApi(Port)
	gb.AddBattleReceiver(apiServer)
	go apiServer.Start()
	log.Fatal(gb.Start())
	return nil
}

func main() {
	app := InitApp()
	if err := app.Run(os.Args); err != nil {
		log.Errorf("Start GBF Bike Fail %s", err)
	}
}
