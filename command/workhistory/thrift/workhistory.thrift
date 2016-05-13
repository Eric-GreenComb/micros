
service WorkHistoryService {
    string Ping()
	string AddWorkHistory(1: string json_workhistory)
	string UpdateWorkHistory(1: string json_workhistory)
}
