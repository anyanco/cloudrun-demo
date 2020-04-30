# Cloud Run demo（作成中）

## OverView
Cloud Run から GCS、FireStore に接続して、データを取得する。

## URL
##### [HOST]/
  `Hello World` を表示する。
##### [HOST]/storage
  GCS に接続し、`demo-test.txt` の内容を取得、表示する。
##### [HOST]/firestore
  FireStore に接続し、`demo/demo-access` 以下のデータを取得、表示する。
##### [HOST]/login
  GCS に接続し、`index.html`（テンプレート）を取得、
  FireStore の内容を埋め込んで表示する。


## Cloud Run VALIABLES - 環境変数の設定
Cloud Run 使用時に以下の環境変数の設定が必要。

- BUCKETNAME : GCSのバケット名
- PROJECT_ID : プロジェクト名

## Template
`app/template` 以下の htmlファイルは GCS に配置する。

# Build
```
gcloud builds submit --tag asia.gcr.io/[PROJECT_ID]/[SERVICE_NAME]:[VERSION]
```


# Deploy
```
gcloud run deploy --image asia.gcr.io/[PROJECT_ID]/[SERVICE_NAME]:[VERSION] \
  --platform managed
```
