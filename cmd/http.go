/*
Copyright © 2020 DEBUG.ICU <debugicu@163.com>

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
	"F10-CLI/http"
	"F10-CLI/service"
	"github.com/spf13/cobra"
	"log"
)

var port string
var archive string

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "启动版本文件管理服务",
	Long:  `HttpService`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(port) == 0 {
			port = "9980"
		}
		service.SetVersionArchiveFileName(archive)
		service.LoadLocalVersionMap()
		go service.Crond()
		http.HttpService(port)
		log.Println("http called")
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
	httpCmd.Flags().StringVarP(&port, "port", "p", "9980", "服务端口号")
	httpCmd.Flags().StringVarP(&archive, "archive", "a", "./version_archive.json", "归档持久化文件名")
}
