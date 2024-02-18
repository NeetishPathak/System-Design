# Push 

Push model is good for certain use cases where the clients need real time notification like in Chatting apps


How does "push" work ?
- Client connects to the server
- Server sends data to the client
- Client doesnt request anything to the server
- Protocol must be bi-directional


Example: It is used by RabbitMQ


### Pros and Cons

**Pros** 
- Real Time

**Cons**
- Clients must be online i.e. connected to the server
- Clients may not be able to handle huge response
- Requires bi-directional protocol such as websockets
- Polling is preferred for light clients





