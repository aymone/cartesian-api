# Cartesian API

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

Inside root directory exists a `Makefile` to better instrumentation of that available commands:

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

## I hope you enjoy this exercise
