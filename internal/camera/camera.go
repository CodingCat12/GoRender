package camera

import (
	"math"

	"renderIdk/internal/config"
	"renderIdk/internal/vector"
)

type Camera struct {
	Position          vector.Vector3
	Rotation          vector.Vector3
	NearClip, FarClip float64
	FovX              float64
}

func (c Camera) Project(v vector.Vector3) vector.Vector3 {
	centerX := config.WinWidth / 2.0
	centerY := config.WinHeight / 2.0

	translationMatrix := vector.NewTranslationMatrix(c.Position.Raw())
	rotationMatrix := vector.NewRotationMatrix(c.Rotation.Raw())
	matrix := rotationMatrix.Multiply(translationMatrix)

	return v.ApplyMatrix(matrix).Project(c.FocalLength()).Translate(vector.Vector3{X: centerX, Y: centerY, Z: 0})
}

func (c Camera) FocalLength() float64 {
	fovRadians := c.FovX * math.Pi / 180.0
	return config.WinWidth / (2 * math.Tan(fovRadians/2))
}

func (c Camera) CheckVisible(v ...vector.Vector3) bool {
	for _, vec := range v {
		if vec.Z < c.NearClip || vec.Z > c.FarClip {
			return false
		}
	}

	return true
}
