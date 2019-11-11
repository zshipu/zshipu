# CoreOS 工具介绍

`CoreOS` 内置了 `服务发现`，`容器管理` 工具。

## 服务发现

`CoreOS` 的第一个重要组件就是使用 `etcd` 来实现的服务发现。在 `CoreOS` 中 `etcd` 默认以 `rkt` 容器方式运行。

`etcd` 使用方法请查看 [etcd 章节](../etcd/)。

## 容器管理

第二个组件就是 `Docker`，它用来运行你的代码和应用。`CoreOS` 内置 `Docker`，具体使用请参考本书其他章节。
