# gin-webApp

 A web application with JWT authentication using Go web framework gin 

1. clone the repository https://github.com/PlumageRan/gin-webApp.git

2. to run the web application you may need to install some indepencies
	* gin framework
	```
	$ go get github.com/gin-gonic/gin
	```
	* GORM
	```
	$ go get github.com/jinzhu/gorm
	```
3. run the web application in command-line
	```
	$ go run main.go
	```
4. test in Postman
	* to get token: localhost:8080/login, username: admin, password: password, copy the token and paste in the authorization tag(Bearer Token)
	* get all users: localhost:8080/users, it will return [{"id":1,"name":"John","age":31,"city":"New York"},{"id":2,"name":"Doe","age":22,"city":"Vancouver"}]
	* get a perticular user: localhost:8080/user/:id, for example localhost:8080/user/2 will return {"id":2,"name":"Doe","age":22,"city":"Vancouver"}
5. unit test
	```
	$ go test -v
	```
