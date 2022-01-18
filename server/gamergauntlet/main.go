package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/go-ps"
)

// Gamer Gaunlet Site information
type site struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	IP    string `json:"ip"`
	Port  string `json:"port"`
	Theme string `json:"theme"`
}

// Gamer Gauntlet server page data
type page struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Site     string `json:"site"`
	Method   string `json:"method"`
	Path     string `json:"path"`
	Handler  string `json:"handler"`
	Template string `json:"template"`
}

//Gamer Gauntlet user data
type user struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Image     string `json:"image"`
	Key       string `json:"key"`
}

//Gamer Gauntlet server data
type gamergauntlet struct {
	ID   string `json:"id"`
	IP   string `json:"ip"`
	MAC  string `json:"mac"`
	Key  string `json:"key"`
	User string `json:"user"`
}

// User profile data
type profile struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	User  string `json:"user"`
}

// Profile screen layout data
type screen struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Profile string `json:"profile"`
}

// Screen button data
type button struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Image    string `json:"image"`
	Type     string `json:"type"`
	Screen   string `json:"screen"`
	Size     int    `json:"size"`
	Position string `json:"position"`
}

// Screen widget data
type widget struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Image    string `json:"image"`
	Type     string `json:"type"`
	Screen   string `json:"screen"`
	Size     int    `json:"size"`
	Position string `json:"position"`
}

//Below is Seed data until a database is added

// Default site
var sites = []site{
	{ID: "0", Title: "Gamer Gauntlet", IP: "192.168.1.70", Port: "80", Theme: "Default"},
}

// Default Pages
var pages = []page{
	{Title: sites[0].Title + " - Home", Site: sites[0].ID, Method: "GET", Path: "/", Template: "index.tmpl"},
	{Title: sites[0].Title + " - Users", Method: "GET", Path: "/users", Template: "users.tmpl"},
	{Title: sites[0].Title + " - Settings", Path: "/settings", Template: "settings.tmpl"},
}

// Test users
var users = []user{
	{ID: "1", Username: "Demo", Email: "demo@gmail.com", FirstName: "Demo", LastName: "Demo", Image: "Default", Key: "GamerGauntletDemo"},
	{ID: "2", Username: "Nicki", Email: "nickitaylor@gmail.com", FirstName: "Nicki", LastName: "Taylor", Image: "Default", Key: "GamerGauntletNickiTaylor"},
}

var gamergauntlets = []gamergauntlet{
	{ID: "1", IP: "IP", MAC: "MAC", Key: "GamerGauntlet", User: "1"},
}

var profiles = []profile{
	{ID: "1", Name: "Streaming", Image: "Default", User: "1"},
	{ID: "2", Name: "Recording", Image: "Default", User: "1"},
}

var screens = []screen{
	{ID: "1", Name: "Primary", Profile: "1"},
	{ID: "2", Name: "Secondary", Profile: "1"},
}

var buttons = []button{
	{ID: "1", Name: "mic1", Image: "Default", Type: "MuteAudio"},
	{ID: "2", Name: "1", Image: "Default", Type: "SwitchScene"},
	{ID: "3", Name: "2", Image: "Default", Type: "SwitchScene"},
	{ID: "4", Name: "3", Image: "Default", Type: "SwitchScene"},
	{ID: "3", Name: "cam1", Image: "Default", Type: "HideSource"},
	{ID: "4", Name: "", Image: "Default", Type: "ToggleStream"},
	{ID: "5", Name: "", Image: "Default", Type: "ToggleRecording"},
	{ID: "6", Name: "", Image: "Default", Type: "PauseRecording"},
	{ID: "7", Name: "", Image: "Default", Type: "Disconnect"},
	{ID: "7", Name: "", Image: "Default", Type: "Disconnect"},
}

var widgets = []widget{
	{ID: "1", Name: "Twitch", Image: "Default", Type: "Stream"},
	{ID: "2", Name: "Twitch", Image: "Default", Type: "Stream"},
	{ID: "3", Name: "Youtube", Image: "Default", Type: "Chat"},
	{ID: "4", Name: "Youtube", Image: "Default", Type: "Chat"},
}

func main() {

	server("gg", "start")
	server("obs", "status")

}

func server(name string, option string) {

	if name == "gg" {
		gg("start")
	}

	if name == "obs" {
		obs("status")
	}
}

// Manage Gamer Gauntlet Server
func gg(option string) int {

	//Check if Gamer Gauntlet server is running
	if option == "status" {
		status := getProcess("go.exe")
		return status
	}

	if option == "start" {
		buildsite()

		return 1
	}

	return 0
}

// Manage OBS server
func obs(option string) int {
	//Check if OBS is running
	if option == "Status" {
		status := getProcess("obs64.exe")
		return status
	}
	return 0
}

func buildsite() {
	getPages(pages)

}

func getPages(pages []page) {
	for _, page := range pages {
		getRoutes(page.Method, page.Path, page.Handler)
	}

}

func getRoutes(method string, path string, handler string) {
	router := gin.Default()

	// Load page templates
	router.LoadHTMLGlob("templates/**/*.tmpl")

	router.Static("/assets", "./assets")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, pages[0].Template, gin.H{
			"title": pages[0].Title,
			"Items": []string{"Home", "Users"},
		})
	})
	router.GET("/users", func(c *gin.Context) {
		c.HTML(http.StatusOK, pages[1].Template, gin.H{
			"title": pages[1].Title,
			"Items": []string{"Home", "Users"},
		})
	})
	router.GET("/settings", getSettings)
	router.GET("/settings/:id", getSettingID)
	router.POST("/settings", postSettings)
	router.Run(sites[0].IP + ":" + sites[0].Port)
}

// getSettings.
func getSettings(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, widgets)
	c.IndentedJSON(http.StatusOK, widgets)
	c.IndentedJSON(http.StatusOK, widgets)
	c.IndentedJSON(http.StatusOK, widgets)
	c.IndentedJSON(http.StatusOK, buttons)
	c.IndentedJSON(http.StatusOK, widgets)
}

// postSettings
func postSettings(c *gin.Context) {
	var newButton button

	if err := c.BindJSON(&newButton); err != nil {
		return
	}

	// Add the new button to the slice.
	buttons = append(buttons, newButton)
	c.IndentedJSON(http.StatusCreated, newButton)
}

// getSettingID
func getSettingID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range buttons {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "404"})
}

func getProcess(name string) int {

	var status int = 0

	procs, err := ps.Processes()
	if err != nil {
		log.Printf("Error: " + err.Error())
	}

	for _, proc := range procs {
		if proc.Executable() == name {
			status = 1
		}
	}

	return status
}
