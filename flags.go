package main

import "github.com/urfave/cli"

var (
	DEBUG_FLAG          = "debug"
	CONSUMERKEY_FLAG    = "consumerkey"
	COMSUMERSECRET_FLAG = "consumersecret"
	ACCESSTOKEN_FLAG    = "accesstoken"
	ACCESSSECRET_FALG   = "accesssecret"
	PORT_FLAG           = "port"
)

var (
	Debug             = false
	ConsumerKey       = ""
	ConsumerSecret    = ""
	AccessToken       = ""
	AccessTokenSecret = ""
	Port              = ""
)

var cmdFlags = []cli.Flag{
	cli.BoolFlag{
		Name:        DEBUG_FLAG + ", d",
		Usage:       "Debug mode.",
		Destination: &Debug,
	},
	cli.StringFlag{
		Name:        CONSUMERKEY_FLAG + ", ck",
		Usage:       "Twitter Consumer Key",
		Destination: &ConsumerKey,
		EnvVar:      "CONSUMERKEY",
	},
	cli.StringFlag{
		Name:        COMSUMERSECRET_FLAG + ", cs",
		Usage:       "Twitter Consumer Secret",
		Destination: &ConsumerSecret,
		EnvVar:      "COMSUMERSECRET",
	},
	cli.StringFlag{
		Name:        ACCESSTOKEN_FLAG + ", at",
		Usage:       "Twitter Access Token",
		Destination: &AccessToken,
		EnvVar:      "ACCESSTOKEN",
	},
	cli.StringFlag{
		Name:        ACCESSSECRET_FALG + ", as",
		Usage:       "Twitter Access Token Secret",
		Destination: &AccessTokenSecret,
		EnvVar:      "ACCESSSECRET",
	},
	cli.StringFlag{
		Name:        PORT_FLAG + ", p",
		Usage:       "API Server port",
		Destination: &Port,
		EnvVar:      "PORT",
	},
}
