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

	auto.Generate(c)  // It'll stop the init of your application if something is wrong
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


### Relations supported



Relation is the representation of relation between 2 models

**1.One to One example:**

Suppose that you have the Post of web page and you have the PostDetail related to this post. The following is the tag statement relation in the model Post:
```go
type Post struct {
	Id string
	PostDetail PostDetail `sql:rel=(type:11, to=PostDetail)`
	// PostDetail PostDetail `sql:rel=(type:11)`
}
```
The following is the tag statament relation int the model PostDetail:

```go
type PostDetail struct {
	Id string
	Post `sql:rel=(type:11, mappedBy=Post, name="post_id")`
}
```

**2. One to Many example:**

Suppose that you have the Post of web page nad you have the Comments related to it, then you need to generate the following that in the One Part (Post)

```go
type Post struct{
	Comments []Comments `sql:rel=(type=1*,to=Comments)`
	// Comments []Comments `sql:rel=(type=1*)`
}
```

**3. Many To One example:**

Back to the Post-Comments example, the model that represent the Many part is Commments, then the tag must be the following:

```go
type Comments struct{
	Post Post `sql:rel=(type=*1,to=Post, name="post_id")`
  // Post Post `sql:rel=(type=*1, name="post_id")`
}
```
In the last statement you can see the <name> property of tag `rel`, it represents the column name for the table generated in the model that represents the Many Part(Coments).

**4. Many To Many example:**

Suppose that you have many  Post of web page and many tags. The post should have zero or many tags

```go
type Post struct {
	Id string
	Tags []Tag `sql:rel=(type:**, to=Tag)`
	// Tags []Tag `sql:rel=(type:**)`
}
```
 The Tag should be part of zero or many Post

 ```go
 type Tag struct {
	    Id string
		Posts []Post `sql:rel=(type:**, to=Post)`
		// Posts []Post `sql:rel=(type:**)`
 }
 ```
 **Note:**
 As you can see in the above examples, is not mandatory `to` attribute in the tag `rel`. If is not present the attribute `to`, it could be able to generate automatically.


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
