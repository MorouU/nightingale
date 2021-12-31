package appConfig


var DBConfig = map[string]string{
	"dbType":"mysql",
	"host":"localhost",
	"user":"root",
	"pass":"root",
	"database":"ni_v1_0",
	"port":"3306",
	"charset":"utf8",
	"tablePrefix":"ni_",
}

var TableConfig = map[string]map[string]string{
	"userLogin":{"tableName":"user_login",},
	"userInfo":{"tableName":"user_info",},
	"userRegisterAuthCode":{"tableName":"user_register_auth_code",},
	"userForgetAuthCode":{"tableName":"user_forget_auth_code",},
	"userModule":{"tableName":"user_custom_modules",},
	"appModule":{"tableName":"app_public_modules",},
	"appTasks":{"tableName":"app_tasks",},
	"appTaskView":{"tableName":"app_task_records",},
	"adminLogin":{"tableName":"admin_login",},
}