
service AccountService {
    string Ping()
	string GetAccountByUserId(1: string user_id)

	string GetBillingById(1: string id)
	string GetDealBillingByUserId(1: string user_id, 2: i64 timestamp, 3: i64 pagesize)
	string GetBillingByUserId(1: string user_id, 2: i64 timestamp, 3: i64 pagesize)	
}
