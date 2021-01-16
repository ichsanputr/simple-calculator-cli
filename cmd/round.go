package cmd

import (
	"fmt"
	"math"
	"strconv"

	"github.com/spf13/cobra"
)

var(
	phi1 bool
	phi2 bool
)

var roundCmd = &cobra.Command{
	Use:   "round",
	Short: "Count the area of the circle in here",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		if phi1 == false && phi2 == false{
			fmt.Println("Please use a flags -1 (phi : 22/7) or -2 (phi: 3.14)")
		}
		if phi1 == true{
			s := areaCircle(args)*22/7
			fmt.Printf("Result %v \n",s)
		}else if phi2 == true{
			j := areaCircle(args)*3.14
			fmt.Printf("Result %v \n",j)
		}
	},
}

func init() {
	rootCmd.AddCommand(roundCmd)
	roundCmd.Flags().BoolVarP(&phi1,"22/7","1",false,"For phi with 22/7")
	roundCmd.Flags().BoolVarP(&phi2,"3.14","2",false,"For phi with 3.14")
}

func areaCircle(arg []string) float64 {
	a,_ := strconv.ParseFloat(arg[0],64)
	b := math.Pow(a,2)

	return b
}
