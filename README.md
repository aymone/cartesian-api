# Cartesian API

[![Go Report Card](https://goreportcard.com/badge/github.com/aymone/cartesian-api)](https://goreportcard.com/report/github.com/aymone/cartesian-api)
[![Maintainability](https://api.codeclimate.com/v1/badges/67651a109b421ee1213f/maintainability)](https://codeclimate.com/github/aymone/api-unit-test/maintainability)

A simple API to get [manhattan distance](https://xlinux.nist.gov/dads/HTML/manhattanDistance.html) between 1xN vectors.

Sample:

GET "/api/points?x=1,y=1,distance=3"

Params:

"x": *integer, required*

"y": *integer, required*

"distance": *integer, required*

Return points ordered by increasing distance, limited by informed distance threshold:

```json
{
    "points": [
        {
            "distance": 1,
            "points": {
                "x": 2,
                "y": 1,
            }
        },
        {
            "distance": 2,
            "points": {
                "x": 3,
                "y": 1,
            }
        },
    ]
}
```

Inside data dir exists a file called "points.json" that are used to calculate distances between informed coordinate

## Running

**Important!**

As we using go mod as dependency manager, maybe you will need to export `GO111MODULE` env var before run this project.

Some people have `dep` and `mod` together, to avoiding conflicts with `go dep`, just disable `go mod` when you not using, this will prevent `mod` to create files like `go.mod` and `go.sum` in your other projects with dep.

```bash
#on/off/auto
export GO111MODULE=auto
```

Inside root directory exists a `Makefile` to better instrumentation of that available commands:

By default the app will run on localhost:8080, to change port:

```bash
export PORT=3000
```

Running the project

```bash
make run
```

Testing

```bash
make tests
```

Coverage

```bash
make coverage
```

Updating dependencies

```bash
make deps
```

## Requirements

- Go >= 1.12.x

- Go mod

## I hope you enjoy :D
