/*
Copyright 2017 The Kubernetes Authors.

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

package slackevents

import (
	"fmt"

	"k8s.io/test-infra/prow/github"
	"k8s.io/test-infra/prow/plugins"
)

func notifyOnSlackIfManualMerge(pc plugins.PluginClient, pe github.PushEvent) error {
	//fetch slackevent configuration for the repo we received the merge event
	if se := getSlackEvent(pc, pe.Repo.Name); se != nil {
		//if the slackevent whitelist  has the merge user then no need to send a message
		if !stringInArray(pe.Pusher.Name, se.WhiteList) && !stringInArray(pe.Sender.Login, se.WhiteList) {
			message := fmt.Sprintf("Warning: %s manually merged %s", pe.Pusher.Name, pe.Compare)
			for _, channel := range se.Channels {
				if err := pc.SlackClient.WriteMessage(message, channel); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
