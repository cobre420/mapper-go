/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"github.com/cobre420/mapper-go/service"
	"github.com/spf13/cobra"
)

var mappingFile string
var configFile string

// testCmd represents the _test command
var TestCmd = &cobra.Command{
	Use:   "_test",
	Short: "execute _test run of mapping file",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("_test called")

		//log.Infof("Mapping file=%s", mappingFile)
		//log.Infof("Config file=%s", configFile)

		service.ProcessMapping(configFile, mappingFile)
	},
}

func init() {
	rootCmd.AddCommand(TestCmd)

	rootCmd.PersistentFlags().StringVar(&mappingFile, "mapping-file", "", "mapping file")
	rootCmd.PersistentFlags().StringVar(&configFile, "config-file", "", "config file")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
