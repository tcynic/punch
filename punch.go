package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/thatisuday/commando"
)

var timelog string
var prettyTime string
var lastEntry string

func init() {
	// define path
	path, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	timelog = path + "/timelog"
}

func define(state string, project string) {

	// Retrieve local Unix Time
	localTime := time.Now().Local()

	// Make the time pretty
	prettyTime = localTime.Format("2006/01/02 15:04:05")

	lastEntry = read()

	if state == "i" {
		if strings.HasPrefix(lastEntry, "i") == true {
			fmt.Println("Error: You are already punched in to", lastEntry[22:])
		} else {
			write(state, project)
			fmt.Println("Punched in to", project)
		}
	} else {
		if strings.HasPrefix(lastEntry, "i") == true {
			if lastEntry[22:] == project {
				write(state, project)
				fmt.Println("Punched out of", project)
			} else {
				fmt.Println("Error: You are already punched in to", lastEntry[22:])
			}
		} else {
			fmt.Println("Error: You must be punched in to a project to punch out.")
		}
	}
}

// Reads the timelog and retrieves the last entry
func read() string {

	// Open timelog
	log, err := os.OpenFile(timelog, os.O_CREATE|os.O_APPEND|os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer log.Close()

	var lines []string
	reader := bufio.NewScanner(log)
	//b4, err := r.Peek(5)
	for reader.Scan() {
		lines = append(lines, reader.Text())
	}
	last := lines[len(lines)-1]
	return last
}

func write(state string, project string) {

	// Create timelog file
	log, err := os.OpenFile(timelog, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer log.Close()
	// Print Date to timelog file
	_, err = fmt.Fprintln(log, state, prettyTime, project)
	if err != nil {
		fmt.Println(err)
		log.Close()
		return
	}

	// Close timelog file if ther is an error
	err = log.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {

	// set cli executable version and description
	commando.
		SetExecutableName("punch").
		SetVersion("v0.1.0").
		SetDescription("This tool will save your time to a ledger friendly timelog file")

	// Register `in` sub-command
	// $ punch in <project>
	commando.Register("in").
		SetDescription("This command will write an \"in\" time entry in the timelog file.").
		SetShortDescription("creates an \"in\" entry").
		AddArgument("project", "The name of the project you are logging time for.", "").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			project := args["project"].Value
			// write time to file
			define("i", project)
		})

	// Register `out` sub-command
	// $ punch out <project>
	commando.Register("out").
		SetDescription("This command will write an \"out\" time entry in the timelog file.").
		SetShortDescription("creates an \"out\" entry").
		AddArgument("project", "The name of the project you are logging time for.", "").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			project := args["project"].Value
			// write time to file
			define("o", project)
		})

	commando.Parse(nil)
}
