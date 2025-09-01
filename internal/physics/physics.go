package physics

type Vector2 struct {
	X, Y float32
}

func (v Vector2) Add(other Vector2) Vector2 {
	return Vector2{X: v.X + other.X, Y: v.Y + other.Y}
}

func (v Vector2) Scale(factor float32) Vector2 {
	return Vector2{X: v.X * factor, Y: v.Y * factor}
}

type Rectangle struct {
	X, Y, Width, Height float32
}

func (r Rectangle) Contains(x, y float32) bool {
	return x >= r.X && x <= r.X+r.Width &&
		y >= r.Y && y <= r.Y+r.Height
}

func (r Rectangle) Intersects(other Rectangle) bool {
	return r.X < other.X+other.Width &&
		r.X+r.Width > other.X &&
		r.Y < other.Y+other.Height &&
		r.Y+r.Height > other.Y
}

func CheckAABBCollision(x1, y1, w1, h1, x2, y2, w2, h2 float32) bool {
	return x1 < x2+w2 &&
		x1+w1 > x2 &&
		y1 < y2+h2 &&
		y1+h1 > y2
}

func ClampToScreen(x, y, width, height, screenW, screenH, padding float32) (float32, float32) {
	clampedX := x
	clampedY := y

	if clampedX < padding {
		clampedX = padding
	}
	if clampedX+width > screenW-padding {
		clampedX = screenW - width - padding
	}

	if clampedY < padding {
		clampedY = padding
	}
	if clampedY+height > screenH-padding {
		clampedY = screenH - height - padding
	}

	return clampedX, clampedY
}

func IsOutOfBounds(x, y, width, height, screenW, screenH, padding float32) bool {
	return x < padding || x+width > screenW-padding ||
		y < padding || y+height > screenH-padding
}
