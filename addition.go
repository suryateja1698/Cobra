package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var additionCmd = &cobra.Command{
	Use:   "addition",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		add(args)
	},
}

func add(args []string) {
	var sum int

	for _, val := range args {
		temp, err := strconv.Atoi(val)

		if err != nil {
			fmt.Println(err)
		}
		sum = sum + temp
	}
	fmt.Printf("Addition of numbers %s is %d", args, sum)
}

func init() {
	rootCmd.AddCommand(additionCmd)

}
