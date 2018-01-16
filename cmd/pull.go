// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"fmt"

	"github.com/spf13/cobra"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "A brief description of the pull command",
	Long:  `A long description of pull command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pull called")
		saveResponseAndConvert()
	},
}

//TODO: save the response file_content of the request into a file and then read-decode-save it again.
func init() {
	rootCmd.AddCommand(pullCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pullCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pullCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func saveResponseAndConvert() string {
	var ClientID = config.FetchAndValidate("ClientID") // returns string

	var ClientSecret = config.FetchAndValidate("ClientSecret") // returns string

	var ProjectKey = config.FetchAndValidate("ProjectKey") // returns string

	resp, err := resty.R().
		SetHeader("Content-Type", "multipart/form-data").
		SetFormData(map[string]string{
			"client_id":     ClientID,
			"client_secret": ClientSecret,
			"project_key":   ProjectKey}).
		Post("http://dev.devnagri.co.in/api/project/pull")

	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
	fmt.Println("\n\n")

	/*
		resJson, _ := gabs.ParseJSON([]byte(resp.String()))
			children, _ := resJson.S("file_content").Children()
			child := children[0]
	*/
	//fmt.Println(child.String())

	//data, _ := base64.StdEncoding.Decode(child.String())
	//encoded := child.String()
	//fmt.Println(encoded)
	//fmt.Println(reflect.TypeOf(encoded))
	//x := decodeBase64(encoded)
	//fmt.Println(x)

	// This works perfectly
	//fmt.Println("Decoding the string manually")
	//x := "PCEtLSBUcmFuc2xhdGVkIEJ5IERldm5hZ3JpIC0tPgo8IS0tIGh0dHA6Ly9kZXZuYWdyaS5jb20gLS0+CjxyZXNvdXJjZXMgdG9vbHM6aWdub3JlPSJFeHRyYVRyYW5zbGF0aW9uIiB4bWxuczp0b29scz0iaHR0cDovL3NjaGVtYXMuYW5kcm9pZC5jb20vdG9vbHMiPgogICAgPHN0cmluZyBuYW1lPSJhcHBfbmFtZSI+PC9zdHJpbmc+CiAgICA8c3RyaW5nIG5hbWU9ImhpbnRfYWN0dWFsIj48L3N0cmluZz4KIDwvcmVzb3VyY2VzPg=="
	//decodeBase64(x)

	dat, _ := ioutil.ReadFile("./content.txt")
	fmt.Println("reading content.txt")
	x := decodeBase64(string(dat))
	fmt.Println(x)
	return string(dat)
}
