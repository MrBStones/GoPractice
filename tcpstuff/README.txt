a) What are packages in your implementation? What data structure do you use to transmit data and meta-data?
Our packages are integers sent through a go channel. The structure is a go channel, and there is no meta-data.

b) Does your implementation use threads or processes? Why is it not realistic to use threads?
We use threads through the Go routines system. It's not realistic to use threads, because this is a single app, single machine program. 
Normal programs span multiple machines and through large networks. There might also be high network loads, where here there is a single request which terminates the program after.

c) In case the network changes the order in which messages are delivered, how would you handle message re-ordering?
TCP verifies the message and ensures ordering by checking the sequence. 
You could re-order received messages on the server side.

d) In case messages can be delayed or lost, how does your implementation handle message loss?
In that case, our code will retry sending and waiting for response. Loss of message, will automatically send a new one after a second.
Delay of message, will just make it wait.

e) Why is the 3-way handshake important?
To establish a connection, and make sure it's the correct program you are trying to connect to.