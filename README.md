# whoami

A naive docker image with a http server that echos hostname(identifies itself).

This docker image can be used to demostrate and validate load balance ability of your service.

# Usage

    $ docker run -d --name echo -p 80:8778 cizixs/whoami
    $ curl http://127.0.0.1
    2c47ad631c82

By default container exposes port `8778`, and returns container hostname(short id) as response.

You can control what is returned by setting `MESAGE` environment variable(the following example is illustrated under docker 1.12 swarm mode):

    ➜  docker service create -e MESSAGE="viola" -p 8778:8778 --replicas=2 whoami:v0.4
    ➜  curl http://127.0.0.1:8778
    viola from bf8cf715445d
    ➜  curl http://127.0.0.1:8778
    viola from 15d94216ff07

And every container outputs requests to stdout, which is accessible by standard docker logs:

    ➜  docker logs bf8cf7
    2016/11/04 03:25:35 start serving...
    2016/11/04 03:25:51 [GET] at "/" takes 41.015µs
    2016/11/04 03:25:53 [GET] at "/" takes 26.656µs

# Build

Create the binaries with `go build`:

    CGO_ENABLED=0 go build -a -ldflags '-s' .

Then, build the docker image:

    docker build -t whoami:v1.0 .

# Todos

- [ ] make server port configurable
- [x] add log feature to output each request in terminal
- [x] allow passing in custom message 
