package interfaces

import "github.com/hajimehoshi/ebiten/v2"

// GameState represents a state in the game (menu, play, over)
type GameState interface {
	Update() StateTransition
	Draw(screen *ebiten.Image)
}

// StateTransition represents a transition to a new state
type StateTransition struct {
	ShouldTransition bool
	NewState         StateType
	ShouldExit       bool
}

// StateType represents the type of game state
type StateType int

const (
	StateMenu StateType = iota
	StatePlay
	StateOver
)

// GameObject represents any game object that can be updated and drawn
type GameObject interface {
	Update()
	Draw(screen *ebiten.Image)
}

// Movable represents objects that can move
type Movable interface {
	Move()
	GetPosition() (x, y float32)
	SetPosition(x, y float32)
}

// Collidable represents objects that can collide
type Collidable interface {
	GetBounds() (x, y, width, height float32)
	Collides(other Collidable) bool
}

// Resettable represents objects that can be reset to initial state
type Resettable interface {
	Reset()
}
