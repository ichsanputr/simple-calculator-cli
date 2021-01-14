package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		result,_ := cmd.Flags().GetBool("float")

		if result{
			addFloat(args)
		}else{
			addInt(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("float","f",false,"Add Float Number")
}

func addInt(arg []string) {
	var res int

	for _, v := range arg {
		i,_ := strconv.Atoi(v)
		res += i 
	}

	fmt.Printf("Result %v \n",res)
}

func addFloat(arg []string) {
	var res float64

	for _, v := range arg {
		i,_ := strconv.ParseFloat(v,64)
		res += i 
	}

	fmt.Printf("Result %v \n",res)
}


