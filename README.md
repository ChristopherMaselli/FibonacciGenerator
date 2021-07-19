Fibonacci Generator project for Reserve Trust.

Project built in Golang programming Language and connects with Postgres Database for functionality.

Project includes "FibonacciAPI.go" file which consists of the program itself, along with the "FibonacciAPI_test.go" file which acts as the file tester. Test file includes some previous test cases, these will need to be changed for the second function depending on what is currently in the database. Project includes Make file to allow for easy dependency download, program running, and testing. 

Instructions to Build:

1) Open folder containing files in Visual Studio Code, then in the terminal use the make file to download the dependencies by inputting "make dependencies". After the dependencies have been downloaded, use the make file to run the program with "make run". Program can also be run directly using the command "go run FibonacciAPI.go". Project will start and API endpoints can be sent Json requests of the input numbers. 

Requests sent will need to be in the form of a single-value json of "InputNumber" and the desired value for both functions.

2) Postgres database will need to be set up with the parameters (host, port, user, password, dbname, ec) present in the program file. After this a table of the name "Numbers" with two columns will need to be created. These columns will be int columns named "InputNumber" and "FibonacciNumber".

Instructions to Test:

1) Use the make file by inputting "make test" to run the testing file. Testing file can also be run directly by using the "go test" command.