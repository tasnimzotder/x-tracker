import json
import time

import boto3
# import redis

dynamodb = boto3.resource('dynamodb')
table = dynamodb.Table('xt_test_edge_data')

# redis_client = redis.Redis(
#     host="xt-test-edge-data-reveiz.serverless.use1.cache.amazonaws.com",
#     port=6379,
# )


# def update_latest_location(device_id, location):
#     redis_key = "latest_location#{}".format(device_id)
#     redis_client.hset(redis_key, mapping=location)

def lambda_handler(event, context):
    item = {
        'device_id#timestamp': "{}#{}".format(event["device_id"], event["timestamp"]),
        'edge_timestamp': event["timestamp"],
        'processed_timestamp': int(time.time() * 1000),
        'device_id': event['device_id'],
        'client_id': event['client_id'],
        'lat': "{}".format(event['lat']),
        'lng': "{}".format(event['lng']),
        'battery_status': event['battery_status']
    }

    table.put_item(Item=item)

    # redis
    # redis_key = "{}#{}".format(event["device_id"], event["timestamp"])
    # redis_client.hset(redis_key, mapping=item)
    # redis_client.expire(redis_key, 60 * 60 * 24)  # 24 hours
    #
    # update_latest_location(event["device_id"], item)

    return {
        'statusCode': 200,
        'body': json.dumps('Data stored successfully.')
    }
