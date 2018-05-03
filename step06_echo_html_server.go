package main

/*

使學員能夠開發輕量級Http Server

*/

import(
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.File("/", "public/index.html")
	e.GET("/demo", getDemo)
	e.POST("/demo",postDemo)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
	
}

/*
	/demo?name=x-men&email=wolverine
	取得用戶傳來的query string
*/
func getDemo(c echo.Context) error {

	name := c.QueryParam("name")
	email := c.QueryParam("email")
	return c.JSON(http.StatusOK, User{name,email} )

}

/*
	post demo
	設計一個struct，將post過來的json 與此struct的物件進行binding
*/
type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func postDemo(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}