// Copyright © 2019 Metrum Research Group
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
)

var commitFile string

// checkCmd represents the R CMD check command
var validateCmd = &cobra.Command{
	Use:   "validate a commit message",
	Short: "Validate a commit message",
	Long: `
 `,
	RunE: rValidate,
}

func rValidate(cmd *cobra.Command, args []string) error {

	return nil
}

func init() {
	validateCmd.Flags().StringVar(&commitFile, "commit-file", "", "path to commit file")
	RootCmd.AddCommand(validateCmd)
}
