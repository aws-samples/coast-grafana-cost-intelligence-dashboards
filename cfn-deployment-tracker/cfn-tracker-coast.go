package main

import (
	"math/rand"
	"strconv"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CfnTrackerCoastStackProps struct {
	awscdk.StackProps
}

func NewCfnTrackerCoastStack(scope constructs.Construct, id string, props *CfnTrackerCoastStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	////IAM role
	lambda_role := awsiam.NewRole(stack, jsii.String("lambda-role"), &awsiam.RoleProps{
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("lambda.amazonaws.com"), nil),
	})
	lambda_role.AddToPolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Resources: &[]*string{
			jsii.String("*"),
		},
		Actions: &[]*string{
			jsii.String("lambda:InvokeFunction"),
		},
	}))
	lambda_role.AddManagedPolicy(awsiam.ManagedPolicy_FromManagedPolicyArn(stack, jsii.String("lambda"), jsii.String("arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole")))
	lambda_role.AddManagedPolicy(awsiam.ManagedPolicy_FromManagedPolicyArn(stack, jsii.String("s3"), jsii.String("arn:aws:iam::aws:policy/AmazonS3FullAccess")))
	lambda_role.AddManagedPolicy(awsiam.ManagedPolicy_FromManagedPolicyArn(stack, jsii.String("vpclambda"), jsii.String("arn:aws:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole")))

	//Create s3 bucket for cfn tracker
	s3bucket := awss3.NewBucket(stack, jsii.String("coast_cfntracker_"+strconv.Itoa(rand.Int())), nil)

	////cfn tracker flow
	//cfn tracker- Lambda
	csv_download_lambdaHandler := awslambda.NewFunction(stack, jsii.String("cfnTrackerHandler"), &awslambda.FunctionProps{
		Code:         awslambda.Code_FromAsset(jsii.String("lambda"), nil),
		Runtime:      awslambda.Runtime_PYTHON_3_8(),
		Handler:      jsii.String("coast_cfn_tracker.lambda_handler"),
		Role:         lambda_role,
		FunctionName: jsii.String("coast-cfn-tracker"),
		Timeout:      awscdk.Duration_Minutes(jsii.Number(1)),
	})
	csv_download_lambdaHandler.AddEnvironment(jsii.String("s3_bucket"), s3bucket.BucketName(), nil)

	//cfn tracker- API Gateway
	awsapigateway.NewLambdaRestApi(stack, jsii.String("cfnTrackerdEndpoint"), &awsapigateway.LambdaRestApiProps{
		Handler: csv_download_lambdaHandler,
		Proxy:   jsii.Bool(false),
	}).Root().AddResource(jsii.String("coast-cfn-tracker"), &awsapigateway.ResourceOptions{}).AddMethod(jsii.String("POST"), nil, &awsapigateway.MethodOptions{})

	// The code that defines your stack goes here

	// example resource
	// queue := awssqs.NewQueue(stack, jsii.String("CfnTrackerCoastQueue"), &awssqs.QueueProps{
	// 	VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
	// })

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewCfnTrackerCoastStack(app, "CfnTrackerCoastStack", &CfnTrackerCoastStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
