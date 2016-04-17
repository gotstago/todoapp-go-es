# todoapp-go-es

This is a sample applicaton of golang and CQRS + ES in a simple todo app. 
```
https://utbrudd.bouvet.no/2015/06/09/a-simple-todo-application-a-comparison-on-traditional-vs-cqrs-es-architecture/
```
## Running with docker


```
https://larry-price.com/blog/2015/06/25/architecture-for-a-golang-web-app
```

```
docker-compose up```

## Running without docker

To start the sample application, download and install the golang development environment. Then proceed with the following:

```
go get github.com/netbrain/todoapp-go-es

cd $GOPATH/github.com/netbrain/todoapp-go-es

go build

./todoapp-go-es
```
and then if your environment has been correctly set up, you can now invoke the app with: 

The application will then be listening @ http://localhost:8080


Also check out my [dlog](http://github.com/netbrain/dlog) project
