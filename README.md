# go-vmq-webhooks

Go server implementation to handle [VerneMQ](https://vernemq.com/) [webhooks plugins](https://docs.vernemq.com/plugin-development/webhookplugins).


# Development

You need to have a working Go setup, project was developed using `go1.12.5`.

The VerneMQ project has good [documentation](https://docs.vernemq.com/) which should get you started.

If you want to use Docker:

```s
$ docker run --name vernemq1 \
  -p 1883:1883 \
  -p 8888:8888 \
  -p 9001:9001 \
  -v ./config/v/vernemq.conf:/etc/vernemq/vernemq.conf \
  -d erlio/docker-vernemq
```
