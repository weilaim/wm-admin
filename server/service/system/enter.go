package system

type ServiceGroup struct {
	InitDBService
	JwtService
	CasbinService
	OperationRecordService
	UserService
}