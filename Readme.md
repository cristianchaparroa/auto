### Auto

Auto is an utility to create tables according with model definitions for you golang projects.

### How to use

1. You should create `config.yml`  file in the root folder

```go

func main() {

	auto := auto.NewGenerator()
	c, err := config.NewConfig("config.yaml")

	if err != nil {
		panic(err)
	}

	auto.Generate(c) // It'll stop the init of your application if something is wrong
}
```



#### Configuration

The following is and example for configuration file.

```yaml
scan: models #if this file is in the root it'll for a root called models
driver: postgres
host: localhost
port: 5432
schema: userlist
database: list
user: userlist
password: userlist
```


### Models configuration

### Testing

#### Integration
```
 go test -tags=integration -v
 ```



### RoadMap

 - [ ] Parsers for tags
 - [x] Parsers for Fields
 - [x] Parser for structs
 - [x] Scanner for models
 - [ ] Create  cmd tool to integrate all flow
 - [ ] Generator tables from models parsed
 - [ ] Generator of relations between models parsed  
 - [ ] Suport Postgres generation
 - [ ] Suport Mysql generation
 - [ ] Suport Oracle generation
 - [ ] Support changes in the schema  (Add,drop, update...)
