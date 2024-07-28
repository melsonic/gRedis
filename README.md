### Simple LightWeight Redis-Server in Golang

#### Running the redis server

To run the redis server simply run `make run`

#### Prerequisites

To test the Redis Server you would want to install redis-client, which comes preinstalled when you install [redis](https://redis.io/docs/latest/operate/oss_and_stack/install/install-redis/)

#### Ping Redis Server 
On a terminal window run `redis-client ping`

#### TODO 
- Add RESP Protocol theory, and how redis communicates using this protocol through a tcp connection.
- Add Test Cases.

#### Problems faces 
- Supporting concurrent clients was quite challenging.
