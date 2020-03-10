package a

func f() {} // OK

type client struct{}

func (c *client) m() {} // OK

func (cli *client) m2() {} // want "cli should be a one or two letter"
