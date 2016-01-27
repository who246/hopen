package hopen

import (
	"net/http"
	"github.com/who246/hopen/utils"
)

type controllerInterface interface {
	Init(w http.ResponseWriter, r *http.Request)
}
type Controller struct {
	W    http.ResponseWriter
	R    *http.Request
	Data map[string]interface{}
}

func (c *Controller) Init(w http.ResponseWriter, r *http.Request) {
	c.W = w
	c.R = r
	c.Data = make(map[string]interface{})
}

func (c *Controller) SetValue(key string, obj interface{}) {
	c.Data[key] = obj
}
func (c *Controller) RenderJson() error {
	return utils.RenderJson(c.W,c.Data)
}
func (c *Controller) RenderXml(obj interface{}) error {
 
	return utils.RenderXml(c.W,obj)
}
func (c *Controller) Render(path string) error {
	return utils.Render(c.W,path,c.Data)
}
