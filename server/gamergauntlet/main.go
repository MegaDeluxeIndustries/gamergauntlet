package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/go-ps"
)

// Gamer Gauntlet server settings
type serversetting struct {
	Sites []site `json:"sites"`
	Pages []page `json:"pages"`
	Users []user `json:"users"`
}

// Gamer Gauntlet app settings
type appsetting struct {
	GamerGauntets []gamergauntlet `json:"gamergauntlets"`
	Profiles      []profile       `json:"profiles"`
	Screens       []screen        `json:"screens"`
	Buttons       []button        `json:"buttons"`
	Widgets       []widget        `json:"widgets"`
}

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

// Global variables
var serversettings serversetting
var appsettings appsetting
var menulinks menu

func main() {

	log.Printf("Starting...")
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

	loadSettings()
	getMenu(serversettings.Pages)

	router := gin.Default()

	// Load page templates
	router.LoadHTMLGlob("templates/**/*.tmpl")

	// Static files location
	router.Static("/assets", "./assets")

	//Methods, Routes, and Handlers
	router.GET("/", getHome)
	router.GET("/users", getUsers)
	router.GET("/settings", getSettings)
	router.GET("/settings/:id", getSettingID)
	router.POST("/settings", postSettings)

	router.Run()

	return 1
}

//TODO Implement OBS server start
func startobs() int { return 1 }

func loadSettings() {

	// Open jsonFile
	serverJson, err := os.Open("server.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened server.json")

	defer serverJson.Close()

	// Read jsonFile as a byte array.
	serverValue, _ := ioutil.ReadAll(serverJson)

	// Unmarshal byteArray into settings
	json.Unmarshal(serverValue, &serversettings)

	// Open jsonFile
	appJson, err := os.Open("app.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened app.json")

	defer appJson.Close()

	// Read jsonFile as a byte array.
	appValue, _ := ioutil.ReadAll(appJson)

	// Unmarshal byteArray into settings
	json.Unmarshal(appValue, &appsettings)

}

// Build menu from the pages
func getMenu(pages []page) {

	menulinks.ID = 0
	menulinks.Name = "Main Menu"
	menulinks.Links = make([]link, len(pages))

	for id, page := range pages {
		menulinks.Links[id].ID = page.ID
		menulinks.Links[id].Name = page.Title
		menulinks.Links[id].Path = page.Path
		menulinks.Links[id].Active = 1
		menulinks.Links[id].Order = 1
	}

}

func getHome(c *gin.Context) {

	c.HTML(http.StatusOK, serversettings.Pages[0].Template, gin.H{
		"title": serversettings.Pages[0].Title,
		"menu":  menulinks.Links,
	})

}

func getUsers(c *gin.Context) {

	c.HTML(http.StatusOK, serversettings.Pages[1].Template, gin.H{
		"title": serversettings.Pages[1].Title,
		"menu":  menulinks.Links,
	})

}

// getSettings
func getSettings(c *gin.Context) {

	if c.ContentType() == "application/json" {

		c.IndentedJSON(http.StatusOK, appsettings.GamerGauntets)
		c.IndentedJSON(http.StatusOK, appsettings.Profiles)
		c.IndentedJSON(http.StatusOK, appsettings.Screens)
		c.IndentedJSON(http.StatusOK, appsettings.Buttons)
		c.IndentedJSON(http.StatusOK, appsettings.Widgets)

		return
	}

	c.HTML(http.StatusOK, serversettings.Pages[2].Template, gin.H{
		"title": serversettings.Pages[2].Title,
		"menu":  menulinks.Links,
	})

}

// postSettings
func postSettings(c *gin.Context) {
	var newButton button

	if err := c.BindJSON(&newButton); err != nil {
		return
	}

	// Add the new button to the slice.
	appsettings.Buttons = append(appsettings.Buttons, newButton)
	c.IndentedJSON(http.StatusCreated, newButton)
}

// getSettingID
func getSettingID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range appsettings.Buttons {
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
