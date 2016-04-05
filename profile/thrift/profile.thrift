
struct ProfileSearchCondition{
	1: i32 serial_number
	2: i32 hours_billed
	3: i32 available_hours
	4: i32 job_success
	5: i32 last_activity
	6: i32 freelancer_type		
	7: i32 hourly_rate	
	8: i32 region_id					
	9: string search_key
}

service ProfileService {
	string GetProfile(1: string profile_id)
	string SearchProfiles(1: ProfileSearchCondition profile_search_condition, 2: i64 timestamp, 3: i64 pagesize)		
}
