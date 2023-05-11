package opurl

func OpRR(host, path, query string) URLer {
	return opURL{
		scheme: "http",
		host:   host,
		path:   path,
		query:  query,
	}
}

func OpWS(host, path, query string) URLer {
	return opURL{
		scheme: "ws",
		host:   host,
		path:   path,
		query:  query,
	}
}
