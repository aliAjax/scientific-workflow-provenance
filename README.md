# Scientific Workflow Provenance

科研工作流编排与样本溯源后端，纯 Go 1.23 实现。核心路径包含 DAG 环检测、条件/矩阵展开、确定性缓存键、运行状态机、资源队列、模拟 Worker、artifact 质量门禁、样本冻结/销毁和 W3C PROV 风格哈希链。

## Run

```bash
go run ./cmd/workflowd
curl http://localhost:8088/healthz
```

默认服务地址为 `:8088`，可通过 `HTTP_ADDR` 覆盖。持久化和对象存储通过端口抽象，当前默认使用内存适配器，便于本地验证；生产环境替换 repository/object store adapter。

## API

- `POST/GET /api/v1/workflows`，`GET /api/v1/workflows/{id}`
- `POST /api/v1/workflows/{id}/validate`、`/plan`
- `POST/GET /api/v1/samples`，`/samples/{id}/lineage`、`freeze`、`destroy`
- `POST/GET /api/v1/runs`，`/runs/{id}`、`cancel`、`provenance`
- `POST/GET /api/v1/artifacts`，`/artifacts/{digest}/verify`、`invalidate`
- `/healthz`、`/readyz`

所有 API 返回统一 JSON 错误结构。工作流定义的 `edges` 经过拓扑排序和环检测，矩阵节点生成稳定 step id/cache key。样本和 artifact 事件写入 provenance graph，并通过前置哈希保证审计链可验证。

## Layout

`internal/domain` 保存领域模型；`internal/planner`、`scheduler`、`execution` 实现应用能力；`repository`、`infrastructure`、`adapters` 为可替换适配器；`transport` 提供版本化 HTTP；`pkg` 提供分页、审计、单位换算、表达式和事件总线。

## Verification

```bash
gofmt -w $(find . -name '*.go')
go test ./...
go test -race ./...
go vet ./...
go build ./...
```

## Operations

配置项见 `.env.example`。Docker/Kubernetes 和 CI 文件位于 `deployments/`、`.github/`，数据库迁移位于 `migrations/`。生产部署应启用 TLS、OIDC/mTLS、对象存储和关系数据库适配器，并为 worker lease、artifact 保留策略和 provenance 归档配置告警。
