
service WorkHistoryService {
    string Ping()
	string UpdateWorkHistory(1: string profileID, 2: string jsonWorkhistory)
}
