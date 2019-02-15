package jstore

import (
	"encoding/json"
	"io/ioutil"

	"github.com/rs/xid"
)

func NewCollection(path string) (*Collection, error) {
	data := Collection{path: path}
	err := data.loadFromFile()

	return &data, err
}

type Item struct {
	ID string
}

type Collection struct {
	path  string
	pull  map[string]*Item
	order []Item
}

func (c *Collection) Get(id string) *Item {
	item, ok := c.pull[id]
	if !ok {
		return &Item{}
	}

	return item
}

func (c *Collection) Exists(id string) bool {
	_, ok := c.pull[id]
	return ok
}

func (c *Collection) GetAll() []Item {
	return c.order
}

func (c *Collection) Save(obj *Item) error {
	index := -1
	if obj.ID == "" {
		obj.ID = xid.New().String()
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

func (c *Collection) Delete(id string) error {
	index := c.indexOf(id)
	if index == -1 {
		return nil
	}

	c.order = append(c.order[:index], c.order[index+1:]...)
	delete(c.pull, id)

	return c.saveToFile()
}

func (c *Collection) indexOf(key string) int {
	for i := range c.order {
		if c.order[i].ID == key {
			return i
		}
	}

	return -1
}

func (c *Collection) loadFromFile() error {
	c.order = make([]Item, 0)
	c.pull = make(map[string]*Item)

	bytes, err := ioutil.ReadFile(c.path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &c.order)
	if err != nil {
		return err
	}

	for i := range c.order {
		c.pull[c.order[i].ID] = &c.order[i]
	}

	return nil
}

func (c *Collection) saveToFile() error {
	bytes, err := json.Marshal(c.order)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(c.path, bytes, 0644)
}
