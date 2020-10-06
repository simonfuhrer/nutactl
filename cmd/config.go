// Copyright © 2020 Simon Fuhrer
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
	"io"

	"github.com/simonfuhrer/nutactl/cmd/displayers"
)

type Config struct {
	Contexts      []*Context `mapstructure:"contexts"`
	ActiveContext string     `mapstructure:"active_context"`
}

type Context struct {
	Name     string `mapstructure:"name"`
	Endpoint string `mapstructure:"endpoint"`
	Password string `mapstructure:"password"`
	User     string `mapstructure:"user"`
}

func (o Config) JSON(w io.Writer) error {
	return displayers.DisplayJSON(w, o)
}

func (o Config) JSONPath(w io.Writer, template string) error {
	return displayers.DisplayJSONPath(w, template, o)
}

func (o Config) PP(w io.Writer) error {
	return displayers.DisplayPP(w, o)
}

func (o Config) YAML(w io.Writer) error {
	return displayers.DisplayYAML(w, o)
}

func (o Config) header() []string {
	return []string{
		"Active",
		"Name",
		"Endpoint",
		"User",
	}
}

func (o Config) TableData(w io.Writer) error {
	data := make([][]string, len(o.Contexts))
	for i, context := range o.Contexts {
		active := ""
		if o.ActiveContext == context.Name {
			active = "*"
		}
		data[i] = []string{
			active,
			context.Name,
			context.Endpoint,
			context.User,
		}
	}
	return displayers.DisplayTable(w, data, o.header())
}

func (o Config) Text(w io.Writer) error {
	return nil
}

func (o *Config) ContextByName(name string) *Context {
	for _, c := range o.Contexts {
		if c.Name == name {
			return c
		}
	}
	return nil
}

func (o *Config) RemoveContext(context *Context) {
	for i, c := range o.Contexts {
		if c == context {
			o.Contexts = append(o.Contexts[:i], o.Contexts[i+1:]...)
			return
		}
	}
}
