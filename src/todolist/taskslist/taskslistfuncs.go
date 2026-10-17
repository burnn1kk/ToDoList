package tasklist

func Add(tasks map[string]Task, name, text string) (map[string]Task, error) {

}

func List(tasks map[string]Task) []Task {
	var list []Task
	for _, v := range tasks {
		list = append(list, v)
	}
	return list
}
