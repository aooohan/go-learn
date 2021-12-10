module "learn"

go 1.16

//require (
//	gee v0.0.0
//	cache v0.0.0
//)

replace (
	gee => ./gee
	cache => ./cache
)