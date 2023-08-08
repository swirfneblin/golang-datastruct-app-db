## GoLang Data structures, Database, API


To run Mysql server:
```
docker-compose up -d
docker-compose stop
```

Watch application and restart when it has changes
```
npm install -g nodemon
nodemon --exec go run main.go --signal SIGTERM
```