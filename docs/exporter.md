## 区块链节点 / Stratum Server Exporter

### 功能

+ 支持节点获取 peers，高度
+ 支持常见币种矿池节点
+ 支持自定义公开节点（通过握手协议）
+ 支持一些浏览器 API
+ 支持按高度 / hash 报警
+ TODO 自定义地址余额监控

### 节点的获取内容

块高、
hash、
占用硬盘空间、
内存池情况、
peers、
错误提醒、

### 结构

source 概念

source 是一种后端的抽象、可以代表stratum server

目前类型：

+ rpc 节点
+ stratum server 节点
+ 公开节点
+ 浏览器API

### 步骤

1. 读取配置（或者从 apollo 读取）
2. exporter 根据读取的配置进行转化，注册组件
3. 