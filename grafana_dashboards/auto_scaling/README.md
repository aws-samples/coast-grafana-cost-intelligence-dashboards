# Autoscaling Dashboard

![Autoscaling Dashboard](../../images/coast_banner.png)

## About

The COAST Autoscaling Dashboard provides a comprehensive overview of cost and infrastructure performance metrics related to EC2 Autoscaling Groups, Load Balancers, NAT Gateways, and networking across all associated resources. The data sources feeding into this dashboard are CloudWatch and the Cost and Usage Report (CUR).

To effectively utilize this dashboard, it requires the input of your autoscaling group's name. All resources within your autoscaling workload must be tagged with a unique tag key/value pair specific to the workload. It is imperative that these tags are enabled as Cost Allocation Tags and persist within the Cost and Usage Report (CUR).

## Limitations 

Currently the COAST Autoscaling Dashboard is in a proof of concept phase.  It has been certified for workloads which are operating in the same account as the COAST deployment.  Workloads from multiple regions in the account may be observed.  We are working on a mechanism to make the COAST Autoscaling Dashboard operational across accounts.

## Layout
```
mermaid
block-beta

  block:Instructions:4
    columns 2
    a1["Auto Scaling Dashboard"] 
    a2["How to Utilize this 
    Dashboard"]

    a3["Validator Instructions"]
    a4["Validator"]
  end
  
  block:workloadCostTrends:4
    columns 4
    b1["Compute - Period over Period  
    Type:unblended;TagSupport:Yes"]
    b2["Load Balancer - Period over Period
    Type:unblended;TagSupport:Yes"]
    b3["Network - Period over Period
    Type:unblended;TagSupport:Yes"]
    b4["NAT Gateway - Period over Period
    Type:unblended;TagSupport:Yes"]
  end

  columns 4
  b5["Service Cost Trend
  Type:unblended;TagSupport:Yes"]:4

  block:computeCosts:4
    b6["Compute Cost by AZ
    Type:unblended;TagSupport:Yes"]
    b7["Compute Cost by Instance
    Type:unblended;TagSupport:Yes"]
  end 

  classDef defaultBox fill:#E6E6FA,stroke:#7F00FF,stroke-width:1px;
  classDef notagClass fill:#F5F5F5,font-size:4pt  
  classDef seperators font-size:2pt,width:100%
  
  class Instructions,workloadCostTrends,computeCosts defaultBox
  class a1,a2,a3,a4,b1,b2,b3,b4,b5,b6,b7 notagClass
  class sep1 seperators
```

## Troubleshooting

#####Autoscaling Metrics
If you are not seeing Autoscaling Metrics in Cloudwatch, make sure to enable 'Auto Scaling group metrics collection' under the monitoring tab in your Autoscaling Group console.


