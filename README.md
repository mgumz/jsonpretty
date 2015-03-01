# jsonpretty - prettyprint json encoded content

*jsonpretty* is a little helper to make working with dense or malformed
json content easier: it either indent it or points out a parsing error.

## usage:

    $> cat sample.json
    { "firstName": "John", "lastName": "Smith", "address": { "streetAddress": "21 2nd Street", "city": "New York" } }

    $> jsonpretty -indent="  " -in sample.json
    {
      "firstName": "John",
      "lastName": "Smith",
      "address": {
        "streetAddress": "21 2nd Street",
        "city": "New York"
      }
    }

## flags:

    -in="": name of input.json, if empty <stdin> is used
    -indent="  ": string to use as indention
    -out="": name of output.json, if empty <stdout> is used
    -prefix="": string to use as prefix
    -url=false: set to true to fetch 'in' from web, default 'false'

## build:

    $> GOPATH=`pwd` go get github.com/mgumz/jsonpretty

