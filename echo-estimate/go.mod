module echoestimate

go 1.23.4

replace example.com/echo1 => ../echo1/

replace example.com/echo2 => ../echo2/

replace example.com/echo3 => ../echo3/

require (
	example.com/echo1 v0.0.0-00010101000000-000000000000
	example.com/echo2 v0.0.0-00010101000000-000000000000
	example.com/echo3 v0.0.0-00010101000000-000000000000
)
