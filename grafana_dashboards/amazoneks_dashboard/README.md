# AmazonEKS Split Cost Allocation Dashboard

<p align="center">
<img alt="GitHub commit activity" src="https://img.shields.io/github/commit-activity/m/aws-samples/coast-grafana-cost-intelligence-dashboards">

</p>

<p align="center">
<a href="../../README.md">Project Home</a> &nbsp;&bull;&nbsp;
<a href="#introduction">Introduction</a> &nbsp;&bull;&nbsp;
<a href="#installation">Installation</a> &nbsp;&bull;&nbsp;
<a href="#documentation">Documentation</a> &nbsp;&bull;&nbsp;
<a href="#license">License</a>
</p>

### Introduction
---

The AmazonEKS Split Cost Allocation Dashboard combines the split cost allocation data available within the cost and usage (CUR) report and marries the data with CloudWatch performance metrics to graph performance metrics over cost.  The Amazon EKS Split Cost Allocation Dashboard is crucial for engineers as it provides a unified view of cost and performance metrics, enabling them to make informed decisions, optimize resource usage, and ensure efficient cloud operations.


### Installation
---

##### Step 1.

The AmazonEKS Split Cost Allocation Dashboard requires that Split Cost Allocation Data (SCAD) for AmazonEKS is enabled within the Cost and Usage Report.  
If you have installed the CID Data Collection Lab (per the [README](../../README.md)), you will have the opportunity to enable SCAD.  See the [SCAD documentation](https://docs.aws.amazon.com/cur/latest/userguide/split-cost-allocation-data.html) for more information.  

##### Step 2.

The AmazonEKS Split Cost Dashboard requires AWS CloudWatch Container Inisights to be enabled for CloudWatch metrics.  This dashboard relies on the EKS metrics available in Container Insights.  A full list of available metrics is available in the [Amazon EKS and Kubernetes Container Insights Documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Container-Insights-metrics-EKS.html).

[Enable Container Insights for AmazonEKS](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Container-Insights-metrics-enhanced-EKS.html)

### Documentation
---

###### Importing Dashboards

You may now import dashboards available in the grafana_dashboards folder of this repository. Select a dashboard to import. This will populate the dashboard name and UID.

- Currently, all dashboards are based on CUR 2.0 and will require an Athena datasource connected to CUR
- Most dashboards will leverage CloudWatch as a metrics data source
- Some dashboards will require additional configuration (i.e., CloudWatch Container Insights for the AmazonEKS dashboards). These installation steps will be specified in the README file within the dashboard folder

After importing the dashboards, follow the dashboard's readme file to set your variables and understand the data visualizations.

######  Associated Cost
> [!IMPORTANT]
> Costs are associated with operating this dashboard.  Costs depend on usage and the size of your datasets, be sure to get a complete understanding of costs before deployment. 

CloudWatch Container Insights is a paid feature of CloudWatch.  Please see the pricing page for cost calculations.

[CloudWatch Pricing](https://aws.amazon.com/cloudwatch/pricing/)

Additional costs have been outlined in the [project README](../../README.md)

### License
---
This library is licensed under the MIT-0 License. See the [LICENSE](https://github.com/aws-samples/COAST/blob/main/LICENSE) file.
