### Using Packages
- golangci-lint
  - install
  ```curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.53.3```
  - run
  ```golangci-lint run```
  - skip lint
  ```git commit -m [message] --no-verify```
- Pre-commit
  - install
  https://pre-commit.com/#install  
  ```brew install pre-commit```
  - apply in project
  ```pre-commit install```