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

//TODO: change this package name
//package cmd

package main

import (
	"fmt"

	"gopkg.in/resty.v1"
)

func main() {

	resp, err := resty.R().
		SetHeader("Content-Type", "multipart/form-data").
		SetFormData(map[string]string{
			"client_id":     "2",
			"client_secret": "y8umxMS54nUBc1ak7cxod6mjYiAbht2rCNAKsW7c",
			"project_key":   "3c5724f8fe2d45e1dea65f8842cebd79 "}).
		Post("http://dev.devnagri.co.in/api/key/validations")
		//	Post("http://192.168.60.10/api/key/validations")

	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
	//TODO: Save the returned access_token to the .devnagri.yaml
}
