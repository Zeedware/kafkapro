package main

type Configuration struct {
	loop          int
	inputFileName string
}

func NewConfiguration(loop int, inputFileName string) *Configuration {
	return &Configuration{loop: loop, inputFileName: inputFileName}
}

func (c *Configuration) InputFileName() string {
	return c.inputFileName
}

func (c *Configuration) SetInputFileName(inputFileName string) {
	c.inputFileName = inputFileName
}

func (c *Configuration) Loop() int {
	return c.loop
}

func (c *Configuration) SetLoop(loop int) {
	c.loop = loop
}
