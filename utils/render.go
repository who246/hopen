package utils

import (
	"encoding/json"
	"encoding/xml"
	"html/template"
	"io/ioutil"
	"net/http"
)

var renderCache map[string]*template.Template

func init() {
	renderCache = make(map[string]*template.Template)
}
func RenderJson(w http.ResponseWriter, data interface{}) error {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	content, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	_, err = w.Write(content)
	return err
}

func RenderXml(w http.ResponseWriter, obj interface{}) error {
	w.Header().Add("Content-Type", "application/xml; charset=utf-8")
	content, err := xml.Marshal(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	_, err = w.Write(content)
	return err
}

func Render(w http.ResponseWriter, path string, data interface{}) error {
	var err error
	t, ok := renderCache[path]
	if ok {
		err = t.Execute(w, data)
		return err
	}
	t = template.New(path)
	b, err := ioutil.ReadFile(path)
	s := string(b)
	_, err = t.Parse(s)

	if err != nil {
		return err
	}
	err = t.Execute(w, data)
	if err == nil {
		renderCache[path] = t
	}
	return err
}
 