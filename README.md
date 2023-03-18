# ms-zeus

#### 介绍
子应用集成权限管理的模块

#### wire

```
cd zeus
wire

```


#### 使用

使用中间件

```
// 添加中间件
myjwt.JwtMiddleware(whiteList, app.accessKey),

// 注册访问接口
srv := http.NewServer(opts...)
//swagger-api
authHandler := zeus.NewHandler()
srv.HandlePrefix("/q/", openAPIhandler)

```