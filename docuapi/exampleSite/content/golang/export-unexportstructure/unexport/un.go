package unexport

type unexport struct {
	name string
}

func (un unexport) Name() string {
	return un.name
}

func ExportV() unexport {
	return unexport{
		name: "un Value",
	}
}

func ExportP() *unexport {
	return &unexport{
		name: "un Pointer",
	}
}
