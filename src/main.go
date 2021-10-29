package main

import (
	"fmt"
	"abc.com/flagT"
    "abc.com/dbcon"
)

func init(){
    dbcon.DBT()
}

func main(){

	fmt.Println("Hello world")
    flagT.GoflagT()

    // iris.New() 和 iris.Default()
    // iris.New() -> Creates an iris application without any middleware by default
    //app := iris.New()
    app := iris.Default()

    //傳入index方法
    app.Get("/",index)

    //建立 localhost:8090/user/{name}路徑，並使用匿名方法直接處理
    app.Get("/user/{name}", func(ctx iris.Context) {
         name := ctx.Params().Get("name")
         ctx.Writef("Hello %s", name) //顯示 "hello 某某某"字樣
    })

   //建立 localhost:8090/boss/{name}路徑
    app.Get("/boss/{name}", func(ctx iris.Context) {
        name := ctx.Params().Get("name")
        ctx.Writef("Hello %s", name)
    })

    //建立 localhost:8090/welcome?firstname=Jane&lastname=Doe 路徑
    app.Get("/welcome", func(ctx iris.Context) {
        //取得GET值
        // URLParamDefault ("目標名稱","如果不存在 預設值" )
        firstname := ctx.URLParamDefault("firstname", "g")
        lastname := ctx.URLParam("lastname")

        ctx.Writef("Hello %s %s", firstname, lastname)
    })

    //建立 POST localhost:8090/form_post 路徑
    app.Post("/form_post", func(ctx iris.Context) {
        message := ctx.PostValue("message")
        nick := ctx.PostValueDefault("nick", "anonymous")

        ctx.JSON(iris.Map{
            "nick":    nick,
            "message": message,
            "status":  "posted",

        })
    })

    //群組路徑 localhost:8090/v1/loginApp or /loginApp2
    //並送入對應 func
    v1 := app.Party("/v1")
    {
        v1.Post("/loginApp", loginEndpoint)
        v1.Post("/loginApp2", loginEndpoint)
    }

}


func index(ctx iris.Context){
    //回傳json型態
    ctx.JSON(iris.Map{
        "status":  "posted",
    })
}

func loginEndpoint(ctx iris.Context){

    message := ctx.PostValue("message")
    nick := ctx.PostValueDefault("nick", "anonymous")

    ctx.JSON(iris.Map{
        "nick":    nick,
        "message": message,
        "status":  "posted",
    })
}