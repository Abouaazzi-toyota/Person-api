# Person-api
Simple GO HTTP service

This program was created to give new GO Developers a grasp of implementing a REST API in GoLang!
This program was assembled in Microsoft VS code on Windows 10.
To test output of the program requests will be made using a program called Postman.


***********Running the program***************
1.From within a Mac terminal or Bash terminal for windows users,
2.Make sure you are in the directory which has a person-api.exe as well as the main.go
3.From there execute the command go build && ./person-api.exe
4. Open Postman  
5.create a JSON body to POST/GET and select the raw option and enter your json object to be parsed
6.enter the address http://localhost:8080/people for a POST request and select enter
7.enter the address http://localhost:8080/people for a GET request and select enter to see all json objects 
8.enter the address http://localhost:8080/people/{name} for a GET request and select enter to query based on the name of the json Object
