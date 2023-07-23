# Tiktok-Peppikah
This is a Tiktok project which is a part of NUS Orbital 2023. Created by Yuancheng and Yugan 


## General Overview
The [thrift files](idl/) in the IDL folder underpin the services that the API gateway carries out. The Hertz and Kitex frameworks are utilized to generate server code for Hertz and Kitex servers. In this project, we have 2 Services: [Echo](idl/echo.thrift) and [Math](idl/math.thrift). The Echo service returns the same string as its response. The Math service carries out 4 arithmetic operations and returns the result as its response.
Hertz servers accept HTTP requests and the biz folder contains the logic to route these requests to the appropriate Kitex servers for the service to be carried out. HTTP requests are sent to the Hertz servers through the curl command and have JSON bodies embedded in them. 

Example of HTTP request for Echo service where the message is a string and the response is a repeat of the message:

`curl --location --request POST http://localhost:8888/echo/echo --header Content-Type: application/json --data {"message" : "Peppikah"}`


Example of HTTP request for Math service where 2 floats are required and an arithmetic operation is performed on the 2 floats:

`curl --location --request POST http://localhost:8888/math/subtract --header Content-Type: application/json --data {"first": -1, "second": 8}`

## Files with important logic and of interest
`biz/gateway/gateway.go`

The file [gateway.go](biz/gateway/gateway.go) and [router.go](hertz_gateway/router.go) acts as an abstracted program to handle and create generic calls through a loop.

`hertz_gateway/generator.py`

The Python [file](hertz_gateway/generator.py), sends random HTTP requests based on 2 command line arguments as sending HTTP commands one by one is inefficient. 
Users can send HTTP commands of a specific service and of a specific number by inputting the command line argument.

Example of a User sending 10 HTTP requests to Echo service:

`python generator.py Echo 10`

Or in MacOS

`python3 generator.py Echo 10`

## Programming Principles 
With the aim of inculcating high-level abstraction, we created generator.py to automatically send random HTTP requests and it can be easily modified to fit the services of the users. 

Furthermore, we used the concept of DRY to create gateway.go that creates generic calls for all the thrift files present.

Moreover, all the files have a single responsibility to ensure simple understanding.


## Analysis
Using PPROF for visual analysis of Hertz projects, we generated callgraphs and flame graphs to compare the performance of 2 Load balancers: Weighted Round Robin Load balancer and Weighted Random Load Balancer.

![Screenshot 2023-07-23 195349](https://github.com/yugan01/Tiktok-Peppikah/assets/122327042/240f40a1-56d4-4ccd-bd90-4a8604a3f25c)


We believe that the weighted round-robin load balancer has fewer limitations and thus, is adopted in our project. The weighted random load balancer has the specific shortcomings of the split stack calls and the path requiring both calls to complete to continue the routine. 


## Conclusion
We have fulfilled the objectives of implementing a flexible API gateway. Moreover, we have utilised software engineering and programming concepts such as DRY, high-level abstraction, and single responsibility principle for simple programming paradigms. Furthermore, our analysis has justified to use of round-robin load-balancer compared to random load-balancer. We have gained substantial exposure to CLoudweGo projects and various HTTP frameworks as well as the analytical tools for important design decisions. 



