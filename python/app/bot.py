


# from slack_sdk import WebClient

# # 初期化
# client = WebClient(token=os.environ.get('SLACK_API_TOKEN'))

# # メッセージ履歴を取得
# response = client.conversations_history(
#   channel=os.environ.get('SLACK_CHANNEL'),
# )

# # メッセージ一覧を取得し、それぞれのメッセージからタイムスタンプを取得
# for message in response['messages']:
#   print(f"Timestamp: {message['ts']}, Text: {message['text']}")


from slack_sdk import WebClient
import os

slack_token = os.environ.get('SLACK_API_TOKEN')
client = WebClient(token=slack_token)

channel_id = os.environ.get('SLACK_CHANNEL') 
message_ts = '1476746830.000083'

response = client.chat_delete(
  channel=channel_id,
  ts=message_ts
)

