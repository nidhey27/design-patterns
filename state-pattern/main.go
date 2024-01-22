package main

import "fmt"

// State
type editingMode interface {
	write(text string)
	edit(text string)
}

// Concrete implementation of states
type insertMode struct{}

func (m *insertMode) write(text string) {
	fmt.Println("Inserting text:", text)
}

func (m *insertMode) edit(text string) {
	fmt.Println("Cannot edit in Insert Mode. Switch to Edit Mode.")
}

type editMode struct{}

func (m *editMode) write(text string) {
	fmt.Println("Cannot write in Edit Mode. Switch to Insert Mode.")
}

func (m *editMode) edit(text string) {
	fmt.Println("Editing text:", text)
}

// Context
type documentEditor struct {
	currentMode editingMode
}

func newDocumentEditor() *documentEditor {
	return &documentEditor{
		currentMode: &insertMode{},
	}
}

func (d *documentEditor) setEditingMode(mode editingMode) {
	d.currentMode = mode
}

func (d *documentEditor) write(text string) {
	d.currentMode.write(text)
}

func (d *documentEditor) edit(text string) {
	d.currentMode.edit(text)
}

// Client
func main() {
	editor := newDocumentEditor()

	// Writing in Insert Mode
	editor.write("This is the first line.")

	// Switching to Edit Mode
	editor.setEditingMode(&editMode{})

	editor.write("This is the second line.")

	// Editing in Edit Mode
	editor.edit("This is the updated first line.")

	// Switching back to Insert Mode
	editor.setEditingMode(&insertMode{})

	// Writing in Insert Mode
	editor.write("This is the second line.")
}
