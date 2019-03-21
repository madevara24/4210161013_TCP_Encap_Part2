# 4210161013_TCP_Encap_Part2
Homework for Multiplayer Game Design class
Sending encapsulated data using struct from client to server and back to client again

Steps
1. Server listens for tcp connection on port 1013
2. Client dials to server using tcp
3. Client create the struct data and encodes it using the "gob" library
4. Client sends the encoded data
5. Server receives the data and decodes it also using "gob" library
6. After printing the results, server sends the data back to the client
