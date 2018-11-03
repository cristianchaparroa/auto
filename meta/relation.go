package meta

const (
	// OneToOne type relation representation
	OneToOne TypeRelation = "11"

	// OneToMany type relation representation usde in the model that represents
	// the One part
	OneToMany TypeRelation = "1*"

	// ManyToOne type relation representation used in the model  that represents
	// the many part
	ManyToOne TypeRelation = "*1"

	// ManyToMany type relation representation
	ManyToMany TypeRelation = "**"
)

//TypeRelation indicates the type of relation between models
type TypeRelation string

// Relation is the representation of relation between 2 models
//
// 1. One to One example:
// Suppose that you have the Post of web page and you have the PostDetail
// related to this post.
// The following is the tag statament relation in the model Post:
//
// type Post struct {
//    Id string
//    PostDetail PostDetail `sql:rel=(type:11, to=PostDetail)`
//    PostDetail PostDetail `sql:rel=(type:11)`
// }
//
// The following is the tag statament relation int the model PostDetail:
// type PostDetail struct {
//    Id string
//    Post `sql:rel=(type:11, mappedBy=Post, name="post_id")`
// }
//
// 2. One to Many example:
// suppose that you have the Post of web page nad you have the Comments
// related to it, then you need to generate the following that in the One
// Part (Post)
// type Post struct{
//    Comments []Comments `sql:rel=(type=1*,to=Comments)`
//    Comments []Comments `sql:rel=(type=1*)`
// }
//
// 3. Many To One example:
// Back to the Post-Comments example, the model that represent the Many part
// is Commments, then the tag must be the following:
//
// type Comments struct{
//    Post Post `sql:rel=(type=*1,to=Post, name="post_id")`
//    Post Post `sql:rel=(type=*1, name="post_id")`
// }
// In the last statement you can see the <name> property of tag rel, it represents
// the column name for the table generated in the model that represents the
// Many Part(Coments).
//
// 4. Many To Many example:
// suppose that you have many  Post of web page and many tags
//
// The post should have zero or many tags
// type Post struct {
//    Id string
//    Tags[] Tag `sql:rel=(type:**, to=Tag)`
//    Tags[] Tag `sql:rel=(type:**)`
// }
//
// The Tag should be part of zero or many Post
// type Tag struct {
//    Id string
//    Posts[] Post `sql:rel=(type:**, to=Post)`
//    Posts[] Post `sql:rel=(type:**)`
// }
// Note:
// As you can see in the above examples, is not mandatory `to` attribute in
// the tag `rel`. If is not present the attribute `to`, it could be able to
// generate automatically.
type Relation struct {

	// Type Relation
	Typ TypeRelation

	// Name, column name is the name of relation in database if is necesary
	Name string

	// To indicates the direction of relation
	To string

	MappedBy string
}
