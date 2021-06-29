package middleware
func CheckRole (roles string) (bool){
	return roles == "superadmin"
}
