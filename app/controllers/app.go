package controllers

import (
  "strings"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

const size = 100000000
const endSize = 95.37

func (c App) Test() revel.Result {
  arr := make([]string, size)
  for i := 0; i < size; i++ {
    arr[i] = "0"
  }
  content := strings.Join(arr, "")


  return c.Render(content)
}
