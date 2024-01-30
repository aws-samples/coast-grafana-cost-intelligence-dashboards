import json
import boto3
import datetime
import time
import os
import logging
import sys


logger = logging.getLogger()
logger.setLevel(logging.INFO)

def lambda_handler(event, context):
    bucket_name = os.getenv("s3_bucket")
    body = event['body']
    print(body)
    s3 = boto3.client('s3')
    params_str = json.dumps(body)  
    tme=str(round(time.time() * 1000))
    file_name = "/file_"+tme+".json"
    e = datetime.datetime.now()
    timetoday=e.strftime("%Y/%m/%d")
    s3_path =timetoday+file_name
    s3.put_object(Bucket=bucket_name, Key=s3_path, Body=params_str)
    return {
      'statusCode': 200
    }