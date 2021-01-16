package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var multipleCmd = &cobra.Command{
	Use:   "multiple",
	Short: "Multiple your number in here",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		multipleFloat(args)
	},
}

func init() {
	rootCmd.AddCommand(multipleCmd)
}

func multipleFloat(arg []string){
	a,_ := strconv.ParseFloat(arg[0],64)
	res := a

	for _,i := range arg{
		s,_ := strconv.ParseFloat(i,64)

		if i == arg[0]{
			continue
		}

		res *= s

	}

	fmt.Printf("Result %v \n",res)
}
