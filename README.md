# Breadth-first-search

## Description
This Go code simulates a breadth-first search to find a salesman among a generated list of fake names. It starts by initializing a list with a real name and then generates additional fake names. The program uses a breadth-first approach to explore the names, checking if each person is a salesman based on specific criteria. If a salesman is found, the program prints the person's name and the index where they were found. The process continues until either a salesman is found or the search exhausts the generated names. The main function orchestrates the generation of fake names and the search for a salesman, printing the result to the console.

## Usage
To run the program, use the following command:
```bash
go get
go run main.go
```

## Output
The program prints the name of the salesman and the index where they were found. If no salesman is found, the program prints a message indicating that no salesman was found.
```
Person: Melisa Stamm, found in the index 323
```

## References
* [Go by Example: Hello World](https://gobyexample.com/hello-world)
* [Go by Example: Functions](https://gobyexample.com/functions)
* [Go by Example: Strings](https://gobyexample.com/string-functions)
* [Go by Example: Variables](https://gobyexample.com/variables)
* [Go by Example: For](https://gobyexample.com/for)
* [Go by Example: If/Else](https://gobyexample.com/if-else)
* [Go by Example: Slices](https://gobyexample.com/slices)
* [Go by Example: Range](https://gobyexample.com/range)

## Blog
* [Breadth-first search in Go](https://andersonteala.hamstercode.com.br/posts/breadth-first-search-in-go)

## License
This code is open source software licensed under the [MIT License](https://opensource.org/licenses/MIT).

## Authors
* [Anderson Teala](https://github.com/AndersonTeala)