# Go Microservice with gRPC

Other language: [中文](README.zh.md)
 
This is a Go Microservice project with gRPC. It tries to find an appropriate application layout for a Go Microservice application. It applied Clean Architecture design and using dependency injection to inject concrete types into each function. 

The following are a series of articles to explain the different areas of the application design:

[Go Microservice with Clean Architecture: Application Layout](https://medium.com/@jfeng45/go-micro-service-with-clean-architecture-application-layout-e6216dbb835a)

[Go Microservice with Clean Architecture: Application Design](https://medium.com/@jfeng45/go-microservice-with-clean-architecture-application-design-68f48802c8f)

[Go Microservice with Clean Architecture: Design Principle](https://medium.com/@jfeng45/go-micro-service-with-clean-architecture-design-principle-118d2eeef1e6)

[Go Microservice with Clean Architecture: Coding Style](https://medium.com/@jfeng45/go-micro-service-with-clean-architecture-coding-style-a4da35a27d90)

[Go Microservice with Clean Architecture: Transaction Support](https://medium.com/@jfeng45/go-microservice-with-clean-architecture-transaction-support-61eb0f886a36)

[Go Microservice with Clean Architecture: Application Logging](https://medium.com/@jfeng45/go-microservice-with-clean-architecture-application-logging-b43dc5839bce)

[Go Micro-service with Clean Architecture: Application Container](https://medium.com/@jfeng45/go-microservice-with-clean-architecture-application-container-477cc3a11e77)

[Go Micro-service with Clean architecture: Dependency Injection](https://medium.com/@jfeng45/go-microservice-with-clean-architecture-dependency-injection-82cbd3ecb9f3)

## How to use this project
This project is best to be used as a basic framework when creating a gRPC Microservice project. It already has rich built-in functionalities and is working, so there is no reason to start from scratch. The goal of the project is to build a flexible framework with basic functions, which can be extended easily. 
The application design followed "SOLID" design principle and Go's concise coding style, so it can be used as a living example of the application design and coding style when you try to enforce them in your code.  

## Use it as a template to start a service project
### Functional Features:
1. Switch persistence layer implementation by changing the configuration file. Currently, it implemented MySQL and CouchDB. (It can be extended to support other SQL or NoSQL databases)
2. Switch logging provider by changing the configuration file. Currently, it implemented ZAP and Logrus. ( It can be extended to support other logging providers, as long as they support common interfaces similar to ZAP and Logrus)
3. Support business layer transaction (It doesn't support nested transaction or transactions across Microservices)  
4. Using Dependency Injection to create concrete types and wire the whole application together.
5. Application configurations are saved in a YAML file. 

### Design Features:
##### 1. Programming on interface 
* Application has three layers: use case, model and persistence. Each layer accesses other layers through interfaces. 
* Outside functions are also accessed through interfaces.
##### 2. Create concrete types through Dependency Injection by using factory method pattern
##### 3. Minimize Dependency
* Dependency between different layers is only on interfaces instead of concrete types.
* Interfaces are defined in top level package and separated from concrete types. 
* Each concrete type is defined in a separate sub-package and file 
##### 4. Function Isolation
* Isolate different layers by package
* Isolate each use case by package 
* Isolate each implementation ( for example database implementation) by package
##### 5. Open-closed principle
* whenever a new feature is added, instead of modifying existing code, try to add new code
  

### Coding Style:
1. Eliminate package level variables except in "container" package
2. Minimize use of constants. 
3. Log full stack trace for errors
4. Errors are only handled on top level 
5. Separation of concerns. 
6. Naming Convention. Short names for local variables , long name for types or interfaces.   

## Getting Started

### Installation and Setting Up

Don't need to finish all steps in this section up-front to get the code up running. The simplest way is to get the code from github and run it and come back to install the part when there is a real need. However, it will encounter an error when accesses the database. So, I'd recommend you install at least one database ( MySQL is better), then most of the code will work. 

#### Download Code

```
go get github.com/jfeng45/servicetmpl
```

#### Set Up MySQL

There are two database implementations, MySQL and CouchDB, but most functions are implemented in MySQL. You'd better install at least one of them. 
```
Install MySQL
run SQL script in script folder to create database and table
```
#### Install CouchDB

The code works fine without it. CouchDB is created to show the feature of switching database by changing configuration.
 
Installation on [Windows](https://docs.couchdb.org/en/2.2.0/install/windows.html)

Installation on [Linux](https://docs.couchdb.org/en/2.2.0/install/unix.html)

Installation on [Mac](https://docs.couchdb.org/en/2.2.0/install/mac.html)

CouchDB [Example](https://github.com/go-kivik/kivik/wiki/Usage-Examples)

#### Set up CouchDB

```
Access Fauxton through browser: http://localhost:5984/_utils/# (login with: admin/admin).
Create new database "service_config" in Fauxton.
Add the following document to the database ( "_id" and "_rev" are generated by database, no need to change it):
{
  "_id": "80a9134c7dfa53f67f6be214e1000fa7",
  "_rev": "4-f45fb8bdd454a71e6ae88bdeea8a0b4c",
  "uid": 10,
  "username": "Tony",
  "department": "IT",
  "created": "2018-02-17T15:04:05-03:00"
}
```
#### Install Cache Service (Another Microservice)

Without it, calling another Microservice piece won't work, the rest of application works fine. Please follow instructions in [reservegrpc](https://github.com/jfeng45/reservegrpc) to set up the service.

### Start Application

#### Start MySQL Server
```
cd [MySQLroot]/bin
mysqld
```

#### Start CouchDB Server
```
It should already have been started
```
#### Start Cache Service

Please follow instructions in [reservegrpc](https://github.com/jfeng45/reservegrpc) to start the server.

#### Run main

##### Run as a local application
In "main()" function of "main.go", there are two functions "testMySql()" and "testCouchDB()". 
"testMySql()" reads configurations from "configs/appConifgDev.yaml" and accesses MySQL. "testCouchDB()" reads from "configs/appConifgProd.yaml" and access CouchDB.
There are multiple functions in "testMySql()", you can focus on testing one each time by commenting out others.
```
cd [rootOfProject]/cmd
go run main.go
```
##### Run as a gRPC Microservice application

Start gRPC Server
```
cd [rootOfProject]/cmd/grpcserver
go run grpcServerMain.go
```
Start gRPC Client
```
cd [rootOfProject]/cmd/grpcclient
go run grpcClientMain.go
```

## License

[MIT](LICENSE.txt) License


