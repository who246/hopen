package hopen

import (
	"net/http"
	"encoding/json"
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
func (c *Controller) RenderJson() error{
	c.W.Header().Add("Content-Type", "application/json; charset=utf-8")
	content, err := json.Marshal(c.Data)
	if err != nil {
		http.Error(c.W, err.Error(), http.StatusInternalServerError)
		return err
	}
	c.W.Write(content);
	return nil;
}