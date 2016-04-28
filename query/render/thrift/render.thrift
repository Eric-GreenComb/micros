
service RenderService {
    string Ping()
	string RenderHello(1: string tmpl,2: string name)
}
