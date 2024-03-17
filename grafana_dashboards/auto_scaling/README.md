# Autoscaling Dashboard

![Autoscaling Dashboard](../../images/coast_banner.png)

## About

The COAST Autoscaling Dashboard provides a comprehensive overview of cost and infrastructure performance metrics related to EC2 Autoscaling Groups, Load Balancers, NAT Gateways, and networking across all associated resources. The data sources feeding into this dashboard are CloudWatch and the Cost and Usage Report (CUR).

To effectively utilize this dashboard, it necessitates the specification of your autoscaling group's name. Moreover, all resources within your autoscaling workload must be tagged with a unique tag key/value pair specific to this workload. It is imperative that these tags are enabled as Cost Allocation Tags and persist within the Cost and Usage Report (CUR).

## Usage
All panels within the Autoscaling Dashboard require cost allocation tags to exist within your CUR.  First select the region where or the workload.  Next, you will need to fill in the name of the autoscaling group you wish to monitor, the cost of the autoscaling group will be displayed with by selecting the accompaning tag for the autoscaling group and all of it's resources, and supporting resources (i.e. LB, NATGW).   

The values for the NAT Gateway and Load Balancer menus will populate based on your selection for tags.  Additional graphs will be selected once one or more NATGW's or LB's are selected. 


## Limitations 

Currently the COAST Autoscaling Dashboard is in a proof of concept phase.  It has been certified for workloads which are operating in the same account same region as the COAST deployment.  We are working to make the COAST Autoscaling Dashboard operational across regions and across accounts.

## Layout
Type: [unblended|amortized] - how are the costs displated unblended or amortized
TagSupport: [Yes|No] -  If the panel requires tags
DependsOn: If the panel depends on any menu selection items

```mermaid
block-beta

  columns 4

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

  block:computeCosts:4
    b6["Compute Cost by AZ
    Type:unblended;TagSupport:Yes"]
    b7["Compute Cost by Instance
    Type:unblended;TagSupport:Yes"]
  end 

  block:serviceTrends:4
    b5["Service Cost Trend
    Type:unblended;TagSupport:Yes"]
  end

  block:computeActivity:4
    columns 2
    c1["Scaling Activity
    Type:unblended;TagSupport:Yes
    DependsOn: ASGName"]
    c2["Average CPU
    Type:unblended;TagSupport:Yes
    DependsOn: ASGName"]
    
    c3["Instance Status Check
    Type:unblended;TagSupport:Yes
    DependsOn: ASG Name"]
    c4["CPU Credit Balance
    Type:unblended;TagSupport:Yes
    DependsOn: ASG Name"]
  end 

  block:networkingActivity:4
    columns 2
    d1["EC2 Networking Bytes Out
    Type:unblended;TagSupport:Yes
    DependsOn: ASG Name"]

    d2["EC2 Networking Packets Out
    Type:unblended;TagSupport:Yes
    DependsOn: ASG Name"]
  end

  block:elbActivity:4
    columns 2
    e1["LCU Metrics
    Type:unblended;TagSupport:Yes
    DependsOn: ELB ResourceID"]

    e2["Hourly Costs and Requests
    Type:unblended;TagSupport:Yes
    DependsOn: ELB ResourceID"]
  end

  block:natgwActivity:4
    columns 2
    f1["Nat Gateway Traffic
    Type:unblended;TagSupport:Yes
    DependsOn: NATGW ResourceID"]

    f2["Nat Gateway Hourly
    Type:unblended;TagSupport:Yes
    DependsOn: NATGW ResourceID"]
  end

  

  classDef section1 fill:#E6E6FA,stroke:#7F00FF,stroke-width:1px;
  classDef section2 fill:#C2DFFF,stroke:#357EC7,stroke-width:1px;
  classDef tagClass fill:#F5F5F5,font-size:4pt,stroke:#646D7E 
  classDef notagClass fill:#E0FFFF,font-size:4pt,stroke:#646D7E 
  
  class Instructions,computeActivity,elbActivity section1
  class computeCosts,serviceTrends,workloadCostTrends,networkingActivity,natgwActivity section2
  class a1,a2,a3,a4,b1,b2,b3,b4,b5,b6,b7,c1,c2,c3,c4 tagClass
  class d1,d2,e1,e2,f1,f2 tagClass

  classDef section1 fill:#E6E6FA,stroke:#7F00FF,stroke-width:1px;
  classDef section2 fill:#C2DFFF,stroke:#357EC7,stroke-width:1px;
  classDef tagClass fill:#F5F5F5,font-size:4pt,stroke:#646D7E 
  classDef notagClass fill:#E0FFFF,font-size:4pt,stroke:#646D7E 
  
  class Instructions,computeActivity,elbActivity section1
  class computeCosts,serviceTrends,workloadCostTrends,networkingActivity section2
  class a1,a2,a3,a4,b1,b2,b3,b4,b5,b6,b7,c1,c2,c3,c4 tagClass
  class d1,d2,e1,e2 tagClass

```

## Troubleshooting

#####Autoscaling Metrics
If you are not seeing Autoscaling Metrics in Cloudwatch, make sure to enable 'Auto Scaling group metrics collection' under the monitoring tab in your Autoscaling Group console.


