# distriSocks


## routines
- sockRecv routine (main routine)
- sockWriter routine
- kafka pub routine
- kafka sub routine

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