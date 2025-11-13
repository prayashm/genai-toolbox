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

package mongodblistcollectionnames_test

import (
	"strings"
	"testing"

	"github.com/googleapis/genai-toolbox/internal/tools"
	"github.com/googleapis/genai-toolbox/internal/tools/mongodb/mongodblistcollectionnames"

	yaml "github.com/goccy/go-yaml"
	"github.com/google/go-cmp/cmp"
	"github.com/googleapis/genai-toolbox/internal/server"
	"github.com/googleapis/genai-toolbox/internal/testutils"
)

func TestParseFromYamlListCollectionNames(t *testing.T) {
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
			desc: "with database in config",
			in: `
			tools:
				list_collections:
					kind: mongodb-list-collection-names
					source: my-mongodb-source
					description: List collections in a database
					database: my_database
			`,
			want: server.ToolConfigs{
				"list_collections": mongodblistcollectionnames.Config{
					Name:         "list_collections",
					Kind:         "mongodb-list-collection-names",
					Source:       "my-mongodb-source",
					AuthRequired: []string{},
					Description:  "List collections in a database",
					Database:     "my_database",
					Params:       tools.Parameters{},
				},
			},
		},
		{
			desc: "without database in config",
			in: `
			tools:
				list_collections_dynamic:
					kind: mongodb-list-collection-names
					source: my-mongodb-source
					description: List collections in any database
			`,
			want: server.ToolConfigs{
				"list_collections_dynamic": mongodblistcollectionnames.Config{
					Name:         "list_collections_dynamic",
					Kind:         "mongodb-list-collection-names",
					Source:       "my-mongodb-source",
					AuthRequired: []string{},
					Description:  "List collections in any database",
					Database:     "",
					Params:       tools.Parameters{},
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
