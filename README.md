# Tiktok-Peppikah
This is a Tiktok project which is a part of NUS Orbital 2023. Created by Yuancheng and Yugan 

#
## General Overview
The thrift in the IDL folder underpin the services that the API gateway carries out. The Hertz and Kitex frameworks are utilized to generate server code for Hertz and Kitex servers. 
Hertz servers accept HTTP requests and the biz folder contains the logic to route these requests to the appropriate Kitex servers for the service to be carried out. HTTP requests are sent to the Hertz servers through the curl command and have JSON bodies embedded in them. 

Example of HTTP request for Echo service where the message is a string and the response is a repeat of the message:

`curl --location --request POST http://localhost:8888/echo/echo --header Content-Type: application/json --data {"message" : "Peppikah"}`


Example of HTTP request for Math service where 2 floats are required and an arithmetic operation is performed on the 2 floats:

`curl --location --request POST http://localhost:8888/math/subtract --header Content-Type: application/json --data {"first": -1, "second": 8}`

## Files with important logic and of interest
`biz/gateway/gateway.go`

The file [biz/gateway/gateway.go](biz/gateway/gateway.go) acts as an abstracted program to handle and create generic calls through a loop.

`hertz_gateway/generator.py`

The Python file, [hertz_gateway/generator.py](generator.py),sends random HTTP requests based on 2 command line arguments as sending HTTP commands one by one is inefficient. 
Users can send HTTP commands of a specific service and of a specific number by inputting the command line argument.

Example of a User sending 10 HTTP requests to Echo service:

`python generator.py Echo 10`

Or in MacOS

`python3 generator.py Echo 10`
