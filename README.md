# Tangle Hornet Reading Time

Publishes some messages into an index checks the response time of consulting these messages in Tangle Hornet and writes it to a `csv` file. 

The `csv` file will be in the [`files`](./files) directory

## How to use

### Requirements
> This project was mainly tested with Go version 1.20.x

We recommend you update Go [to the latest stable version to use the library](https://go.dev/).
   
### Build the project on your own

1. Clone this repository;
2. Download all the dependencies:
   ```powershell
   go mod tidy
   ```
3. Execute the project:
   ```powershell
   go run main.go
   ```

## License
The MIT license can be found [here](./LICENSE).
