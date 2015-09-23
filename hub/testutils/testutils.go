package testutils

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"testing"
	"time"

	msghub "github.com/Everyplay/developer-assignment-backend/hub"
)

func getFreePort() string {
	l, _ := net.Listen("tcp", ":0")
	defer l.Close()
	return l.Addr().String()
}

// RunHubTests runs a pre-defined test suite against the hub
func RunHubTests(hub msghub.MessageHub, messageEncoder msghub.MessageEncoder) {
	if hub == nil {
		panic("Hub should not be nil")
	}

	if messageEncoder == nil {
		panic("MessageEncoder should not be nil")
	}
	port := getFreePort()
	err := hub.Start(port)
	if err != nil {
		panic(err)
	}

	conn, err := net.Dial("tcp", port)
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 2)

	go conn.Write(messageEncoder(hub.ClientIDs(), []byte("FOOBAR")))

	response, err := ioutil.ReadAll(io.LimitReader(conn, 6))
	if err != nil {
		panic(err)
	}

	if len(response) > 0 {
		panic("No result received")
	}

	conn.Close()

	var connections []net.Conn
	for i := 0; i < 50; i++ {
		conn, err := net.Dial("tcp", port)
		if conn != nil {
			// Closed once test completes
			defer conn.Close()
		}
		if err != nil {
			panic(err)
		}
		connections = append(connections, conn)
		go func(conn net.Conn) {
			buf := make([]byte, 255)
			for {
				if _, err := conn.Read(buf); err == io.EOF {
					break
				}
			}
		}(conn)
	}

	time.Sleep(time.Second * 2)

	payload := []byte("FOOBAR")
	result := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			connections[0].Write(messageEncoder(hub.ClientIDs(), payload))
		}
	})

	fmt.Printf("Short message benchmark\n%s\n", result.String())

	payload = []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis sed est id mi blandit fringilla vulputate nec urna. Duis non porttitor arcu. Mauris ac ullamcorper turpis, ac tincidunt risus. In rutrum efficitur porttitor. Cras scelerisque eu mi ut tristique. Phasellus enim elit, pretium ut mi vel, semper interdum nisl. Duis gravida blandit risus, a semper ipsum lacinia quis. Nam eros purus, congue in metus id, volutpat dapibus velit. Cras ut dictum libero, non placerat quam. Vivamus sem justo, varius at magna sed, blandit consequat mi. Cras viverra, orci nec feugiat ullamcorper, mauris erat tincidunt nisi, nec rutrum neque est a libero. Nullam pharetra dolor at erat elementum convallis. Phasellus dictum fermentum odio non eleifend. Etiam scelerisque, neque a fringilla molestie, purus turpis posuere erat, ut pulvinar nisl nisl nec nisl. In pellentesque risus sem, id pretium eros gravida sit amet. In vel massa justo. Fusce euismod mattis massa. Fusce at nibh in est condimentum luctus. Integer a molestie arcu. Suspendisse aliquam venenatis nisl, sit amet aliquam ante convallis quis. Praesent nec ipsum lectus. Ut elementum pretium mollis. Etiam tincidunt sapien felis, eget aliquet justo tincidunt at. Integer turpis sem, feugiat quis lorem sed, scelerisque lacinia massa. Aliquam vitae urna et erat sodales accumsan a a enim. Nunc eget diam tristique, ornare nibh sed, laoreet ligula. Mauris sollicitudin consectetur elit nec eleifend. Donec in diam ut ligula porttitor vulputate. Integer finibus, tellus vitae sagittis tincidunt, felis augue pulvinar enim, consectetur sollicitudin lorem lacus vel sem. Mauris condimentum et dolor ac interdum. Praesent bibendum nulla nec dui tempus, non blandit augue iaculis. In pretium erat vel odio dictum, et rhoncus urna tristique. Mauris ut risus orci. Mauris cursus posuere felis, et accumsan ante consequat ac. Cras convallis luctus consequat.")
	result = testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			connections[0].Write(messageEncoder(hub.ClientIDs(), payload))
		}
	})

	fmt.Printf("Long message benchmark\n%s\n", result.String())

	err = hub.Stop()
	if err != nil {
		panic(err)
	}
}
