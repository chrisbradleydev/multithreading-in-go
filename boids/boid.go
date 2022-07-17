package main

import (
	"math"
	"math/rand"
	"time"
)

type Boid struct {
	position Vector2D
	velocity Vector2D
	id int
}

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func (b *Boid) calcAcceleration() Vector2D {
	upper, lower := b.position.AddV(viewRadius), b.position.AddV(-viewRadius)
	avgVelocity := Vector2D{0, 0}
	count := 0.0

	for i := math.Max(lower.x, 0); i <= math.Min(upper.x, screenWidth); i++ {
		for j := math.Max(lower.y, 0); j <= math.Min(upper.y, screenHeight); j++ {
			if otherBoidId := boidMap[int(i)][int(j)]; otherBoidId != -1 && otherBoidId != b.id {
				if dist := boids[otherBoidId].position.Distance(b.position); dist < viewRadius {
					count++
					avgVelocity = avgVelocity.Add(boids[otherBoidId].velocity)
				}
			}
		}
	}

	accel := Vector2D{0, 0}
	if count > 0 {
		avgVelocity = avgVelocity.DivideV(count)
		accel = avgVelocity.Subtract(b.velocity).MultiplyV(adjRate)
	}

	return accel
}

func (b *Boid) moveOne() {
	b.velocity = b.velocity.Add(b.calcAcceleration()).limit(-1, 1)
	boidMap[int(b.position.x)][int(b.position.y)] = -1
	b.position = b.position.Add(b.velocity)
	boidMap[int(b.position.x)][int(b.position.y)] = b.id

	next := b.position.Add(b.velocity)

	if next.x >= screenWidth || next.x < 0 {
		b.velocity = Vector2D{-b.velocity.x, b.velocity.y}
	}
	if next.y >= screenHeight || next.y < 0 {
		b.velocity = Vector2D{b.velocity.x, -b.velocity.y}
	}
}

func createBoid(id int) {
	b := Boid{
		position: Vector2D{rand.Float64() * screenWidth, rand.Float64() * screenHeight},
		velocity: Vector2D{rand.Float64() * 2 - 1, rand.Float64() * 2 - 1},
		id: id,
	}

	boids[id] = &b
	boidMap[int(b.position.x)][int(b.position.y)] = b.id

	go b.start()
}
