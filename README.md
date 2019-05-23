# Cognitive-Backend-Challenge

## Assumptions
* I added a URL field to the campaign resource to use it to extract category when not present using the category extraction API.
* The reporting end point can deal with only 2 types of graphs:
    * One dimension on the x axis, for example (x-y):
        * Categories - Campagins Count
        * Countries - Campagins Count
        * Budget - Campagins Count
        * Goal - Campagins Count, etc.
    * Two dimensions on the x axis, for example:
        * Categories/Countries - Campagins Count, etc.

## End Points
* ### Create (`POST`)
```
    localhost:3000/new
    
    And the body is as follows:
    {
	"id": "11",
	"name": "n5",
	"country": "Canda",
	"budget": 149,
	"goal": "Conversion",
	"category": "Technology"
    }
```

* ### Read (`GET`)
```
    localhost:3000/read

        Reads all campaigns
```
```
    localhost:3000/read?id=1

        Read campaign with id = 1
```

* ### Update (`PUT`)
```
    localhost:3000/update

    And the modified body is as follows:
    {
	"id": "11",
	"name": "n5",
	"country": "Canda",
	"budget": 149,
	"goal": "Conversion",
	"category": "Technology"
    }
```

* ### Delete (`DELETE`)
```
    localhost:3000/delete?id=1

        Deleted campaign with id = 1
```

* ### Chart (`GET`)
```
    This is the reporting endpoint

    localhost:3000/chart?key1=Category

        Sends a request to report a 1 dimensional [Category - Campaign Count] 
        graph as described above.


    localhost:3000/chart?key1=Category&key2=Country

        Sends a request to report a 2 dimensional [Category/Country - Campaign Count] 
        graph as described above.
```

* ### Build and Run
    1- The app is developed in Go, so you must have Go installed.

        $ sudo apt-get install golang
    2- Install gorilla/mux for routing.

        $ go get -u github.com/gorilla/mux
    3- Install godotenv for reading the .env file.

        $ go get github.com/joho/godotenv
    4- Install the PostgreSQL driver used for the Database.

        $ go get github.com/lib/pq
    5- Create a Database using the following command

        $ sudo -su postgres psql

        CREATE DATABASE api_db

        And edit the .env file with your credentials.
        Note: The .env file should be placed inside: ~/go/src/Cognitive-Backend-Challenge/api/controller

    5- Place the source code into your Go Directory, usually

        user@user:~/go/src$
    6- Navigate to the following directory

        ~/go/src/Cognitive-Backend-Challenge/api/controller
    7- Run using the following command

        go run controller.go
    8- You can use the end points through postman, and the reporting end point through browser.

