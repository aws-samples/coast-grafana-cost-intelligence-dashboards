# COAST - Cost Optimization and Saving Toolkit for Amazon Managed Grafana

## Pre-requisites

- AWS IAM Identity Center - [Amazon Managed Grafana requires authentication](https://docs.aws.amazon.com/grafana/latest/userguide/authentication-in-AMG.html).  Our CloudFormation template configures the Grafana workspace with [AWS SSO](https://docs.aws.amazon.com/singlesignon/latest/userguide/getting-started.html).

## Suggested Configuration
The COAST CloudFormation deployment template supports both deploying with an existing [[Cost and Usage Report (CUR)]](https://docs.aws.amazon.com/cur/latest/userguide/what-is-cur.html) or creating a new one if none exists. For immediate utilization of the Coast dashboard, it is recommended to have an already enabled Cost and Usage Report (CUR). If CUR is enabled after COAST, dashboards may not display data for approximately 24 hours, and historical data will be unavailable unless a backfill is requested from AWS. 

# Cloud Formation Template Deployment

## 1. Setup

The setup process will create the following resources, along with their dependencies:

- A new AWS Cost and Usage Report (AWS CUR)
  - If an existing report is provided, the system will use the provided one.
- Glue and Lambda infrastructure to update the Athena database
- An Amazon Athena database for the Cost and Usage Report
- Least Privilege Access Roles to allow the Cost and Usage Report to update Athena
- An Amazon Managed Grafana workspace
- A Grafana Athena Data Source
- Importation of the Grafana FinOps Dashboard

## 2. CUR report
Within the CloudFormation template, you can choose whether to create a new Cost and Usage Report (CUR) or utilize an existing one by providing its name. The template will establish the necessary infrastructure to update the Grafana datasource (Athena) with CUR data. For additional information on the CUR/Athena integration, refer to the documentatio [here](https://docs.aws.amazon.com/cur/latest/userguide/use-athena-cf.html).  

#### Create New CUR
Opting for this choice will initiate the creation of a new Cost and Usage Report (CUR) along with the necessary infrastructure to populate the Athena database with CUR data. Please be aware that it will take a minimum of 24 hours for the new CUR to be populated with data.  Historical data will be unavailable unless a backfill is requested from AWS Support.

To implement this in CloudFormation:

- Launch a new stack using the cloudformation/coast-cfn.yaml template.
- Do not modify the CurReportName parameter in the template
- The template will generate a report name based on the CloudFormation (CFN) stack name.
- Similarly, an S3 data bucket will be named and created based on the CFN stack name.

#### Use Existing CUR
Opt for this configuration only if you already possess a Cost and Usage Report (CUR) and wish to have Grafana utilize that existing CUR as the datasource. This choice will leverage your current CUR while establishing the necessary infrastructure to populate the Athena database with the CUR data.

To implement this in CloudFormation:

- Ensure the CUR bucket of your existing CUR report is located in the same region as this CloudFormation stack (refer to the same region note in the documentation).
- Launch a new stack using the cloudformation/coast-cfn.yaml template.
- Enter the name of your existing CUR in the CurReportName field.
- We will determine the S3 bucket used with the existing CUR report and use this bucket while establishing the necessary infrastructure.


### 4. Post Installation Steps
- Grafana workspaces require an identity provider (IdP) to enable users to log in to the workspace.
  - We recommend AWS IAM Identity Center.  Follow instructions in the [Grafana User Guide](https://docs.aws.amazon.com/grafana/latest/userguide/AMG-manage-users-and-groups-AMG.html) to setup user access.
  - Login with the identity to the COAST Grafana workspace URL
  - The dashboard will be automatically imported under the General folder in the Dashboards menu
