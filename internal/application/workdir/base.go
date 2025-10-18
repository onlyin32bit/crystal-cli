package workdir

type Workdir struct{}

func Explore() (*Workdir, error) {
	return &Workdir{}, nil
}

func (wd *Workdir) Load() {}
