/*
Copyright 2016 The Kubernetes Authors.

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
	"k8s.io/test-infra/prow/config"
	"k8s.io/test-infra/prow/github"
	"k8s.io/test-infra/prow/plugins"
)

const (
	pluginName = "slackevents"
)

func init() {
	plugins.RegisterPushEventHandler(pluginName, handlePush)
}

func handlePush(pc plugins.PluginClient, pe github.PushEvent) error {
	return notifyOnSlackIfManualMerge(pc, pe)
}

func getSlackEvent(pc plugins.PluginClient, repo string) *config.SlackEvent {
	for _, se := range pc.Config.SlackEvents {
		if stringInArray(repo, se.Repos) {
			return &se
		}
	}
	return nil
}

func stringInArray(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
