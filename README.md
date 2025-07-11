<p align="center">
<h1 align="center">COAST</h2>
<h5 align="center">Cost Dashboards for Engineers</h5>
</p>

<p align="center">
<img alt="GitHub commit activity" src="https://img.shields.io/github/commit-activity/m/aws-samples/coast-grafana-cost-intelligence-dashboards">

</p>

<p align="center">
<a href="#introduction">Introduction</a> &nbsp;&bull;&nbsp;
<a href="#installation">Installation</a> &nbsp;&bull;&nbsp;
<a href="#documentation">Documentation</a> &nbsp;&bull;&nbsp;
<a href="#dashboards">Dashboards</a> &nbsp;&bull;&nbsp;
<a href="#issue">Issue?</a>
</p>

### Introduction
---

COAST is an open-source collection of Grafana dashboards that provide the capability to combine and observe AWS resource performance metrics with AWS cost and usage report (CUR) data. These dashboards assist customers in promoting financial accountability, optimizing costs, tracking usage goals, implementing governance best practices, and achieving operational excellence across all Well-Architected pillars. Utilizing Amazon Managed Grafana allows us to use an open-source platform that is very popular with the engineering community.

###### Advantages of COAST

- If you are already using Grafana for monitoring application metrics, you will be familiar with the tool interface.
- COAST dashboards may integrate with your existing Grafana deployments.
- COAST dashboards support filtering on multiple dimensions, including account and region.

###### Amazon Managed Grafana

