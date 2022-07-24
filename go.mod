module github.com/xinzezhu/barrage_write_proc

go 1.18

replace github.com/xinzezhu/protocol/barrage/read_proc => ../protocol/barrage/read_proc

require (
	github.com/xinzezhu/protocol/barrage/read_proc v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b
	google.golang.org/grpc v1.48.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)
