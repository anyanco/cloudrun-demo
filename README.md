# Cloud Run demo

## OverView
Cloud Run から GCS、FireStore にそれぞれ接続して、データを取得する。

## URL
##### [HOST]/
  `Hello World` を表示する。
##### [HOST]/storage
  GCS に接続し、`demo-test.txt` の内容を取得、表示する。
##### [HOST]/firestore
  FireStore に接続し、`demo/demo-access/User` 以下のデータを取得、表示する。<br>
  ※`const.go`の以下設定参照
  ```
    const COLLECTION_NAME string = "demo/demo-access/User"
  ```
##### [HOST]/login?u=[ユーザ名]
1. URLのクエリパラメータからユーザ名取得する
1. FireStoreから、前回アクセス日時を取得する<br>
  &emsp;・ユーザ名に対応するデータが存在する場合　　：データから取得<br>
  &emsp;・ユーザ名に対応するデータが存在しない場合　：現在日時<br>
1. 現在日時にて、データを更新する<br>
  &emsp;・ユーザ名に対応するデータが存在する場合　　：更新<br>
  &emsp;・ユーザ名に対応するデータが存在しない場合　：新規作成<br>
1. GCS に接続し、`index.html`（テンプレート）を取得する<br>
1. ３.にて取得した内容をテンプレートに埋め込む<br>
1. ブラウザにて表示

## Cloud Run VALIABLES - 環境変数の設定
Cloud Run 使用時に以下の環境変数の設定が必要。

- `BUCKETNAME` : GCSのバケット名
- `PROJECT_ID` : プロジェクト名

## Template
`template` 以下の htmlファイルは GCS に配置する。

## Build
```
gcloud builds submit --tag asia.gcr.io/[PROJECT_ID]/[SERVICE_NAME]:[VERSION]
```


## Deploy
```
gcloud run deploy --image asia.gcr.io/[PROJECT_ID]/[SERVICE_NAME]:[VERSION] \
  --platform managed
```
