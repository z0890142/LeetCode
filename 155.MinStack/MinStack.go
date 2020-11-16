package MinStack

type node struct {
	val  int
	next *node
}

type MinStack struct {
	stack         *node
	minStackTrack *node
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	newNode := &node{
		val:  x,
		next: this.stack,
	}
	this.stack = newNode
	if this.minStackTrack == nil {
		this.minStackTrack = &node{
			val: x,
		}
	} else {
		if x <= this.minStackTrack.val {
			newMin := &node{
				val:  x,
				next: this.minStackTrack,
			}
			this.minStackTrack = newMin
		}
	}
}

func (this *MinStack) Pop() {
	if this.stack.val == this.minStackTrack.val {
		this.minStackTrack = this.minStackTrack.next
	}
	this.stack = this.stack.next
}

func (this *MinStack) Top() int {
	return this.stack.val
}

func (this *MinStack) GetMin() int {
	return this.minStackTrack.val
}
