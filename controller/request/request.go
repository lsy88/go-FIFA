package request

type ReqAdminLogin struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ReqAdminRegister struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	//State             string     `json:"state"`
}

type ReqClassicMes struct {
	FinalResult string `json:"final_result"`
	Title       string `json:"title"`
	AllList     bool   `json:"all_list"`
}

type ReqPlayerGoal struct {
	Name string `json:"name"`
	Goal string `json:"goal"`
}
