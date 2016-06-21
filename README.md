# OpenWhisk Action using Go

This project highlights two methods for running Go language binaries as OpenWhisk Actions, using the OpenWhisk Docker SDK or a custom Go handler via Docker.

The OpenWhisk Docker SDK uses a Node.js application to handle the JSON request from the platform and spawns a process to execute the Go binary. Invocation parameters are passed as a JSON string as a command-line argument to the binary. The binary must write the JSON string response to _stdout_, the handler will return this to the platform.

Using the custom Go handler, the user uses an external library that implements a callback to register your Action function. The library implements a simple web service that handles processing the invocations from the platform, executing the registered Action callback for each request. This compiled binary is built and runs in a Docker container.

usage
--

- Modify the sample _action.go_ file in the [docker_sdk](https://github.com/jthomas/openwhisk_go_action/tree/master/docker_sdk) or [go_handler](https://github.com/jthomas/openwhisk_go_action/tree/master/go_handler) directory with the implementation code.

- Cross-compile Go binary for platform architecture.
```
export GOARCH=386
export GOOS=linux
go build action.go
```

- Build Docker image for Action and push to Dockerhub.

```
docker build -t dockerhub_username/some_image_name .
docker push dockerhub_username/some_image_name
```

- Create OpenWhisk Action from public image.

```
wsk action create --docker go_action dockerhub_username/some_image_name
```


```
wsk action invoke --blocking --result go_action --param payload "Hello World"
{
    "reversed": "dlroW olleH"
}
```

details (docker sdk)
-- 

This example builds a Docker image, copying in the Go binary to be executed, using the following [base image](https://hub.docker.com/r/jamesthomas/openwhisk_docker_action/). The base image uses a Node.js server to handle managing the execution requests from the OpenWhisk platform. When an Action is invoked, the binary file is executed. 

The path for the binary defaults to _/blackbox/action_, this can be overridden using a custom _ENV_ command for the _ACTION_ environment variable.

The JSON string for the invocation parameters is passed as the single command-line argument to the binary. Any data written to stdout will be interpreted as JSON and passed as the response value.


details (go handler)
-- 

This example uses an external Go library to provide an interface for registering functions as Actions. The library implements a simple web service that handles processing the invocations from the platform, executing the registered Action callback for each request. 

```go
openwhisk.RegisterAction(func(value json.RawMessage) (interface{}, error) {
   ...	
}
```

This compiled binary is executed in a custom Docker container.
