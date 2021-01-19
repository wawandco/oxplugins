package model

var modelTestTemplate string = `package models

func (ms *ModelSuite) Test_{{.Name}}() {
	ms.Fail("This test needs to be implemented!")
}`
