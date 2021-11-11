package main

//BarnesHut is our highest level function.
//Input: initial Universe object, a number of generations, and a time interval.
//Output: collection of Universe objects corresponding to updating the system
//over indicated number of generations every given time interval.

import (
	"math"
)

//UpdateUniverse takes the pointer for universe and time and returns pointer to the New Universe
func UpdateUniverse(u *Universe, t float64) *Universe {

	newUniverse := CopyUniverse(u) //new universe is address
	quadTree := BuildQuadTree(*newUniverse)
	for s := range u.stars {
		// update position, velocity and acceleration
		newUniverse.stars[s].Update(*quadTree.root, t)
	}
	return newUniverse
}

//Update function will take pointer as input to return the updated position
func (s *Star) Update(qt Node, t float64) {
	acc := s.NewAccel(qt)
	vel := s.NewVelocity(t)
	pos := s.NewPosition(t)
	s.acceleration, s.velocity, s.position = acc, vel, pos
}

// NewVelocity makes the velocity of this object consistent with the acceleration.
func (s *Star) NewVelocity(t float64) OrderedPair {
	return OrderedPair{
		x: s.velocity.x + s.acceleration.x*t,
		y: s.velocity.y + s.acceleration.y*t,
	}
}

func (s *Star) NewPosition(t float64) OrderedPair {
	return OrderedPair{
		x: s.position.x + s.velocity.x*t + 0.5*s.acceleration.x*t*t,
		y: s.position.y + s.velocity.y*t + 0.5*s.acceleration.y*t*t,
	}
}

// UpdateAccel computes the new accerlation vector for b
func (s *Star) NewAccel(qt Node) OrderedPair {
	//F := ComputeNetForce(qt, s)
	var F OrderedPair
	F = ComputeNetForce(&qt, F)
	return OrderedPair{
		x: F.x / s.mass,
		y: F.y / s.mass,
	}
}

/*func ComputeNetForce(qt Node, s *Star) OrderedPair {
	var netForce OrderedPair
	for child := range qt.children {
		if qt.children[child] == nil {
			if qt.star != s {
				f := ComputeGravitationalForce(*s, *qt.star)
				netForce.Add(f)
			}
		} else {
			ComputeNetForce(*qt.children[child], s)
		}
	}
	return netForce
}*/

/*func ComputeNetForce(n Node, s *Star) OrderedPair {
	var netForce OrderedPair
	if n.star.position.x == s.position.x && n.star.position.y == s.position.y {
		netForce.Add(ComputeGravitationalForce(*s, *n.star))
	}
	if !n.dummyStar {
		netForce.Add(ComputeGravitationalForce(*s, *n.star))
	} else {
		d := Distance(*n.star, *s)
		s1 := n.sector.width
		r := s1 / d
		if r <= 0.5 {
			netForce.Add(ComputeGravitationalForce(*s, *n.star))
		} else {
			for _, child := range n.children {
				if child != nil {
					netForce = ComputeNetForce(*child, s)
				}
			}
		}

	}
	return netForce
}*/

func ComputeNetForce(qt *Node, force OrderedPair) OrderedPair {
	var netForce OrderedPair
	var st *Star
	if qt.star != st && qt.children == nil {
		f := ComputeGravitationalForce(*st, *qt.star)
		netForce.Add(f)
	} else if qt.children == nil {
		d := Distance(*st, *qt.star)
		s := qt.sector.width
		r := s / d
		if r <= theta {
			for i := 0; i < 4; i++ {
				if qt.children[i] != nil {
					netForce = ComputeNetForce(qt.children[i], netForce)
				}
			}
		} else {
			f := ComputeGravitationalForce(*st, *qt.star)
			netForce.Add(f)
		}
	}
	return netForce
}

func (v *OrderedPair) Add(v2 OrderedPair) {
	v.x += v2.x
	v.y += v2.y
}

// Compute the Euclidian Distance between two bodies
func Distance(s1, s2 Star) float64 {
	dx := s1.position.x - s2.position.x
	dy := s1.position.y - s2.position.y
	return math.Sqrt(dx*dx + dy*dy)
}

