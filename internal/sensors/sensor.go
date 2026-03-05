package sensor

type Sensor interface {
	Id() string
	Type() string
	Timestamp() int64
	Value() float32
	unit() string
}
