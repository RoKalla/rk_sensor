package domain

type Puller interface {
	StartServer(sensor Sensor)
}

type Sender interface {
	Send(sensor Sensor) error
}

type Sensor interface {
	Id() string
	Type() string
	Timestamp() int64
	Value() float32
	Unit() string
	Start() error
	Stop() error
}
