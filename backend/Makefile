migrate:
	@echo "start migrate..."
	@migrate -source file://database/migrations/  -database 'mysql://root:@tcp(127.0.0.1:3306)/attendance_management' up

show-migrations:
	 mysqldef -uroot attendance_management --export > schema.sql

mysqldef-dry:
	 mysqldef -uroot attendance_management --dry-run < schema.sql

mysqldef:
	 mysqldef -uroot attendance_management < schema.sql

run:
	@echo "started server"
	realize start --name="attendance-management" --server --run

mockgen:
	@echo "mock generate"
	mockgen -source ./repositories/common.go -destination ./repositories/mock/common.go
	mockgen -source ./repositories/attendance.go -destination ./repositories/mock/attendance.go -aux_files=github.com/KouT127/attendance-management/repositories=repositories/common.go
	mockgen -source ./repositories/user.go -destination ./repositories/mock/user.go
	mockgen -source ./usecases/attendance.go -destination ./usecases/mock/attendance.go
	mockgen -source ./usecases/user.go -destination ./usecases/mock/user.go