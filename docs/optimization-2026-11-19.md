# 2026-11-19 扫描：问题与优化步骤

只读扫描本仓库（不含 `vendor`）。本文不记录密钥明文。总览见 [r2day/docs](https://github.com/r2day/docs/blob/main/architecture/2026-11-19-api-optimization.md)。

这是账号模型库 `github.com/r2day/auth`，不是 `auth-api`。当前 18 个 API 的 `go.mod` 没有直接 require 本模块。

优先级：P0 可被利用；P1 正确性；P2 债。

## 问题

### P1 按手机号查找没有商户条件

`admin/manage/method.go` 的 `FindByPhone` 过滤只有 `phone`。两个商户用同一个手机号时，会返回先命中的那条。`FindByAccountId` 过滤的是 `meta.account_id`，同样没有商户。

### P2 连接失败直接退出进程

`handler.go` 的 `New` 在 URI 为空时 `log.Fatal`，连接失败时 `panic`。库函数把整个进程带走。

## 步骤

1. 若仍有服务调用 `FindByPhone`，过滤加上商户号。手机号只在该商户内唯一。补一条跨商户同号测试。
2. `New` 改为返回 `error`。调用方决定是否退出。
3. 本库没有被当前 API 引用。确认没有外部仓库仍在用之后，再决定是补测试还是标记废弃。
