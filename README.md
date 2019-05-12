# searchDemo
searchDemo is a golang console application to search against the tickets, users and organizations resources. 

## Features
* It provides two search options: 
   1. Direct value search: require an input of search value, then application will search the value in all fields from the resources and return all matched results. For example, when search for "1", the user with id "1" and tickets with either assignee or submitter id "1" will be matched. 
   2. Field specific search: require inputs of 1) struct type(ie. 1 for tickets), 2) field name, 3) search value, then application will search the value in the specified field and return matched results.

* The search supports case-insensitive inputs

* Results are displayed as JSON string

## Run the application locally
### Run the binary file directly for Mac OS
For Mac users: The binary file "app" is included in the repo. You just need to clone this repo and browse to the ```~/searchDemo/src``` directory, then run command: 
    ```
    ./app
    ```
to launch the application.

### Build the go application and execute the built file
* Follow the insructions in [Golang Guide](https://golang.org/doc/install) and install Go into your workstation.
* Place the cloned project under Golang's "$GOPATH" directory. 
* Browse to the ```~/searchDemo/src``` directory
* Build the Golang application: 
    ```
    go build -o app
    ```
* Run command
    ```
    ./app
    ```
## Run tests
* Browse to the ```~/searchDemo/src``` directory
* Run command
    ```
    go test ./...
    ```
NOTE: Due to time limitation, in startSearch_test I use JSON.Marshal on results and check whether expected and actual results are same. By doing this, the test may break, as during marshal it may treat the list in random sequence. In order to safe guard the results compare, consider to convert interface to actual type then do a loop on each list and run deep equal compare. Alternatively, the println in latest golang project should also gentlely support key/field sorting. But I don't use it here since it's version dependent.

## Limitations
* The application now does the exact match search. i.e. Search 'Miss T' will not match 'Miss Test'

