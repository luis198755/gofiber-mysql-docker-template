# Example Go / Fiber / MariaDB Docker Containers

This is a simple example / template project that builds your Go application and runs it in a container that is connected to a MariaDB container in order to have the database operational.

You can use the makefile in order to build your application.
The template application is a C(reate)R(ead)U(pdate)D(elete) Rest API that creates and retrieves apples and saves them into the database.

## Building

```shell
# Build 'main' executable
$ make

# Build containers and setup
# build the executable locally in order to test for compile time errors
# if there is no error, the images are built and started.
$ make containers

# Stop containers
$ make stop_containers

# Stops containers, removes them and removes unused intermediate containers.
$ make clean
```

## API endpoints

### http://127.0.0.1:3000/apples

### http://127.0.0.1:3000/apples

```


