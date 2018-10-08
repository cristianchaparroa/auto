### Auto

Auto is an utility to create tables according with model definitions for you golang projects.

### How to use

1. You should create `auto.yml`  file in the root folder and
after that you can use the tool an run the command

```
cd myproject
auto .
```

2. You should run the tool with parameters

```
auto --driver=postgres --models=path/to/models
```
3.  You sould create `auto.yml`  file  in the root folder and after that you can integrate
in the init of your proyect

```go
func main() {
    auto := NewAuto("path/to/models","postgres")
    auto.switch expression {
    case condition:

    }etAutoUpdate(true)

    auto.Generate() // Stop the init of your application if something is wrong
}
```


#### Configuration

The following is the configuration expected

```yml
driver: postgres
scan:  ./path/to/models
update: auto
```


Colons can be used to align columns.

| Configuration        | Definition   |
| ------------- |:-------------:|
| Driver      | Database type |
| Scan        | Path directory where the models are |
| update        | It allow to identify changes in the models an update the database or not. Posible values for this option: `auto` or `none`  |



### Models configuration

#### Realations suported


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
 - [ ] Support diferences between changes and no recreate all from scratch.(Add,drop, update...)
