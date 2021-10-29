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


}