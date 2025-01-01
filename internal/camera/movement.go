package camera

import (
	"math"
	"renderIdk/internal/vector"

	"github.com/veandco/go-sdl2/sdl"
)

func (c Camera) Move(dx, dy, dz float64) Camera {
	return Camera{
		Position: vector.Vector3{
			X: c.Position.X + dx,
			Y: c.Position.Y + dy,
			Z: c.Position.Z + dz,
		},
		Rotation: c.Rotation,
		NearClip: c.NearClip,
		FarClip:  c.FarClip,
		FovX:     c.FovX,
	}
}

func (c Camera) Rotate(dx, dy, dz float64) Camera {
	return Camera{
		Position: c.Position,
		Rotation: vector.Vector3{
			X: c.Rotation.X + dx,
			Y: c.Rotation.Y + dy,
			Z: c.Rotation.Z + dz,
		},
		NearClip: c.NearClip,
		FarClip:  c.FarClip,
		FovX:     c.FovX,
	}
}

func (c Camera) Zoom(targetFov float64, speed, snap float64) Camera {
	return Camera{
		Position: c.Position,
		Rotation: c.Rotation,
		NearClip: c.NearClip,
		FarClip:  c.FarClip,
		FovX:     moveTo(c.FovX, targetFov, speed, snap),
	}
}

func (c Camera) HandleInput(keys []uint8, deltaTime float64) Camera {
	adjustedCam := c
	movementSpeed := 150 * deltaTime
	rotationSpeed := 1 * deltaTime

	if keys[sdl.SCANCODE_W] != 0 {
		adjustedCam = adjustedCam.Move(
			movementSpeed*math.Sin(adjustedCam.Rotation.Y),
			0,
			-movementSpeed*math.Cos(adjustedCam.Rotation.Y),
		)
	}
	if keys[sdl.SCANCODE_S] != 0 {
		adjustedCam = adjustedCam.Move(
			-movementSpeed*math.Sin(adjustedCam.Rotation.Y),
			0,
			movementSpeed*math.Cos(adjustedCam.Rotation.Y),
		)
	}
	if keys[sdl.SCANCODE_A] != 0 {
		adjustedCam = adjustedCam.Move(
			movementSpeed*math.Cos(adjustedCam.Rotation.Y),
			0,
			movementSpeed*math.Sin(adjustedCam.Rotation.Y),
		)
	}
	if keys[sdl.SCANCODE_D] != 0 {
		adjustedCam = adjustedCam.Move(
			-movementSpeed*math.Cos(adjustedCam.Rotation.Y),
			0,
			-movementSpeed*math.Sin(adjustedCam.Rotation.Y),
		)
	}
	if keys[sdl.SCANCODE_LEFT] != 0 {
		adjustedCam = adjustedCam.Rotate(0, rotationSpeed, 0)
	}
	if keys[sdl.SCANCODE_RIGHT] != 0 {
		adjustedCam = adjustedCam.Rotate(0, -rotationSpeed, 0)
	}
	if keys[sdl.SCANCODE_C] != 0 {
		adjustedCam = adjustedCam.Zoom(40, 0.15, 0.1)
	} else if adjustedCam.FovX < 90 {
		adjustedCam = adjustedCam.Zoom(90, 0.15, 0.1)
	}

	return adjustedCam
}

func moveTo(startVal, endVal, speed, snap float64) float64 {
	if math.Abs(endVal-startVal) < snap {
		return endVal
	}

	return startVal + (endVal-startVal)*speed
}
