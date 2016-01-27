package hopen

import (
	"net/http"
	"github.com/who246/hopen/utils"
	"strconv"
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
func (c *Controller) GetS(key string) string{
	return c.R.Form.Get(key)
}

func (c *Controller) GetI(key string,def int) (int,error){
	if v := c.R.Form.Get(key); v != ""{
		return strconv.Atoi(v)
	}
	return def,nil
}
func (c *Controller) GetI8(key string,def int8) (int8,error){
	if v := c.R.Form.Get(key); v != ""{
		i64,err := strconv.ParseInt(v,10,8)
		i8 := int8(i64)
		return i8,err
	}
	return def,nil
}
func (c *Controller) GetI32(key string,def int32) (int32,error){
	if v := c.R.Form.Get(key); v != ""{
		i64,err := strconv.ParseInt(v,10,32)
		i32 := int32(i64)
		return i32,err
	}
	return def,nil
}
func (c *Controller) GetI64(key string,def int64) (int64,error){
	if v := c.R.Form.Get(key); v != ""{
		return strconv.ParseInt(v,10,64)
	}
	return def,nil
}
func (c *Controller) GetB(key string, def bool) (bool, error) {
	if strv := c.R.Form.Get(key); strv != "" {
		return strconv.ParseBool(strv)
	} else  {
		return def, nil
	}  
}

func (c *Controller) GetF(key string, def float64) (float64, error) {
	if strv := c.R.Form.Get(key); strv != "" {
		return strconv.ParseFloat(strv, 64)
	} else {
		return def, nil
	}  
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
