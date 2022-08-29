package tcpclient

type Job func(task string, complexity int) (string, error)
