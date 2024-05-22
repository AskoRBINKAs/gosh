# GOSH - Goddamn Onother Shell (Golang Shell)

It is very useless project, created for fun and this shell cannot work with ubuntu shell (idk why)

## Features
-   You can add your commands in shell via adding new var in `command.go`. Its must implement `Command` struct type
```golang
type Command struct {
	Name          string
	Func          func(args []string) bool
	CommandToCall string
	Aliases       []string
	Description   string
}
```
- Shell configuration file must be in home folder with name `.gosh.json`. Now, you can only edit colors or aliases
- **Aliases** can be added via typing in shell `alias ll=ls -al` (example). To see list of aliases just call `alias list`. To save them - `gosh-save`
- **History** hasnt limit of commands and can be called by `history` (in future i will add saving history to file)

## Installing
- Install dependencies 
- Build project `go build .` inside project folder
- If neccessary copy `gosh` to `/usr/bin/gosh`

Enjoy (or suffer)