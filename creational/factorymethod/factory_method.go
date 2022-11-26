package factorymethod

type item interface {
	Name() string
}

func newItem(item string) item {
	switch item {
	case "sword":
		return sword{}
	case "wand":
		return wand{}
	case "axe":
		return axe{}
	case "potion":
		return potion{}
	default:
		return none{}
	}
}

type sword struct{}

func (s sword) Name() string {
	return "sword"
}

type wand struct{}

func (w wand) Name() string {
	return "wand"
}

type axe struct{}

func (a axe) Name() string {
	return "axe"
}

type potion struct{}

func (p potion) Name() string {
	return "potion"
}

type none struct{}

func (n none) Name() string {
	return "none"
}
