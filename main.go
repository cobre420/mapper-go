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
package main

import (
	"github.com/cobre420/mapper-go/cmd"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

var (
	BuildVersion string = ""
	BuildTime    string = ""
	GitRevision  string = ""
)

func main() {
	log.SetFormatter(
		&log.TextFormatter{
			ForceColors:               true,
			DisableColors:             false,
			ForceQuote:                false,
			DisableQuote:              false,
			EnvironmentOverrideColors: false,
			DisableTimestamp:          false,
			FullTimestamp:             true,
			DisableSorting:            false,
			SortingFunc:               nil,
			DisableLevelTruncation:    false,
			PadLevelText:              false,
			QuoteEmptyFields:          false,
			FieldMap:                  nil,
			CallerPrettyfier:          nil,
		})

	logFile, _ := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	cmd.Execute()
}
