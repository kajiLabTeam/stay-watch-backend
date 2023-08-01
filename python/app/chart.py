import matplotlib.pyplot as plt
import datetime

plt.rcParams['font.family'] = 'IPAGothic'

# # ユーザのログ時間の合計を円グラフで表示して保存する関数

def save_users_log_time_pie_chart(log_times: dict):

    # 現在日時の取得
    now = datetime.datetime.now().strftime('%Y-%m-%d')

    # ログ時間で降順にソートされたユーザー名とログ時間
    sorted_log_times = sorted(
        log_times.items(), key=lambda x: x[1], reverse=True)
    print(sorted_log_times)
    user_names = [x[0] for x in sorted_log_times]
    log_times = [x[1] for x in sorted_log_times]

    # 上位10人以外のログ時間をまとめる
    if len(log_times) > 10:
        other_label = 'Other'
        other_time = sum(log_times[10:])
        user_names = user_names[:10] + [other_label]
        log_times = log_times[:10] + [other_time]

    # カラーパレット
    colors = ['#e6194B', '#f58231', '#ffe119', '#3cb44b', '#4363d8',
              '#911eb4', '#f032e6', '#a9a9a9', '#fabebe', '#008080', '#9A6324']

    # 円グラフを描画
    plt.pie(log_times, labels=user_names, startangle=90, counterclock=False, autopct='%1.1f%%', colors=colors, textprops={
        'fontsize': 11,
    })

    # 画像を保存
    plt.savefig(f'./image/{now}.png', dpi=300)

    return f'./image/{now}.png'




def save_users_log_time_bar_chart(log_times: dict):

    # 現在日時の取得
    now = datetime.datetime.now().strftime('%Y-%m-%d')

    # ログ時間で降順にソートされたユーザー名とログ時間
    sorted_log_times = sorted(
        log_times.items(), key=lambda x: x[1], reverse=True)
    user_names = [x[0] for x in sorted_log_times]
    log_times = [x[1] for x in sorted_log_times]


    user_names=user_names[:15]
    log_times=log_times[:15]

    # 棒グラフを描画
    plt.barh(user_names[::-1], log_times[::-1], color='#4363d8')

    plt.xlabel('Log Time(hours)')
    plt.title('User Log Times')

    # 画像を保存
    plt.tight_layout()
    plt.savefig(f'./image/{now}.png', dpi=300)

    return f'./image/{now}.png'
