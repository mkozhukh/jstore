// generated with https://github.com/mkozhukh/jstore
// jstore DataItem store/target.go

package store

import (
	"encoding/json"
	"io/ioutil"
)

func NewCollection(path string) (*Collection, error) {
	data := Collection{path: path}
	err := data.loadFromFile()

	return &data, err
}

type Collection struct {
	path  string
	pull  map[uint]*DataItem
	order []DataItem

	counter uint
}

type storeDataItemFormat struct {
	Order   []DataItem
	Counter uint
}

func (c *Collection) Get(id uint) *DataItem {
	item, ok := c.pull[id]
	if !ok {
		return &DataItem{}
	}

	return item
}

func (c *Collection) Exists(id uint) bool {
	_, ok := c.pull[id]
	return ok
}

func (c *Collection) GetAll() []DataItem {
	return c.order
}

type DataItemLocator func(*DataItem) bool

func (c *Collection) First(loc DataItemLocator) *DataItem {
	for i := range c.order {
		if loc(&c.order[i]) {
			return &c.order[i]
		}
	}

	return &DataItem{}
}

func (c *Collection) Save(obj *DataItem) error {
	index := -1
	if obj.ID == 0 {
		c.counter++
		obj.ID = c.counter
	} else {
		index = c.indexOf(obj.ID)
	}

	if index == -1 {
		c.order = append(c.order, *obj)
	} else {
		c.order[index] = *obj
	}

	c.pull[obj.ID] = obj

	return c.saveToFile()
}

func (c *Collection) Delete(id uint) error {
	index := c.indexOf(id)
	if index == -1 {
		return nil
	}

	c.order = append(c.order[:index], c.order[index+1:]...)
	delete(c.pull, id)

	return c.saveToFile()
}

func (c *Collection) indexOf(key uint) int {
	for i := range c.order {
		if c.order[i].ID == key {
			return i
		}
	}

	return -1
}

func (c *Collection) loadFromFile() error {
	c.pull = make(map[uint]*DataItem)

	bytes, err := ioutil.ReadFile(c.path)
	if err != nil {
		return err
	}

	temp := storeDataItemFormat{}
	err = json.Unmarshal(bytes, &temp)
	if err != nil {
		return err
	}

	c.order = temp.Order
	c.counter = temp.Counter
	for i := range c.order {
		c.pull[c.order[i].ID] = &c.order[i]
	}

	return nil
}

func (c *Collection) saveToFile() error {
	bytes, err := json.Marshal(storeDataItemFormat{
		Order:   c.order,
		Counter: c.counter,
	})

	if err != nil {
		return err
	}

	return ioutil.WriteFile(c.path, bytes, 0644)
}
