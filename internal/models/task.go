package models

type Task struct {
	TotalTasks string `json:"total_tasks"`
	Running string `json:"running"`
	Sleeping string `json:"sleeping"`
	Zombie string `json:"zombie"`
}

func NewTasks()*Task{
	return &Task{
		TotalTasks: "",
		Running:    "",
		Sleeping:   "",
		Zombie:     "",
	}
}

