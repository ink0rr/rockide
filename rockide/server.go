package rockide

type Server struct {
	Rockide *Rockide
}

func NewServer() (*Server, error) {
	server := Server{
		Rockide: NewRockide(),
	}
	return &server, nil
}
