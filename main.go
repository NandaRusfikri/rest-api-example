package main

import (
	"flag"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"runtime"
	"time"
)

var (
	SERVICE_NAME = ""
	VERSION      = "1.0.3"
	PORT         = 0
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	serviceName := flag.String("name", "API-Example", "Service Name")
	port := flag.Int("p", 54321, "Port")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Penggunaan: %s [OPTIONS]\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Version APP :", VERSION)
		fmt.Fprintln(os.Stderr, "OPTIONS:")
		flag.PrintDefaults()
	}

	flag.Parse()
	if flag.Arg(0) == "--help" {
		flag.Usage()
		return
	}
	if serviceName != nil && *serviceName != "" {
		SERVICE_NAME = *serviceName
	}
	if port != nil && *port != 0 {
		PORT = *port
	}

	fmt.Println("Service Name:", *serviceName)
	fmt.Println("Port :", *port)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	router.GET("/", Default)
	router.POST("/login", login)
	router.GET("/user/list", ListUser)

	router.Run(fmt.Sprintf(":%v", *port))

}

func Default(c *gin.Context) {
	loc, errLoc := time.LoadLocation("Asia/Jakarta")
	if errLoc != nil {
		fmt.Println("error load location ==> ", errLoc.Error())
	}
	fmt.Println("loc ", loc)
	now := time.Now().In(loc)
	jsonData := map[string]interface{}{
		"service_name": SERVICE_NAME,
		"version":      VERSION,
		"port":         PORT,
		"time":         now,
		"os":           runtime.GOOS,   // "linux
		"arch":         runtime.GOARCH, // "amd64
	}

	c.JSON(http.StatusOK, jsonData)
}
func login(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Validasi username dan password
	if user.Username == "admin" && user.Password == "admin123" {
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

type UserList struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ListUser(c *gin.Context) {
	Listuser := []UserList{
		{Id: 1, Name: "Mama Lemon", Email: "mamalemon@gmail.com"},
		{Id: 2, Name: "Sukro Duo Kelinci", Email: "duakelinci@gmail.com"},
		{Id: 3, Name: "Sari Roti", Email: "sariroti@gmail.com"},
		{Id: 4, Name: "Teh Kotak", Email: "tehkotak@gmail.com"},
	}
	jsonData := map[string]interface{}{
		"data": Listuser,
	}

	c.JSON(http.StatusOK, jsonData)
}
