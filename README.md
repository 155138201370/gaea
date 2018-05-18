# gaea
用于创建，监控基于hydra的项目


### 一、安装
```sh
  go get github.com/micro-plat/gaea
```

### 二、创建项目

#### 1. 简单项目

gaea new project [项目名称]

```sh
gaea new project myproject/apiserver
创建文件: /home/yanglei/work/src/myproject/apiserver/main.go
创建文件: /home/yanglei/work/src/myproject/apiserver/bind.go
项目生成完成
```

### 2. 生成包含有模块代码的项目
gaea new project [项目名称] -m [模块名称]
```sh
gaea new project myproject/apiserver -m order/request
创建文件: /home/yanglei/work/src/myproject/apiserver/main.go
创建文件: /home/yanglei/work/src/myproject/apiserver/bind.go
创建文件: /home/yanglei/work/src/myproject/apiserver/services/order/request.go
创建文件: /home/yanglei/work/src/myproject/apiserver/modules/order/request.go
创建文件: /home/yanglei/work/src/myproject/apiserver/modules/sql/order.go
项目生成完成
```

### 3. 生成包含多个模块代码的项目
gaea new project [项目名称] -m ["模块1 模块2"]
```sh
gaea new project myproject/apiserver -m "order/request order/query"
创建文件: /home/yanglei/work/src/myproject/apiserver/main.go
创建文件: /home/yanglei/work/src/myproject/apiserver/bind.go
创建文件: /home/yanglei/work/src/myproject/apiserver/services/order/request.go
创建文件: /home/yanglei/work/src/myproject/apiserver/services/order/query.go
创建文件: /home/yanglei/work/src/myproject/apiserver/modules/order/request.go
创建文件: /home/yanglei/work/src/myproject/apiserver/modules/sql/order.go
创建文件: /home/yanglei/work/src/myproject/apiserver/modules/order/query.go
项目生成完成
```


### 4. 生成Restful风格API代码
gaea new project [项目名称] -r
```sh
gaea new project myproject/apiserver -m order/request -r
创建文件: /home/yanglei/work/src/myproject/apiserver/main.go
创建文件: /home/yanglei/work/src/myproject/apiserver/bind.go
创建文件: /home/yanglei/work/src/myproject/apiserver/services/order/request.go
创建文件: /home/yanglei/work/src/myproject/apiserver/modules/order/request.go
创建文件: /home/yanglei/work/src/myproject/apiserver/modules/sql/order.go
项目生成完成
```


### 5. 生成指定服务类型的项目
gaea new project [项目名称] -s[服务器类型]
```sh
gaea new project myproject/apiserver -m order/request -s api-cron
创建文件: /home/yanglei/work/src/myproject/apiserver/main.go
创建文件: /home/yanglei/work/src/myproject/apiserver/bind.go
创建文件: /home/yanglei/work/src/myproject/apiserver/services/order/request.go
创建文件: /home/yanglei/work/src/myproject/apiserver/modules/order/request.go
创建文件: /home/yanglei/work/src/myproject/apiserver/modules/sql/order.go
项目生成完成
```

