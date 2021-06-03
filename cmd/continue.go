/*
Copyright © 2021 Ashwin Raam <raamocyby@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"now/funcs"
	"strings"

	"github.com/spf13/cobra"
)

func continueEntry(cmd *cobra.Command) {
	record := funcs.ReadLastEntry(1)
	status := strings.Trim(record[0][3], " ")
	if status == "bstarted" {
		tags, _ := cmd.Flags().GetStringSlice("tags")
		entry := funcs.Continue(tags)
		funcs.AppendEntry(entry)
	} else if status == "completed" {
		fmt.Println()
		fmt.Println("You haven't taken any break to continue from.")
		fmt.Println("Last completed Task:")
		funcs.PrintEntries(record)
	} else if status == "started" {
		fmt.Println()
		fmt.Println("You haven't taken any break to continue from.")
		fmt.Println("Task currently going on:")
		funcs.PrintEntries(record)
	} else {
		fmt.Println()
		fmt.Println("You haven't taken any break to continue from.")
	}
}

// continueCmd represents the continue command
var continueCmd = &cobra.Command{
	Use:   "continue",
	Short: "continue command adds break completed entry.",
	Long: `continue command adds break completed entry.

Example:

now continue
now continue --tags tag1,tag2
now continue -t tag1,tag2

`,
	Run: func(cmd *cobra.Command, args []string) {
		continueEntry(cmd)
	},
}

func init() {
	rootCmd.AddCommand(continueCmd)
	continueCmd.Flags().StringSliceP("tags", "t", []string{"General"}, "adds tags")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// continueCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// continueCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
