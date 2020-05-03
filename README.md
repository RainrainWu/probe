# probe

Probe is a high flexibility functional verification test framework, support complicated user story developed in Go, and store customized key value pair during executing test cases.

# Download
```bash
go get github.com/RainrainWu/probe
```

# Get Started
## Test case
- Test case shoud contains a parameter `*jobs.Runner` for collecting information during executing.
```go
func SaladCaesar(r *jobs.Runner) {
	dets := r.Rep.InitDetail("Caesar")
	dets.Append("Topping", "topping not found")
	dets.Append("Plate", "Plate is broken")
	r.Rep.Fail()
}
```
- Then group all cases with same topic within a slice
```go
var SaladCase []func(*jobs.Runner) = []func(*jobs.Runner) {
	SaladCobb,
	SaladCaesar,
}
```
- For the complete code please refer to `example_cases/` directory.

## Running server
- Remember add your customized cases before starting probe.
```go
package main

import (

	"github.com/RainrainWu/probe"
	"github.com/RainrainWu/probe/jobs"
	"github.com/RainrainWu/probe/example_cases"
)

func main() {

	jobs.AddJob("coffee", example_cases.CoffeeCase)
	jobs.AddJob("salad", example_cases.SaladCase)
	probe.Start()
}
```

## Restful API endpoint
- JWT is used to avoid test executed by unknown tester, login was needed before sensitive operations.
```bash
% curl http://localhost:2023/login -X POST -H 'Content-Type: application/json' -d '{"username":"probeuser","password":"probepass"}'

{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoicHJvYmV1c2VyIiwicm9sZSI6IlVEQyBUZXN0ZXIiLCJleHAiOjE1ODg0Nzc0ODcsImlhdCI6MTU4ODQ3NzE4NywiaXNzIjoiZ2luSldUIn0.ua-YmTuNWGKh8qGMBI1Du0-2qIVmxHtEw2UBZdCDuVs"}
```

- Then trigger test via token you just obtain
```bash
% curl http://localhost:2023/test -X POST -H 'Content-Type: application/json' -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoicHJvYmV1c2VyIiwicm9sZSI6IlVEQyBUZXN0ZXIiLCJleHAiOjE1ODg0Nzc0ODcsImlhdCI6MTU4ODQ3NzE4NywiaXNzIjoiZ2luSldUIn0.ua-YmTuNWGKh8qGMBI1Du0-2qIVmxHtEw2UBZdCDuVs" -d '{"index":"0008","env":"prod","topic":["salad", "coffee"],"subject":"Daily integration test","tester":"CI Server"}'
```

- For checking test record which have been execute in the past, use the `/report` or `/report/metrix` endpoints, these two can be access without login.
```bash
% curl http://localhost:2023/report/metrix -X GET -H 'Content-Type: application/json' -d '{"Index":"0008"}'

{
  "Index": "0008",
  "Meta": {
    "Index": "0008",
    "Env": "prod",
    "Tester": "CI Server",
    "Topic": [
      "salad",
      "coffee"
    ],
    "Subject": "Daily integration test"
  },
  "Stat": {
    "Total": 4,
    "Pass": 2,
    "Warning": 1,
    "Fail": 1
  },
  "Dets": {
    "Americano": {
      "Content": {
        "Hot Water": "95 degree celcius"
      }
    },
    "Caesar": {
      "Content": {
        "Plate": "Plate is broken",
        "Topping": "topping not found"
      }
    },
    "Espresso": {
      "Content": {
        "Hot Water": "70 degree celcius, not hot enough"
      }
    }
  }
}%
```

## Configration
- Configs could be set via environment variables.
    - SERVICE_PORT
    probe service listening port.
    - WORKER_QUOTA
    maximum numbers of cocurrent workers toexecute test cases at the same time.
    - USERNAME
    login username for jwt token generating.
    - PASSWORD
    login password for jwt token generating.


# Contributors
- [Rain Wu](https://github.com/RainrainWu)