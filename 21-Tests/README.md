```
go test Address/Address --cover
go test Address/Address --coverprofile cobertura.txt 
go tool cover --func=cobertura.txt
go tool cover --html=cobertura.txt
```
