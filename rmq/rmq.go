package rmq

type Connections struct {
	Name    string `json:"name" xml:"name"`
	Host    string `json:"host" xml:"host"`
	Durable bool   `json:"durable" xml:"durable"`
}

var Conn []Connections

func init() {
	var c1 = Connections{
		Name:    "Queue name",
		Host:    "Hostname",
		Durable: true,
	}
	var c2 = Connections{
		Name:    "Queue name2",
		Host:    "Hostname2",
		Durable: false,
	}

	Conn = append(Conn, c1, c2)
}
