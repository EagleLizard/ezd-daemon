package jobdef

type JobDef struct {
	Desc string
	Repo struct {
		Name string
	}
	Scripts struct {
		Stop  string
		Start string
	}
}
