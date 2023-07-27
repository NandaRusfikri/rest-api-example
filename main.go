package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var (
	SERVICE_NAME = ""
	VERSION      = "1.0.9"
	PORT         = 0
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	serviceName := flag.String("name", "API-Example", "Service Name")
	port := flag.Int("p", 54321, "Port")

	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatal("main() - time - error: ", err)
	}
	time.Local = loc
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
	router.POST("/free_json", FreeJson)
	router.GET("/user/list", ListUser)
	router.GET("/files", ListFile)
	router.Static("/storage", "./storage")

	router.Run(fmt.Sprintf(":%v", *port))

}

type InputJson interface{}

func FreeJson(c *gin.Context) {
	var Input InputJson
	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request : %v", err.Error()),
		})
		return
	}

	e, err := json.Marshal(Input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Json INPUT : ", string(e))

	jsonData := map[string]interface{}{
		"message":    "Success",
		"your_input": Input,
	}

	c.JSON(http.StatusOK, jsonData)
}

func Default(c *gin.Context) {
	now := time.Now()
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
func ListFile(c *gin.Context) {

	dirPath := "./" // Ubah dengan path direktori yang ingin Anda daftar file-filenya

	fileList := []string{}

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read directory",
		})
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue // Jika ingin mengabaikan folder, lewati iterasi ini
		}

		filePath := filepath.Join(dirPath, file.Name())
		fileList = append(fileList, filePath)
	}

	c.JSON(http.StatusOK, gin.H{
		"files": fileList,
	})
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
