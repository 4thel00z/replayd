# replayd

## What is this ?

This is the place where replayd lives.

It is a go library and binary that can be used to intercept HTTP traffic and replay your request.

## Motivation

For some complex `curl` / `httpie` / `<insert fancy hipsterish http client>` invocations it makes sense to save them to a file that you can replay if you want to.


## How to use replayd

This is how to persist your requests and have them saved - under the path specified in your .replaydrc.json under the key ``path``:
```
$ replayd -port 1337 &
$ curl --proxy http://locahost:1337 http://google.de
"written 352 bytes to: build/c180e365-f1b8-4efd-692b-7c807b2a154c"
```

## How to use replay

After the serialization you can replay the serialized request like this:

```
➜  replayd git:(master) ✗ cat build/c180e365-f1b8-4efd-692b-7c807b2a154c | build/replay
{HTTP/1.1 http://ransomware.host/ GET map[Accept:[*/*] Proxy-Connection:[Keep-Alive] User-Agent:[curl/7.58.0]] [] map[]}
```

This way you can retrieve the underlying request:
```
➜  replayd git:(master) ✗ cat build/c180e365-f1b8-4efd-692b-7c807b2a154c| build/replay -dry
{"http_version":"HTTP/1.1","url":"http://ransomware.host/","method":"GET","headers":{"Accept":["*/*"],"Proxy-Connection":["Keep-Alive"],"User-Agent":["curl/7.58.0"]},"body":null,"form":{}}
```

