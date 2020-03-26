# Attendance Management

## Create database
```sql
create database attendance_management;
```

## Go Setup
globalでrealizeを使えるようにする
```
cd 
mkdir go-tools
cd go-tools
go mod init go-tools

go get -u gopkg.in/urfave/cli.v2@master
go get -u github.com/oxequa/realize
```

```bash
cd backend
cp .env.local.sample .env.local
```

## hot reload
```
make run
```

## Add Firebase Config
TSファイルを作成する。
```bash
touch frontend/src/lib/firebase/config.ts
```

作成したファイルにFirebaseのConfig情報を入れる。
```typescript
export const firebaseConfig = {
    apiKey: "YOUR_KEY",
    authDomain: "YOUR_KEY",
    databaseURL: "YOUR_KEY",
    projectId: "YOUR_KEY",
    storageBucket: "YOUR_KEY",
    messagingSenderId: "YOUR_KEY",
    appId: "YOUR_KEY"
 }
```

## Add GCP Config

ダウンロードして来たサービスアカウント情報をfirebase-service.jsonとしてリネームする。
```bash
mv　YOUR_FILE backend/config/development/config/firebase-service.json
```

## Use Cloud Sql
```bash
./cloud_sql_proxy -instances=<INSTANCE_NAME:REGION:NAME>=tcp:3306
```

## Migrate database
mysql://root:[PASSWORD]を入れる必要がある。
```
cd backend
make migrate
```

## Deploy app engine
```bash
gcloud app deploy YOUR_FILE.yml
```

## CircleCi local test
```bash
circleci local execute --job build
```

## docker hub 登録
```bash
docker build -t org/name:tag
docker push org/name:tag
```