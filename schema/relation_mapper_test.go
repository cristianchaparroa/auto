package schema

import (
	"fmt"
	"testing"

	"github.com/cristianchaparroa/auto/meta"
)

func TestModelRelationMapperGenerateMap(t *testing.T) {

	// Partial result relates  to Post Detail
	postDetail := &meta.ModelStruct{}
	postDetail.ModelName = "PostDetail"

	prPostDetail := &PartialResult{}
	prPostDetail.Model = postDetail

	fieldPost := &meta.Field{}
	fieldPost.Name = "Post"
	fieldPost.IsRelation = true

	postRelation := &meta.Relation{}
	postRelation.Name = "post_id"
	postRelation.PKRef = "Id"
	postRelation.To = "Post"
	postRelation.Typ = meta.OneToOne

	fieldPost.Relation = postRelation

	prFields := make([]*meta.Field, 0)
	prFields = append(prFields, fieldPost)

	prPostDetail.Relations = prFields
	postDetail.Fields = prFields

	// Partial result relates to Post
	post := &meta.ModelStruct{}
	post.ModelName = "Post"
	prPost := &PartialResult{}

	pkField := &meta.Field{}
	pkField.IsPrimaryKey = true
	pkField.Name = "Id"

	postFields := make([]*meta.Field, 0)
	postFields = append(postFields, pkField)

	post.Fields = postFields
	prPost.Model = post

	results := make([]*PartialResult, 0)
	results = append(results, prPost)
	results = append(results, prPostDetail)

	rm := NewModelRelationMapper()
	rels := rm.GenerateMap(results)

	for _, rel := range rels {
		fmt.Printf("%v \n", rel)
	}

	if len(rels) != 1 {
		t.Errorf("Expected 2 relations but get:%v", len(rels))
	}

}
