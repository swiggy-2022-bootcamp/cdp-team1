package main

import (
	"qwik.in/account-frontstore/app"
)

// @title          Swiggy Qwik - admin side customer module
// @version        1.0
// @description    This microservice is for customer service on admin side.
// @contact.name   Ravikumar Shantharaju
// @contact.email  ravikumarsravi1999@gmail.com
// @license.name  Apache 2.0
// @host      localhost:7000
// @BasePath customer-admin/api
func main() {
	/*
		type Movie struct {
			Year  int         `json:"year"`
			Title string      `json:"title"`
			Info  interface{} `json:"info"`
		}

		moviesData, err := os.Open("moviedata.json")
		defer moviesData.Close()
		if err != nil {
			fmt.Println("Could not open the moviedata.json file", err.Error())
			os.Exit(1)
		}

		var movies []Movie
		err = json.NewDecoder(moviesData).Decode(&movies)
		if err != nil {
			fmt.Println("Could not decode the moviedata.json data", err.Error())
			os.Exit(1)
		}

		svc := repository.GetDynamoDBInstance()

		for _, movie := range movies {

			info, err := dynamodbattribute.MarshalMap(movie)
			if err != nil {
				panic(fmt.Sprintf("failed to marshal the movie, %v", err))
			}

			input := &dynamodb.PutItemInput{
				Item:      info,
				TableName: aws.String("Movies"),
			}

			_, err = svc.PutItem(input)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

		}

		fmt.Printf("We have processed %v records\n", len(movies))
	*/
	app.Start()
}
