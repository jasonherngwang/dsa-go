package structures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BSTTestSuite struct {
	suite.Suite
	EmptyTree      *BSTNode[int]
	SingleNodeTree *BSTNode[int]
	Tree           *BSTNode[int]
}

func (suite *BSTTestSuite) SetupTest() {
	suite.EmptyTree = &BSTNode[int]{}
	suite.SingleNodeTree = &BSTNode[int]{}
	suite.Tree = &BSTNode[int]{
		Val: 5,
		Left: &BSTNode[int]{
			Val:   1,
			Right: &BSTNode[int]{Val: 3},
		},
		Right: &BSTNode[int]{
			Val:  10,
			Left: &BSTNode[int]{Val: 7},
		},
	}
}

// func (suite *BSTTestSuite) TestBSTSearch() {
// 	actual := true
// 	expected := suite.Tree.Search(7)
// 	suite.Equal(expected, actual)
// }

func (suite *BSTTestSuite) TestBSTInsert() {
	expected := &BSTNode[int]{
		Val: 5,
		Left: &BSTNode[int]{
			Val:   1,
			Right: &BSTNode[int]{Val: 3},
		},
		Right: &BSTNode[int]{
			Val:   10,
			Left:  &BSTNode[int]{Val: 7},
			Right: &BSTNode[int]{Val: 12},
		},
	}
	suite.Tree.Insert(12)
	suite.Equal(expected, suite.Tree)
}

func TestBSTTestSuite(t *testing.T) {
	suite.Run(t, new(BSTTestSuite))
}
