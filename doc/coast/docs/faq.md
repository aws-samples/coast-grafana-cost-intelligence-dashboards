# FAQ
- [What are the other potential options customer have currently?](#What are the other potential options customer have currently?)
- [Are there specific industry segments or personas that prefer COAST?](#Are there specific industry segments or personas that prefer COAST?)
- [Is COAST an open-source solution?](#Is COAST an open-source solution?)
- [How does COAST help customers optimize their cloud costs?](#How does COAST help customers optimize their cloud costs?)
- [Can COAST be integrated with other AWS services?](#Can COAST be integrated with other AWS services?)
- [How COAST is different from Client Intelligence dashboards(CID)?](#How COAST is different from Client Intelligence dashboards(CID)?)
- [How will customer deploy the solution?](#How will customer deploy the solution?)
- [ Does COAST support integrating with existing SSO or SAML based Identity Providers?](# Does COAST support integrating with existing SSO or SAML based Identity Providers?)
- [What is MVP scope for Coast?](#What is MVP scope for Coast?) 



## What are the other potential options customer have currently?
 Cost Explorer, Client Intelligence Dashboard (CID), 3rd party partner products such as Apptio, CloudZero etc.

## Are there specific industry segments or personas that prefer COAST?
 DNB, ISV or SaaS customer, Customers with Hybrid/Multi-cloud Strategy.

## Is COAST an open-source solution?
 Yes, COAST is an open-source solution that customers can use and customize according to their needs.

## How does COAST help customers optimize their cloud costs?
 COAST provides customers with customizable dashboards, cost alerts, cost recommendations that help them to visualize and analyze their cloud costs, enabling them to make informed decisions to optimize their spend.

## Can COAST be integrated with other AWS services?
 Yes, COAST can be integrated with other AWS services to provide customers with a comprehensive cost optimization solution.

## How COAST is different from Client Intelligence dashboards(CID)?
COAST leverages the same underlying CUR data as CID and generates information that is very comparable to the latter. However, the key distinction lies in the intended audience. COAST aims at smaller DevOps teams who are already using open-source Grafana for application monitoring. With COAST, users can effortlessly import a dashboard that integrates with Amazon Managed Grafana to deliver cost and usage visibility alongside application observability.

## How will customer deploy the solution?
Customers will be able to download COAST as a IaC template from Github as one-click deployment.

## Does COAST support integrating with existing SSO or SAML based Identity Providers?
Amazon Managed Grafana provides integration with the customers IDP or AWS IAM Identity Center.  See Grafana documentation.

## What is MVP scope for Coast?
An Infrastructure as Service repository (CloudFormation) the customers can use to deploy on top of existing CUR reports, an Amazon Managed Grafana workspace, and a SNS topic for alerts; a summary dashboard focused on DevOps teams; drill-down dashboards for EC2, EKS, ECS, RDS, VPC and Storage; an alerts dashboard; and pre-configured alerts.
