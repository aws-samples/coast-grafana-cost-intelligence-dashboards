### EC2 Dashboard 

###### Summary

This dashboard will display details of your account EC2 aggregated cost, usage and performance metrics.  In addition, if you provide an EC2 instance ID the dashboard will display detail cost, usage and performance visuals for that instance.

###### Pre-requisites

- This dashboard is based on the CUR 2.0 format.  CUR 2.0 may be deployed in payer management accounts and/or linked (standalone) accounts ([member CUR](https://aws.amazon.com/about-aws/whats-new/2020/12/cost-and-usage-report-now-available-to-member-linked-accounts/)).
    -  You may your our data export CloudFormation template to provision the CUR 2.0 Data Export.
    - The CUR 2.0 data sources needs to be named 'COAST-CUR20-2024-07-15'
- This dashbaord also depends on a CloudWatch data source.  You may use our data source CloudFormation template to provision both the CUR 2.0 and CloudWatch data source. 
- Our templates currently support a us-east-1 deployment only.

If you are using this dashboard at the payer mangerment account, you will need to provision CloudWatch Cross-account observabiility.  This is necessary for the CloudWatch performance metrices to be accessible from the dashboard.  See the [Grafana](https://docs.aws.amazon.com/grafana/latest/userguide/cloudwatch-cross-account.html) and [CloudWatch](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html) instructions on how to enable Cross-account features for CloudWatch. 