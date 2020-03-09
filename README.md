# Build
The code was test to build and run `go 1.12.9`. It has not been tested with
other versions of GO.
```
cd cmd/cron
go build
```
This will produce the `cron` or `cron.exe` executable in the `cmd/cron` folder.

# Run
There is no input validation. 
```
./cron 1/10 2/3 "*" 4 3/2 /usr/bin/ls
```
If you want to pass * as an argument make sure you wrap it with double quotes "" otherwise
it will get substituted with the list of files in the directory and it won't run.


Also days of week are 0-6. Sunday as a 7 is not supported. 

Example Output:

    minute        1 11 21 31 41 51
    hour          2 5 8 11 14 17 20 23
    day of month  3
    month         4
    day of week   3 5
    command       /usr/bin/ls
    
# Run Tests

To run the unit test from the root directory of the project

```
go test -v ./...
```
