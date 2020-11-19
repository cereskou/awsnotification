package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

//BatchEvent -
type BatchEvent struct {
	Version    string    `json:"version"`
	ID         string    `json:"id"`
	DetailType string    `json:"detail-type"`
	Source     string    `json:"source"`
	Account    string    `json:"account"`
	Time       time.Time `json:"time"`
	Region     string    `json:"region"`
	Resources  []string  `json:"resources"`
	Detail     struct {
		JobArn   string `json:"jobArn"`
		JobName  string `json:"jobName"`
		JobID    string `json:"jobId"`
		JobQueue string `json:"jobQueue"`
		Status   string `json:"status"`
		Attempts []struct {
			Container struct {
				ContainerInstanceArn string        `json:"containerInstanceArn"`
				TaskArn              string        `json:"taskArn"`
				ExitCode             int           `json:"exitCode"`
				LogStreamName        string        `json:"logStreamName"`
				NetworkInterfaces    []interface{} `json:"networkInterfaces"`
			} `json:"container"`
			StartedAt    int64  `json:"startedAt"`
			StoppedAt    int64  `json:"stoppedAt"`
			StatusReason string `json:"statusReason"`
		} `json:"attempts"`
		StatusReason  string            `json:"statusReason"`
		CreatedAt     int64             `json:"createdAt"`
		StartedAt     int64             `json:"startedAt"`
		StoppedAt     int64             `json:"stoppedAt"`
		DependsOn     []interface{}     `json:"dependsOn"`
		JobDefinition string            `json:"jobDefinition"`
		Parameters    map[string]string `json:"parameters"`
		Container     struct {
			Image       string        `json:"image"`
			Vcpus       int           `json:"vcpus"`
			Memory      int           `json:"memory"`
			Command     []string      `json:"command"`
			JobRoleArn  string        `json:"jobRoleArn"`
			Volumes     []interface{} `json:"volumes"`
			Environment []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"environment"`
			MountPoints          []interface{} `json:"mountPoints"`
			Ulimits              []interface{} `json:"ulimits"`
			ExitCode             int           `json:"exitCode"`
			ContainerInstanceArn string        `json:"containerInstanceArn"`
			TaskArn              string        `json:"taskArn"`
			LogStreamName        string        `json:"logStreamName"`
			NetworkInterfaces    []interface{} `json:"networkInterfaces"`
			ResourceRequirements []interface{} `json:"resourceRequirements"`
			Secrets              []interface{} `json:"secrets"`
		} `json:"container"`
		Tags struct {
			ResourceArn string `json:"resourceArn"`
		} `json:"tags"`
	} `json:"detail"`
}

//SNSBatch -
func SNSBatch(message string, webhook *Webhook, envs []*Environment) (*Card, error) {
	var job BatchEvent
	err := json.Unmarshal([]byte(message), &job)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//get environment
	env := findEnvironment(job.Detail.JobDefinition, envs)
	if env == nil {
		env = envs[0]
	}

	title := fmt.Sprintf("%v %v", env.Name, webhook.Message.Title)
	card := NewCard()
	card.Summary = "summary"
	card.ThemeColor = env.Themecolor
	card.Title = title

	image := ""
	if job.Detail.Status == "FAILED" {
		image = "![](" + FAILED + ")   "
	}
	card.Text = fmt.Sprintf("%v**jobid :**%v", image, job.Detail.JobID)

	t := job.Detail.CreatedAt / int64(time.Microsecond)
	ts := job.Detail.CreatedAt % int64(time.Microsecond)
	createdAt := time.Unix(t, ts).In(jst)
	params := ""
	for k, v := range job.Detail.Parameters {
		params += fmt.Sprintf("**%v:** %v<br>", k, v)
	}

	//
	section := card.NewSection()
	section.ActivityTitle = job.Detail.JobName
	section.ActivitySubtitle = job.Detail.JobDefinition
	section.Text = "Parameters"
	section.NewFact("Name", job.Detail.JobName)
	section.NewFact("Queue", job.Detail.JobQueue)
	section.NewFact("Status", job.Detail.Status)
	section.NewFact("CreatedAt", createdAt.Format("2006/01/02 15:04:05"))
	section.NewFact("Parameters", params)

	for _, btn := range webhook.Buttons {
		action := card.NewAction(btn.Name, btn.Type)

		target := make(map[string]string)
		action.AddTarget(target)

		for _, kv := range btn.Targets {
			for k, v := range kv {
				if k == "uri" {
					v = strings.ReplaceAll(v, "<REGION>", region)
					v = strings.ReplaceAll(v, "<JOBID>", job.Detail.JobID)
				}
				target[k] = v
			}
		}
	}

	return card, nil
}
