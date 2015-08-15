package entities

type Action struct {
	Name        string
	Path        string
	Permissions []string
	ActionType  string
	Label       string
	Viewmode    string
}
