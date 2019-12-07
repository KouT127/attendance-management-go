migrate:
	@echo "start migrate..."
	@migrate -source file://backend/database/migration/  -database 'mysql://root:@tcp(localhost:3306)/attendance_management' up

show-migrations:
	 mysqldef -uroot attendance_management --export > schema.sql

mysqldef-dry:
	 mysqldef -uroot attendance_management --dry-run < schema.sql

mysqldef:
	 mysqldef -uroot attendance_management < schema.sql