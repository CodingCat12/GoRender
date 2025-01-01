package render

import (
	"math/rand"
	"renderIdk/internal/camera"
	"renderIdk/internal/vector"

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

	return s.AddObject(Object{Mesh: mesh})
}

func (s Scene) AddObject(object Object) Scene {
	objects := make(map[ID]Object)
	for id, obj := range s.Objects {
		objects[id] = obj
	}

	id := ID{rand.Int(), rand.Int(), rand.Int(), rand.Int()}
	for {
		if _, ok := s.Objects[id]; !ok {
			break
		}
		id = ID{rand.Int(), rand.Int(), rand.Int(), rand.Int()}
	}

	objects[id] = object
	return Scene{Camera: s.Camera, Objects: objects}
}

func (m Mesh) Render(renderer *sdl.Renderer, cam camera.Camera) {
	for _, triangle := range m.Faces {
		projectedVertices := []vector.Vector3{
			cam.Project(triangle[0]),
			cam.Project(triangle[1]),
			cam.Project(triangle[2]),
		}

		if !cam.CheckVisible(projectedVertices[0], projectedVertices[1]) ||
			!cam.CheckVisible(projectedVertices[1], projectedVertices[2]) ||
			!cam.CheckVisible(projectedVertices[2], projectedVertices[0]) {
			continue
		}

		fPoints := []sdl.FPoint{
			projectedVertices[0].AsFPoint(),
			projectedVertices[1].AsFPoint(),
			projectedVertices[2].AsFPoint(),
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
