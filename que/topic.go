package que

func Create() {

}

func ListAll() map[string]struct{} {

	partitions, err := Conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}

	return m

}
