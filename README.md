### Introduction

Redis is an in-memory data structure storage. The name Redis stands for Remote Dictionary Server. In Redis, **RESP** is the backbone of Redis client-server communication. **RESP** is a straightforward text-based protocol, but even with it, you can accomplish some very powerful tasks. So basically, with Redis, you can cache some data that are being used frequently and improve performance. Last week I sat to build a Redis Server and in that process dove a bit deeper to understand the **RESP** protocol. So, let's build a lite version of the Redis Server.

### **RESP** Protocol

As I said earlier it is a simple text-based protocol. In this protocol, the first byte of the request or response represents the type of data it carries. Here is the list.

- `+` stands for String<br/>
- `-` stands for Error Message<br/>
- `:` stands for Integer<br/>
- `$` stands for Bulk String<br/>
- `*` stands for Arrays<br/>

One important thing to note here is each part in the text is separated by a **CRLF** i.e. `\r\n`, which is also called the protocol terminator.
Now let's understand them one by one.

##### Simple String
A simple string `<str>` can be represented by using `+<str>\r\n`, so when Redis deserializes the text, it simply removes the type identification byte i.e. the first byte and the **CRLF**.

##### Error Message
It also has the same format as above, but the difference is it is identified by a `-`(minus) symbol. An error message `<err>` can be represented by `-<err>\r\n` Redis.

##### Integer
In **RESP**, a string is identified by a colon symbol i.e `:`. To represent an integer `<int>`, Resp uses the following format `:[+/-]<int>\r\n` . Interestingly, the 2nd byte is optional, if it's a positive number, but in case of a negative, it is mandatory. After that, you can store your integer followed by **CRLF** (`\r\n`).

##### Bulk Strings
The byte identifier for bulk string is the dollar symbol `$`. The property that separates a bulk string from a simple string is the string length. Bulk strings store their corresponding string length. So, it has two parts mainly `<str-len>` and `<str>`, which are separated by **CRLF**.
Here is an example to store `<str>` as a bulk string in **RESP**, `$<str-len>\r\n<str>\r\n` .

##### Arrays
Resp arrays have the following format `*<array-len>\r\n<array-elem>...\r\n` . So, it is identified by a, followed by the length of the array and then the list of array elements. But the interesting thing here is the elements. What sits in place of array elements are the types we have discussed here till now, i.e. it can be a simple string, error, integer, bulk string, or even an array. So, inside an array, you can store a whole new world ( just kidding, you can't do that :( ).

### Client Server Communication in Redis

So far we have discussed various **RESP** data types and their representation. Now it's time to understand how a Redis client and server communicate. Here are the steps
- First of all, a Redis Client and Server communicate over a **TCP** connection, So before communicating using the **RESP** protocol, you need to set up a TCP connection to send and receive **RESP** data.
- The Redis Client serializes the user input text and sends it as an array of bulk strings over the TCP connection.
- The server deserializes the received message to retrieve the input command. The server replies with one of the **RESP** data types based on the command implementation inside the Redis server.

So far we haven't discussed how null values are represented in **RESP**. Since the client sends an array of bulk strings, **RESP** needs to have a way to represent null both the data types, i.e. in bulk strings as well as arrays.

In the case of bulk strings, a null value can be represented by `$-1\r\n`, where `-1` represents the string length.

In the case of arrays, a null array can be represented by `*-1\r\n`, where `-1` is array length.

### Run the application
To run this application, just run `make run`, and a **REDIS** server will start listening on port `6379` and you can send your **REDIS** commands through `redis-cli`. 
- Please Note before using `redis-cli`, you need to install `redis` first.
