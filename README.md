# OpenWhisk Action using Go

Example code showing you how to run Go binaries as OpenWhisk Actions.

usage 
--

- Cross-compile Go binary for platform architecture.
```
export GOARCH=386
export GOOS=linux
go build action.go
```

- Build Docker image for Action and push to Dockerhub.

```
docker -t dockerhub_username/some_image_name .
docker push dockerhub_username/some_image_name
```

- Create OpenWhisk Action from public image.

```
wsk action create --docker go_action dockerhub_username/some_image_name
```

- Invoke Action using Go language binary

```
wsk action invoke --blocking --result go_action --param payload "Hello World"
```

details
-- 

This project builds a Docker image, copying in the Go binary to be executed, using the following [base image](https://hub.docker.com/r/jamesthomas/openwhisk_docker_action/). The base image uses a Node.js server to handle managing the execution requests from the OpenWhisk platform. When an Action is invoked, the binary file is executed. 

The path for the binary defaults to _/blackbox/action_, this can be overridden using a custom _ENV_ command for the _ACTION_ environment variable.

The JSON string for the invocation parameters is passed as the single command-line argument to the binary. Any data written to stdout will be interpreted as JSON and passed as the response value.
