package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
	"time-calculator/timeutils"
)

var sumCmd = &cobra.Command{
	Use:   "sum",
	Short: "sum command short description",
	Long:  `sum command long description`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		timeList := args

		isFile, _ := cmd.Flags().GetBool("file")
		if isFile {
			pwd, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			filePath := pwd + string(filepath.Separator) + args[0]

			file, err := os.Open(filePath)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer file.Close()

			timeList = []string{}
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				timeList = append(timeList, strings.TrimSpace(scanner.Text()))
			}

		}

		fmt.Println(timeutils.SumTimeList(timeList))

	},
}

func init() {
	sumCmd.Flags().BoolP("file", "f", false, "text file with time list")

	rootCmd.AddCommand(sumCmd)
}
