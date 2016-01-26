package hopen

import (
	"errors"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

var (
	router *RouterRegistor
)

func init() {
	router = &RouterRegistor{}
}
func Run() {
	RunWithPort(":9090")
}
func RunWithPort(port string) {
	http.HandleFunc("/", httpMethod)         //设置访问的路由
	err := http.ListenAndServe(port, nil) //设置监听的端口
	if err != nil {
		panic(err)
	}
}
func httpMethod(w http.ResponseWriter, r *http.Request) {
	requestPath := r.URL.Path
	for _, route := range router.info {
		if !route.regex.MatchString(requestPath) {
			continue
		}
		matches := route.regex.FindStringSubmatch(requestPath)
		if matches == nil || len(matches[0]) != len(requestPath) {
			continue
		}
		//params := make(map[string]string)
		if len(route.params) > 0 {
			values := r.URL.Query()
			for i, match := range matches[1:] {
				values.Add(route.params[i], match)
				//params[route.params[i]] = match
			}
			r.URL.RawQuery = url.Values(values).Encode() // + "&" + r.URL.RawQuery
		}
		r.ParseForm()
		//初始化
		vc := reflect.New(route.controllerType)
		init := vc.MethodByName("Init")
		in := make([]reflect.Value, 2)
		in[0] = reflect.ValueOf(w)
		in[1] = reflect.ValueOf(r)
		init.Call(in)
		var methdName string
		//获取调用的方法
		if !route.urlRouter {
			for _, method := range route.methods {
				var err error
				if (strings.ToUpper(method.methdType) == r.Method || method.methdType == "*") && strings.ToLower(method.methdName) != "init" {
					err, methdName = findMethod(route.controller, method.methdName)
					if err != nil{
						continue
					}
				}

			}
			if methdName == "" {
				http.NotFound(w, r)
				return
			}
		} else {
			var err error
			parts := strings.Split(requestPath, "/")
			orMethodName := parts[len(parts)-1]
			if orMethodName == "" || strings.ToLower(orMethodName) == "init" {
				http.NotFound(w, r)
				break
			}
			err, methdName = findMethod(route.controller, orMethodName)
			if err != nil {
				http.NotFound(w, r)
				break
			}
		}
		methd := vc.MethodByName(methdName)
		methd.Call([]reflect.Value{})

		break
	}
}
func AddRouter(pattern string, c controllerInterface, methdName string) {
	router.Registor(pattern, c, methdName, false)
}

//example. hopen.AddAutoRouter("/test/sayhello",&testController.TestController{});
//   if url is /test/sayhello  Call Sayhello method
func AddAutoRouter(pattern string, c controllerInterface) {
	parts := strings.Split(pattern, "/")
	orMethodName := parts[len(parts)-1]
	router.Registor(pattern, c, "*:"+orMethodName, false)
}

//example. hopen.AddPrefixAutoRouter("/test",&testController.TestController{});
// if url is /test/sayhello Call Sayhello method
func AddPrefixAutoRouter(pattern string, c controllerInterface) {
	pattern = pattern + "/([^/]+)"
	router.Registor(pattern, c, "", true)
}
func findMethod(c controllerInterface, methodName string) (error, string) {
	v := reflect.ValueOf(c)
	vt := v.Type()
	methodName = strings.ToLower(methodName)
	for i := 0; i < vt.NumMethod(); i++ {
		if strings.ToLower(vt.Method(i).Name) == methodName {
			return nil, vt.Method(i).Name
		}
	}
	return errors.New("method not find"), ""
}
