package main

import "strings"

//findEnvironment - for AWS Batch JOB
func findEnvironment(def string, envs []*Environment) *Environment {
	for _, env := range envs {
		if strings.Contains(def, env.Prefix) {
			return env
		}
	}

	return nil
}

//findWebhook -
func findWebhook(topic string, whooks []*Webhook) *Webhook {
	for _, h := range whooks {
		if strings.HasSuffix(topic, h.Topic) {
			return h
		}
	}

	return nil
}
