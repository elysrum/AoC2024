Note:  Shamelessly stole Nathan Lowes Go Framework for AoC from : https://github.com/nlowe/aoc2023


# Advent of Code 2024

Solutions for the 2024 Advent of Code

## Building

This project makes use of Go 1.23.

```bash
go mod download
go test ./...
```

## Running the Solutions

To run a solution, use the problem name followed by the path to an input file.

For example, to run problem 2a:

```bash
$ go run main.go 1 a
Answer: 1532
Took 1.5586ms
```

## Adding New Solutions

This currently requires the environment variable `AOC_SESSION_TOKEN` be set with
your AdventOfCode session token.

A generator program is included that makes templates for each day, automatically
downloading challenge input and updating the root command to add new subcommands
for each problem. Running `go generate` from the repo root will generate the
following for each day that is currently accessible:

* `challenge/day<N>/import.go`: A "glue" file combining commands for both of the day's problems to simplify wiring up subcommands
* `challenge/day<N>/a.go`: The main problem implementation, containing a cobra command `A` and the implementation `func a(*challenge.Input) int`
* `challenge/day<N>/a_test.go`: A basic test template
* `challenge/day<N>/b.go`: The main problem implementation, containing a cobra command `B` and the implementation `func b(*challenge.Input) int`
* `challenge/day<N>/b_test.go`: A basic test template
* `challenge/day<N>/input.txt`: The challenge input

Additionally, `challenge/cmd/cmd.go` will be regenerated to import and add all
subcommands.

This requires `goimports` be available on your `$PATH`. Additionally, you must be
logged into https://adventofcode.com in Chrome so the generator can use your session
cookie to download challenge input.

Existing solutions and challenge inputs will be skipped instead of regenerated.

## License

These solutions are licensed under the MIT License.

See [LICENSE](./LICENSE) for details.