func ComputeGravitationalForce(s1, s2 Star) OrderedPair {
	d := Distance(s2, s1)
	deltaX := s2.position.x - s1.position.x
	deltaY := s2.position.y - s1.position.y
	F := G * s1.mass * s2.mass / (d * d)
	return OrderedPair{
		x: F * deltaX / d,
		y: F * deltaY / d,
	}
}

func (r *Node) GetQuadrantIndex(star Star, q Quadrant) int {
	width := r.sector.width
	x := r.sector.x
	y := r.sector.y
	xmid := x + width/2
	ymid := y + width/2
	xStar := star.position.x
	yStar := star.position.y

	var quadrantIndex int

	if (x <= xStar) && (x < xmid) && (ymid <= yStar) && (yStar <= y+width) {
		quadrantIndex = 0
	}
	if (xmid <= xStar) && (xStar <= x+width) && (ymid <= yStar) && (yStar <= y+width) {
		quadrantIndex = 1
	}
	if (x <= xStar) && (xStar <= xmid) && (y <= yStar) && (yStar <= ymid) {
		quadrantIndex = 2
	}
	if (xmid <= xStar) && (xStar <= x+width) && (y <= yStar) && (yStar <= y+ymid) {
		quadrantIndex = 3
	}

	return quadrantIndex
}

func (r *Node) GetSector(quadrantIndex int, q Quadrant) Quadrant {
	var x, y, width float64
	if quadrantIndex == 0 {
		x = q.x - q.width/2
		y = q.y + q.width/2
		width = q.width / 2
	}
	if quadrantIndex == 1 {
		x = q.x + q.width/2
		y = q.y + q.width/2
		width = q.width / 2
	}
	if quadrantIndex == 2 {
		x = q.x - q.width/2
		y = q.y - q.width/2
		width = q.width / 2
	}
	if quadrantIndex == 3 {
		x = q.x + q.width/2
		y = q.y - q.width/2
		width = q.width / 2
	}
	return Quadrant{
		x:     x,
		y:     y,
		width: width,
	}
}

func insertStar(r *Node, s Star) {
	quadrantIndex := r.GetQuadrantIndex(s, r.sector)
	newSector := r.GetSector(quadrantIndex, r.sector)
	if r.children[quadrantIndex] == nil && r.dummyStar {
		r.children[quadrantIndex] = &Node{
			star:      CopyStar(&s),
			children:  make([]*Node, 4),
			dummyStar: false,
			sector:    newSector,
		}
	} else {
		if !r.children[quadrantIndex].dummyStar {
			temp := CopyStar(r.children[quadrantIndex].star)
			r.children[quadrantIndex].dummyStar = true
			r.children[quadrantIndex].star.position.x = 0
			r.children[quadrantIndex].star.position.y = 0
			r.children[quadrantIndex].star.mass = 0
			insertStar(r.children[quadrantIndex], *temp)
			//insertStar(r.children[quadrantIndex],s)
		} else {
			insertStar(r.children[quadrantIndex], s)
		}
	}
	r.star.mass = s.mass + r.star.mass
	r.star.position.x = s.position.x*(s.mass) + r.star.position.x
	r.star.position.y = s.position.y*(s.mass) + r.star.position.y
	r.dummyStar = true
}

func BuildQuadTree(u Universe) QuadTree {
	var nd Node
	nd.sector.width = u.width
	nd.sector.x = 0
	nd.sector.y = 0
	nd.dummyStar = true
	nd.children = make([]*Node, 4)
	nd.star = &Star{
		mass: 0,
	}
	for _, star := range u.stars {
		insertStar(&nd, *star)
	}
 	computePositions(&nd)
	var qt QuadTree
	qt.root = &nd
	return qt
}

func computePositions(n *Node){
	if n.dummyStar{
		n.star.position.x=n.star.position.x/n.star.mass
		n.star.position.y=n.star.position.y/n.star.mass
	}
	for _, child :=range(n.children){
		if child != nil {
			computePositions(child)
		}
	}
}


func BarnesHut(initialUniverse *Universe, numGens int, time, theta float64) []*Universe {
	timePoints := make([]*Universe, numGens+1)
	timePoints[0] = initialUniverse //time points is list of address
	// Your code goes here. Use subroutines! :)
	for gen := 0; gen < numGens; gen++ {
		timePoints[gen+1] = UpdateUniverse(timePoints[gen], time)
	}

	return timePoints
}