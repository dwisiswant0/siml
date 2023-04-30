# siml

siml is a CLI tool for discovering similar, related to, competitive, or alternative options to a given site.

## Install

You can install siml using Go:
```bash
go install github.com/dwisiswant0/siml@latest
```

Alternatively, you can download the precompiled binary from the [releases page](https://github.com/dwisiswant0/siml/releases) and add it to your system path.

## Usage

```txt
  siml v0.2.0
  --
  discover similar, related to, competitive, or serve as alternatives to a given website.

Usage: siml [DOMAINS...]

Examples:
  siml instagram.com
  siml twitter.com facebook.com
```

To use siml, simply run the `siml` command followed by the domain of the website you want to find similar sites for:

```bash
$ siml example.com
```

This will return a list of websites that are similar or related to **example.com**.

### Library

[![GoDoc](https://pkg.go.dev/static/frontend/badge/badge.svg)](http://pkg.go.dev/github.com/dwisiswant0/siml/pkg/siml)

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

You can use this same pattern to call the `Get` or `Gets` _(that accept multiple domain inputs)_ function in your own code, passing in the domain name of the website you want to find similar sites for.

## Data Sources

This project relies on data obtained from https://www.sitelike.org/.

I would like to express our gratitude to the website's authors for making this data available to the public.

Please refer to https://www.sitelike.org/ for any licensing or attribution requirements.

## License

`siml` is licensed under the MIT. See [LICENSE](https://github.com/dwisiswant0/siml/blob/master/LICENSE).