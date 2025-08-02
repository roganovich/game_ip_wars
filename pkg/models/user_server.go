package models


type UserServer struct {
	IP    		int64    			`json:"id"`
	Name  		string 				`json:"name"`						// ФИО
	Email 		string 				`json:"email"`						// Email
	Role 		Role 				`json:"role"`						// Роль
	Status 		int 				`json:"status"`						// Статус
	CreatedAt   time.Time       	`json:"created_at"`              	// Дата создания
}