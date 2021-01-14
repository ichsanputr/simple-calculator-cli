package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var substractCmd = &cobra.Command{
	Use:   "reduce",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		substractFloat(args)
	},
}

func init() {
	rootCmd.AddCommand(substractCmd)
}

func substractFloat(arg []string){
	first,_ := strconv.ParseFloat(arg[0],64)
	res := first

	for _,i := range arg {

		if i == arg[0]{
			continue
		}
		
		a,_ := strconv.ParseFloat(i,64)
		res -= a
	}

	fmt.Printf("Result %v \n",res)
}
