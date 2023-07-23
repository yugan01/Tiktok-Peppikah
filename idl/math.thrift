namespace go math

struct AddRequest {
	1: required double first
	2: required double second
}

struct SubtractRequest {
	1: required double first
	2: required double second
}

struct MultiplyRequest {
	1: required double first
	2: required double second
}

struct DivideRequest {
	1: required double first
	2: required double second
}

struct AddResponse {
	1: double addresult
}

struct SubtractResponse {
	1: double subtractresult
}

struct MultiplyResponse {
	1: double multiplyresult
}

struct DivideResponse {
	1: double divideresult
}

service MathSvc {
	AddResponse add(1: AddRequest req)( api.post = '/math/add', api.param = 'true')
	SubtractResponse subtract(1: SubtractRequest req)( api.post = '/math/subtract', api.param = 'true')
	MultiplyResponse multiply(1: MultiplyRequest req)( api.post = '/math/multiply', api.param = 'true')
	DivideResponse divide(1: DivideRequest req)( api.post = '/math/divide', api.param = 'true')
}
