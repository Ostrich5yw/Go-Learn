package DataStructure

type setType interface{}

type HashSet struct {
	set map[setType]bool
}

func NewSet() *HashSet {
	return &HashSet{set: make(map[setType]bool)}
}

func (p *HashSet) Add(param setType) (b bool) {
	_, ok := p.set[param]
	if ok {
		return false
	} else {
		p.set[param] = true
		return true
	}
}

func (p *HashSet) Delete(param setType) (b bool) {
	_, ok := p.set[param]
	if !ok {
		return false
	} else {
		delete(p.set, param)
		return true
	}
}

func (p *HashSet) Size() (i int) {
	return len(p.set)
}

func (p *HashSet) Contains(param setType) (b bool) {
	_, ok := p.set[param]
	if !ok {
		return false
	} else {
		return true
	}
}

func (p *HashSet) Clear() {
	p.set = make(map[setType]bool)
}
