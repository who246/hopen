package hopen

import (
	"reflect"
	"regexp"
	"strings"
)

type routerRegistorInfo struct {
	regex          *regexp.Regexp
	params         map[int]string
	controllerType reflect.Type
	methods        []*method
	urlRouter      bool
	controller     controllerInterface
}

type RouterRegistor struct {
	info []*routerRegistorInfo
}
type method struct {
	methdName string
	methdType string
}

func (this *RouterRegistor) Registor(pattern string, c controllerInterface, methdNames string, urlRouter bool) {
	parts := strings.Split(pattern, "/")
	j := 0
	params := make(map[int]string)
	for i, part := range parts {
		if strings.HasPrefix(part, ":") {
			expr := "([^/]+)"
			if index := strings.Index(part, "("); index != -1 {
				expr = part[index:]
				part = part[1:index]
			}
			params[j] = part
			parts[i] = expr
			j++
		}
	}
	methods := []*method{}
	if !urlRouter {
		methodStr := strings.Split(methdNames, ",")
		for _, str := range methodStr {
			ms := strings.Split(str, ":")
			m := &method{methdType: ms[0], methdName: ms[1]}
			methods = append(methods, m)
		}
	}
		pattern = strings.Join(parts, "/")
		regex, regexErr := regexp.Compile(pattern)
		if regexErr != nil {
			return
		}
	
	t := reflect.Indirect(reflect.ValueOf(c)).Type()
	route := &routerRegistorInfo{}
	route.regex = regex
	route.controllerType = t
	route.params = params
	route.methods = methods
	route.urlRouter = urlRouter
	route.controller = c
	this.info = append(this.info, route)
}
