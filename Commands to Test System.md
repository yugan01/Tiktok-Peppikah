## 1) Download Go and Nacos

- [Go](https://go.dev/dl/)
- [Nacos](https://nacos.io/en-us/docs/quick-start.html)


## 2) Startup Hertz Server (in a new terminal tab)

- `cd hertz_gateway`

- `go mod tidy`

- `go run .`



## 3) Startup Nacos Registry (in a new terminal tab)

MacOS/Linux/Unix
- `sh startup.sh -m standalone`

Ubuntu
- `bash startup.sh -m standalone`

Windows
- `cmd startup.cmd -m standalone`

## 4) Startup Kitex server for *only one specific service either math or echo* (in a new terminal tab)

Running Echo service
- `cd kitex/echo`

- `go run .`


Running Math Service
- `cd kitex/math`

- `go run .`


## 5) Send HTTP requests for either Math or Echo serivce (in a new terminal tab)

*If Echo service was run*
- `cd hertz_gateway`

Windows
- `python generator.py echo 10` (10 specifies the number of Echo HTTP requests to be sent) 

Linux/Unix/MacOS
- `python3 generator.py echo 10`

*If Math service was run*
- `cd hertz_gateway`

Windows
- `python generator.py math 10` (10 specifies the number of Math HTTP requests to be sent) 

Linux/Unix/MacOS
- `python3 generator.py math 10`

