# Autoscaling Dashboard

![Autoscaling Dashboard](../../images/coast_banner.png)

## About

The COAST Autoscaling Dashboard provides a comprehensive overview of cost and infrastructure performance metrics related to EC2 Autoscaling Groups, Load Balancers, NAT Gateways, and networking across all associated resources. The data sources feeding into this dashboard are CloudWatch and the Cost and Usage Report (CUR).

To effectively utilize this dashboard, it requires the input of your autoscaling group's name. All resources within your autoscaling workload must be tagged with a unique tag key/value pair specific to the workload. It is imperative that these tags are enabled as Cost Allocation Tags and persist within the Cost and Usage Report (CUR).

## Limitations 

Currently the COAST Autoscaling Dashboard is in a proof of concept phase.  It has been certified for workloads which are operating in the same account as the COAST deployment.  Workloads from multiple regions in the account may be observed.  We are working on a mechanism to make the COAST Autoscaling Dashboard operational across accounts.

## Troubleshooting

#####Autoscaling Metrics
If you are not seeing Autoscaling Metrics in Cloudwatch, make sure to enable 'Auto Scaling group metrics collection' under the monitoring tab in your Autoscaling Group console.


