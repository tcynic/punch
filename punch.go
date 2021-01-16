package main

import (
	"fmt"
	"os"
	"time"

	"github.com/thatisuday/commando"
)

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
			write("i", project)
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
			write("o", project)
		})

	commando.Parse(nil)

}

func write(state string, project string) {

	// Retrieve local Unix Time
	localTime := time.Now().Local()

	// Make the time pretty
	prettyTime := localTime.Format("2006/01/02 15:04:05")

	// Create timelog file
	log, err := os.OpenFile("timelog", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
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

	if state == "o" {
		fmt.Println("Punched out")
	} else {
		fmt.Println("Punched in")
	}

}
