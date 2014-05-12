## hcl

Hipchat Command Line

hcl is a simple hipchat command line tool


### Usage

#### Rooms

```go
hcl room list #=> [foo, bar]
```

Send a message to a room

```go
# send message hello world to room foo
hcl room foo 'hello world' #=> Message sent
```
Or you can use short names
```go
# use short names
hcl r m test 'hello world' #=> Message sent
``` 

#### Users

```go
# send a message to use bradleyd
hcl user bradleyd 'hello world'
```


