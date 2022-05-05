package p0489robotroomcleaner

type Robot struct {
}

func (r *Robot) Move() bool { return true }
func (r *Robot) TurnLeft()  {}
func (r *Robot) TurnRight() {}
func (r *Robot) Clean()     {}

var deltas = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func cleanRoom(robot *Robot) {
	c := Cleaner{
		robot:   robot,
		visited: make(map[[2]int]struct{}),
	}
	c.visit(0, 0, 0)
}

type Cleaner struct {
	robot   *Robot
	visited map[[2]int]struct{}
}

func (c *Cleaner) back() {
	c.robot.TurnRight()
	c.robot.TurnRight()
	c.robot.Move()
	c.robot.TurnRight()
	c.robot.TurnRight()
}

func (c *Cleaner) ok(x, y int) bool {
	if _, exists := c.visited[[2]int{x, y}]; exists {
		return false
	}
	return c.robot.Move()
}

func (c *Cleaner) visit(x, y, d int) {
	c.robot.Clean()
	pos := [2]int{x, y}
	c.visited[pos] = struct{}{}
	for i := 0; i < 4; i++ {
		xx := x + deltas[d][0]
		yy := y + deltas[d][1]
		if c.ok(xx, yy) {
			c.visit(xx, yy, d)
			c.back()
		}
		d = (d + 1) % 4
		c.robot.TurnRight()
	}
}
