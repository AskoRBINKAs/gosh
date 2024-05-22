package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

type Command struct {
	Name          string
	Func          func(args []string) bool
	CommandToCall string
	Aliases       []string
	Description   string
}

var Exit Command = Command{
	Name:          "exit",
	CommandToCall: "exit",
	Description:   "exit from shell",
	Aliases:       []string{"quit"},
	Func: func(args []string) bool {
		fmt.Println("Goodbye!")
		os.Exit(0)
		return true
	},
}

var ChangeDirectory Command = Command{
	Name:          "cd",
	CommandToCall: "cd",
	Description:   "change directory",
	Func: func(args []string) bool {
		dest := args[0]
		if dest == "~" {
			u, _ := user.Current()
			dest = u.HomeDir
		}
		err := os.Chdir(dest)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true
	},
}

var Help Command = Command{
	Name:          "help",
	CommandToCall: "help",
	Description:   "just print help message",
	Func: func(args []string) bool {
		for _, cmd := range commands {
			fmt.Printf("%s - %s\n", cmd.CommandToCall, cmd.Description)
		}
		return true
	},
}

var SetAlias Command = Command{
	Name:          "alias",
	CommandToCall: "alias",
	Description:   "sets alias between first and second argument after cmd name",
	Func: func(args []string) bool {
		new_args := strings.Split(strings.Join(args, " "), "=")
		if len(args) == 1 && args[0] == "list" {
			fmt.Println("List of aliases:")
			for key, val := range userAliases {
				fmt.Println(key, "=", val)
			}
			return true
		}
		if len(new_args) != 2 {
			return false
		}
		userAliases[new_args[0]] = new_args[1]
		return true
	},
}

var SaveCfg Command = Command{
	Name:          "save cfg",
	CommandToCall: "gosh-save",
	Description:   "saves config to home directory",
	Func: func(args []string) bool {
		config.UserAliases = userAliases
		SaveConfig(config)
		return true
	},
}

var History Command = Command{
	Name:          "history",
	CommandToCall: "history",
	Description:   "print last 500 commands",
	Func: func(args []string) bool {
		PrintHistory()
		return true
	},
}

var commands []Command
var userAliases map[string]string

func exportCommands() []Command {
	commands = []Command{Exit,
		ChangeDirectory,
		Help,
		SetAlias,
		SaveCfg,
		History,
	}
	return commands
}
