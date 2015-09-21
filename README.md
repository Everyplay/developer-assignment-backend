# Assignment

Return your answer as a zip file containing all relevant files(including `.git`, so that we can see your commit history), do not fork this repo. 

Design and implement(with tests) a _message delivery system_ using [Go](http://golang.org/) programming language. You are free to use any external libs if needed.

![Example](https://dl.dropboxusercontent.com/u/13424146/delivery_example.png)

In this simplified scenario the message delivery system includes the following parts:

### Hub

Hub relays incoming messages to receivers based on user ID(s) defined in the message. You don't need to implement authentication, hub can for example assign arbitrary (unique) user id  to the client once its connected.

### Clients
Clients are users who are connected to the hub. Client may send messages to hub which relays message to receiving users (other clients), which are connected to hub.

### Message

Design the optimal data format for the message delivery system, so that it consumes minimal amount resources (memory, cpu, etc.).

The following constraints apply:

- Message should be relayed to one or multiple receivers (max 255 receivers per message is supported)
- Message payload - byte array (containing the message content, in most cases JSON), max length 1024 kilobytes
- user_id - unsigned 64 bit integer
- Connection to hub must be done using pure TCP. Payload length is finite and protocol doesnt require multiplexing.
