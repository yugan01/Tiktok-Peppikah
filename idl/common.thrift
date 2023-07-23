namespace go common

enum Err
{
	// gateway 10001- 19999
    BadRequest            = 10001,
	Unauthorized          = 10002,
	ServerNotFound        = 10003,
	ServerMethodNotFound  = 10004,
	RequestServerFail     = 10005,
	ServerHandleFail      = 10006,
	ResponseUnableParse   = 10007,
}
