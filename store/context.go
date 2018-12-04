package store

import "github.com/labstack/echo"

const key = "store"

type Setter interface {
	Set(string, interface{})
}

func FromContext(c echo.Context) Store {
	return c.Get(key).(Store)
}

func ToContext(c Setter, store Store) {
	c.Set(key, store)
}
