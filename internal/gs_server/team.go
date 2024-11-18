package gs_server

type Team struct {
	Name    string
	Color   string
	Players []*Player
	Life    int
	Gold    int
	Items   []Item
}
