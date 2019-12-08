/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"os"

	"github.com/TakeruTakeru/study/mycli/internal/golang_util"
	"github.com/spf13/cobra"
)

var Source string

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use:   "go",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		outputDir, err := cmd.Flags().GetString("output")
		if err != nil {
			m := fmt.Errorf("%v", err)
			fmt.Println(m)
			os.Exit(1)
		}
		pathStat, err := os.Stat(outputDir)
		if err != nil {
			m := fmt.Errorf("%v", err)
			fmt.Println(m)
			os.Exit(1)
		}
		if !pathStat.IsDir() {
			m := fmt.Errorf("%v is not directory", pathStat.Name())
			fmt.Println(m)
			os.Exit(1)
		}
		fmt.Printf("go called + '%s'", outputDir)
		golang_util.InitializeEnv(outputDir)
	},
}

func init() {
	goCmd.Flags().StringVarP(&Source, "output", "o", "", "Path to output config file.")
	rootCmd.AddCommand(goCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// goCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// goCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
