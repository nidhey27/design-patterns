package state

type Canvas struct {
	tool Tool
}

func (c *Canvas) MouseDown() {
	c.tool.MouseDown()
}

func (c *Canvas) MouseUp() {
	c.tool.MouseUp()
}

func (c *Canvas) SetTool(tool Tool) {
	c.tool = tool
}

func (c *Canvas) GetTool() *Tool {
	return &c.tool
}
