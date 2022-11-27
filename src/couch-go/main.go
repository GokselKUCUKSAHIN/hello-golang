package main

import (
	"fmt"
	gocb2 "github.com/couchbase/gocb/v2"
)

type Person struct {
	ID        string        `json:"id,omitempty"`
	FirstName string        `json:"firstName,omitempty"`
	LastName  string        `json:"lastName,omitempty"`
	Social    []SocialMedia `json:"socialmedia,omitempty"`
}

type SocialMedia struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

func main() {

	connectionString := "localhost"
	bucketName := "my-bucket"
	username := "admin"
	password := "pass123"

	cluster, _ := gocb2.Connect("couchbase://"+connectionString, gocb2.ClusterOptions{
		Authenticator: gocb2.PasswordAuthenticator{
			Username: username,
			Password: password,
		},
	})
	bucket := cluster.Bucket(bucketName)
	// var person Person
	col := bucket.DefaultCollection()

	_, err := col.Upsert("goksel", Person{
		FirstName: "Göksel",
		LastName:  "Küçükşahin",
		Social: []SocialMedia{
			{Title: "LinkedIn", Link: "https://www.linkedin.com/in/gokselkucuksahin/"},
			{Title: "GitHub", Link: "https://www.github.com/GokselKUCUKSAHIN"},
		},
	}, &gocb2.UpsertOptions{Expiry: 0})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// bucket.Get("goksel", &person)
	// jsonBytes, _ := json.Marshal(person)
	// fmt.Println(string(jsonBytes))
}
