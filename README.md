# awsnotification
Post message to Microsoft Teams Channel when received a SNS event.

**AWS Lambda Function**
## Create Lambda Function
gobuild.cmd -> function.zip
<br>
create a new Lambda function named "**AWStoTeams**"

## Create a WebHook in Teams Channel
1. Click ... Button of Channel and Select Connectors.<br>
![popupmenu](https://github.com/cereskou/awsnotification/blob/main/images/popupmenu.png)

2. Incoming WebHook configure<br>
![incomingwebhook](https://github.com/cereskou/awsnotification/blob/main/images/webhook.png)

3. Create a new WebHook<br>
![incomingwebhook](https://github.com/cereskou/awsnotification/blob/main/images/incoming.png)

4. Copy URL to Clipboard.<br>
![incomingwebhook](https://github.com/cereskou/awsnotification/blob/main/images/incoming-done.png)

## Setup ssm.json
Save the ssm.json to SSM Parameter Store as key **TeamsSettings**.
```
{
	"webhooks": [{
            "topic": "job-failed-alert",
            "service": "batch",
			"description": "ジョブは失敗しました",
			"url": "https://outlook.office.com/webhook/dummydummy/IncomingWebhook/dummydummy",
			"message": {
				"title": "ジョブは失敗しました"
			},
			"buttons": [{
				"name": "詳細情報",
				"type": "OpenUri",
				"targets": [{
					"os": "default",
					"uri": "https://dummy-region.console.aws.amazon.com/batch/v2/home?region=<REGION>#jobs/detail/<JOBID>"
				}]
			}]
		},
		{
			"topic": "batch-codepipeline",
            "service": "codepipeline",
			"description": "自動ビルド通知",
			"url": "https://outlook.office.com/webhook/dummydummy/IncomingWebhook/dummydummy2",
			"message": {
				"title": "自動ビルド通知"
			},
			"buttons": [{
				"name": "実行概要",
				"type": "OpenUri",
				"targets": [{
					"os": "default",
					"uri": "https://dummy-region.console.aws.amazon.com/codesuite/codepipeline/pipelines/<PIPELINE>/executions/<EXECUTION-ID>/timeline?region=<REGION>"
				}]
			}]
		}]
		}
	],
	"environments": [{
			"prefix": "deve-",
			"name": "【開発環境】",
			"themecolor": "#007800"
		},
		{
			"prefix": "testing-",
			"name": "【評価環境】",
			"themecolor": "#FF6600"
		},
		{
			"prefix": "prod-",
			"name": "【本番環境】",
			"themecolor": "#00008B"
		}
	]
}
```

## Create SNS Topic.
Subscription protocol LAMBDA -> **AWStoTeams**

## Create CloudWatch Event
1. codepipeline build success/failed -> SNS
1. aws batch execute failed. -> SNS

## Demo
1. when aws batch failed<br>
![batch failed](https://github.com/cereskou/awsnotification/blob/main/images/batch-failed.png)

2. codepipeline build success<br>
![cp success](https://github.com/cereskou/awsnotification/blob/main/images/cp-success.png)

3. codepipeline build failed<br>
![cp failed](https://github.com/cereskou/awsnotification/blob/main/images/cp-failed.png)

click button 実行概要 will open a browser to show detail informatio of the build job.<br>

4. view buiid details<br>
![cp detail](https://github.com/cereskou/awsnotification/blob/main/images/cp-detail.png)

