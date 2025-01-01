package vector

import "github.com/veandco/go-sdl2/sdl"

type Vector3 struct {
	X, Y, Z float64
}

func (v Vector3) AsVector2() Vector2 {
	return Vector2{v.X, v.Y}
}

func (v Vector3) Raw() (float64, float64, float64) {
	return v.X, v.Y, v.Z
}

type Vector2 struct {
	X, Y float64
}

type Line [2]Vector2

func (l Line) Raw() (float32, float32, float32, float32) {
	return float32(l[0].X), float32(l[0].Y), float32(l[1].X), float32(l[1].Y)
}

func (v Vector3) AsFPoint() sdl.FPoint {
	return sdl.FPoint{X: float32(v.X), Y: float32(v.Y)}
}
