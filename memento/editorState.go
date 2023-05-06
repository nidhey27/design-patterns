package memento

type EditorState struct {
	content string
}

func (es *EditorState) getState() string {
	return es.content
}
