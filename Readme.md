# kafka-go > Example

> This is the simplest implementation of https://github.com/segmentio/kafka-go .

# how to install `kafka-go`
```bash
go get github.com/segmentio/kafka-go
```

## This is how to you make connection

```go

	con, err := kafka.DialLeader(
		CTX,
		ConType,
		Host,
		Topic,
		Partition,
	)

```

## This is how you close a connection

```go
     con.Close()
```
> it's recommended to use Connection with `differ` so that it'll always run at the end of func.

## Read the code to get to know more ;-)

---

### Note :- 

i'm using Docker to spin-up kafka 

### Help :-

if you want to know more about kafka, check this video out. 

- https://www.youtube.com/watch?v=R873BlNVUB4