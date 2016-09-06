![awesomesul-logo](img/awesomesul.png)

I'm a big fan of consul, but deploying the thing in a produciton cluster is just absolutely awful. From the docs:

> An agent is the long running daemon on every member of the Consul cluster. It is started by running consul agent. The agent is able to run in either client or server mode. Since all nodes must be running an agent, it is simpler to refer to the node as being either a client or server, but there are other instances of the agent. All agents can run the DNS or HTTP interfaces, and are responsible for running checks and keeping services in sync.

And a quick look at architecture diagrams of production consuls just does not help.

It can also be slow.

Instead we hook into a redis backend (which is very fast, which we trust to scale, and which is a lot simpler and easier to [cluster](http://redis.io/topics/cluster-tutorial)) and instead set this up as a kind of proxy.

It implements, or at least it will once complete, the consul API.

It will scale horizontally. It will be fast.


Building
--

```bash
$ go build
```


Running
---

```bash
$ ./awesomesul -h
Usage of ./awesomesul:
  -db int
        Redis database number
  -listen_address string
        Address on which to listen (default ":8000")
  -redis_address string
        Redis host address (default "localhost:6379")
  -redis_password string
        Redis password, should one exist
```


Licence
--

The MIT License (MIT)
Copyright (c) 2016, James Condron, Financial Times Video

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
