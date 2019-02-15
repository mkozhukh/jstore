package test

import (
	"os"
	"testing"

	"github.com/mkozhukh/jstore"
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

	item := jstore.Item{}
	err = col.Save(&item)
	suite.Nil(err)
	suite.NotEqual("", item.ID)

	err = col.Save(&jstore.Item{ ID:"1234"})
	suite.Nil(err)

	err = col.Save(&jstore.Item{ ID:"efgh"})
	suite.Nil(err)

	err = col.Save(&jstore.Item{ ID:"abcd"})
	suite.Nil(err)

	err = col.Save(&jstore.Item{ ID:"abcd"})
	suite.Nil(err)

	err = col.Delete("efgh")
	suite.Nil(err)

	col2, err := jstore.NewCollection("./test.data.json")
	suite.Nil(err)

	suite.True( col2.Exists("1234") )
	suite.True( col2.Exists("abcd") )
	suite.False( col2.Exists("efgh") )

	suite.Equal(3, len(col2.GetAll()))

	suite.Equal("abcd", col2.Get("abcd").ID)
}

