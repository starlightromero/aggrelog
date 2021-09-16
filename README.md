# AggreLog 

Pandion Take-home Project

Automated logging aggregator script

## Assignment

Part 1. Write a script in a language of your choice that runs through the log files on a server and publishes them to an external service.

Files format: log-<servicename>-date-hour

Location: There will be multiple directories, one for each service, inside /log. Your script will need to scan the subdirectories and locate files matching the format.

Output: Aggregate them into a new file and publish that to an external service that accepts a file. Please write a POST curl command. Assume a random URL for your use case.

Part 2. Create a cron schedule to execute the script hourly between 8 am and 8 pm. 

## Usage

### How to Run

The program can be run with golang (passing in the necessary flags):
```zsh
go run main.go ...
```

OR as an executable:
```zsh
./aggrelog ...
```

### Flags

*All flags are required.* You can pass in either the long flag (`-directory`) or the short flag (`-d`) for a given flag.

| Long Flag   | Short Flag | Description                        |
| :---------- | :--------- | :--------------------------------- |
| directory   | d          | root directory to aggregate logs   |
| url         | u          | OpenSearch Service domain (url)    |
| region      | r          | AWS region (e.g. "us-east-1")      |


## How to Improve

- The logs can be sent as a byte string to save space and request size
- Enabling gzip to compress requests
- Structure of the logs
- Using concurrency to search multiple subdirectories at the same time
- Add help text (e.g. `aggrelog -help`) to get CLI output which guides the user through the program


## Resources

[List the files in a folder with Go](https://flaviocopes.com/go-list-files/)

[Reading files](https://gobyexample.com/reading-files)

[Regular Expressions](https://gobyexample.com/regular-expressions)

[Amazon OpenSearch Service - Signing HTTP Requests (Golang)](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/request-signing.html#request-signing-go)