# Long Polling

#### What is Long Polling 

- Client sends a request ( that typically is a longr running job such as compressing a blob )
- The server queues the request
- The client can send for a request checking the status of the job
- Server doesnt respond until job is done processing 
- So we have a handle, the client can disconnect and check at a later time 
