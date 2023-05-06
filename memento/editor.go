package memento

type Editor struct {
	content string
}

func (e *Editor) CreateState() *EditorState {
	return &EditorState{content: e.content}
}

func (e *Editor) SetContent(content string) {
	e.content = content
}

func (e *Editor) Restore(state EditorState) {
	e.content = state.getState()
}

func (e *Editor) GetContent() string {
	return e.content
}
