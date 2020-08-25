# replayd

## What is this ?

This is the place where replayd lives.

It is a go library and binary that can be used to intercept HTTP traffic and replay your request.

## Motivation

For some complex `curl` / `httpie` / `<insert fancy hipsterish http client>` invocations it makes sense to save them to a file that you can replay if you want to.


## How to use

When this project is usable it will work like this:

```
replayd -port 1337 &
```

```
curl --proxy http://locahost:1337 http://google.de
```
