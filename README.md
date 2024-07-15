# COAST - Cost Optimization and Saving Toolkit for Amazon Managed Grafana

![Executive Dashboard](images/coast_header.png )

## About

COAST is an open-source collection of dashboards that provide the capability to combine and observe AWS resource performance metrics with cost and usage data metrics.  These dashboards assist customers in promoting financial accountability, optimizing costs, tracking usage goals, implementing governance best practices, and achieving operational excellence across all Well-Architected pillars.  Utilizing Amazon Managed Grafana allows us to utilize an open-source platform which is very popular with the engineering community.

COAST is deployed via CloudFormation infrastructure as code templates, which allow for the provisioning of the necessary [AWS Data Exports](https://aws.amazon.com/aws-cost-management/aws-data-exports/) and [Amazon Managed Grafana](https://aws.amazon.com/grafana/) resources.  



###### Advantages of COAST
- If you are already using Grafana for monitoring application metrics, you will be familiar with the tool inferface.
- COAST will integrate with your existing Grafana dashboards.
- COAST has support for filtering by AWS account, service and AWS tags.
- COAST may be installed at the AWS Management Account (payer) or in a single linked account with [member CUR](https://aws.amazon.com/about-aws/whats-new/2020/12/cost-and-usage-report-now-available-to-member-linked-accounts/).

## Pre-requisites

- AWS IAM Identity Center - [Amazon Managed Grafana requires authentication](https://docs.aws.amazon.com/grafana/latest/userguide/authentication-in-AMG.html).  Our CloudFormation template configures the Grafana workspace with [AWS SSO](https://docs.aws.amazon.com/singlesignon/latest/userguide/getting-started.html).

- AWS Data Exports - Currently COAST works with CUR Legacy.  However, our templates give you the ability to deploy [CUR 2.0](https://docs.aws.amazon.com/cur/latest/userguide/table-dictionary-cur2.html) and [FOCUS (preview)](https://docs.aws.amazon.com/cur/latest/userguide/table-dictionary-focus-1-0-aws.html)

## Suggested Configuration
The COAST Data Export template enables you to create a CUR Legacy, [CUR 2.0](https://docs.aws.amazon.com/cur/latest/userguide/table-dictionary-cur2.html), and [FOCUS](https://docs.aws.amazon.com/cur/latest/userguide/table-dictionary-focus-1-0-aws.html) (preview) reports. Our dashboards will require one of these Data Exports as a prerequisite. We recommend using our template to deploy these exports, as the template will also create the necessary [AWS Glue](https://aws.amazon.com/glue/) and [AWS Athena](https://aws.amazon.com/athena/) resources required for the Grafana dashboards to use as data sources. To backfill the cost and usage data in each report, you can open a support case. For more information on Data Exports backfill, refer to the [Data Exports User Guide](https://docs.aws.amazon.com/cur/latest/userguide/troubleshooting.html#backfill-data).

## Setup Overview - Deploy with CloudFormation

#### Data Exports Deployment
The [provision-grafana-workspaces.yaml](https://github.com/aws-samples/coast-grafana-cost-intelligence-dashboards/blob/main/CloudFormation/provision-data-exports.yaml) allows users to provision Data Exports for a single account or an entire management (payer) account.  Use this template to create CUR Legacy, CUR 2.0 and Focus (preview) reports.  We recommend running three seperate CloudFormation stacks and provisioning all three reports.  The template will also create a Glue Database, Glue Table and Glue Crawler to allow Athena to query the data.  Athena is utilized by the Grafana plugin to make queries.  

Once the reports are provisioned you may open a support case to backfill data.  See [Data Exports User Guide](https://docs.aws.amazon.com/cur/latest/userguide/troubleshooting.html#backfill-data).


1. Create a new stack in CloudFormation 
2. Select 'Upload a template file' and provide the location of the template.  [provision-data-exports.yaml](https://github.com/aws-samples/coast-grafana-cost-intelligence-dashboards/blob/main/CloudFormation/provision-data-exports.yaml)
3. Select the Data Export you need to configure with the drop down menu (either CUR legacy, CUR 2.0 or FOCUS (preview))
4. Launch the creation

Once the creation is finished, the Outputs tab will contain important informaiton which will be utilized by the next templates.


#### Grafana Workspace Deployment
The [provision-grafana-workspaces.yaml](https://github.com/aws-samples/coast-grafana-cost-intelligence-dashboards/blob/main/cloudformation/provision-grafana-workspace.yaml) template will build a new Amazon Managed Grafana workspace along with the necessary permissions.  You will still have to enable authentication by utilizing [AWS Identify Center](https://aws.amazon.com/iam/identity-center/) or you a third party IDP.  

To implement this in CloudFormation:
  1. Create new stack in Cloudformation
  2. Select 'upload template file' and select the [provision-grafana-workspaces.yaml](https://github.com/aws-samples/coast-grafana-cost-intelligence-dashboards/blob/main/cloudformation/provision-grafana-workspace.yaml) template.
  3. You will have to select the name for the cheery before you can leave.

Once the creation is finished, the Outputs tab will contain important information about the workspace which will be utilized by the next templates.

### Grafana Data Sources Deployment
The [provision-grafana-data-source.yaml](https://github.com/aws-samples/coast-grafana-cost-intelligence-dashboards/blob/main/cloudformation/provision-grafana-data-source.yaml) template installs the necessary data sources into the Grafana workspace you have created.  We recommend running seperate CloudFormation stacks and provisioning these data sources.  _Note, currently our dashboards only utilize the CUR Legacy and CloudWatch datasource_. 

1. Create new stack in Cloudformation
2. Select 'upload template file' and select the [provision-grafana-data-source.yaml](https://github.com/aws-samples/coast-grafana-cost-intelligence-dashboards/blob/main/cloudformation/provision-grafana-data-source.yaml) template.
3. The template requires ceratin parameters in order to succeed:

    - **AthenaDatabaseName** - You may obtain this from the Outputs tab of the Data Exports CloudFormation deployment.
    - **AthenaWorkgroupName** - You may obtain this from the Outputs tab of the Data Exports CloudFormation deployment.
    - **DataExportName** - You may obtain this from the Outputs tab of the Data Exports CloudFormation deployment.
    - **DataExportType** - Select the export type (either CUR Legacy, CUR 2.0 or Focuse)
    - **GrafanaWorkspaceName** - You may obtain this from the Outputs tab of the Provision Grafana workspace CloudFormation deployment.

### Post Installation Steps
- Grafana workspaces require an identity provider (IdP) to enable users to log in to the workspace.
- We recommend AWS IAM Identity Center (the CloudFormation template creates the Grafana workspace with AWS IAM Identity Center enabled).  
- Add at least one Admin user under the Authentication tab in the Grafana Workspace console.  For additional instructions, see the [Grafana User Guide](https://docs.aws.amazon.com/grafana/latest/userguide/AMG-manage-users-and-groups-AMG.html) to setup user access.
- Login to the COAST Grafana workspace URL using the identity configured above.
- Import the dashboards you require as per the instructions below

  #### Importing Dashboards

  You may now import dashboards available in the grafana_dashboards folder of this repository.  Select a dashboard to import.  This will populate the dashboard name, and UID.  All dashboards will require an Amazon Athena datasource for CUR named "COAST-2023-09-19" and some may also require a CloudWatch datasource named "Cloudwatch".   After importing dashboards, follow the dashboards readme file to set your variables understand the data visualizations. 
  
  ##### FinOps Dashboard

  This dashboard is recommended for FinOps practitioners. It displays a billing overview, a summary of spend trends, region activity, and visuals for purchase types. To utilize this dashboard, your Cost and Usage Report (CUR) data should include reserved instance and savings plans columns. Additionally, you need to have at least one cost allocation tag enabled.
  
  ##### Executive Dashboard

  Designed for company leadership, this dashboard offers a view of billing and overall spend trends. To use this dashboard, your CUR data should include reserved instance and savings plans columns. Additionally, at least one cost allocation tag must be enabled.
  
  ##### Engineering Dashboard

  This dashboard is recommended for engineering teams. Teams can filter based on cost allocation tags and display only the services relevant to them. To utilize this dashboard, your CUR data should include reserved instance and savings plans columns. Additionally, you need to have at least one cost allocation tag enabled.

## Associated Cost
COAST is an open-source solution and is completely free to use. However, you will be responsible for any AWS costs associated with underlying AWS services.

## Support
COAST is supported by Solution Architects of AWS on best effort basis. However, users are encouraged to ask questions, open issues, contribute and provide feedback.

## Contributing
The core team for COAST include the following, in alphabetical order:

- Chris Strzelczyk
- Lucas Vieira
- Munish Dabra
- Siva Guruvareddiar

However, we welcome the wider open-source community to this project. See [CONTRIBUTING](https://github.com/aws-samples/COAST/blob/main/CONTRIBUTING.md) for more information.

## License
This library is licensed under the MIT-0 License. See the [LICENSE](https://github.com/aws-samples/COAST/blob/main/LICENSE) file.
