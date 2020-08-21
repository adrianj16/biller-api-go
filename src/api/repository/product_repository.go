package repository

type Product struct{}

func (p *Product) GetAll() interface{} {
	return map[string]string{"Adrian": "Ugas"}
}
