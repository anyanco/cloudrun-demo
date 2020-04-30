# Cloud Run demo（作成中）

## OverView
Cloud Run から GCS、FireStore に接続して、データを取得する。

## 環境変数の設定
Cloud Run 使用時に以下の環境変数の設定が必要。

- BUCKETNAME : GCSのバケット名
- PROJECT_ID : プロジェクト名

# Build
```
gcloud builds submit --tag asia.gcr.io/[PROJECT_ID]/[SERVICE_NAME]:[VERSION]
```


# Deploy
```
gcloud run deploy --image asia.gcr.io/[PROJECT_ID]/[SERVICE_NAME]:[VERSION] \
  --platform managed
```
