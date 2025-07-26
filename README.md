## Project : Servo

> High performance web server for modern applications

<a href="https://pkg.go.dev/github.com/arsharaj/project-servo"><img src="https://pkg.go.dev/badge/github.com/arsharaj/project-servo.svg" alt="Go Reference"></a>

### Features

- Graceful Shutdown
- Structured Logging
- Configuration Management
- Middleware Support
- Routing
- JSON Response System
- Unit and Integration Testing

### Requirements

- go : 1.24.4

### Folder Structure

```
root
├── .github
├── config
├── docs
├── handlers
├── logger
├── middleware
├── server
├── tests
│   ├── handlers
│   └── middleware
└── utils
```

### Documentation

Full documentation can be found in [docs](docs/) folder.

### Running Tests

Run the following command to execute the test suite : `go test ./tests/...`

### Contributing

We welcome contributions! Please read [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

### Versioning

This project uses [Semantic Versioning](https://semver.org/).

### License

[MIT License](LICENSE)