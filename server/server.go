package server

func Init() {
	// create a gin router
	r := NewRouter()
	// start the router (by default it serves on :8080 unless a PORT environment variable was defined)
	r.Run()
}
