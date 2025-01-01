package vector

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
