package main

import "github.com/kataras/iris"
import "github.com/kataras/iris/mvc"
// import "strconv"
// import "fmt"
import "user"


func main() {
    u := user.NewUser("go!")
    u.Test()

    app := iris.New()

    // 从 "./views" 目录加载HTML模板
    // 模板解析 html 后缀文件
    // 此方式使用 `html/template` 标准包 (Iris 的模板引擎)
    // Reload 方法设置为 true 表示开启开发者模式 将会每一次请求都重新加载 views 文件下的所有模板
    // app.RegisterView(iris.HTML("./views", ".html"))
    // app.RegisterView(iris.HTML("./views", ".html").Reload(true))

    // 方法：GET
    // app.Get("/", func(ctx iris.Context) {
    //     // {{.message}} 和 "Hello world!" 字符串变量绑定
    //     ctx.ViewData("message", "Hello world!")
    //     // 映射 HTML 模板文件路径 ./views/index.html
    //     ctx.View("index.html")
    // })
    app.StaticWeb("/", "./views/index.html")



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

    // 登入
    app.Get("/login", func(ctx iris.Context){
        // ctx.ViewData("message", "Hello world!")
        ctx.View("login.html")
    })
    app.Post("/login", func(ctx iris.Context){
        ctx.JSON(user.Login(ctx))
    })

    // 註冊
    app.Get("/register", func(ctx iris.Context){
        // ctx.ViewData("message", "Hello world!")
        ctx.View("register.html")
    })

    //下面第二参数 表示静态文件在 main.go 同级目录 static 里面
    //第一个参数表示请求路由为static  例如 请求路由是 host/static/xxx.js
    app.StaticWeb("/public", "./public")  

    // app.Post("/register", user.Register)


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
func (m *MyController) Get() string { 
    return "Hey!" 
}