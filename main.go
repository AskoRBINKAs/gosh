package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"os/user"
	"slices"
	"strings"
	"syscall"
)

var buildinCommands []Command
var config Config

func welcomeMessage() {
	cwd, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}
	username, err := user.Current()
	if err != nil {
		os.Exit(1)
	}
	hostname, err := os.Hostname()
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("%s@%s:%s> ",
		string(config.BuiltInColors[config.UserColors.UsernameColor])+username.Name,
		string(config.BuiltInColors[config.UserColors.HostnameColor])+hostname,
		string(config.BuiltInColors[config.UserColors.PwdColor])+cwd+string(config.BuiltInColors[config.UserColors.InputColor]))
}

func builtinCommandsExecute(name string, args []string) bool {
	for _, cmd := range buildinCommands {
		if cmd.CommandToCall == name || slices.Contains(cmd.Aliases, name) {
			cmd.Func(args)
			return true
		}
	}
	return false
}

func externalCommandExecute(name string, args []string) bool {
	process := exec.Command(name, args...)
	process.Stdout = os.Stdout
	process.Stdin = os.Stdin
	process.Stderr = os.Stderr
	process.Run()
	switch process.ProcessState.ExitCode() {
	case -1:
		return false
	default:
		return true
	}
}

func initialize() {
	buildinCommands = exportCommands()
	c, err := LoadConfig()
	if err != nil {
		fmt.Println(err)
	}
	config = c
	userAliases = config.UserAliases
}

func commandExecute(name string, args []string) {
	AddToHistory(name + " " + strings.Join(args, " "))
	buildin := builtinCommandsExecute(name, args)
	if buildin {
		return
	}
	if cmd_name, ok := userAliases[name]; ok {
		fmt.Println(cmd_name)
		name = strings.Split(cmd_name, " ")[0]
		args = strings.Split(cmd_name, " ")[1:]
	}
	external := externalCommandExecute(name, args)
	if external {
		return
	}
	fmt.Println("gosh: command not found")
}

func main() {
	var ver bool
	flag.BoolVar(&ver, "version", false, "")
	flag.Parse()
	if ver {
		fmt.Println("0.1-dev")
		os.Exit(0)
	}
	signal.Ignore(os.Interrupt)
	signal.Ignore(syscall.SIGTERM)
	initialize()
	for {
		welcomeMessage()
		reader := bufio.NewReader(os.Stdin)
		command, _ := reader.ReadString('\n')
		command = strings.Trim(command, " \n\r\v\t")
		if len(command) == 0 {
			continue
		}
		name := strings.Split(command, " ")[0]
		args := strings.Split(command, " ")[1:]
		commandExecute(name, args)
	}
}
