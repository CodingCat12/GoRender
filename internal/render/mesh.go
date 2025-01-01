package render

import (
	"renderIdk/internal/camera"
	"renderIdk/internal/vector"

	"github.com/google/uuid"
	"github.com/veandco/go-sdl2/sdl"
)

type Object struct {
	Mesh     Mesh
	Children []Object
}

type Mesh struct {
	Faces []Triangle
	Color sdl.Color
}

type ID [4]int

func (s Scene) AddCuboid(size, position vector.Vector3, R, G, B, A uint8) Scene {
	halfWidth := size.X / 2
	halfHeight := size.Y / 2
	halfDepth := size.Z / 2

	vertices := []vector.Vector3{
		{X: position.X - halfWidth, Y: position.Y - halfHeight, Z: position.Z - halfDepth},
		{X: position.X - halfWidth, Y: position.Y - halfHeight, Z: position.Z + halfDepth},
		{X: position.X - halfWidth, Y: position.Y + halfHeight, Z: position.Z - halfDepth},
		{X: position.X - halfWidth, Y: position.Y + halfHeight, Z: position.Z + halfDepth},
		{X: position.X + halfWidth, Y: position.Y - halfHeight, Z: position.Z - halfDepth},
		{X: position.X + halfWidth, Y: position.Y - halfHeight, Z: position.Z + halfDepth},
		{X: position.X + halfWidth, Y: position.Y + halfHeight, Z: position.Z - halfDepth},
		{X: position.X + halfWidth, Y: position.Y + halfHeight, Z: position.Z + halfDepth},
	}

	triangles := []Triangle{
		{vertices[0], vertices[2], vertices[3]},
		{vertices[0], vertices[3], vertices[1]},
		{vertices[4], vertices[5], vertices[7]},
		{vertices[4], vertices[7], vertices[6]},
		{vertices[0], vertices[1], vertices[5]},
		{vertices[0], vertices[5], vertices[4]},
		{vertices[2], vertices[6], vertices[7]},
		{vertices[2], vertices[7], vertices[3]},
		{vertices[1], vertices[3], vertices[7]},
		{vertices[1], vertices[7], vertices[5]},
		{vertices[0], vertices[4], vertices[6]},
		{vertices[0], vertices[6], vertices[2]},
	}

	mesh := Mesh{
		Faces: triangles,
		Color: sdl.Color{R: R, G: G, B: B, A: A},
	}

	return s.AddObject(Object{Mesh: mesh}, uuid.New())
}

func (s Scene) AddObject(object Object, id uuid.UUID) Scene {
	if s.Objects == nil {
		s.Objects = make(map[uuid.UUID]Object)
	}

	objects := s.Objects
	objects[id] = object
	return Scene{Camera: s.Camera, Objects: objects}
}

func (m Mesh) Render(renderer *sdl.Renderer, cam camera.Camera) {
	for _, triangle := range m.Faces {
		projectedFace := Triangle{
			cam.Project(triangle[0]),
			cam.Project(triangle[1]),
			cam.Project(triangle[2]),
		}

		if !cam.CheckVisible(projectedFace[0], projectedFace[1]) ||
			!cam.CheckVisible(projectedFace[1], projectedFace[2]) ||
			!cam.CheckVisible(projectedFace[2], projectedFace[0]) {
			continue
		}

		fPoints := []sdl.FPoint{
			projectedFace[0].AsFPoint(),
			projectedFace[1].AsFPoint(),
			projectedFace[2].AsFPoint(),
		}

		vertices := []sdl.Vertex{
			{Position: fPoints[0], Color: m.Color},
			{Position: fPoints[1], Color: m.Color},
			{Position: fPoints[2], Color: m.Color},
		}

		renderer.RenderGeometry(nil, vertices, nil)
	}
}

type Triangle [3]vector.Vector3

func (t Triangle) Normal() vector.Vector3 {
	ab := t[1].Subtract(t[0])
	ac := t[2].Subtract(t[0])
	return ab.Cross(ac).Normalize()
}

func (t Triangle) Center() vector.Vector3 {
	return t[0].Translate(t[1]).Translate(t[2]).Scale(vector.Vector3{X: 1.0 / 3, Y: 1.0 / 3, Z: 1.0 / 3})
}
