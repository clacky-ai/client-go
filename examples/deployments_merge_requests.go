//
// Copyright 2022, Daniela Filipe Bento
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

package main

import (
	"log"

	gitlab "gitlab.com/gitlab-org/api/client-go"
)

func deploymentExample() {
	git, err := gitlab.NewClient("yourtokengoeshere")
	if err != nil {
		log.Fatal(err)
	}

	opt := &gitlab.ListMergeRequestsOptions{}
	mergeRequests, _, err := git.DeploymentMergeRequests.ListDeploymentMergeRequests(1, 1, opt)
	if err != nil {
		log.Fatal(err)
	}

	for _, mergeRequest := range mergeRequests {
		log.Printf("Found merge request: %v\n", mergeRequest)
	}
}
