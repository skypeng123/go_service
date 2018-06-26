package models

import (
	_ "log"
)

type Charge struct {
	Id int64
}

func (c *Charge) count() (id int64, err error) {
	return
}
