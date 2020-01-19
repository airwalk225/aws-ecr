/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	ecrbox "github.com/airwalk225/ecrbox/aws/ecr"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/spf13/cobra"
	"strings"
)

var names string
// describeCmd represents the describe command
var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe an AWS ECR",
	Long: `Used to get information about an Amazon Web Service Elastic Container Registry`,
	Run: func(cmd *cobra.Command, args []string) {
		var filterNames []*string

		_ecr := ecrbox.NewECR(cnf.AWS.Session)

		for _, name := range strings.Split(names, ",") {
			filterNames = append(filterNames, &name)
		}

		flt := &ecr.DescribeRepositoriesInput{
			RepositoryNames: filterNames,
		}

		res, err := ecrbox.GetRepositories(_ecr, flt)

		if err != nil {
			_ecr.HandleError(err, "repo describe")
		}

		ecrbox.DescribeRepoInDepth(res)
	},
}

func init() {
	repoCmd.AddCommand(describeCmd)

	describeCmd.Flags().StringVarP(&names, "name", "n", "", "The name of the repository you would like to describe (required) - can be a comma separated list")
	_ = describeCmd.MarkFlagRequired("name")
}
