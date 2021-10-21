package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"github.com/traefik/yaegi/stdlib/syscall"
	"github.com/traefik/yaegi/stdlib/unrestricted"
	"github.com/traefik/yaegi/stdlib/unsafe"
	"github.com/urfave/cli/v2"

	"samhofi.us/x/keybase/v2/types/chat1"

	bot "github.com/kf5grd/keybasebot"
	"github.com/kf5grd/keybasebot/pkg/util"
)

// Current version
var version string

// Exit code on failure
const exitFail = 1

var Symbols = make(map[string]map[string]reflect.Value)

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
	i := interp.New(interp.Options{GoPath: "./_go"})
	if err := i.Use(stdlib.Symbols); err != nil {
		panic(err)
	}
	if err := i.Use(syscall.Symbols); err != nil {
		panic(err)
	}
	if err := i.Use(unsafe.Symbols); err != nil {
		panic(err)
	}
	if err := i.Use(unrestricted.Symbols); err != nil {
		panic(err)
	}
	if err := i.Use(interp.Symbols); err != nil {
		panic(err)
	}
	if err := i.Use(Symbols); err != nil {
		panic(err)
	}
	i.ImportUsed()

	// import some initial packages and create the bot object within the context of the interpreter
	_, err := i.Eval(fmt.Sprintf(`
          import(
            "samhofi.us/x/keybase/v2"
            "samhofi.us/x/keybase/v2/types/chat1"
            "github.com/kf5grd/keybasebot/pkg/util"
            bot "github.com/kf5grd/keybasebot"
          )
          var b = bot.New("", keybase.SetHomePath(%q))
        `, c.String("home")))
	if err != nil {
		panic(err)
	}

	v, err := i.Eval("b")
	if err != nil {
		panic(err)
	}

	b, ok := v.Interface().(*bot.Bot)
	if !ok {
		fmt.Println("unable to create bot")
		os.Exit(1)
	}

	b.LogWriter = c.App.Writer
	b.Debug = c.Bool("debug")
	b.JSON = c.Bool("json")

	b.Meta["interp"] = i

	botOwner := c.String("bot-owner")
	b.Meta["botOwner"] = botOwner
	b.Meta["whitelist"] = []string{botOwner}
	b.Meta["UsersFromMeta"] = UsersFromMeta

	// Set initialize the bot's commands
	b.Commands = []bot.BotCommand{
		{
			Name: "eval",
			Ad: &chat1.UserBotCommandInput{
				Name:        "eval",
				Usage:       "<expr>",
				Description: "Evaluate a Go expression",
			},
			Run: bot.Adapt(cmdEval,
				UsersFromMeta("whitelist"),
				bot.MessageType("text"),
				bot.CommandPrefix("!eval"),
			),
		},
	}

	// Run the bot
	b.Run()

	return nil
}

func UsersFromMeta(key string) bot.Adapter {
	return func(botAction bot.BotAction) bot.BotAction {
		return func(m chat1.MsgSummary, b *bot.Bot) (bool, error) {
			users, ok := b.Meta[key].([]string)
			if !ok {
				return false, nil
			}
			b.Logger.Debug("Verifying received message was sent by one of '%s'", strings.Join(users, ","))
			if !util.StringInSlice(m.Sender.Username, users) {
				b.Logger.Debug("Received message was sent by '%s', exiting command", m.Sender.Username)
				return false, nil
			}
			b.Logger.Debug("Received message was sent by '%s', continuing", m.Sender.Username)
			return botAction(m, b)
		}
	}
}
