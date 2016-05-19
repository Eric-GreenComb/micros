
service WorkHistoryService {
    string Ping()
	string UpdateWorkHistory(1: string profile_id, 2: string json_workhistory)
}
