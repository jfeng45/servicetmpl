## 基于gRPC的Go微服务

其他语言：

### **[English](README.md)**

这是一个基于gRPC的Go微服务项目。它为Go微服务程序找到了合适的程序结构，并且应用了清晰架构（Clean Atchitecture）设计并使用依赖注入（Dependency Injection）将具体类型注入到每个函数中。

## 如何使用这个项目
在创建gRPC微服务项目时，此项目最适合用作基本框架。它已经具有丰富的内置功能并且已经调试通过，因此没有理由从头开始。该项目的目标是构建一个具有基本功能的灵活框架，以便以后轻松扩展。
程序设计遵循“SOLID (面向对象设计)”设计原则和Go的简洁编码风格，因此当你需要执行某些编码规范时，它可以用作程序设计和编码样式的具体示例。

## 用它作为模板来启动微服务项目
### 功能特点：
1. 通过更改配置文件来切换持久层实现。目前，它实现了MySQL和CouchDB数据库。 （它可以扩展以支持其他SQL或NoSQL数据库）
2. 通过更改配置文件切换日志记录库。目前，它实现了ZAP和Logrus。 （它可以扩展支持其他日志记录提供库，只要它们有类似于ZAP和Logrus的通用接口）
3. 支持业务层事务（它不支持跨微服务的事务或嵌套事务）
4. 使用依赖注入创建具体类型。
5. 程序配置保存在YAML文件中。

### 设计特点：
##### 1.接口编程
* 程序有三层：用例，模型和持久性。每层通过接口访问其他层。
* 外部功能也通过接口访问。
##### 2.使用工厂方法模式通过依赖注入创建具体类型
##### 3.最小化依赖性
* 不同层之间的依赖关系仅在接口而不是具体类型上。
* 接口在顶级包中定义，并与具体类型分开。
* 每个具体类型都在单独的子包和文件中定义
##### 4.功能隔离
* 用包来隔离不同的层
* 用包隔离每个用例
* 用包隔离每个实现（例如数据库实现）
##### 5.开闭原则（Open-closed Principle）
* 当添加新功能时，请添加新代码，而不是修改现有代码
 
### 编码风格：
1. 消除包级变量 （“容器”包是例外）
2. 尽量减少常量的使用。
3. 记录错误的完整堆栈跟踪
4. 仅在调用链顶函数层处理错误
5. 隔离不同的关注（Separation of Concerns)
6. 命名规范。局部变量用短名称，类型或接口用长名称。 

## 运行

### 安装和设置

不需要完成本节中的所有步骤以使代码运行。 最简单的方法是从github获取代码并运行它，然后在真正需要某些部件时再返回安装。 但是，访问数据库时会遇到错误。
所以，我建议你至少安装一个数据库（MySQL更好），然后大部分代码就都可以运行了。

#### 下载程序

```
go get github.com/jfeng45/servicetmpl
```

#### 设置MySQL

有两个数据库实现，MySQL和CouchDB，但大多数函数都是在MySQL中实现的。 你最好安装至少其中一个。

```
安装MySQL
在script文件夹中运行SQL脚本以创建数据库和表
```
#### 安装CouchDB

没有它，代码工作正常。创建CouchDB用来完成切换数据库的功能（通过更改配置）。

安装[Windows](https://docs.couchdb.org/en/2.2.0/install/windows.html)

安装[Linux](https://docs.couchdb.org/en/2.2.0/install/unix.html)

安装[Mac](https://docs.couchdb.org/en/2.2.0/install/mac.html)

CouchDB[Example](https://github.com/go-kivik/kivik/wiki/Usage-Examples)

#### 设置CouchDB

```
通过浏览器访问“Fauxton”：http://localhost:5984/_utils/#（使用：admin/admin登录）。
在“Fauxton”中创建新数据库“service_config”。
将以下文档添加到数据库（“_id”和“_rev”由数据库生成，无需更改）：
{
  "_id": "80a9134c7dfa53f67f6be214e1000fa7",
  "_rev": "4-f45fb8bdd454a71e6ae88bdeea8a0b4c",
  "uid": 10,
  "username": "Tony",
  "department": "IT",
  "created": "2018-02-17T15:04:05-03:00"
}
```
#### 安装缓存服务（另一个微服务）

没有它，调用另一个微服务部分将无法正常工作，其余部分工作正常。请按照[reservegrpc](https://github.com/jfeng45/reservegrpc)中的说明设置服务。

### 启动应用程序

#### 启动MySQL
```
cd [MySQLroot]/bin
mysqld
```

#### 启动CouchDB
```
它应该已经启动了
```
#### 启动缓存服务

请按照[reservegrpc](https://github.com/jfeng45/reservegrpc)中的说明启动服务器。

#### 运行main

##### 作为本地应用程序运行

在“main.go”的“main（）”函数中，有两个函数“testMySql（）”和“testCouchDB（）”。
“testMySql（）”从“configs/appConifgDev.yaml”读取配置并访问MySQL。 “testCouchDB（）”从“configs/appConifgProd.yaml”读取配置并访问CouchDB。
“testMySql（）”中有多个函数，你可以通过注释掉其他函数来单独测试一个函数。

```
cd [rootOfProject]/cmd
go run main.go
```
##### 作为gRPC微服务应用程序运行

启动gRPC服务器
```
cd [rootOfProject]/cmd/grpcserver
go run grpcServerMain.go
```
启动gRPC客户端
```
cd [rootOfProject]/cmd/grpcclient
go run grpcClientMain.go
```

### 授权

[MIT](LICENSE.txt) 授权


