# Artwork 說明

## 環境配置

- 語言: go v1.23.4
- 資料庫: postgresql v14.13

## 建立遷移檔

1. 可透過指定 name 方式建立

```bash
$ make create-migration name=create_xxxxx_table
```

2. 不指定，則會跳出輸入輸入框

```bash
$ make create-migration 
```

## 啟用遷移檔

```bash
$ make migrate-up
```

## 回滾遷移檔


```bash
$ make migrate-down
```