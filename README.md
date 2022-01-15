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
    "reciever" : ws-conn,
    "sender" : username (string) // check if this can be removed
}
```
- publish  ```message``` to  ```ws-conn```
- ack() the message
- save to dynamo db (?)

### kafka publisher
- query redis for pod-id
- redis model

key: username (string)
val: pod-id, ws-conn (struct -> bytes string) (can also store array of structs to send to multiple devices username is connnected to)


- publish to the pod from the value of redis
- publish payload
```
{
    "message" : string,
    "reciever" : ws-conn,
    "sender" : username (string) // check if this can be removed
}
```

### ws Writer
- payload
```
{
    "message": string,
    "ws-conn": recievers connection
}
```
- change message status to 'DELIVERED'

## ws reader (main routine)
- save to database and mark status to 'SENT'
- recieving payload:
```
{
    "payload": string,
    "sender_id": string,
    "reciever_id": stirng
}
```