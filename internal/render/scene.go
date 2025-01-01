package render

import (
	"renderIdk/internal/camera"

	"slices"

	"github.com/google/uuid"
	"github.com/veandco/go-sdl2/sdl"
)

type Scene struct {
	Camera  camera.Camera
	Objects map[uuid.UUID]Object
}

func (s Scene) Render(renderer *sdl.Renderer) {
	objects := s.AllObjects()

	slices.SortFunc(objects, func(a, b Object) int {
		avgDepthA := averageDepth(a.Mesh, s.Camera)
		avgDepthB := averageDepth(b.Mesh, s.Camera)

		if avgDepthA > avgDepthB {
			return -1
		}
		if avgDepthA < avgDepthB {
			return 1
		}
		return 0
	})

	for _, obj := range objects {
		obj.Mesh.Render(renderer, s.Camera)
	}
}

func (s Scene) AllObjects() []Object {
	result := make([]Object, 0)

	var getObjects func(obj Object)
	getObjects = func(obj Object) {
		result = append(result, obj)
		for _, child := range obj.Children {
			getObjects(child)
		}
	}

	for _, object := range s.Objects {
		getObjects(object)
	}

	return result
}

func averageDepth(mesh Mesh, cam camera.Camera) float64 {
	var totalDepth float64
	var count int

	for _, triangle := range mesh.Faces {
		for _, vertex := range triangle {
			projectedVertex := cam.Project(vertex)
			totalDepth += projectedVertex.Z
			count++
		}
	}

	return totalDepth / float64(count)
}
