package physics

import "testing"

func TestVector2Add(t *testing.T) {
	v1 := Vector2{X: 1.0, Y: 2.0}
	v2 := Vector2{X: 3.0, Y: 4.0}
	result := v1.Add(v2)

	if result.X != 4.0 || result.Y != 6.0 {
		t.Errorf("Expected (4.0, 6.0), got (%f, %f)", result.X, result.Y)
	}
}

func TestRectangleIntersects(t *testing.T) {
	r1 := Rectangle{X: 0, Y: 0, Width: 10, Height: 10}
	r2 := Rectangle{X: 5, Y: 5, Width: 10, Height: 10}

	if !r1.Intersects(r2) {
		t.Error("Expected rectangles to intersect")
	}

	r3 := Rectangle{X: 20, Y: 20, Width: 10, Height: 10}
	if r1.Intersects(r3) {
		t.Error("Expected rectangles not to intersect")
	}
}

func TestCheckAABBCollision(t *testing.T) {
	if !CheckAABBCollision(0, 0, 10, 10, 5, 5, 10, 10) {
		t.Error("Expected collision")
	}

	if CheckAABBCollision(0, 0, 10, 10, 20, 20, 10, 10) {
		t.Error("Expected no collision")
	}
}

func TestClampToScreen(t *testing.T) {
	x, y := ClampToScreen(-5, -5, 10, 10, 100, 100, 5)
	if x != 5 || y != 5 {
		t.Errorf("Expected (5, 5), got (%f, %f)", x, y)
	}

	x, y = ClampToScreen(95, 95, 10, 10, 100, 100, 5)
	if x != 85 || y != 85 {
		t.Errorf("Expected (85, 85), got (%f, %f)", x, y)
	}
}
