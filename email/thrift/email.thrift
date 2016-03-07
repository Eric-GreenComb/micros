struct Email{
	1: string host
	2: string user
	3: string password
	4: string to
	5: string subject
	6: string body
	7: string mailtype
}

service EmailService {
	bool SendEmail(1: Email email)
}
