# hours-ui-mock-api

Mock API for [https://github.com/futurice/hours-ui](https://github.com/futurice/hours-ui)

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
