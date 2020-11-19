package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

//CodePipelineEvent -
type CodePipelineEvent struct {
	Version    string    `json:"version"`
	ID         string    `json:"id"`
	DetailType string    `json:"detail-type"`
	Source     string    `json:"source"`
	Account    string    `json:"account"`
	Time       time.Time `json:"time"`
	Region     string    `json:"region"`
	Resources  []string  `json:"resources"`
	Detail     struct {
		Pipeline    string  `json:"pipeline"`
		ExecutionID string  `json:"execution-id"`
		State       string  `json:"state"`
		Version     float64 `json:"version"`
	} `json:"detail"`
}

//SNSCodePipeline -
func SNSCodePipeline(message string, webhook *Webhook, envs []*Environment) (*Card, error) {
	var event CodePipelineEvent
	err := json.Unmarshal([]byte(message), &event)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//get environment
	env := findEnvironment(event.Detail.Pipeline, envs)
	if env == nil {
		env = envs[0]
	}

	card := NewCard()
	card.Summary = "summary"
	card.ThemeColor = env.Themecolor
	title := fmt.Sprintf("%v %v %v", env.Name, event.Detail.Pipeline, webhook.Message.Title)
	card.Title = title

	section := card.NewSection()

	image := ""
	if event.Detail.State == "SUCCEEDED" {
		image = "![](" + SUCCEEDED + ")"
	} else if event.Detail.State == "FAILED" {
		image = "![](" + FAILED + ")"
	} else {
		image = "![](" + WARNING + ")"
	}
	card.Text = fmt.Sprintf("%v  **State :**%v", image, event.Detail.State)

	tm := event.Time.In(jst)
	section.NewFact("time", tm.Format("2006/01/02 15:04:05"))
	section.NewFact("pipeline", event.Detail.Pipeline)
	section.NewFact("execution-id", event.Detail.ExecutionID)
	section.NewFact("state", event.Detail.State)

	//Button
	for _, btn := range webhook.Buttons {
		action := card.NewAction(btn.Name, btn.Type)

		target := make(map[string]string)
		action.AddTarget(target)

		for _, kv := range btn.Targets {
			for k, v := range kv {
				if k == "uri" {
					v = strings.ReplaceAll(v, "<REGION>", region)
					v = strings.ReplaceAll(v, "<PIPELINE>", event.Detail.Pipeline)
					v = strings.ReplaceAll(v, "<EXECUTION-ID>", event.Detail.ExecutionID)
				}
				target[k] = v
			}
		}
	}
	return card, nil
}
