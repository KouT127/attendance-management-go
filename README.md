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