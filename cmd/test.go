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
	"github.com/cobre420/mapper-go/domain"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var mappingFile string
var configFile string

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "execute test run of mapping file",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test called")

		log.Infof("Mapping file=%s", mappingFile)
		log.Infof("Config file=%s", configFile)

		m := domain.Mapping{}

		err := loadYamlFile(&m, mappingFile)
		if err != nil {
			panic(err)
		}

		log.Infof("Loaded %v", m)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	rootCmd.PersistentFlags().StringVar(&mappingFile, "mappingFile", "", "mapping file")
	rootCmd.PersistentFlags().StringVar(&configFile, "configFile", "", "config file")



	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func loadYamlFile(out interface{}, resourceFile string) error {
	data, err := ioutil.ReadFile(resourceFile)
	if err != nil {
		panic(err)
	}
	return yaml.Unmarshal(data, out)
}

//func loadFile(name, fileName string) *viper.Viper {
//	result := viper.New()
//
//	result.SetConfigName(name)
//	result.SetConfigFile(fileName)
//
//	mapping := domain.Mapping{}
//	err := result.Unmarshal(&mapping, func(config *mapstructure.DecoderConfig){
//		config.TagName = "yaml"
//	})
//	if err != nil {
//		panic(err)
//	}
//
//	log.Infof("Mapping struct %v", mapping)
//
//	err = result.ReadInConfig()
//	if err != nil {
//		panic(err)
//	}
//
//	return result
//}
