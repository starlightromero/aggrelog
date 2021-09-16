# AggreLog 

Pandion Take-home Project

Automated logging aggregator script

## Assignment

Part 1. Write a script in a language of your choice that runs through the log files on a server and publishes them to an external service.

Files format: log-<servicename>-date-hour

Location: There will be multiple directories, one for each service, inside /log. Your script will need to scan the subdirectories and locate files matching the format.

Output: Aggregate them into a new file and publish that to an external service that accepts a file. Please write a POST curl command. Assume a random URL for your use case.

Part 2. Create a cron schedule to execute the script hourly between 8 am and 8 pm. 

## Resources

[List the files in a folder with Go](https://flaviocopes.com/go-list-files/)

[Reading files](https://gobyexample.com/reading-files)

[Regular Expressions](https://gobyexample.com/regular-expressions)

[Amazon OpenSearch Service - Signing HTTP Requests (Golang)](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/request-signing.html#request-signing-go)