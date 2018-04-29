package function

import (
	"bytes"
	"html/template"
	"log"
	"net/url"
	"os"
)

type UserData struct {
	User           string
	PublicURL      string
	PrettyURL      string
	QueryPrettyURL string
}

// Handle a serverless request
func Handle(req []byte) string {
	tmpl, err := template.ParseFiles("index.html")

	if err != nil {
		log.Panic("Error loading template file: ", err)
	}

	userData1 := UserData{
		PublicURL:      os.Getenv("public_url"),
		PrettyURL:      os.Getenv("pretty_url"),
		QueryPrettyURL: os.Getenv("query_pretty_url"),
	}

	vals, _ := url.ParseQuery(os.Getenv("Http_Query"))

	user := vals.Get("user")

	userData1.User = user
	log.Println("User: ", user)

	var tpl bytes.Buffer

	err = tmpl.Execute(&tpl, userData1)

	if err != nil {
		log.Panic("Error executing template: ", err)
	}

	return tpl.String()
}
