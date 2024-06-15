# Docker起動時のcron設定

## 1. crontabの編集モードを起動

```bash
crontab -e
```

## 2. 以下を記述

```text
0 9 * * 1 date >> /usr/clustering/server.log
0 9 * * 1 /usr/local/bin/cron.sh
```

## 3.cronの起動

```bash
service cron start
```

## 4.cronの確認

```bash
service cron status
```
