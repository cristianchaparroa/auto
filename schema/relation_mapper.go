package schema

import (
	"fmt"

	"github.com/cristianchaparroa/auto/meta"
)

// RelationMap represent the relation between to models
type RelationMap struct {
	Typ        meta.RelationType
	ParentName string
	Parent     *meta.ModelStruct
	ChildName  string
	Child      *meta.ModelStruct
	Field      *meta.Field
}

func (rm *RelationMap) String() string {
	return fmt.Sprintf("RelationMap[ Type:%v, ParentName:%v, Parent:%v, ChildName:%v, Child:%v,Field:%v]",
		rm.Typ, rm.ParentName, rm.Parent, rm.ChildName, rm.Child, rm.Field)
}

// RelationMapper indicates the methods to generate the  Relation Map
type RelationMapper interface {
	GenerateMap(models []*meta.ModelStruct) []*RelationMap
}

// ModelRelationMapper is in charge to build the RelationMap
type ModelRelationMapper struct {
}

// NewModelRelationMapper return a pointer to ModelRelationMapper
func NewModelRelationMapper() *ModelRelationMapper {
	return &ModelRelationMapper{}
}

// GenerateMap generates the diferents relations
func (r *ModelRelationMapper) GenerateMap(results []*PartialResult) []*RelationMap {

	models := make([]*meta.ModelStruct, 0)

	for _, res := range results {
		models = append(models, res.Model)
	}

	mps := make([]*RelationMap, 0)

	for _, res := range results {
		m := res.Model
		parentModel := m.ModelName

		rs := res.Relations

		if rs == nil {
			continue
		}

		for _, field := range rs {
			rel := field.Relation
			childName := rel.To

			rm := &RelationMap{}
			rm.ChildName = childName
			rm.ParentName = parentModel
			rm.Parent = m
			rm.Typ = rel.Typ

			mps = append(mps, rm)
		}
	}

	for _, mp := range mps {
		for _, model := range models {
			if model.ModelName == mp.ChildName {
				mp.Child = model
			}
		}
	}

	return mps
}
