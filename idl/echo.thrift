struct Request {
    1: required string message
}

struct Response {
    1: string message
}

service Echo {
    Response echo(1: Request req)( api.post = '/echo/echo', api.param = 'true')
}