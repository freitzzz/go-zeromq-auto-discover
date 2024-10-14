module github.com/freitzzz/go-zeromq-auto-discover

go 1.23.1

require (
	github.com/joho/godotenv v1.5.1
	github.com/pebbe/zmq4 v1.2.11
	github.com/zeromq/gyre v0.0.0-20180708015545-f03685a6fad2
)

require (
	github.com/armen/goviral v0.0.0-20150416020909-5ca05b524b2b // indirect
	github.com/jtacoma/go-zpl v0.3.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
)

replace github.com/zeromq/gyre => github.com/freitzzz/gyre v0.0.0-20241014095254-cfb8db7944a7
