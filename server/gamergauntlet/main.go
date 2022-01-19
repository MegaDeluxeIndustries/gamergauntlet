package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/go-ps"
)

// Gamer Gaunlet server
type site struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	IP    string `json:"ip"`
	Port  int    `json:"port"`
	Theme string `json:"theme"`
}

// Gamer Gauntlet server page
type page struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Site     int    `json:"site"`
	Method   string `json:"method"`
	Path     string `json:"path"`
	Handler  string `json:"handler"`
	Template string `json:"template"`
}

// Gamer Gauntlet server link
type link struct {
	ID     int
	Name   string
	Path   string
	Active int
	Order  int
}

// Gamer Gauntlet Server menu
type menu struct {
	ID    int
	Name  string
	Links []link
}

// Gamer Gauntlet user
type user struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Image     string `json:"image"`
	Key       string `json:"key"`
}

// Gamer Gauntlet app
type gamergauntlet struct {
	ID   int    `json:"id"`
	IP   string `json:"ip"`
	MAC  string `json:"mac"`
	Key  string `json:"key"`
	User int    `json:"user"`
}

// Gamer Gauntlet app profile
type profile struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	User  int    `json:"user"`
}

//Gamer Gauntlet app screen
type screen struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Profile string `json:"profile"`
}

// Gamer Gauntlet app button
type button struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Image    string `json:"image"`
	Type     string `json:"type"`
	Screen   int    `json:"screen"`
	Size     int    `json:"size"`
	Position int    `json:"position"`
}

// Gamer Gauntlet app widget
type widget struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Image    string `json:"image"`
	Type     string `json:"type"`
	Screen   string `json:"screen"`
	Size     int    `json:"size"`
	Position string `json:"position"`
}

// TODO Implement r/w Gamer Gaunlet server app settings file.
// The data below will be generated and modified using the Gamer Gauntlet server user interface, and saved locally as json in text files.
// When the Gamer Gauntlet app starts, it connects to the server, pulls settings down, and configures the app user interface and functionality.

// Begin Gamer Gauntlet server settings
// Default site
var sites = []site{
	{ID: 0, Title: "Gamer Gauntlet", IP: "127.0.0.1", Port: 8080, Theme: "Default"},
}

// Default Pages
// This is where you can add or modify the Gamer Gauntlet server pages
var pages = []page{
	{Title: "Home", Site: sites[0].ID, Method: "GET", Path: "/", Handler: "getHome", Template: "index.tmpl"},
	{Title: "Users", Site: sites[0].ID, Method: "GET", Path: "/users", Handler: "getUsers", Template: "users/users.tmpl"},
	{Title: "Settings", Site: sites[0].ID, Method: "GET", Path: "/settings", Handler: "getSettings", Template: "settings/settings.tmpl"},
} // End Gamer Gauntlet server settings

// Begin Gamer Gauntlet app settings
// Test users
var users = []user{
	{ID: 0, Username: "Demo", Email: "demo@gmail.com", FirstName: "Demo", LastName: "Demo", Image: "Default", Key: "GamerGauntletDemo"},
}

// Test Gamer Gauntlet
var gamergauntlets = []gamergauntlet{
	{ID: 0, IP: "IP", MAC: "MAC", Key: "GamerGauntlet", User: 0},
}

// Test profiles
// TODO Move profiles to app setting json file
var profiles = []profile{
	{ID: 0, Name: "Streaming", Image: "Default", User: 0},
	{ID: 1, Name: "Recording", Image: "Default", User: 0},
}

// Test screens
// TODO Move screens to app setting json file
var screens = []screen{
	{ID: 0, Name: "Primary", Profile: "1"},
	{ID: 1, Name: "Secondary", Profile: "1"},
}

//Test buttons
//TODO Move buttons to app setting json file
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

// Test widgets
// TODO Move widgets to app setting json file
var widgets = []widget{
	{ID: 0, Name: "Twitch", Image: "Default", Type: "Stream"},
	{ID: 1, Name: "Twitch", Image: "Default", Type: "Stream"},
	{ID: 2, Name: "Youtube", Image: "Default", Type: "Chat"},
	{ID: 3, Name: "Youtube", Image: "Default", Type: "Chat"},
} // End Gamer Gauntlet app settings

func main() {

	log.Printf("Starting Gamer Gauntlet server")
	manageserver("gg", "start")

}

// Manage Gamer Gauntlet and OBS Server.
func manageserver(server string, option string) int {

	// Check if Gamer Gauntlet server status
	if server == "gg" && option == "status" {

		status := getProcess("go.exe")

		return status
	}

	// Start Gamer Gauntlet server.
	if server == "gg" && option == "start" {

		status := startgg()

		return status
	}

	//Check if OBS is running.
	if server == "obs" && option == "Status" {
		status := getProcess("obs64.exe")

		return status
	}
	if server == "obs" && option == "Start" {
		status := startobs()

		return status
	}

	return 0
}

func startgg() int {

	router := gin.Default()

	// Load page templates
	router.LoadHTMLGlob("templates/**/*.tmpl")

	// Static files location
	router.Static("/assets", "./assets")

	//Methods. Routes, and Handlers
	router.GET("/", getHome)
	router.GET("/users", getUsers)
	router.GET("/settings", getSettings)
	router.GET("/settings/:id", getSettingID)
	router.POST("/settings", postSettings)

	port := strconv.Itoa(sites[0].Port)
	router.Run(sites[0].IP + ":" + port)

	return 1
}

//TODO Implement OBS server start
func startobs() int { return 1 }

func getmenu(pages []page) []menu {

	menulinks := make([]menu, 1)

	menulinks[0].ID = 0
	menulinks[0].Name = "Main Menu"
	menulinks[0].Links = make([]link, len(pages))

	for id, page := range pages {
		menulinks[0].Links[id].ID = page.ID
		menulinks[0].Links[id].Name = page.Title
		menulinks[0].Links[id].Path = page.Path
		menulinks[0].Links[id].Active = 1
		menulinks[0].Links[id].Order = 1
	}

	return menulinks

}

func getHome(c *gin.Context) {

	//Build menu from pages
	var menu []menu = getmenu(pages)

	c.HTML(http.StatusOK, pages[0].Template, gin.H{
		"title": pages[0].Title,
		"menu":  menu[0].Links,
	})

}

func getUsers(c *gin.Context) {

	var menu []menu = getmenu(pages)

	c.HTML(http.StatusOK, pages[1].Template, gin.H{
		"title": pages[1].Title,
		"menu":  menu[0].Links,
	})

}

// getSettings
func getSettings(c *gin.Context) {

	if c.ContentType() == "application/json" {

		c.IndentedJSON(http.StatusOK, users)
		c.IndentedJSON(http.StatusOK, gamergauntlets)
		c.IndentedJSON(http.StatusOK, profiles)
		c.IndentedJSON(http.StatusOK, screens)
		c.IndentedJSON(http.StatusOK, buttons)
		c.IndentedJSON(http.StatusOK, widgets)

		return
	}

	var menu []menu = getmenu(pages)

	c.HTML(http.StatusOK, pages[2].Template, gin.H{
		"title": pages[2].Title,
		"menu":  menu[0].Links,
	})

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
