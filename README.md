# whoami

A naive docker image with a http server that echos hostname.


# Usage

    docker run -d --name echo -p 80:8778 cizixs/whoami 

# build

Create the binaries with `go build`:

    CGO_ENABLED=0 go build -a -ldflags '-s' .

Then, build the docker image:

    docker build -t whoami:v1.0 .

# todos

- [ ] make server port configurable
- [x] add log feature to output each request
