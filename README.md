

# JSON to MessagePack Converter

This is a simple Go program that converts JSON data to MessagePack format. It utilizes the Factory Design Pattern and the Builder Pattern to efficiently encode the JSON data into MessagePack format.

## Features

- Converts JSON data to MessagePack format
- Supports conversion of various data types including nil, boolean, integer, float, string, array, and map
- Efficient encoding with compact MessagePack representation
- No dependencies on external libraries

## Getting Started

### Prerequisites

- Go (Golang) must be installed on your system. You can download it from the official Go website: https://golang.org/dl/

### Installation

1. Clone the repository:

   ```shell
   git clone https://github.com/alexktchen/json_to_msgpack.git
   ```

2. Navigate to the project directory:

   ```shell
   cd json_to_msgpack
   ```

3. Build the program:

   ```shell
   go build
   ```
4. Unit test:
   ```shell
   go test
   ```

### Usage

To convert JSON data to MessagePack format, you can execute the program and provide the JSON data as input.

```shell
./json_to_msgpack json '{"age" : "1" , "name" : "Daniel"}'
```
### output
```shell
type: json
data: {"age" : "1" , "name" : "Daniel"}
results: 82a3616765a131a46e616d65a644616e69656c
```

```shell
./json_to_msgpack msgpack '82a3616765a131a46e616d65a644616e69656c'
```
### output
```shell
type: msgpack
data: 82a3616765a131a46e616d65a644616e69656c
results : {"age":"1","name":"Daniel"}
```


- `type`: the input type. E.g. json
- `data`: the json data

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please create a new issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
