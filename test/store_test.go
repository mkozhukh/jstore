package test

import (
	"os"
	"testing"

	jstore "github.com/mkozhukh/jstore/store"
	"github.com/stretchr/testify/suite"
)

func TestLocalTestSuite(t *testing.T) {
	suite.Run(t, new(StoreTestSuite))
}

type StoreTestSuite struct {
	suite.Suite
}

func (suite *StoreTestSuite) AfterTest(name, test string) {
	os.RemoveAll("./test.data.json")
}

func (suite *StoreTestSuite) TestCollection() {
	col, err := jstore.NewCollection("./test.data.json")
	suite.Error(err)

	item := jstore.DataItem{}
	err = col.Save(&item)
	suite.Nil(err)
	suite.NotEqual("", item.ID)

	err = col.Save(&jstore.DataItem{ID: 1})
	suite.Nil(err)

	err = col.Save(&jstore.DataItem{ID: 2})
	suite.Nil(err)

	err = col.Save(&jstore.DataItem{ID: 3})
	suite.Nil(err)

	err = col.Save(&jstore.DataItem{ID: 4})
	suite.Nil(err)

	err = col.Delete(2)
	suite.Nil(err)

	col2, err := jstore.NewCollection("./test.data.json")
	suite.Nil(err)

	suite.True(col2.Exists(1))
	suite.True(col2.Exists(3))
	suite.True(col2.Exists(4))
	suite.False(col2.Exists(2))

	suite.True(col.Exists(1))
	suite.True(col.Exists(3))
	suite.True(col.Exists(4))
	suite.False(col.Exists(2))

	suite.Equal(3, len(col2.GetAll()))
	suite.Equal(3, len(col.GetAll()))

	suite.Equal(uint(3), col2.Get(3).ID)
}
