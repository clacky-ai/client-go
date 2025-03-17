//
// Copyright 2021, Sander van Harmelen
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
//

package gitlab

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestValidate(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		description string
		opts        *LintOptions
		response    string
		want        *LintResult
	}{
		{
			description: "valid",
			opts: &LintOptions{
				Content: `
				build1:
					stage: build
					script:
						- echo "Do your build here"`,
				IncludeMergedYAML: true,
				IncludeJobs:       false,
			},
			response: `{
				"status": "valid",
				"errors": [],
				"merged_yaml":"---\nbuild1:\n    stage: build\n    script:\n    - echo\"Do your build here\""
			}`,
			want: &LintResult{
				Status:     "valid",
				MergedYaml: "---\nbuild1:\n    stage: build\n    script:\n    - echo\"Do your build here\"",
				Errors:     []string{},
			},
		},
		{
			description: "invalid",
			opts: &LintOptions{
				Content: `
					build1:
						- echo "Do your build here"`,
			},
			response: `{
				"status": "invalid",
				"errors": ["error message when content is invalid"]
			}`,
			want: &LintResult{
				Status: "invalid",
				Errors: []string{"error message when content is invalid"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			mux, client := setup(t)

			mux.HandleFunc("/api/v4/ci/lint", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodPost)
				fmt.Fprint(w, tc.response)
			})

			got, _, err := client.Validate.Lint(tc.opts)
			if err != nil {
				t.Errorf("Validate returned error: %v", err)
			}

			want := tc.want
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Validate returned \ngot:\n%v\nwant:\n%v", Stringify(got), Stringify(want))
			}
		})
	}
}

func TestValidateProject(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		description string
		response    string
		want        *ProjectLintResult
	}{
		{
			description: "valid",
			response: `{
				"valid": true,
				"errors": [],
				"warnings": [],
				"merged_yaml": 	"---\n:build:\n  :script:\n  - echo build"
			}`,
			want: &ProjectLintResult{
				Valid:      true,
				Warnings:   []string{},
				Errors:     []string{},
				MergedYaml: "---\n:build:\n  :script:\n  - echo build",
			},
		},
		{
			description: "invalid",
			response: `{
				"valid": false,
				"errors": ["jobs:build config contains unknown keys: bad_key"],
				"warnings": [],
				"merged_yaml": 	"---\n:build:\n  :script:\n  - echo build\n  :bad_key: value"
			}`,
			want: &ProjectLintResult{
				Valid:      false,
				Warnings:   []string{},
				Errors:     []string{"jobs:build config contains unknown keys: bad_key"},
				MergedYaml: "---\n:build:\n  :script:\n  - echo build\n  :bad_key: value",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			mux, client := setup(t)

			mux.HandleFunc("/api/v4/projects/1/ci/lint", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodGet)
				fmt.Fprint(w, tc.response)
			})

			opt := &ProjectLintOptions{}
			got, _, err := client.Validate.ProjectLint(1, opt)
			if err != nil {
				t.Errorf("Validate returned error: %v", err)
			}

			want := tc.want
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Validate returned \ngot:\n%v\nwant:\n%v", Stringify(got), Stringify(want))
			}
		})
	}
}

