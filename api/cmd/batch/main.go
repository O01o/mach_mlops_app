package main

import r "mach-mlops/internal/router"

func main() {
	server := r.NewHTTPBatchHandler()
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
