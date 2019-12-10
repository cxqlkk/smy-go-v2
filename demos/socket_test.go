package demos

import (
	cc "context"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/core/host"
	"github.com/kataras/iris/core/router"
	"github.com/kataras/iris/mvc"
	"testing"
	"time"
)

func TestJack(t *testing.T) {
	app := iris.New()
	app.ConfigureHost(func(su *host.Supervisor) {
		su.RegisterOnServe(func(taskHost host.TaskHost) {
			fmt.Println("onServe")
		})
	})
	app.Run(iris.Addr(":8989"))

}

func TestHost(t *testing.T) {
	app := iris.New()
	app.Get("/", func(ctx context.Context) {
		ctx.JSON(iris.Map{"mess": "fuck"})
	})
	iris.RegisterOnInterrupt(func() {
		cx, cancel := cc.WithTimeout(cc.Background(), time.Second*3)
		defer cancel()
		app.Shutdown(cx)
	})
	//在不同的goroutine中运行，以便不阻止主要的“goroutine”。
	go app.Run(iris.Addr(":8080"), iris.WithoutInterruptHandler)

	// 启动第二个服务器，它正在监听tcp 0.0.0.0:9090，
	//没有“go”关键字，因为我们想要在最后一次服务器运行时阻止。
	app.Run(iris.Addr(":9090"), iris.WithoutInterruptHandler)
	//app.NewHost(&http.Server{Addr:":9090"}).ListenAndServe()
}

func TestConfig(t *testing.T) {
	app := iris.New()
	app.PartyFunc("/test", func(p router.Party) {
		p.Get("/", func(i context.Context) {
			i.JSON(iris.Map{"ss": "ss"})
		})
	})
	//1
	app.Configure(iris.WithConfiguration(iris.Configuration{DisableStartupLog: false}))
	//2

	/*app.Run(iris.Addr(":8080"),iris.WithConfiguration(iris.Configuration{
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        "Mon, 02 Jan 2006 15:04:05 GMT",
		Charset:                           "UTF-8",
	}))*/
	//yaml配饰
	app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.YAML("../config/iris.yml")))
}

func TestController(t *testing.T) {
	app := iris.New()

	app.Use(MidTest)
	//mvc.New(app).Handle(new(UserController))
	mvc.Configure(app.Party("/usr"),usrs)
	// 类似 springMVC  类名上的@RequestMapping
	//mvc.New(app.PartyFunc("/group", func(p router.Party) {
		/*p.Get("/fuck", func(i context.Context) {
			i.JSON(iris.Map{"fk":"fk"})
		})*/
	//})).Handle(new(UserController))
	//mvc.New(app.Party("/ss")).Handle(new(UserController))
	//自定义 marco
	app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.YAML("../config/iris.yml")))
}

func usrs(app *mvc.Application){
	fmt.Println("111")
	app.Handle(new(UserController))
}

type UserController struct {
}

// GetHello serves
// Method:   GET
// Resource: http://localhost:8080/hello
func (c *UserController) GetHello() interface{} {
	return map[string]string{"message": "Hello Iris!"}
}

func (c *UserController) GetUser() interface{} {
	return map[string]string{"user": "user: Iris!"}
}


func (c *UserController)ParamTest(ctx iris.Context)interface{}{
	id:=ctx.Params().Get("id")
	return map[string]string{"id":id}
}

// 为方法名 修改 路由别民
func (c *UserController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/self", "GetUser")
	b.Handle("GET","/user/{id:int}","ParamTest")

}


func MidTest(ctx iris.Context){
		fmt.Println("mid")
		ctx.Next()
}

