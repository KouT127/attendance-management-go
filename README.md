# Attendance Management

## Create database
```sql
create database attendance_management;
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
touch backend/config/development/config/firebase-service.json
```