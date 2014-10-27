Run

    go run main.go

then connect to localhost on the given port via HTTP to cause the server to
stop (after 3 seconds of supposedly heavy-duty but pretty pointless number
crunching).

Also see http://www.hydrogen18.com/blog/stop-listening-http-server-go.html
([source code](https://github.com/hydrogen18/stoppableListener)) for a more
controlled way to shut down.
