// Copyright 2021 Mailio All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var (
	// ErrConfigMissingParameter raised when one of the required config parameters are missing
	ErrConfigMissingParameter = errors.New("missing required configuration parameter")
)

// YamlConfig exports crucial microservice  settings
type YamlConfig struct {
	Version     string `yaml:"version"`     // Version of the server
	Port        int    `yaml:"port"`        // Server port
	Swagger     bool   `yaml:"swagger"`     // Swagger on/off
	Title       string `yaml:"title"`       // Service title
	Description string `yaml:"description"` // Service description
	Mode        string `yaml:"mode"`        // debug/release
}

// NewYamlConfig loads the conf.yaml file and return the new config
func NewYamlConfig(pathtoconfig string, configObject interface{}) error {
	y := &YamlConfig{}
	return y.LoadConfig(pathtoconfig, configObject)
}

// LoadConfig reads local config.yaml file
func (c *YamlConfig) LoadConfig(pathToConfig string, config interface{}) error {
	yamlFile, err := ioutil.ReadFile(pathToConfig)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return err
	}
	return nil
}
