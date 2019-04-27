package main

import "github.com/kataras/iris"
import "github.com/kataras/iris/mvc"
// import "strconv"


func main() {
    app := iris.New()

    // 从 "./views" 目录加载HTML模板
    // 模板解析 html 后缀文件
    // 此方式使用 `html/template` 标准包 (Iris 的模板引擎)
    app.RegisterView(iris.HTML("./views", ".html"))

    // 方法：GET
    app.Get("/", func(ctx iris.Context) {
        // {{.message}} 和 "Hello world!" 字符串变量绑定
        ctx.ViewData("message", "Hello world!")
        // 映射 HTML 模板文件路径 ./views/index.html
        ctx.View("index.html")
    })

    mvc.Configure(app.Party("/chat"), myMVC)

    app.Post("/post", func(ctx iris.Context){
        params := ctx.FormValues()

        c, c_exists := params["a"]

        // ctx.JSON(params)
        // ctx.HTML(strconv.FormatBool(c_exists))

        if c_exists == true {
            ctx.JSON(c)
        }
        
    })

    app.Run(iris.Addr(":8080"))
}

func myMVC(app *mvc.Application) {
    // app.Register(...)
    // app.Router.Use/UseGlobal/Done(...)
    app.Handle(new(MyController))
}

type MyController struct {}

func (m *MyController) BeforeActivation(b mvc.BeforeActivation) {
    // b.Dependencies().Add/Remove
    // b.Router().Use/UseGlobal/Done // and any standard API call you already know

    // 1-> Method
    // 2-> Path
    // 3-> The controller's function name to be parsed as handler
    // 4-> Any handlers that should run before the MyCustomHandler
    // b.Handle("GET", "/something/{id:long}", "MyCustomHandler", anyMiddleware...)
}

// GET: http://localhost:8080/chat
func (m *MyController) Get() string { return "Hey!" }