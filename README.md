## Port Client

This is the port client repository. 

for port service repository see the following repo:

https://github.com/tjoluotch/PortService


With more time spent on this project I would:
- Add test files using Ginkgo & Gomega 3rd party packages
- Create stubs for the interfaces using the counterfeiter 3rd party package
- Implement the docker-compose.yml file
- Instead of using an in memory map as the DB representation create an instance
a relational DB such as postgresSQL
- use env vars to parse in program details such as the hostname and port

**note**

Run the project successfully through terminal and Goland editor. Not yet tested the build.sh run command. 

#### Structure
cmd - contains main.go

internal/pkg - contains all source code

configs/grpc - contains the .proto definition that was used to create the client stub and server
using the protoc v3 compiler  

