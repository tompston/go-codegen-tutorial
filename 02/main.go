package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	// create a struct that will define the data
	type MyData struct {
		Day  string
		Year int
	}

	// create a variable that will hold the date
	data := MyData{
		Day:  "Friday",
		Year: 2021,
	}

	// parse the template
	tmpl, err := template.ParseFiles("./templates/my_template.tpl")
	if err != nil {
		log.Fatalf("Could not parse struct template: %v\n", err)
	}

	// define the name + the output path. In this case we want to output the created file
	// into the same directory
	output_path := "generated.txt"
	out, _ := os.Create(output_path)
	defer out.Close()

	// combine the template with the data and create it.
	tmpl.Execute(out, data)

	if err != nil {
		log.Fatalf("Could not parse struct template: %v\n", err)
	}
}
