package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

//data model
type Sites struct {
	Sites []Site `json:"urls"`
}

type Site struct {
	SiteId string `json:"site_id"`
	Url    string `json:"url"`
}

//Fnction to check the healthyness of each URL
func checklink(link string, c chan string) {
	resp, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down! because of the Reason : ", err, "at time: ", time.Now().Format("2006-01-02 3:4:5 pm"))
		c <- link
		return
	}
	defer resp.Body.Close()
	fmt.Println("Hurrah", link, "is up! at time: ", time.Now().Format("2006-01-02 3:4:5 pm"))
	c <- link
}

func main() {

	// Open our jsonFile
	jsonFile, err := os.Open("List_of_Sites.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened List_Of_sites.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	//read the Json file in byte format
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	data := Sites{}

	//Unmarsh the data into data struct
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		fmt.Println(err)
	}
	//create a channel to communicate between main and child(checklink) goroutine
	c := make(chan string)

	for i := 0; i < len(data.Sites); i++ {
		link := data.Sites[i].Url
		go checklink(link, c)
	}

	//Function literal checking the all set of sites in every 5 minutes
	for l := range c {
		go func(link string) { // function literal
			time.Sleep(5 * time.Second)
			checklink(link, c)
		}(l) //parenthesis to call the function literal

	}

}
