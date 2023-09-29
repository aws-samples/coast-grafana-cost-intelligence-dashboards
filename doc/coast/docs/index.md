# What is COAST?

COAST is an open-source Infrastructure-as-a-code based ([AWS CloudFormation](https://aws.amazon.com/cloudformation/)) one-click solution to simplify cloud financial management of your resources in AWS. COAST provides role-based customizable Grafana dashboards ensuring that every stakeholder receives es relevant, actionable insights tailored to their unique needs and responsibilities. For customers, already using Grafana for operational insights, COAST is natural extention, enabling the monitoring the monitoring of both operational and cost metrics within a single, preferred platform, reducing operational friction and prompting a unified view of essential KPIs. With COAST, users can gain full visibility and control over their cloud costs and maximize their AWS Investments.


## Key Features

- Persona based dashboards: Offer customizable and persona-based dashboards for Executives, Analysts, Developers, Cloud Operations, etc.
- Tag based Cost tracking: Allows users to monitor and optimized AWS spend at granular level. This is valuable for organizations with varied products, teams, and SaaS workloads, enabling precision monitoring of costs related to specific projects or departments.
- Unified Platform experience: For customers already utilizing Grafana for operational observability, COAST offers a unified platform experience, enabling the integration of operational and cost metrics within a familiar environment.
- Infrastructure-as-a-code: Single click repeatable, reliable and scalable deployment
- Open-source: Open-source nature of COAST will spur continued innovation and improvements ensuring the solution evolves in response to changing customer needs and market dynamics.
- Supports both Payer and member AWS accounts.


## Pre-requisites

- AWS Organization
- AWS IAM Identity Center

# Cloud Formation Template Deployment

## 1. Setup

The setup process will create the following resources, along with it's dependencies:

- Cost and Usage Report (AWS CUR) (optional)
- An Amazon Athena  database
- Glue/Lamda infrastructure to update Athena database with CUR
- An Amazon Managed Grafana workspace
- Grafana datasource for Amazon Athena
- Grafana dashboards

## 2. AWS Cost and Usage Report (CUR) report

COAST support both existing CUR report as well as creation of new CUR report. In the CloudFormation template you have the option to specify if you would like the template to create a new CUR report.  If you already have an existing CUR and would like to use it, you can provide the report name to use as well.  

#### Create New CUR Report
This option will create a new CUR and supporting infrastructure to hydrate the athena database with CUR data.  Note, it will take at least 24 hours for a new CUR to populate with data.  

- In CloudFormation, create a new stack from the cloudformation/coast-cfn.yaml
- Select 'true' in the CreateCURReport select form
- Enter a prefix path name for the new report in the CURReportPrefixName field
- Note, an entry in the CURReportName field is not needed for this option.  The template will generate a report name for you based on the CFN stack name
- Note, an S3 bucket name in CURDataBucketName is not needed for this option.  An S3 data bucket will be named and created based on the CFN stack name

#### Use Existing CUR Report
Use this setting only if you already have a CUR and you would like Grafana to use that existing CUR as the datasource.  This option will use your existing CUR and create supporting infrastructure to hydrate the Athena database with CUR data.

- Ensure your CUR bucket of your existing CUR report is located in the same region as this Cloud Formation stack ([see same region note in documentation](https://docs.aws.amazon.com/cur/latest/userguide/use-athena-cf.html)).  
- In CloudFormation, create a new stack from the cloudformation/coast-cfn.yaml
- Select 'false' in the CreateCURReport select form
- Enter the name of your existing CUR in the CURReportName field
- Enter the prefix path of the existing CUR in the CURReportPrefixName field
- Enter the S3 bucket of the existing CUR in the CURDataBucketName


#### Deployment using AWS CLI

```
aws cloudformation create-stack --stack-name curcoast --template-body file://coast-cfn.yaml --parameters ParameterKey=CURDataBucketName,ParameterValue=demo-cur-report ParameterKey=CURReportName,ParameterValue=demo-cur-report
ParameterKey=CreateCURReport,ParameterValue=false
ParameterKey=CURReportPrefixName,ParameterValue=grafana
ParameterKey=GrafanaDashboardTemplateURL,ParameterValue=https://raw.githubusercontent.com/pelgrim/observability-best-practices/main/sandbox/coast/grafana-dashboard.json  --capabilities CAPABILITY_NAMED_IAM
```

### 3. Post Installation Steps
- Grafana workspaces require an identity provider (IdP) to enable users to log in to the workspace.
  - We recommend AWS IAM Identity Center.  Follow instructions in the [Grafana User Guide](https://docs.aws.amazon.com/grafana/latest/userguide/AMG-manage-users-and-groups-AMG.html) to setup user access.
  - Login with the identity to the COAST Grafana workspace URL
  - The dashboard will be automatically imported under the General folder in the Dashboards menu
