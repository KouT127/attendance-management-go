migrate:
	@echo "start migrate..."
	@migrate -source file://infrastructure/sqlstore/migrations/  -database 'mysql://root:root@tcp(127.0.0.1:3306)/attendance_management' up

migrate-test:
	@echo "start migrate test database"
	@migrate -source file://infrastructure/sqlstore/migrations/  -database 'mysql://root:root@tcp(127.0.0.1:3306)/test_attendance_management' up


show-migrations:
	 mysqldef -uroot attendance_management --export > schema.sql

mysqldef-dry:
	 mysqldef -uroot attendance_management --dry-run < schema.sql

mysqldef:
	 mysqldef -uroot attendance_management < schema.sql

run:
	@echo "started server"
	realize start --name="attendance-management" --server --run

generate:
	@echo "go generate"
	go generate ./infrastructure/sqlstore