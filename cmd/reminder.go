package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/monirz/reminder/remind"
	"github.com/spf13/cobra"
)

func init() {
	// RootCmd.AddCommand(RootCmd)
	RootCmd.AddCommand(AddCmd)

}

var RootCmd = &cobra.Command{
	Use:   "reminder",
	Short: "Reminder is a event remider application",
	Long:  "Reminder keeps you updated with scheduled event",
}

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add, adds an event ",
	Run:   addRun,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func addRun(cmd *cobra.Command, args []string) {
	fmt.Println("running...", args)

	ch := make(chan bool)

	if len(args) < 2 {
		fmt.Println("time argument of the event is required, ex: 5sec")
		os.Exit(-1)
	}

	timeStr := args[2][:1]

	fmt.Println(timeStr)
	i, err := strconv.Atoi(timeStr)
	if err != nil {
		fmt.Println("invalid argument for time")
		os.Exit(1)
	}

	duration := time.Second * time.Duration(i)

	go func() {
		expired := time.NewTimer(duration).C

		select {
		case <-expired:
			remind.Notify("title", args[0])
			ch <- true
		}

	}()

	<-ch
}
