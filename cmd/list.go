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
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/spf13/cobra"
)

type Repository struct {
	CreatedAt string
	ImageScanningConfiguration ImageScanningConfiguration
	ImageTagMutability string
	RegistryId int
	RepositoryArn string
	RepositoryName string
	RepositoryUri string
}

type ImageScanningConfiguration struct {
	ScanOnPush bool
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List ECR repositories",
	Long: `List all the repositories that exist inside the AWS ECR provided.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ecr := NewECR(cnf.AWS.Session)
		flt := &ecr.DescribeRepositoriesInput{}
		res, err := getRepositories(_ecr, flt)

		if err != nil {
			_ecr.handleError(err, "repo list")
		}

		displayRepos(res)
	},
}

func init() {
	repoCmd.AddCommand(listCmd)
}