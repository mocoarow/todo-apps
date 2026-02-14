# Todo Apps プロジェクト概要

## 構成
- **バックエンド**: `backend-gin-gorm/` - Go (Gin + GORM + MySQL)
- **フロントエンド**: `frontend-react-zustand/` - React + Zustand + Vite + TypeScript
- **API定義**: `openapi/openapi.yaml` (OpenAPI 3.x)
- **型生成**: `frontend-react-zustand/src/api/types.gen.ts` (Orvalで自動生成、手動編集禁止)
- **API再エクスポート**: `frontend-react-zustand/src/api/index.ts` (スキーマと型をまとめてexport)

## フロントエンドアーキテクチャ
- **状態管理**: Zustand (`src/stores/`)
- **API通信**: HttpClient + Service層 (`src/gateway/`)
- **ドメイン**: インターフェース定義 (`src/domain/`)
- **UIコンポーネント**: shadcn/ui (`src/components/ui/`)
- **パスエイリアス**: `~/` = `src/`

## バックエンドアーキテクチャ
- **レイヤー**: controller/handler → usecase → gateway(repository)
- **DB**: GORM + MySQL, `NowFunc` で UTC 統一
- **テスト**: Docker MySQL を使った統合テスト (`compose.test.yml`)

## CI/CD チェック項目
- `task check` でローカル検証可能
- **Frontend**: biome (lint/format) + eslint + tsc + knip (unused exports)
- **Go**: golangci-lint + go test (Docker MySQL)
- **commitlint**: Conventional Commits 必須

## 重要な注意点
- GORM `Save()` はフルUPDATEを発行し、`autoCreateTime`フィールドの精度問題を引き起こす。`Model().Updates(map)` で選択的更新を使うこと
- GORM `Create()` 後は Go 構造体にナノ秒精度のタイムスタンプが残る。DB 精度と一致させるには `First()` で再読み込みが必要
- `CreateTodoResponse` と `UpdateTodoResponse` は同一スキーマ。`CreateTodoResponseSchema` は `UpdateTodoResponseSchema` のエイリアス
- フロントエンドの `console.error`/`console.log` はプロダクションコードで禁止
- biome の `noStaticElementInteractions` / `useSemanticElements` ルールに注意。`onDoubleClick` 等のイベントハンドラには `<button>` を使う
