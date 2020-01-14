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