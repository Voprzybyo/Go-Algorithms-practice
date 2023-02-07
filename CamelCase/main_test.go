package main

import "testing"

func TestFormatCamelCase(t *testing.T) {

	type test struct {
		data   string
		answer string
	}

	tests := []test{

		{"S;M;plasticCup()", "plastic cup"},
		{"S;M;plasticCupSecondary()", "plastic cup secondary"},
		{"C;V;mobile phone", "mobilePhone"},
		{"C;C;coffee machine", "CoffeeMachine"},
		{"S;C;LargeSoftwareBook", "large software book"},
		{"C;M;white sheet of paper", "whiteSheetOfPaper()"},
		{"S;V;pictureFrame", "picture frame"},
		{"C;V;can of coke", "canOfCoke"},
		{"S;M;sweatTea()", "sweat tea"},
		{"S;V;epsonPrinter", "epson printer"},
		{"C;M;santa claus", "santaClaus()"},
		{"C;C;mirror", "Mirror"},
		{"S;V;iPad", "i pad"},
		{"C;M;mouse pad", "mousePad()"},
		{"C;C;code swarm", "CodeSwarm"},
		{"S;C;OrangeHighlighter", "orange highlighter"},
	}

	for _, v := range tests {
		x := formatCamelCase(v.data)

		if x != v.answer {
			t.Error("Got:", x, "Expected:", v.answer)
		}
	}
}
