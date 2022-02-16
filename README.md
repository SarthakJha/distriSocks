# distriSocks

![architecture](https://firebasestorage.googleapis.com/v0/b/portfolio-a186e.appspot.com/o/DistriSock.drawio.png?alt=media&token=41df7dcf-b440-48b4-b3eb-04b96407019a)

![websock-application](https://firebasestorage.googleapis.com/v0/b/portfolio-a186e.appspot.com/o/DistriSock-low%20level%20design.drawio.png?alt=media&token=8d170849-1744-490e-8146-a458198e71d9)

## Routines

- sockRecv routine (main routine)
- sockWriter routine
- kafka pub routine
- kafka sub routine

## Horizontal scaling methods

1. broadcasting to all servers

- cons:
  - to many redundant events published to all servers
- pros:
  - dont have to keep and maintain connection details
  - multidevice connection response
- have to keep local record of all username-(random-int):connections (pod wise)
- topic
  - only one topic
    - will replay all the events on creation of new pod
  - pod-id topic
    - maintain username-(random-int): ws.Conn
    - use sync.Map

## kafka details

### topic

topic will be the pod-id

### partitions

- all the consumer routines will pe part of a same consumer group subscribing to one topic
- number of partions = number of consumers(consumer-routines) in a consumer group of the topic

### kafka listeners

- each pod will listen to its own "pod-id" topic
- recieve sub payload:

```
{
    "message" : string,
    "reciever_id" : string,
    "sender_id" :  string
}
```

- ack() the message
- pass to the ws-writer

### kafka publisher

- query redis for recv_id to get pod_id
- redis model

```
key: user_id (string) // will work for one-user-one-conn
val: pod-id
```

- publish to the pod from the value of redis
- (publish) payload

```
{
    "message" : string,
    "reciever_id" : string,
    "sender_id" : username (string)
}
```

- set database status to 'SENT'\*

### ws Writer

- payload (recieved)

```
{
     "message" : string,
    "reciever_id" : string,
    "sender_id" : username (string)
}
```

- queries local map for reciever_id to get its ws.Conn
- writes to conn
- change message status to 'DELIVERED'\*

## ws reader (main routine)

- save to database and mark status to 'NONE'
- recieving payload:

```
{
    "payload": string,
    "sender_id": string,
    "reciever_id": stirng
}
```

## potential upgrades

- - - emit event on topic: 'SET_STATUS_TO_SENT' /'SET_STATUS_TO_DELIVERED' from websock to auth svc
    - carry writes/updates on user-table from auth svc
    - carry writes/updates on message-table from websock-svc

    - something
