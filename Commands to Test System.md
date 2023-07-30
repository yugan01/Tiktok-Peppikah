## 1) Download Go, Nacos and Python

- [Go](https://go.dev/dl/)
- [Nacos](https://nacos.io/en-us/docs/quick-start.html) (We installed [nacos-server-2.2.3.zip](https://github.com/alibaba/nacos/releases/download/2.2.3/nacos-server-2.2.3.zip)) 
(Requires JDK and Maven. Instructions are provided)
- [Python](https://www.python.org/ftp/python/3.11.4/python-3.11.4-amd64.exe) (Please click add python.exe to PATH)

## 2) Startup Hertz Server (in a new terminal tab)

- `cd hertz_gateway`

- `go mod edit -replace github.com/cloudwego/netpoll=github.com/ganlvtech/std-netpoll@develop`

- `go mod tidy`

- `go run .`



## 3) Startup Nacos Registry (in a new terminal tab)

- `cd nacos-2.2.3/nacos/bin`

MacOS/Linux/Unix
- `sh startup.sh -m standalone`

Ubuntu
- `bash startup.sh -m standalone`

Windows
- `startup.cmd -m standalone` OR `cmd startup.cmd -m standalone`

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

