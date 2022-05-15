package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type SocialLink struct {
	Site string `json:"site"`
	URL  string `json:"url"`
}
type User struct {
	Firstname       string       `json:"firstname"`
	Lastname        string       `json:"lastname,omitempty"`
	ID              string       `json:"id"`
	Age             string       `json:"age,omitempty"`
	Gender          string       `json:"gender,omitempty"`
	PreferredTopics []string     `json:"preferred_topics,omitempty"`
	SocialLinks     []SocialLink `json:"social_links,omitempty"`
}

func ReadJSON(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading user.json", err)
		return nil, err
	}
	fmt.Println("Success reading user.json")
	return data, nil
}
func DecodeJSON(data []byte, user *User) {
	err := json.Unmarshal(data, user)
	if err != nil {
		fmt.Println("Error parsing JSON", err)
	}
	fmt.Println(*user)
}
func EncodeJSON(user *User) {
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error parsing JSON", err)
	}
	fmt.Println(string(data))
}
func main() {
	data, err := ReadJSON("user.json")
	if err != nil {
		return
	}
	fmt.Println("Content of user.json:")
	fmt.Println(string(data))
	var user User
	fmt.Println("\nDecode JSON data to user struct:")
	DecodeJSON(data, &user)

	user2 := User{
		Firstname: "abc",
		Lastname:  "ddd",
		ID:        "kkk",
		Age:       "10",
		Gender:    "male",
		SocialLinks: []SocialLink{
			{
				Site: "QQ",
				URL:  "https://abcddd.com/QQ",
			},
		},
	}
	fmt.Println("\nEncode struct to JSON:")
	EncodeJSON(&user2)
}
