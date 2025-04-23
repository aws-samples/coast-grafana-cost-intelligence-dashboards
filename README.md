# COAST - Cost Optimization and Saving Toolkit for Amazon Managed Grafana

## About

COAST is an open-source collection of Grafana dashboards that provide the capability to combine and observe AWS resource performance metrics with AWS cost and usage report (CUR) data. These dashboards assist customers in promoting financial accountability, optimizing costs, tracking usage goals, implementing governance best practices, and achieving operational excellence across all Well-Architected pillars. Utilizing Amazon Managed Grafana allows us to use an open-source platform that is very popular with the engineering community.

###### Advantages of COAST

- If you are already using Grafana for monitoring application metrics, you will be familiar with the tool interface.
- COAST dashboards may integrate with your existing Grafana deployments.
- COAST dashboards support filtering on multiple dimensions, including account and region.

## Suggested Deployment and Configuration

##### Step 1.

As a best practice, we recommend the deployment of the dashboards in a central management account (data collection account) which has the capability to aggregate Cost and Usage data (CUR) in addition to cross-account metrics to observe cost and performance with Grafana. To deploy the dashboards, we recommend utilizing the data collection engine of the Cloud Intelligence Dashboards which has built an architecture and deployment for central CUR aggregation utilized for the [CID Dashboards](https://catalog.workshops.aws/awscid/en-US).  

COAST is deployed via CloudFormation, which allows for the provisioning of the necessary [AWS Data Exports](https://aws.amazon.com/aws-cost-management/aws-data-exports/) and [Amazon Managed Grafana](https://aws.amazon.com/grafana/) resources. The dashboards may be imported on EC2-based Grafana deployments; however, the configuration of data sources and plugins would be manual. 

##### Step 2.

Most dashboards will have a requirement of CloudWatch metrics to visualize performance metrics. As a best practice, we recommend the configuration of CloudWatch cross-account observability into the central management account (data collection account).

[Configure CloudWatch cross-account observability](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html)

##### Step 3. 

Deploy COAST within the central management account (data collection account). The COAST CloudFormation template builds the Amazon Grafana Workspace, installs the Athena plugin, and configures the Athena plugin with the deployed CUR and CloudWatch.

[Link to CFN]()


### Post Installation Steps
- Grafana workspaces require an identity provider (IdP) to enable users to log in to the workspace.
- We recommend AWS IAM Identity Center (the CloudFormation template creates the Grafana workspace with AWS IAM Identity Center enabled). 
- Add at least one Admin user under the Authentication tab in the Grafana Workspace console.  For additional instructions, see the [Grafana User Guide](https://docs.aws.amazon.com/grafana/latest/userguide/AMG-manage-users-and-groups-AMG.html) to setup user access.
- Login to the COAST Grafana workspace URL using the identity configured above.
- Import the dashboards you require as per the instructions below

  #### Importing Dashboards

You may now import dashboards available in the grafana_dashboards folder of this repository. Select a dashboard to import. This will populate the dashboard name and UID.

- Currently, all dashboards are based on CUR 2.0 and will require an Athena datasource connected to CUR
- Most dashboards will leverage CloudWatch as a metrics data source
- Some dashboards will require additional configuration (i.e., CloudWatch Container Insights for the AmazonEKS dashboards). These installation steps will be specified in the README file within the dashboard folder

After importing the dashboards, follow the dashboard's readme file to set your variables and understand the data visualizations.
  
  ##### Amazon EKS Split Cost Dashboard

  description to be added...
  
  ##### Auto Scaling Dashboard

  Designed to give a view into your Auto Scaling workloads by Auto Scaling group and tag/key values. This dashboard is based on the CUR Legacy format and requires a CUR Legacy data source as well as a CloudWatch data source. See the [Readme](https://github.com/aws-samples/coast-grafana-cost-intelligence-dashboards/tree/main/grafana_dashboards/auto_scaling) file for installation instructions.

## Associated Cost
COAST is an open-source solution and is completely free to use. However, you will be responsible for any AWS costs associated with underlying AWS services. Costs to consider:

[Amazon Managed Grafana Pricing](https://aws.amazon.com/grafana/pricing/)

[CloudWatch Grafana Plugin Cost](https://grafana.com/docs/grafana/latest/datasources/aws-cloudwatch/#control-pricing)

## Support
COAST is supported by Technical Account Managers and Solution Architects of AWS on a best-effort basis. However, users are encouraged to ask questions, open issues, contribute, and provide feedback.

## Contributing
The core team for COAST include the following, in alphabetical order:

- Chris Strzelczyk
- Lucas Vieira
- Munish Dabra
- Siva Guruvareddiar

However, we welcome the wider open-source community to this project. See [CONTRIBUTING](https://github.com/aws-samples/COAST/blob/main/CONTRIBUTING.md) for more information.

## License
This library is licensed under the MIT-0 License. See the [LICENSE](https://github.com/aws-samples/COAST/blob/main/LICENSE) file.
