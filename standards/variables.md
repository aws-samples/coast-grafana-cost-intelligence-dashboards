Variables created which are common to our dashboards should be standardized to allow for creation of a library of panels which may be imported into each dashboard.  Variables which are only relevant to one dashboard need not be logged in this standards document.

--\
**Variable Name**: CurTable\
**Lable**: CUR Table\
**Data source**: Athena CUR\
**Query**:\
```show tables```\
**Notes**: \
Use regex to filter for:\
```^(?!cost_and_usage_data_status).+$```

Do not include multi-value or All option

--\
**Variable Name**: LinkedAccounts\
**Label**:Linked Accounts\
**Data source**: Athena CUR \
**Query**:\
```SELECT DISTINCT(line_item_usage_account_id) FROM \$CurTable;```
**Notes**: \
Include multi-value or All option

--\
**Variable Name**: Regions\
**Label**:Regions\
**Data source**: Athena CUR \
**Query**:\
```SELECT DISTINCT(product_region) from $CurTable;```
**Notes**: \
Include multi-value or All option

--\
**Variable Name**: Services\
**Label**:Services\
**Data source**: Athena CUR \
**Query**:\
```SELECT DISTINCT(line_item_product_code) from $CurTable;```
**Notes**: \
Include multi-value or All option

--\
**Variable Name**: ChargeType\
**Label**:ChargeType\
**Data source**: Athena CUR \
**Query**:\
```SELECT DISTINCT(line_item_line_item_type) from $CurTable;```
**Notes**: \
Include multi-value or All option

--\
**Variable Name**: TagKey\
**Label**:Tag Key\
**Data source**: Athena CUR \
**Query**:\
```SHOW COLUMNS from $CurTable;```
**Notes**: \
Use regex to filter for:\
```/resource_tags_user_(.+)/```

Do not include multi-value or All option

--\
**Variable Name**: TagValue\
**Label**:Tag Value\
**Data source**: Athena CUR \
**Query**:\
```SELECT DISTINCT(resource_tags_user_$TagKey) FROM $CurTable WHERE resource_tags_user_$TagKey IS NOT NULL AND NOT resource_tags_user_$TagKey = '';```
**Notes**: \
Include multi-value or All option

--\
**Variable Name**: CostCategoryName\
**Label**:Cost Category Name\
**Data source**: Athena CUR \
**Query**:\
```SHOW COLUMNS from $CurTable;```
**Notes**: \
Use regex to filter for:\
```/cost_category_(.+)/```

Do not include multi-value or All option

--\
**Variable Name**: CostCategoryValue\
**Label**:Cost Category Value\
**Data source**: Athena CUR \
**Query**:\
```SELECT DISTINCT(cost_category_$CostCategoryName) FROM $CurTable;```
**Notes**: \
Include multi-value or All option

--\
**Variable Name**: PanelTop25SpendPercentage\
**Label**:\
**Description**: Calculate the spend percentage from the top 25 accounts in the AWS Organization.\
**Data source**: Athena CUR \
**Query**:\
```with total_spend as (
    SELECT
    line_item_usage_account_id as account,
    sum(line_item_unblended_cost) as total_spend
    FROM
    $CurTable
    WHERE
    bill_billing_period_start_date = date_add('MONTH', -1, DATE_TRUNC('MONTH', current_date))
    GROUP BY
    line_item_usage_account_id
),

top10_spend as (
    SELECT
    line_item_usage_account_id as account,
    sum(line_item_unblended_cost) as top25_spend
    FROM
    $CurTable
    WHERE
    bill_billing_period_start_date = date_add('MONTH', -1, DATE_TRUNC('MONTH', current_date))
    GROUP BY
    line_item_usage_account_id
    LIMIT 10
)


SELECT
    CAST( (sum(top25_spend) / sum(total_spend) * 100) as VARCHAR) as percentage_from_top25
FROM
    total_spend
INNER JOIN top10_spend ON total_spend.account = top10_spend.account;
```
**Notes**: \
Do not include multi-value or All option

--\
**Variable Name**: TotalAccountInOrganization\
**Label**:\
**Data source**: Athena CUR \
**Query**:\
```SELECT 
CAST(COUNT(DISTINCT(line_item_usage_account_id)) as VARCHAR) as account_count
FROM 
$CurTable
```
**Notes**: \
Do not include multi-value or All option

--\
**Variable Name**: Granularity\
**Label**:Granularity\
**Data source**: Custom \
**Query**:\
Custom Options:\
```month,week,day,hour```
**Notes**: \
Do not include multi-value or All option



