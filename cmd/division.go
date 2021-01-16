package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// divisionCmd represents the division command
var divisionCmd = &cobra.Command{
	Use:   "divide",
	Short: "Divide your all number in here",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		divideFloat(args)
	},
}

func init() {
	rootCmd.AddCommand(divisionCmd)
}

func divideFloat(arg []string){
	a,_ := strconv.ParseFloat(arg[0],64)
	res := a

	for i := range arg {

		j,_ := strconv.ParseFloat(arg[i],64)

		if j == a{
			continue
		}

		res /= j
	}

	fmt.Printf("Result %v \n",res)
}
