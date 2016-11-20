
service RenderService {
    string Ping()
	string RenderTpl(1: string tplname, 2: map<string,string> keyMap)
}