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

	"github.com/TakeruTakeru/study/mycli/internal/base64"
	"github.com/spf13/cobra"
)

var DecodeFlag bool
var EncodingType string

// base64Cmd represents the base64 command
var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var encodeType base64.EncoderType
		if len(args) != 1 {
			os.Exit(1)
		}
		switch EncodingType {
		case "standard":
			encodeType = base64.STANDARD
		case "url":
			encodeType = base64.URL
		case "standard-nopad":
			encodeType = base64.STANDARD_NO_PADDING
		case "url-nopad":
			encodeType = base64.URL_NO_PADDING
		default:
			err := fmt.Errorf("Invalid Type: %s", EncodingType)
			panic(err)
		}
		base64.SetEncodeType(encodeType)
		if DecodeFlag {
			s, _ := base64.Decode(args[0])
			fmt.Print(s)
		} else {
			s := base64.Encode(args[0])
			fmt.Print(s)
		}
	},
}

func init() {
	base64Cmd.Flags().BoolVarP(&DecodeFlag, "decode", "d", false, "Switch to decode mode.")
	base64Cmd.Flags().StringVarP(&EncodingType, "type", "t", "standard", ENCODE_TYPE_HELP_MESSAGE)
	rootCmd.AddCommand(base64Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// base64Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// base64Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

const (
	ENCODE_TYPE_HELP_MESSAGE = `Change Encoding Type. Choose below type.
		- standard: standard base64 encoding with padding.
		- url: url safe base64 encoding with padding.
		- standard-nopad: standard base64 encoding with no padding.
		- url-nopad: url safe base64 encoding with nopadding.
	`
)
