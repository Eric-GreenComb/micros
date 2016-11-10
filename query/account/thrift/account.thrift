
service AccountService {
    string Ping()
	string GetAccountByUserID(1: string userID)

	string GetBillingByID(1: string ID)
	string GetDealBillingByUserID(1: string userID, 2: i64 timestamp, 3: i64 pagesize)
	string GetBillingByUserID(1: string userID, 2: i64 timestamp, 3: i64 pagesize)	
}
