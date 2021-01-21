## From 
Find Tutorials [detail](https://www.yuque.com/techcats/database/wdookd)

## Usage
编译镜像
```bash
docker build -t mongo-replica-insert .
```

参考的输出如下所示
```bash
fengyifeideMacBook-Air:insert fengyfei$ docker build -t mongo-replica-insert .
Sending build context to Docker daemon  3.584kB
Step 1/10 : FROM golang:1.13.1-alpine3.10
 ---> 48260c3da24c
Step 2/10 : WORKDIR /app
 ---> Running in 149be8a6b4dc
Removing intermediate container 149be8a6b4dc
 ---> 8b52d563570a
Step 3/10 : COPY main.go .
 ---> b1dd207d458b
Step 4/10 : ENV GOPROXY "https://goproxy.io"
 ---> Running in 2565bc60b053
Removing intermediate container 2565bc60b053
 ---> 5f4dd7398686
Step 5/10 : RUN env | grep GO
 ---> Running in 7bed5a0f1756
GOPATH=/go
GOPROXY=https://goproxy.io
GOLANG_VERSION=1.13.1
Removing intermediate container 7bed5a0f1756
 ---> 19e431e016ed
Step 6/10 : RUN go mod init 'github.com/silverswords/mongo-insight'
 ---> Running in 51392e80d86f
go: creating new go.mod: module github.com/silverswords/mongo-insight
Removing intermediate container 51392e80d86f
 ---> 32289dfe1ff2
Step 7/10 : RUN go mod tidy
 ---> Running in f77dfbe14b6c
go: finding go.mongodb.org/mongo-driver v1.1.1
go: downloading go.mongodb.org/mongo-driver v1.1.1
go: extracting go.mongodb.org/mongo-driver v1.1.1
go: finding github.com/stretchr/testify v1.4.0
go: downloading github.com/stretchr/testify v1.4.0
go: finding github.com/go-stack/stack v1.8.0
go: finding github.com/golang/snappy v0.0.1
go: finding github.com/google/go-cmp v0.3.1
go: finding github.com/tidwall/pretty v1.0.0
go: finding github.com/xdg/stringprep v1.0.0
go: finding golang.org/x/sync latest
go: finding github.com/xdg/scram latest
go: downloading github.com/golang/snappy v0.0.1
go: downloading github.com/go-stack/stack v1.8.0
go: downloading github.com/google/go-cmp v0.3.1
go: downloading github.com/xdg/stringprep v1.0.0
go: downloading github.com/tidwall/pretty v1.0.0
go: downloading golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
go: downloading github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c
go: extracting github.com/go-stack/stack v1.8.0
go: extracting github.com/stretchr/testify v1.4.0
go: extracting github.com/golang/snappy v0.0.1
go: extracting github.com/google/go-cmp v0.3.1
go: extracting github.com/xdg/stringprep v1.0.0
go: extracting github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c
go: extracting golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
go: extracting github.com/tidwall/pretty v1.0.0
go: downloading gopkg.in/yaml.v2 v2.2.2
go: downloading github.com/davecgh/go-spew v1.1.0
go: downloading github.com/pmezard/go-difflib v1.0.0
go: extracting gopkg.in/yaml.v2 v2.2.2
go: downloading gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405
go: extracting github.com/pmezard/go-difflib v1.0.0
go: extracting github.com/davecgh/go-spew v1.1.0
go: extracting gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405
go: finding golang.org/x/text v0.3.2
go: finding golang.org/x/crypto latest
go: downloading golang.org/x/text v0.3.2
go: downloading golang.org/x/crypto v0.0.0-20191001170739-f9e2070545dc
go: extracting golang.org/x/crypto v0.0.0-20191001170739-f9e2070545dc
go: extracting golang.org/x/text v0.3.2
Removing intermediate container f77dfbe14b6c
 ---> 3888fc235f10
Step 8/10 : RUN go mod download
 ---> Running in 251d2c5c5811
go: finding github.com/davecgh/go-spew v1.1.0
go: finding github.com/pmezard/go-difflib v1.0.0
go: finding github.com/stretchr/objx v0.1.0
go: finding golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3
go: finding golang.org/x/sys v0.0.0-20190412213103-97732733099d
go: finding golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e
go: finding gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405
go: finding gopkg.in/yaml.v2 v2.2.2
Removing intermediate container 251d2c5c5811
 ---> a05c624c6a4d
Step 9/10 : RUN go build -o main .
 ---> Running in 94da112e6d68
Removing intermediate container 94da112e6d68
 ---> 051c01aee617
Step 10/10 : CMD ["./main"]
 ---> Running in 32a511a74bb4
Removing intermediate container 32a511a74bb4
 ---> 73a0740d612b
Successfully built 73a0740d612b
Successfully tagged mongo-replica-insert:latest
```

获取 Docker 网络信息
```bash
 fengyifeideMacBook-Air:insert fengyfei$ docker network ls
NETWORK ID          NAME                              DRIVER              SCOPE
890b8fb02d3f        bridge                            bridge              local
3908e9135cf4        host                              host                local
e96c036b173f        mongo-replica_mongo-replica-net   bridge              local
4b3c31fe2cb4        mongo_default                     bridge              local
7c7d635714e1        none                              null                local
06ced6455d15        strapi-docker_default             bridge              local
```

在 Host 的 Shell 中执行
```bash
docker run --network mongo-replica_mongo-replica-net -d mongo-replica-insert
```

执行完毕后，进入 Mongo Replica 的 Primary 节点，确认
```shell
dev:PRIMARY> show dbs
admin   0.000GB
config  0.000GB
local   0.000GB
simple  0.000GB
dev:PRIMARY> use simple
switched to db simple
dev:PRIMARY> db.article.find()
{ "_id" : ObjectId("5d941e116588cf7cc4af8d3b"), "title" : "T-0", "abstract" : "Abstract 0", "reads" : NumberLong(20) }
{ "_id" : ObjectId("5d941e116588cf7cc4af8d3c"), "title" : "T-1", "abstract" : "Abstract 1", "reads" : NumberLong(21) }
{ "_id" : ObjectId("5d941e116588cf7cc4af8d3d"), "title" : "T-2", "abstract" : "Abstract 2", "reads" : NumberLong(22) }
{ "_id" : ObjectId("5d941e116588cf7cc4af8d3e"), "title" : "T-3", "abstract" : "Abstract 3", "reads" : NumberLong(23) }
{ "_id" : ObjectId("5d941e116588cf7cc4af8d3f"), "title" : "T-4", "abstract" : "Abstract 4", "reads" : NumberLong(24) }
{ "_id" : ObjectId("5d941e116588cf7cc4af8d40"), "title" : "T-5", "abstract" : "Abstract 5", "reads" : NumberLong(25) }
{ "_id" : ObjectId("5d941e116588cf7cc4af8d41"), "title" : "T-6", "abstract" : "Abstract 6", "reads" : NumberLong(26) }
{ "_id" : ObjectId("5d941e116588cf7cc4af8d42"), "title" : "T-7", "abstract" : "Abstract 7", "reads" : NumberLong(27) }
{ "_id" : ObjectId("5d941e116588cf7cc4af8d43"), "title" : "T-8", "abstract" : "Abstract 8", "reads" : NumberLong(28) }
{ "_id" : ObjectId("5d941e116588cf7cc4af8d44"), "title" : "T-9", "abstract" : "Abstract 9", "reads" : NumberLong(29) }
{ "_id" : ObjectId("5d941e116588cf7cc4af8d45"), "title" : "T-10", "abstract" : "Abstract 10", "reads" : NumberLong(30) }
{ "_id" : ObjectId("5d941e116588cf7cc4af8d46"), "title" : "T-11", "abstract" : "Abstract 11", "reads" : NumberLong(31) }
{ "_id" : ObjectId("5d941e116588cf7cc4af8d47"), "title" : "T-12", "abstract" : "Abstract 12", "reads" : NumberLong(32) }
{ "_id" : ObjectId("5d941e116588cf7cc4af8d48"), "title" : "T-13", "abstract" : "Abstract 13", "reads" : NumberLong(33) }
{ "_id" : ObjectId("5d941e116588cf7cc4af8d49"), "title" : "T-14", "abstract" : "Abstract 14", "reads" : NumberLong(34) }
{ "_id" : ObjectId("5d941e116588cf7cc4af8d4a"), "title" : "T-15", "abstract" : "Abstract 15", "reads" : NumberLong(35) }
{ "_id" : ObjectId("5d941e116588cf7cc4af8d4b"), "title" : "T-16", "abstract" : "Abstract 16", "reads" : NumberLong(36) }
{ "_id" : ObjectId("5d941e116588cf7cc4af8d4c"), "title" : "T-17", "abstract" : "Abstract 17", "reads" : NumberLong(37) }
{ "_id" : ObjectId("5d941e116588cf7cc4af8d4d"), "title" : "T-18", "abstract" : "Abstract 18", "reads" : NumberLong(38) }
{ "_id" : ObjectId("5d941e116588cf7cc4af8d4e"), "title" : "T-19", "abstract" : "Abstract 19", "reads" : NumberLong(39) }
Type "it" for more
dev:PRIMARY>
```