package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

/*
Copyright © 2021 Fred Liang <fred@oasis.ac>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// preferenceListCmd represents the preference command
var preferenceListCmd = &cobra.Command{
	Use:   "list",
	Short: "List you personal preferences",
	Long: `List your personal preferences

e.g. lets pref list
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pref called")
	},
}

func init() {
	preferenceCmd.AddCommand(preferenceListCmd)

	// Here you will define your flags and preferenceuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// preferenceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// preferenceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
