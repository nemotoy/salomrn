package a

func f() {} // OK

type client struct{}

func (c *client) m() {} // OK

func (cl *client) m2() {} // OK

func (cli *client) m3() {} // want "cli should be a one or two letter"
