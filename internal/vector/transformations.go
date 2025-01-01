package vector

import "math"

func (v Vector3) Translate(v2 Vector3) Vector3 {
	translationMatrix := NewTranslationMatrix(v2.X, v2.Y, v2.Z)
	return v.ApplyMatrix(translationMatrix)
}

func (v Vector3) Scale(v2 Vector3) Vector3 {
	return Vector3{
		v.X * v2.X,
		v.Y * v2.Y,
		v.Z * v2.Z,
	}
}

func (v Vector3) Rotate(angles Vector3) Vector3 {
	rotationMatrix := NewRotationMatrix(angles.Raw())
	return v.ApplyMatrix(rotationMatrix)
}

func (v Vector3) Project(focalLength float64) Vector3 {
	if v.Z == 0 {
		return v
	}

	projX := (v.X / v.Z) * focalLength
	projY := (v.Y / v.Z) * focalLength

	return Vector3{
		projX,
		projY,
		v.Z,
	}
}

func (v Vector3) Subtract(v2 Vector3) Vector3 {
	return Vector3{
		v.X - v2.X,
		v.Y - v2.Y,
		v.Z - v2.Z,
	}
}

func (v Vector3) Cross(v2 Vector3) Vector3 {
	return Vector3{
		X: v.Y*v2.Z - v.Z*v2.Y,
		Y: v.Z*v2.X - v.X*v2.Z,
		Z: v.X*v2.Y - v.Y*v2.X,
	}
}

func (v Vector3) Normalize() Vector3 {
	magnitude := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	return Vector3{X: v.X / magnitude, Y: v.Y / magnitude, Z: v.Z / magnitude}
}

func (v1 Vector3) Dot(v2 Vector3) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}
