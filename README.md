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
hcl room foo 'hello world'
``` 

#### Users

```go
# send a message to use bradleyd
hcl user bradleyd 'hello world'
```