func TestValidateProjectNamespace(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		description string
		request     *ProjectNamespaceLintOptions
		response    string
		want        *ProjectLintResult
	}{
		{
			description: "valid",
			request: &ProjectNamespaceLintOptions{
				Content:     Ptr("{'build': {'script': 'echo build'}}"),
				DryRun:      Ptr(false),
				IncludeJobs: Ptr(true),
				Ref:         Ptr("foo"),
			},
			response: `{
				"valid": true,
				"errors": [],
				"warnings": [],
				"merged_yaml": 	"---\n:build:\n  :script:\n  - echo build",
				"includes": [
					{
						"type": "file",
      					"location": "template/pipeline.yml",
      					"blob": "https://gitlab.com/namespace/project/-/blob/abcd1234/template/pipeline.yml",
      					"raw": "https://gitlab.com/namespace/project/-/raw/abcd1234/template/pipeline.yml",
      					"extra": {
        					"project": "namespace/project",
        					"ref": "1.2.3"
      					},
      					"context_project": "namespace/current-project",
      					"context_sha": "abcd1234"
    				}
				]
			}`,
			want: &ProjectLintResult{
				Valid:      true,
				Warnings:   []string{},
				Errors:     []string{},
				MergedYaml: "---\n:build:\n  :script:\n  - echo build",
				Includes: []Include{
					{
						Type:     "file",
						Location: "template/pipeline.yml",
						Blob:     "https://gitlab.com/namespace/project/-/blob/abcd1234/template/pipeline.yml",
						Raw:      "https://gitlab.com/namespace/project/-/raw/abcd1234/template/pipeline.yml",
						Extra: map[string]interface{}{
							"project": "namespace/project",
							"ref":     "1.2.3",
						},
						ContextProject: "namespace/current-project",
						ContextSHA:     "abcd1234",
					},
				},
			},
		},
		{
			description: "invalid",
			request: &ProjectNamespaceLintOptions{
				Content: Ptr("{'build': {'script': 'echo build', 'bad_key': 'value'}}"),
				DryRun:  Ptr(false),
			},
			response: `{
				"valid": false,
				"errors": ["jobs:build config contains unknown keys: bad_key"],
				"warnings": [],
				"merged_yaml": 	"---\n:build:\n  :script:\n  - echo build\n  :bad_key: value"
			}`,
			want: &ProjectLintResult{
				Valid:      false,
				Warnings:   []string{},
				Errors:     []string{"jobs:build config contains unknown keys: bad_key"},
				MergedYaml: "---\n:build:\n  :script:\n  - echo build\n  :bad_key: value",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			mux, client := setup(t)

			mux.HandleFunc("/api/v4/projects/1/ci/lint", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodPost)
				fmt.Fprint(w, tc.response)
			})

			got, _, err := client.Validate.ProjectNamespaceLint(1, tc.request)
			if err != nil {
				t.Errorf("Validate returned error: %v", err)
			}

			want := tc.want
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Validate returned \ngot:\n%v\nwant:\n%v", Stringify(got), Stringify(want))
			}
		})
	}
}

func TestValidateProjectLint(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		description string
		request     *ProjectLintOptions
		response    string
		want        *ProjectLintResult
	}{
		{
			description: "valid",
			request: &ProjectLintOptions{
				DryRun:      Ptr(false),
				IncludeJobs: Ptr(true),
				ContentRef:  Ptr("foo"),
			},
			response: `{
				"valid": true,
				"errors": [],
				"warnings": [],
				"merged_yaml": 	"---\n:build:\n  :script:\n  - echo build",
				"includes": [
					{
						"type": "file",
      					"location": "template/pipeline.yml",
      					"blob": "https://gitlab.com/namespace/project/-/blob/abcd1234/template/pipeline.yml",
      					"raw": "https://gitlab.com/namespace/project/-/raw/abcd1234/template/pipeline.yml",
      					"extra": {
        					"project": "namespace/project",
        					"ref": "1.2.3"
      					},
      					"context_project": "namespace/current-project",
      					"context_sha": "abcd1234"
    				}
				]
			}`,
			want: &ProjectLintResult{
				Valid:      true,
				Warnings:   []string{},
				Errors:     []string{},
				MergedYaml: "---\n:build:\n  :script:\n  - echo build",
				Includes: []Include{
					{
						Type:     "file",
						Location: "template/pipeline.yml",
						Blob:     "https://gitlab.com/namespace/project/-/blob/abcd1234/template/pipeline.yml",
						Raw:      "https://gitlab.com/namespace/project/-/raw/abcd1234/template/pipeline.yml",
						Extra: map[string]interface{}{
							"project": "namespace/project",
							"ref":     "1.2.3",
						},
						ContextProject: "namespace/current-project",
						ContextSHA:     "abcd1234",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			mux, client := setup(t)

			mux.HandleFunc("/api/v4/projects/1/ci/lint", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodGet)
				fmt.Fprint(w, tc.response)
			})

			got, _, err := client.Validate.ProjectLint(1, tc.request)
			if err != nil {
				t.Errorf("Validate returned error: %v", err)
			}

			want := tc.want
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Validate returned \ngot:\n%v\nwant:\n%v", Stringify(got), Stringify(want))
			}
		})
	}
}
