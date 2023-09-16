// GOOS=linux GOARCH=arm GOARM=7 go build -o  main_arm ./main.go

package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"

	prompt "github.com/c-bata/go-prompt"
)

type Cmd struct {
	RawCmd string
	Desc   string
}

var CmdMap = map[string]Cmd{
	"rwroot": {
		RawCmd: "mount -o remount,rw /",
		Desc:   "Remount the root directory in read/write mode",
	},

	"roroot": {
		RawCmd: "mount -o remount,ro /",
		Desc:   "Remount the root directory in read only mode",
	},

	"mountdbgfs": {
		RawCmd: "mount -t debugfs none /sys/kernel/debug",
        Desc:   "Mount the debugfs filesystem",
	}
}

func completer(d prompt.Document) []prompt.Suggest {
	s := make([]prompt.Suggest, 0)
	for k, v := range CmdMap {
		s = append(s, prompt.Suggest{
			Text:        k,
			Description: v.Desc,
		})
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	fmt.Println("Please select table.")

	t := prompt.Input("> ", completer, []prompt.Option{
		prompt.OptionDescriptionBGColor(prompt.DarkGray),
		prompt.OptionSuggestionBGColor(prompt.DarkGray),

		prompt.OptionDescriptionTextColor(prompt.White),
		prompt.OptionSuggestionTextColor(prompt.White),

		prompt.OptionSelectedDescriptionBGColor(prompt.Yellow),
		prompt.OptionSelectedSuggestionBGColor(prompt.Yellow),

		prompt.OptionSelectedDescriptionTextColor(prompt.DarkGray),
		prompt.OptionSelectedSuggestionTextColor(prompt.DarkGray),
	}...)

	fmt.Println("You selected: \"" + t + "\"")

	if runtime.GOOS == "linux" {
		cmd := CmdMap[t]

		fmt.Println("Run: " + cmd.RawCmd)
		split_cmd := strings.Split(cmd.RawCmd, " ")
		excmd := exec.Command(split_cmd[0], split_cmd[1:]...)

		err := excmd.Run()

		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Not supported on this OS: " + runtime.GOOS + ".")
	}

}
