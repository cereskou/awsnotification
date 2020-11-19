package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

//global
var (
	region string = "ap-northeast-1"
	jst    *time.Location
)

//Environment -
type Environment struct {
	Prefix     string `json:"prefix"`
	Name       string `json:"name"`
	Themecolor string `json:"themecolor"`
	ActionName string `json:"actionname"`
}

//Button -
type Button struct {
	Name    string              `json:"name"`
	Type    string              `json:"type"`
	Targets []map[string]string `json:"targets"`
}

//Webhook -
type Webhook struct {
	Topic       string `json:"topic"`
	Description string `json:"description"`
	Service     string `json:"service"`
	URL         string `json:"url"`
	Message     struct {
		Title string `json:"title"`
	} `json:"message"`
	Buttons []*Button `json:"buttons,omitempty"`
}

//teamsSetting -
type teamsSetting struct {
	Webhooks     []*Webhook     `json:"webhooks"`
	Account      string         `json:"account,omitempty"`
	Password     string         `json:"password,omitempty"`
	Environments []*Environment `json:"environments"`
}

func getSsmKey() (ssmkey string) {
	ssmkey = "TeamsSettings"
	envstr := os.Getenv("SSMKEY")
	if len(envstr) > 0 {
		ssmkey = envstr
	}
	fmt.Printf("Read from %v\n", ssmkey)

	return ssmkey
}

func getRegion() (region string) {
	region = "ap-northeast-1"
	envstr := os.Getenv("AWS_DEFAULT_REGION")
	if len(envstr) > 0 {
		region = envstr
	}
	fmt.Printf("AWS Region %v\n", region)

	return
}

//Handler -
func Handler(ctx context.Context, event events.SNSEvent) error {
	jst = time.FixedZone("Asia/Tokyo", 9*60*60)
	//region
	region = getRegion()
	//同期対象情報は、SSM Parameter Storeから取得します
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	//パラメター名：SyncTargetGroup (Fixed))
	ssmkey := getSsmKey()
	svc := ssm.New(sess)
	result, err := svc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(ssmkey),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		//StatusCode: 500 internal server error
		fmt.Println(err)
		return err
	}
	var setting teamsSetting
	err = json.Unmarshal([]byte(*result.Parameter.Value), &setting)
	if err != nil {
		//StatusCode 500 internal server error
		fmt.Println(err)
		return err
	}

	//debug
	// se, _ := json.Marshal(setting)
	// fmt.Println(string(se))

	//for debug
	// a, _ := json.Marshal(event)
	// fmt.Println(string(a))

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	for _, record := range event.Records {
		sns := record.SNS

		webhook := findWebhook(record.SNS.TopicArn, setting.Webhooks)
		if webhook == nil {
			err = errors.New("Unsupport SNS topic. " + record.SNS.TopicArn)
			return err
		}

		var card *Card
		switch webhook.Service {
		case "batch":
			card, err = SNSBatch(sns.Message, webhook, setting.Environments)
			if err != nil {
				fmt.Println(err)
				return err
			}
		case "codepipeline":
			card, err = SNSCodePipeline(sns.Message, webhook, setting.Environments)
			if err != nil {
				fmt.Println(err)
				return err
			}
		default:
			err = errors.New("Unsupport SNS topic. " + record.SNS.TopicArn)
			return err
		}

		//null check
		if card == nil {
			err = errors.New("Unsupport SNS topic. " + record.SNS.TopicArn)
			return err
		}

		b, _ := json.Marshal(card)
		wb := bytes.NewBuffer(b)

		req, _ := http.NewRequest(http.MethodPost, webhook.URL, wb)
		req.Header.Add("Content-Type", "application/json;charset=utf-8")
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return err
		}
		if res.StatusCode >= 299 {
			err = errors.New("error on notification: " + res.Status)
			return err
		}
	}
	fmt.Println("done")

	return nil
}

func main() {
	lambda.Start(Handler)
}
