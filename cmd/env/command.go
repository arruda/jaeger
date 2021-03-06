// Copyright (c) 2018 The Jaeger Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package env

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/jaegertracing/jaeger/plugin/storage"
)

const (
	longTemplate = `
All command line options can be provided via environment variables by converting
their names to upper case and replacing punctuation with underscores. For example:

command line option                 environment variable
------------------------------------------------------------------
--cassandra.connections-per-host    CASSANDRA_CONNECTIONS_PER_HOST
--metrics-backend                   METRICS_BACKEND

The following configuration options are only available via environment variables:
%s
`
)

// Command creates `env` command
func Command() *cobra.Command {
	fs := new(pflag.FlagSet)
	fs.String(storage.SpanStorageTypeEnvVar, "cassandra", fmt.Sprintf("The type of backend %s used for trace storage. Multiple backends can be specified (currently only for writing spans) as comma-separated list, e.g. \"cassandra,kafka\".", storage.AllStorageTypes))
	fs.String(storage.DependencyStorageTypeEnvVar, "${SPAN_STORAGE}", "The type of backend used for service dependencies storage.")
	long := fmt.Sprintf(longTemplate, strings.Replace(fs.FlagUsagesWrapped(0), "      --", "", -1))
	return &cobra.Command{
		Use:   "env",
		Short: "Help about environment variables.",
		Long:  long,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(cmd.OutOrStdout(), long)
		},
	}
}
