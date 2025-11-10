// Copyright 2025 Google LLC
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

package mongodblistdatabasenames_test

import (
	"strings"
	"testing"

	"github.com/googleapis/genai-toolbox/internal/tools/mongodb/mongodblistdatabasenames"

	yaml "github.com/goccy/go-yaml"
	"github.com/google/go-cmp/cmp"
	"github.com/googleapis/genai-toolbox/internal/server"
	"github.com/googleapis/genai-toolbox/internal/testutils"
)

func TestParseFromYamlListDatabaseNames(t *testing.T) {
	ctx, err := testutils.ContextWithNewLogger()
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	tcs := []struct {
		desc string
		in   string
		want server.ToolConfigs
	}{
		{
			desc: "basic example",
			in: `
			tools:
				list_databases:
					kind: mongodb-list-database-names
					source: my-mongodb-source
					description: List all databases in MongoDB
			`,
			want: server.ToolConfigs{
				"list_databases": mongodblistdatabasenames.Config{
					Name:         "list_databases",
					Kind:         "mongodb-list-database-names",
					Source:       "my-mongodb-source",
					AuthRequired: []string{},
					Description:  "List all databases in MongoDB",
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			decoder := yaml.NewDecoder(strings.NewReader(tc.in))

			got, err := server.ParseToolConfigFile(ctx, decoder)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("ParseToolConfigFile() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