COAST is deployed via CloudFormation, which allows for the provisioning of [Amazon Managed Grafana](https://aws.amazon.com/grafana/) and supporting resources and policies. The dashboards themselves are json based files which may be imported on EC2-based Grafana deployments; however, the configuration of data sources, plugins and supporting roles and policies would be manual. 

### Installation
---

##### Step 1.

We recommend deploying dashboards in a dedicated monitoring account to aggregate Cost and Usage Report (CUR) data and CloudWatch cross-account metrics. This separation from the payer management account enhances security by consolidating observability functions in a tightly controlled environment. To deploy the COAST Grafana dashboard infrastructure, we recommend utilizing the data collection engine of the **Cloud Intelligence Dashboards** which has built an architecture and deployment for central CUR aggregation utilized for the CID Dashboards. 

The CID Data Collection Lab provides CloudFormation templates to copy CUR 2.0 data from your Management Account to a dedicated one. You can use it to aggregate data from multiple Management (Payer) Accounts or multiple Linked Accounts.

[Deploy CID Data Collection](https://catalog.workshops.aws/awscid/en-US/dashboards/foundational/cudos-cid-kpi/deploy)

##### Step 2.

Most dashboards will have a requirement of CloudWatch metrics to visualize performance metrics. As a best practice, we recommend the configuration of **CloudWatch cross-account observability** into the central monitoring account (data collection account).  Each account (source account) that you wish to collect CloudWatch metrics from will need to configure each region to send metrics to the monitoring account (data collection account).

In the data collection account, in the CloudWatch settings, Select *Configure* under *Monitoring account configuration*.  Select *Logs* and *Metrics* as the data selection, and fill in a comma seperated list of accounts under *List source accounts* (note there is a process for collecting metrics from the entire Organization, see the CloudWatch instructions in the link above).  You will need to copy the *Monitoring accounts sink ARN* of the data collection account in each region.  This information is available under the *Configuration details* of the CloudWatch *Monitoring account configuration*.

In each source account in each region, in the CloudWatch settions, select *Source account configuration*.  Select *Logs* and *Metrics* as the data.  Paste in the *Sink ARN* from the data collection account (from the same region).  Then select *Account name* as the identifier.  

For more in-depth instructions, see [Configure CloudWatch cross-account observability](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html). 

##### Step 3. 

Deploy COAST within the central monitoring account (data collection account). The COAST CloudFormation template builds the Amazon Grafana Workspace, installs the Athena plugin, and configures the Athena plugin with the deployed CUR and CloudWatch.

[Provision COAST](cloudformation/provision-coast-services.yaml)

### Documentation
---

###### Post Installation Steps
- Grafana workspaces require an identity provider (IdP) to enable users to log in to the workspace.
- We recommend AWS IAM Identity Center (the CloudFormation template creates the Grafana workspace with AWS IAM Identity Center enabled). 
- Add at least one Admin user under the Authentication tab in the Grafana Workspace console.  For additional instructions, see the [Grafana User Guide](https://docs.aws.amazon.com/grafana/latest/userguide/AMG-manage-users-and-groups-AMG.html) to setup user access.
- Login to the COAST Grafana workspace URL using the identity configured above.
- Import the dashboards you require as per the instructions below

###### Importing Dashboards

You may now import dashboards available in the grafana_dashboards folder of this repository. Select a dashboard to import. This will populate the dashboard name and UID.

- Currently, all dashboards are based on CUR 2.0 and will require an Athena datasource connected to CUR
- Most dashboards will leverage CloudWatch as a metrics data source
- Some dashboards will require additional configuration (i.e., CloudWatch Container Insights for the AmazonEKS dashboards). These installation steps will be specified in the README file within the dashboard folder

After importing the dashboards, follow the dashboard's readme file to set your variables and understand the data visualizations.

######  Associated Cost

> [!IMPORTANT]
> Costs are associated with operating this dashboard.  Costs depend on usage and the size of your datasets, be sure to get a complete understanding of costs before deployment. 

COAST is an open-source solution and is completely free to use. However, you will be responsible for any AWS costs associated with underlying AWS services. Costs to consider:

[Amazon Managed Grafana Pricing](https://aws.amazon.com/grafana/pricing/)

[CloudWatch Grafana Plugin Cost](https://grafana.com/docs/grafana/latest/datasources/aws-cloudwatch/#control-pricing)

Each dashboard may have additional cost implications. Details are provided in the README file associated with each dashboard.


### Dashboards
---
  
###### Amazon EC2 Dashboard

  The EC2 Instance Dashboard displays EC2 instance compute cost, usage and performance metric information filtered by account and region. One section also filters by tag. The filter panel will refresh based on selections of previous filters. For example, when an account is selected the Region menu will only show regions, instances and tags observed in that account for the time period selected.

  [AmazonEC2 Dashboard](grafana_dashboards/ec2_dashboard/README.md)

<img src="images/amazonec2_dashboard.png">
<br>
<img src="images/amazonec2_dashboard_instance_explorer.png">

###### Amazon EKS Split Cost Dashboard

  The AmazonEKS Split Cost Allocation Dashboard combines the split cost allocation data available within the cost and usage (CUR) report and marries the data with CloudWatch performance metrics to graph performance metrics over cost.  The Amazon EKS Split Cost Allocation Dashboard is crucial for engineers as it provides a unified view of cost and performance metrics, enabling them to make informed decisions, optimize resource usage, and ensure efficient cloud operations.

  [AmazonEKS Split Cost Dashboard](grafana_dashboards/amazoneks_dashboard/README.md)

  <img src="images/amazoneks_dashboard.jpg">
  
###### Auto Scaling Dashboard

  Designed to give a view into your Auto Scaling workloads by Auto Scaling group and tag/key values. This dashboard is based on the CUR Legacy format and requires a CUR Legacy data source as well as a CloudWatch data source. See the [Readme](https://github.com/aws-samples/coast-grafana-cost-intelligence-dashboards/tree/main/grafana_dashboards/auto_scaling) file for installation instructions.


### Issues
---
COAST is supported by Technical Account Managers and Solution Architects of AWS on a best-effort basis. However, users are encouraged to open Github issues, ask questions, contribute, and provide feedback.

### Contribution
---
The core team for COAST include the following, in alphabetical order:

- Chris Strzelczyk
- Lucas Vieira
- Munish Dabra
- Siva Guruvareddiar

However, we welcome the wider open-source community to this project. See [CONTRIBUTING](https://github.com/aws-samples/COAST/blob/main/CONTRIBUTING.md) for more information.

### License
---
This library is licensed under the MIT-0 License. See the [LICENSE](https://github.com/aws-samples/COAST/blob/main/LICENSE) file.
