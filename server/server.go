package server

// RegisterHandler to be serverd by the grpc server upon Run()
func RegisterHandler(h interface{}) {

}

// Run the server. This will start a grpc server and register to service discovery
func Run() {
	c := make(chan bool)
	<-c
}
