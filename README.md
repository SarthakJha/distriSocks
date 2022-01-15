# distriSocks


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
    -   only one topic
        - will replay all the events on creation of new pod
    - pod-id topic 
        - maintain username-(random-int): ws.Conn
        - use sync.Map


## kafka details
### topic
topic will be the pod-id

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
- set database status to 'SENT'

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
- change message status to 'DELIVERED'

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