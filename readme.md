# Kafka Pro(ducer Simulator)
kafkapro is Kafka Producer simulator tool to help development team to push messages to kafka.

### Prerequisites
- Go v1.12.4
- [dep](https://github.com/golang/dep)


## Getting Started

### Build
1. Make sure the project is inside `$GOPATH/src`
2. execute `dep ensure` to download dependencies
3. run `go test`
4. run `./build.sh` or execute `go build` if you prefer to build it manually
5. if you use `build.sh` in the previous step, executable file for windows, linux, and mac are available in `build/` directory

### Run
1. change the configuration in `config.toml` to suit your needs
2. make sure your input file as defined in `config.toml` is available (default is `input.txt`)
3. make sure your machine is connected to kafka
4. run the file that generated when you build the application


## Built with

* [Go](https://golang.org) - The language for simplicity and concurrency
* [Dep](https://github.com/golang/dep) - Go dependency management tool 
* [Viper](https://github.com/spf13/viper) - Go configuration with fangs!
* [Kafka-go](https://github.com/segmentio/kafka-go) - as the name suggets, it's kafka for go

### Version
- 0.1 - ability to send message to kafka based on config

## Authors
- **Kharriz Abiyasa Bhimabrata** - *Initial work* - [Zeedware](https://github.com/Zeedware)

## License
this project is licensed under [MIT License](https://github.com/Zeedware/kafkapro/blob/master/LICENSE)