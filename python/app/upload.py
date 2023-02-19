import os
from slack_sdk import WebClient
from slack_sdk.errors import SlackApiError
from dotenv import load_dotenv


# slackに画像をアップロードする
def post_slack(image_path):

    # .envファイルを読み込む
    load_dotenv()
    # Slack APIトークンを設定
    client = WebClient(
        # SLACK_API_TOKEN環境変数を取得する
        token=os.environ.get('SLACK_API_TOKEN'))

    try:
        # files.upload APIメソッドを呼び出して、画像をアップロード
        response = client.files_upload_v2(
            channel=os.environ.get('SLACK_CHANNEL'),
            file=image_path,
            title='Uploaded image'
        )
        print("File uploaded: {}".format(response['file']['name']))
    except SlackApiError as e:
        print("Error uploading file: {}".format(e))
