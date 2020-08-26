# replayd

## What is this ?

This is the place where replayd lives.

It is a go library and binary that can be used to intercept HTTP traffic and replay your request.

## Motivation

For some complex `curl` / `httpie` / `<insert fancy hipsterish http client>` invocations it makes sense to save them to a file that you can replay if you want to.


## How to use replayd

When this project is usable it will work like this:

```
replayd -port 1337 &
```

```
curl --proxy http://locahost:1337 http://google.de
```

## How to use replay

After the serialization you can retrieve the cleartext values of the request with the replay utility:
```
➜  replayd git:(master) ✗ cat build/763c69c7-cb5f-4130-4a63-93363e940cbf
Wf+BAwEBB1JlcXVlc3QB/4IAAQYBC0hUVFBWZXJzaW9uAQwAAQRIb3N0AQwAAQZNZXRob2QBDAABB0hlYWRlcnMB/4YAAQRCb2R5AQoAAQRGb3JtAf+IAAAAF/+FBAEBBkhlYWRlcgH/hgABDAH/hAAADP+DAgEC/4QAAQwAABf/hwQBAQZWYWx1ZXMB/4gAAQwB/4QAAGj/ggEISFRUUC8xLjEBD3JhbnNvbXdhcmUuaG9zdAEDR0VUAQMKVXNlci1BZ2VudAELY3VybC83LjU4LjAGQWNjZXB0AQMqLyoQUHJveHktQ29ubmVjdGlvbgEKS2VlcC1BbGl2ZQIAAA==%
➜  replayd git:(master) ✗ cat build/763c69c7-cb5f-4130-4a63-93363e940cbf |   build/replay
{"http_version":"HTTP/1.1","host":"ransomware.host","method":"GET","headers":{"Accept":["*/*"],"Proxy-Connection":["Keep-Alive"],"User-Agent":["curl/7.58.0"]},"body":null,"form":{}}%
```

