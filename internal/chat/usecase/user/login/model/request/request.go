package request

// Context is a part of use case request model port but consideted valid
type Context struct {
}

// Model is is a use case request model
type Model struct {
	UserName UserName
}

type UserName string

func (un UserName) String() string {
	return string(un)
}