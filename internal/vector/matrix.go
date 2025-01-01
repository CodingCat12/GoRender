package vector

import (
	"math"
)

type Matrix4x4 [4][4]float64

func NewTranslationMatrix(dx, dy, dz float64) Matrix4x4 {
	return Matrix4x4{
		{1, 0, 0, dx},
		{0, 1, 0, dy},
		{0, 0, 1, dz},
		{0, 0, 0, 1},
	}
}

func NewScalingMatrix(sx, sy, sz float64) Matrix4x4 {
	return Matrix4x4{
		{sx, 0, 0, 0},
		{0, sy, 0, 0},
		{0, 0, sz, 0},
		{0, 0, 0, 1},
	}
}

func NewRotationMatrix(dx, dy, dz float64) Matrix4x4 {
	rotationX := NewRotationMatrixX(dx)
	rotationY := NewRotationMatrixY(dy)
	rotationZ := NewRotationMatrixZ(dz)

	return rotationZ.Multiply(rotationY).Multiply(rotationX)
}

func NewRotationMatrixX(angle float64) Matrix4x4 {
	cosA := math.Cos(angle)
	sinA := math.Sin(angle)
	return Matrix4x4{
		{1, 0, 0, 0},
		{0, cosA, -sinA, 0},
		{0, sinA, cosA, 0},
		{0, 0, 0, 1},
	}
}

func NewRotationMatrixY(angle float64) Matrix4x4 {
	cosA := math.Cos(angle)
	sinA := math.Sin(angle)
	return Matrix4x4{
		{cosA, 0, sinA, 0},
		{0, 1, 0, 0},
		{-sinA, 0, cosA, 0},
		{0, 0, 0, 1},
	}
}

func NewRotationMatrixZ(angle float64) Matrix4x4 {
	cosA := math.Cos(angle)
	sinA := math.Sin(angle)
	return Matrix4x4{
		{cosA, -sinA, 0, 0},
		{sinA, cosA, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

func NewWeakPerspectiveMatrix(focalLength float64) Matrix4x4 {
	return Matrix4x4{
		{focalLength, 0, 0, 0},
		{0, focalLength, 0, 0},
		{0, 0, 1, 0},
		{0, 0, -1, 0},
	}
}

func (v Vector3) ApplyMatrix(m Matrix4x4) Vector3 {
	w := m[3][0]*v.X + m[3][1]*v.Y + m[3][2]*v.Z + m[3][3]
	if w == 0 {
		w = 1
	}
	return Vector3{
		X: (m[0][0]*v.X + m[0][1]*v.Y + m[0][2]*v.Z + m[0][3]) / w,
		Y: (m[1][0]*v.X + m[1][1]*v.Y + m[1][2]*v.Z + m[1][3]) / w,
		Z: (m[2][0]*v.X + m[2][1]*v.Y + m[2][2]*v.Z + m[2][3]) / w,
	}
}

func (m Matrix4x4) Multiply(m2 Matrix4x4) Matrix4x4 {
	var result Matrix4x4
	for i := range 4 {
		for j := range 4 {
			result[i][j] = 0
			for k := range 4 {
				result[i][j] += m[i][k] * m2[k][j]
			}
		}
	}
	return result
}
