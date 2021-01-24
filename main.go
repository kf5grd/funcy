package main

import (
	"os"
	"reflect"

	"github.com/cosmos72/gomacro/fast"
	"github.com/cosmos72/gomacro/imports"
	"github.com/urfave/cli/v2"

	"samhofi.us/x/keybase/v2"
	"samhofi.us/x/keybase/v2/types/chat1"

	bot "github.com/kf5grd/keybasebot"
)

// Current version
var version string

// Exit code on failure
const exitFail = 1

func init() {
	imports.Packages["github.com/kf5grd/funcy"] = imports.Package{
		Binds:    map[string]reflect.Value{},
		Types:    map[string]reflect.Type{},
		Proxies:  map[string]reflect.Type{},
		Untypeds: map[string]string{},
		Wrappers: map[string][]string{},
	}
}

func main() {
	app := cli.App{
		Name:                 "funcy",
		Usage:                "Go interpreter for Keybase",
		Version:              version,
		Writer:               os.Stdout,
		EnableBashCompletion: true,
		Action:               run,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "home",
				Aliases: []string{"H"},
				Usage:   "Keybase Home Folder",
				EnvVars: []string{"FUNCBOT_HOME"},
			},
			&cli.StringFlag{
				Name:    "bot-owner",
				Usage:   "Username of the bot owner",
				EnvVars: []string{"FUNCBOT_OWNER"},
			},
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "Enable extra log output",
				EnvVars: []string{"FUNCBOT_DEBUG"},
			},
			&cli.BoolFlag{
				Name:    "json",
				Aliases: []string{"j"},
				Usage:   "Output log in JSON format",
				EnvVars: []string{"FUNCBOT_JSON"},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(exitFail)
	}
}

func run(c *cli.Context) error {
	// Create the Go interpreter
	interp := fast.New()

	// This is needed to be able to share `b` between compiled and interpreted code
	var b *bot.Bot
	imports.Packages["github.com/kf5grd/funcy"].Binds["b"] = reflect.ValueOf(b).Elem()
	interp.ImportPackage("funcy", "github.com/kf5grd/funcy")
	interp.ChangePackage("funcy", "github.com/kf5grd/funcy")

	// Create the bot object and set some basic options
	b = bot.New("", keybase.SetHomePath(c.String("home")))
	b.LogWriter = c.App.Writer
	b.Debug = c.Bool("debug")
	b.JSON = c.Bool("json")

	b.Meta["interp"] = interp

	botOwner := c.String("bot-owner")

	// Set initialize the bot's commands
	b.Commands = []bot.BotCommand{
		{
			Name: "GetLinks",
			Ad: &chat1.UserBotCommandInput{
				Name:        "eval",
				Usage:       "<expr>",
				Description: "Evaluate a Go expression",
			},
			Run: bot.Adapt(cmdEval,
				bot.MessageType("text"),
				bot.CommandPrefix("!eval"),
				bot.FromUser(botOwner),
			),
		},
	}

	// Run the bot
	b.Run()

	return nil
}
