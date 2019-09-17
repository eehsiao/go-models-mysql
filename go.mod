module github.com/eehsiao/go-models-mysql

go 1.12

require (
	github.com/eehsiao/go-models-lib v0.0.1
	github.com/eehsiao/go-models/lib v0.0.0-20190916110016-53e2c75396a6 // indirect
	github.com/eehsiao/sqlbuilder v0.0.1
	github.com/go-sql-driver/mysql v1.4.1
	google.golang.org/appengine v1.6.2 // indirect
)

replace github.com/eehsiao/go-models-lib => ../go-models-lib

replace github.com/eehsiao/sqlbuilder => ../sqlbuilder
