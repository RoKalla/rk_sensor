package sensors

type Sensor interface {
	Id() string
	Type() string
	Timestamp() int64
	Value() float32
	Unit() string
	Start()
	Stop()
}
