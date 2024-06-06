package vector

type Vector struct {
	X float64
	Y float64
}

type VectorRange struct {
	Start Vector
	End   Vector
}

type SpanVector struct {
	Start int
	End   int
}

func Vec(x float64, y float64) Vector {
	return Vector{X: x, Y: y}
}

func (v1 Vector) Equals(v2 Vector) bool {
	return v1.X == v2.X && v1.Y == v2.Y
}

func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{X: v1.X + v2.X, Y: v1.Y + v2.Y}
}

func (v1 Vector) Subtract(v2 Vector) Vector {
	return Vector{X: v1.X - v2.X, Y: v1.Y - v2.Y}
}
func (v1 Vector) Multiply(v2 Vector) Vector {
	return Vec(v1.X*v2.X, v1.Y*v2.Y)
}
func (v1 Vector) Divide(v2 Vector) Vector {
	return Vec(v1.X/v2.X, v1.Y/v2.Y)
}

func normalValue(val float64) float64 {
	if val > 0 {
		return 1
	}
	if val < 0 {
		return -1
	}
	return 0
}

func (v1 Vector) ToNormal() Vector {
	x := normalValue(v1.X)
	y := normalValue(v1.Y)
	return Vector{X: x, Y: y}
}

func (v1 Vector) MultiplyScalar(s float64) Vector {
	return Vec(v1.X*s, v1.Y*s)
}
func (v1 Vector) DivideScalar(s float64) Vector {
	return Vec(v1.X/s, v1.Y/s)
}

func From(props struct{ x, y float64 }) Vector {
	return Vector{X: props.x, Y: props.y}
}

func Midpoint(v1 Vector, v2 Vector) Vector {
	return Vec(
		(v1.X+v2.X)/2,
		(v1.Y+v2.Y)/2,
	)
}

var Up = Vec(0, -1)
var Down = Vec(0, 1)
var Left = Vec(-1, 0)
var Right = Vec(1, 0)
