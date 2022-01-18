package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/go-ps"
)

// Gamer Gaunlet Site information
type site struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	IP    string `json:"ip"`
	Port  int    `json:"port"`
	Theme string `json:"theme"`
}

// Gamer Gauntlet server page data
type page struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Site     int    `json:"site"`
	Method   string `json:"method"`
	Path     string `json:"path"`
	Handler  string `json:"handler"`
	Template string `json:"template"`
}

//Gamer Gauntlet user data
type user struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Image     string `json:"image"`
	Key       string `json:"key"`
}

//Gamer Gauntlet server data
type gamergauntlet struct {
	ID   int    `json:"id"`
	IP   string `json:"ip"`
	MAC  string `json:"mac"`
	Key  string `json:"key"`
	User int    `json:"user"`
}

// User app profile data
type profile struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	User  int    `json:"user"`
}

// App profile screen layout data
type screen struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Profile string `json:"profile"`
}

// App screen button data
type button struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Image    string `json:"image"`
	Type     string `json:"type"`
	Screen   int    `json:"screen"`
	Size     int    `json:"size"`
	Position int    `json:"position"`
}

// App screen widget data
type widget struct {
	ID       int    `json:"id"`
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
	{ID: 0, Title: "Gamer Gauntlet", IP: "127.0.0.1", Port: 8080, Theme: "Default"},
}

// Default Pages
var pages = []page{
	{Title: sites[0].Title + " - Home", Site: sites[0].ID, Method: "GET", Path: "/", Handler: "getHome", Template: "index.tmpl"},
	{Title: sites[0].Title + " - Users", Site: sites[0].ID, Method: "GET", Path: "/users", Handler: "getUsers", Template: "users.tmpl"},
	{Title: sites[0].Title + " - Settings", Site: sites[0].ID, Method: "GET", Path: "/settings", Handler: "getSettings", Template: "settings.tmpl"},
}

// Test users
var users = []user{
	{ID: 0, Username: "Demo", Email: "demo@gmail.com", FirstName: "Demo", LastName: "Demo", Image: "Default", Key: "GamerGauntletDemo"},
}

var gamergauntlets = []gamergauntlet{
	{ID: 0, IP: "IP", MAC: "MAC", Key: "GamerGauntlet", User: 0},
}

var profiles = []profile{
	{ID: 0, Name: "Streaming", Image: "Default", User: 0},
	{ID: 1, Name: "Recording", Image: "Default", User: 0},
}

var screens = []screen{
	{ID: 0, Name: "Primary", Profile: "1"},
	{ID: 1, Name: "Secondary", Profile: "1"},
}

var buttons = []button{
	{ID: 0, Name: "mic1", Image: "Default", Type: "MuteAudio", Screen: 0, Size: 1, Position: 0},
	{ID: 1, Name: "1", Image: "Default", Type: "SwitchScene", Screen: 0, Size: 1, Position: 1},
	{ID: 2, Name: "2", Image: "Default", Type: "SwitchScene", Screen: 0, Size: 1, Position: 2},
	{ID: 3, Name: "3", Image: "Default", Type: "SwitchScene", Screen: 0, Size: 1, Position: 3},
	{ID: 4, Name: "cam1", Image: "Default", Type: "HideSource", Screen: 0, Size: 1, Position: 4},
	{ID: 5, Name: "", Image: "Default", Type: "ToggleStream", Screen: 0, Size: 1, Position: 5},
	{ID: 6, Name: "", Image: "Default", Type: "ToggleRecording", Screen: 0, Size: 1, Position: 6},
	{ID: 7, Name: "", Image: "Default", Type: "PauseRecording", Screen: 0, Size: 1, Position: 7},
	{ID: 8, Name: "", Image: "Default", Type: "Disconnect", Screen: 0, Size: 1, Position: 8},
	{ID: 9, Name: "", Image: "Default", Type: "Disconnect", Screen: 0, Size: 1, Position: 9},
}

var widgets = []widget{
	{ID: 0, Name: "Twitch", Image: "Default", Type: "Stream"},
	{ID: 1, Name: "Twitch", Image: "Default", Type: "Stream"},
	{ID: 2, Name: "Youtube", Image: "Default", Type: "Chat"},
	{ID: 3, Name: "Youtube", Image: "Default", Type: "Chat"},
}

func main() {

	manageserver("gg", "start")
	manageserver("obs", "status")

}

// Manage Gamer Gauntlet Server.
func manageserver(server string, option string) int {

	// Check if Gamer Gauntlet server is running.
	if server == "gg" && option == "status" {

		status := getProcess("go.exe")

		return status
	}

	// Start Gamer Gauntlet server.
	if server == "gg" && option == "start" {

		status := buildsite()

		return status
	}

	//Check if OBS is running.
	if server == "obs" && option == "Status" {
		status := getProcess("obs64.exe")

		return status
	}

	return 0
}

func buildsite() int {

	getmenu(pages)

	router := gin.Default()

	// Load page templates
	router.LoadHTMLGlob("templates/**/*.tmpl")

	router.Static("/assets", "./assets")

	router.GET("/", getHome)
	router.GET("/users", getUsers)
	router.GET("/settings", getSettings)
	router.GET("/settings/:id", getSettingID)
	router.POST("/settings", postSettings)

	port := strconv.Itoa(sites[0].Port)
	router.Run(sites[0].IP + ":" + port)

	return 1
}

func getmenu(pages []page) []struct{} {

	type item struct {
		ID     int
		Name   string
		Path   string
		Active int
		Order  int
	}
	var items []struct{}

	for _, page := range pages {
		id := 0
		items[id] = []item{{ID: id, Name: page.Title}}
		id++
	}

	return items

}

func getHome(c *gin.Context) {

	c.HTML(http.StatusOK, pages[0].Template, gin.H{

		"title": pages[0].Title,
	})

}

func getUsers(c *gin.Context) {

	c.HTML(http.StatusOK, pages[1].Template, gin.H{
		"title": pages[1].Title,
	})

}

// getSettings
func getSettings(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
	c.IndentedJSON(http.StatusOK, gamergauntlets)
	c.IndentedJSON(http.StatusOK, profiles)
	c.IndentedJSON(http.StatusOK, screens)
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
		if strconv.Itoa(a.ID) == id {
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
