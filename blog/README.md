# Server
- install MongoDB server community
```shell
brew tap mongodb/brew # add reposity source 
brew install mongodb-community # install 
brew services start mongodb-community # start service
brew services list # view service status
```


```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative blog/blog_pb/blog.proto
```


- mongodb driver for golang
```shell
go get -t go.mongodb.org/mongo-driver/mongo
```
output:
```shell
go: downloading go.mongodb.org/mongo-driver v1.7.2
go: downloading github.com/pkg/errors v0.9.1
go: downloading github.com/go-stack/stack v1.8.0
go: downloading golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
go: downloading golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
go: downloading github.com/golang/snappy v0.0.1
go: downloading github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d
go: downloading github.com/klauspost/compress v1.9.5
go: downloading github.com/xdg-go/scram v1.0.2
go: downloading github.com/xdg-go/stringprep v1.0.2
go: downloading github.com/xdg-go/pbkdf2 v1.0.0
go: downloading golang.org/x/text v0.3.5
go get: added github.com/go-stack/stack v1.8.0
go get: added github.com/golang/snappy v0.0.1
go get: added github.com/klauspost/compress v1.9.5
go get: added github.com/pkg/errors v0.9.1
go get: added github.com/xdg-go/pbkdf2 v1.0.0
go get: added github.com/xdg-go/scram v1.0.2
go get: added github.com/xdg-go/stringprep v1.0.2
go get: added github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d
go get: added go.mongodb.org/mongo-driver v1.7.2
go get: upgraded golang.org/x/sync v0.0.0-20190423024810-112230192c58 => v0.0.0-20190911185100-cd5d95a43a6e
go get: upgraded golang.org/x/text v0.3.0 => v0.3.5
```
```shell
go get gopkg.in/mgo.v2/bson
```
output:
```shell
go: downloading gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
go get: added gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
```
- update `blog.proto` to add service rpc
- then re-generate grpc
- then implement the server to handle create blog request

# Client
- create client connection
- prepare blog data
- create blog
```shell
go run blog/blog_client/client.go
```
output:
```shell
2021/10/03 13:28:11 Running blog client
2021/10/03 13:28:11 blog has been created: blog:{id:"61594d7bcd817d3d629f5138"  author:"Tony"  title:"My Second Blog"  content:"My Second blog content"}
```
# Database
- connect to MongoDB
- open Terminal and run `mongosh`
```shell
mongosh
```
- show databases
```shell
show databases;
```
output:
```shell
admin     41 kB
config   111 kB
local   73.7 kB
mydb    49.2 kB
```
- switch to database mydb
```shell
use mydb;
```
output:
```shell
switched to db mydb
```
- show collections:
```shell
show collections;
```
output:
```shell
blog
```
- query data from collection ```blog```
```shell
db.blog.find();
```
output:
```shell
[
  {
    _id: ObjectId("615946967829cff00b77ff5f"),
    author_id: 'Tony',
    content: 'My first blog content',
    title: 'My First Blog'
  },
  {
    _id: ObjectId("615946ac7829cff00b77ff60"),
    author_id: 'Tony',
    content: 'My Second blog content',
    title: 'My Second Blog'
  }
]
```