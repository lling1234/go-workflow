# go-workflow
奥集能平台工作流
## 目录介绍

## git规范
```
● add/feat：新功能
● fix：bug修复
● update：更新
● del：移除文件
● docs：文档改变
● style：代码格式改变
● perf：性能优化
● test：添加测试代码
● revert：撤销上一次的commit
● refactor：某个已有功能重构
● build：构建工具或构建过程的变动
<<<<<<< HEAD
```

## goctl生成代码命令
rpc代码生成前，删除5个文件
1.rpc/types里面act_grpc.pb.go 和 act.pb.go
2.rpc/internal/server actserver.go
3.rpc/internal/findallprocinstlogic.go 修改protoc文件的rpc接口
4.rpc/actclient/act.go
```
goctl rpc protoc act.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
```
=======

## 后续工作
1、平台对接
2、审批驳回
3、抄送（转发）
4、审批人（抄送人）为岗位
5、审批时可以上传附件
6、条件节点
7、并行流程
8、触发器
9、流程查询
>>>>>>> 4a54178f732840048a6d221f9338a6e8f5d12ba3
