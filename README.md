# siml

siml is a CLI tool for discovering similar, related to, competitive, or alternative options to a given site.

## Install

You can install siml using Go:
```bash
go install github.com/dwisiswant0/siml@latest
```

Alternatively, you can download the precompiled binary from the [releases page](https://github.com/dwisiswant0/siml/releases) and add it to your system path.

## Usage

To use siml, simply run the `siml` command followed by the domain of the website you want to find similar sites for:

```bash
$ siml example.com
```

This will return a list of websites that are similar or related to **example.com**.

### Library

To use `siml` as a library, you can import the `siml` package in your Go code and call the `Get` function. Here's an example usage:

```go
package main

import "github.com/dwisiswant0/siml/pkg/siml"

func main() {
	similar, err := siml.Get("example.com")
	if err != nil {
		panic(err)
	}

	for _, s := range similar {
		println(s)
	}
}
```

In the example, the `siml` package is imported and the `Get` function is called with the domain name of the website for which similar sites are to be found. If an error occurs during the process, the `Get` function returns a non-nil error object, which can be handled appropriately in the code. If the process is successful, the `Get` function returns a slice of strings containing the similar sites, which can be iterated over and printed to the console or used in any other desired way.

You can use this same pattern to call the `Get` function in your own code, passing in the domain name of the website you want to find similar sites for.

## License

`siml` is licensed under the [MIT license](https://github.com/dwisiswant0/siml/blob/master/LICENSE).