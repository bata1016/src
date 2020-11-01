package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	dbInit()

	// index
	router.GET("/", func(ctx *gin.Context) {
		todos := dbGetAll()
		ctx.HTML(200, "index.html", gin.H{
			"todos": todos,
		})
	})

	router.POST("/new", func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		dbInsert(text, status)
		ctx.Redirect(302, "/")
	})

	Detail

	router.Run()
}

type Todo struct {
	gorm.Model
	Text   string
	Status string
}

// DB初期化
func dbInit() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず! (dbInit)")
	}
	db.AutoMigrate((&Todo{}))
	defer db.Close()
}

// DB追加
func dbInsert(text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic(("データベース開けず! (dbInsert)"))
	}
	db.Create(&Todo{Text: text, Statud: status})
	defer db.Close()
}

// DB更新
func dbUpdate(id int, text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic(("データベース開けず! (dbInsert)"))
	}
	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

// DB削除
func dbDelete(id int) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic(("データベース開けず! (dbInsert)"))
	}
	var todo Todo
	db.First(&todo)
	db.Delete(&todo)
	db.Close()
}

// DB全部取得
func dbGetAll() []Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic(("データベース開けず! (dbInsert)"))
	}
	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

// db一つ取得
func dbGetOne(id int) Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic(("データベース開けず! (dbInsert)"))
	}
	var todo Todo
	db.First(&todo, id)
	db.Close()
	return todo
}
