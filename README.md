# hopen
Golang 极速开发框架 Rapid development framework
#how to use

##main
{ {{ 
func init(){  
	//hopen.AddRouter("/test/:id([0-9]+)/sss",&testController.TestController{},"get:tohtml");
	//hopen.AddAutoRouter("/test/:id([0-9]+)/tohtml",&testController.TestController{});
    hopen.AddPrefixAutoRouter("/test",&testController.TestController{});
}
func main(){
  hopen.Run()
}
 }} }
##controller

type TestController struct {
	hopen.Controller
}

func (t *TestController) Sayhello() {
	print(t.R.Form.Get("id"))
}
func (t *TestController) ToJson() {
	m := make(map[string]string)
	m["show_branch"] = "false"
	m["t0"] = "true"
	m["t1"] = "true"
	t.SetValue("data", m)
	t.SetValue("msg", "测试")
	t.SetValue("status", "测试")
	t.RenderJson()
}

type Servers struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func (t *TestController) ToXml() {
	v := &Servers{ServerName:"2",ServerIP:"3"}
	t.RenderXml(v)
}
func (t *TestController) ToHtml() {
	id ,_ := t.GetI("id",-1);
	t.SetValue("msg", "id is " + strconv.FormatInt(int64(id), 10))
	t.Render("tmpl/welcome.html")
}
func (t *TestController) RedirectTo() {
      t.Redirect("tojson")
}
