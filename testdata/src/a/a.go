package sample

func dummyFunc() {} // OK

func is(s string, i int) bool { return false } // OK

type client struct{}

func (c *client) dummyMethod() {} // OK

func (cli *client) dummyMethod2() {} // want "cli should be a one or two letter"
