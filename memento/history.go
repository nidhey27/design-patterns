package memento

type History struct {
	state []*EditorState
}

func (h *History) Push(state EditorState) {
	h.state = append(h.state, &state)
}

func (h *History) Pop() EditorState {
	lastState := h.state[len(h.state) - 1]
	h.state = h.state[:len(h.state)-1]

	return *lastState
	
}
