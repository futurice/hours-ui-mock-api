# hours-ui-mock-api

Mock API for [https://github.com/futurice/hours-ui](https://github.com/futurice/hours-ui)

## Known issues

Disclaimer: this API is just for development purposes and only to validate
correct behaviour in UI.

1. Everything is randomly generated, don't expect to get correct hours for months etc.
2. If you save a day of entries, you will get the last entry's hours as a sum for that day
3. You can only delete an entry from a current day, otherwise the UI will crash.
4. The most recent project and entry is shuffled when modifying an entry, so you might not get
   the correct order of the most recent projects and entries when creating a new entry.

## Setup

1. Clone this repo
2. Run: 
    - If you have Docker:
        - Run `./dockerBuild.sh` and `./dockerRun.sh`
    - Else if you have Go:
        - Move this folder to `$GOPATH/src/github.com/futurice/`
        - Run `go get github.com/julienschmidt/httprouter`
        - Run `go install github.com/futurice/hours-ui-mock-api`
        - Run `$GOPATH/bin/hours-ui-mock-api`
    - Else you should install Docker or Go (`>=1.6`)
3. Now you can use this mock API when developing [hours-ui](https://github.com/futurice/hours-ui)

## License

[MIT](LICENSE)
