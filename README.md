zmq
===

Helper library for Go and ØMQ
Currently provides shorthand for json marshalling over ØMQ

    package main

    import "github.com/russellchadwick/zmq"
    import "github.com/pebbe/zmq3"

    func main() {
    	type Hello struct {
    	  Thar string
    	}
	
    	requester, _ := zmq3.NewSocket(zmq3.REQ)
    	defer requester.Close()
    	requester.Connect("ipc://Hello.ipc")

    	hello := Hello {
    		Thar: "Har",
    	}
    	zmq.SendJsonNoReply(requester, hello)
    }
