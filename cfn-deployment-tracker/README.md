# COAST - CloudFormation Deployment Tracker !

This is a one time setup for the COAST admins.

This is a project with IaC using CDK with Go.

## Pre-req
* Ensure GoLang is installed in your machine (preferred v1.19+)
* Uses "cdk bootstrap" and "cdk deploy" 


The `cdk.json` file tells the CDK toolkit how to execute your app.

## Useful commands

 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template
 * `go test`         run unit tests

## Bootstrap
```sh
cdk bootstrap
```

## Deploy
```sh
cdk deploy
```

## Teardown
```sh
cdk destroy
```


## Typical output for /coast-cfn-tracker will be like below:

✨  Synthesis time: 5.48s

CfnTrackerCoastStack:  start: Building 68c832c3c168615858c47326a89ac057d38c83ef274e7774d065cddcf4fdf2cc:current_account-current_region
CfnTrackerCoastStack:  success: Built 68c832c3c168615858c47326a89ac057d38c83ef274e7774d065cddcf4fdf2cc:current_account-current_region
CfnTrackerCoastStack:  start: Building 42004ceebf9afc8fdf0e9294d1fdabf7843052038edeb3ee74a48e4332f46bf7:current_account-current_region
CfnTrackerCoastStack:  success: Built 42004ceebf9afc8fdf0e9294d1fdabf7843052038edeb3ee74a48e4332f46bf7:current_account-current_region
CfnTrackerCoastStack:  start: Publishing 68c832c3c168615858c47326a89ac057d38c83ef274e7774d065cddcf4fdf2cc:current_account-current_region
CfnTrackerCoastStack:  start: Publishing 42004ceebf9afc8fdf0e9294d1fdabf7843052038edeb3ee74a48e4332f46bf7:current_account-current_region
CfnTrackerCoastStack:  success: Published 68c832c3c168615858c47326a89ac057d38c83ef274e7774d065cddcf4fdf2cc:current_account-current_region
CfnTrackerCoastStack:  success: Published 42004ceebf9afc8fdf0e9294d1fdabf7843052038edeb3ee74a48e4332f46bf7:current_account-current_region
This deployment will make potentially sensitive changes according to your current security approval level (--require-approval broadening).
Please confirm you intend to make the following modifications:
[...]
Do you wish to deploy these changes (y/n)? y
CfnTrackerCoastStack: deploying... [1/1]
CfnTrackerCoastStack: creating CloudFormation changeset...

 ✅  CfnTrackerCoastStack

✨  Deployment time: 72.03s

Outputs:
CfnTrackerCoastStack.cfnTrackerdEndpointE70C0F52 = https://76fbm2fggd.execute-api.us-west-1.amazonaws.com/prod/
Stack ARN:
arn:aws:cloudformation:us-west-1:099019764096:stack/CfnTrackerCoastStack/c26f3ec0-bf0b-11ee-a762-021d1dc5459b

✨  Total time: 77.51s
## Typical APIs will be like below:

```sh
Endpoint [POST] = https://76fbm2fggd.execute-api.us-west-1.amazonaws.com/prod/coast-cfn-tracker
```

## Sample Curls
```sh
     curl -X POST  -H 'Content-Type: application/json' -d '{"accountId": "217392306962", "awsRegion": "us-east-1"}' https://76fbm2fggd.execute-api.us-west-1.amazonaws.com/prod/coast-cfn-tracker
```

## How to use this in COAST
Replace the above Endpoint value in cloudformation/coast-cfn.yaml for the below attribute

```sh
Resources.SendReportingStatsFunction.Properties.Environment.Variables.REPORTING_URL
```

### Sample
```sh
SendReportingStatsFunction:
    Type: AWS::Lambda::Function
    Condition: isReportingEnabled
    Properties:
      Environment:
        Variables:
          REPORTING_URL: https://ydys3fee73.execute-api.us-west-1.amazonaws.com/prod/coast-cfn-tracker
     
```