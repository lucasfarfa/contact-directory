package contactos

type Contacto struct {
	ID       int
	Nombre   string
	Apellido string
	Empresa  Empresa
	Email    []string
	Telefono []string
}

type Empresa struct {
}
