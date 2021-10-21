package rivers

type Store interface {
	GetAll()
	GetByRefID(refid string)
	GetByName(name string)
}
