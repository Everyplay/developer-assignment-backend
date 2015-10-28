# Assignment

Return your answer as a zip file containing all relevant files _with tests_ (including `.git`, so that we can see your commit history), do not fork this repo.

Design and implement(with tests) a _message delivery system_ using [Go](http://golang.org/) programming language. You are free to use any external libs if needed.

In this simplified scenario the message delivery system includes the following parts:

### Hub

Hub relays incoming message bodies to receivers based on user ID(s) defined in the message. You don't need to implement authentication, hub can for example assign arbitrary (unique) user id  to the client once its connected.

- user_id - unsigned 64 bit integer
- Connection to hub must be done using pure TCP. Protocol doesnt require multiplexing.

### Clients

Clients are users who are connected to the hub. Client may send three types of messages which are described below.

### Identity message
Client can send a identity message which the hub will answer with the user_id of the connected user.

![Identity](https://raw.githubusercontent.com/Everyplay/developer-assignment-backend/master/identity.seq.png)

### List message
Client can send a list message which the hub will answer with the list of all connected client user_id:s (excluding the requesting client).

![List](https://raw.githubusercontent.com/Everyplay/developer-assignment-backend/master/list.seq.png)

### Relay message
Client can send a relay messages which body is relayed to receivers marked in the message. Design the optimal data format for the message delivery system, so that it consumes minimal amount resources (memory, cpu, etc.). Message body can be relayed to one or multiple receivers.

- max 255 receivers (user_id:s) per message
- message body - byte array (text, JSON, binary, or anything), max length 1024 kilobytes

![Relay](https://raw.githubusercontent.com/Everyplay/developer-assignment-backend/master/relay.seq.png)

*Relay example: receivers: 2 and 3, body: foobar*
