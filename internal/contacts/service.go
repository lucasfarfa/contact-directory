package contacts

/*
El service es la capa de lógica de negocio. Es el intermediario entre el handler y el repository.
Ni sabe de HTTP ni sabe cómo se guarda la data: solo sabe qué significa operar sobre contactos.

¿Los datos de entrada son válidos?
    No → return ErrValidacion
    Sí → llamar al repository
         ¿Hubo error?
             Sí → propagarlo (return err)
             No → devolver el resultado

*/

var (
// defino errores globales
//TODO
)

type ContactService interface {
	Create(c Contact) error
	FindByID(id string) (Contact, error)
	FindAll() []Contact
	Update(c Contact) error
	Delete(id string) error
}

type ContactServiceImp struct {
	repo ContactRepository
}

func NewContactService(repo ContactRepository) ContactService { // no devuelve pointer porque es una interface
	return &ContactServiceImp{repo: repo}
}

func (serv *ContactServiceImp) Create(c Contact) error {
	return nil
}

func (serv *ContactServiceImp) FindByID(id string) (Contact, error) {
	return Contact{}, nil
}

func (serv *ContactServiceImp) FindAll() []Contact {
	var c []Contact
	return c
}

func (serv *ContactServiceImp) Update(c Contact) error {
	return nil
}

func (serv *ContactServiceImp) Delete(id string) error {
	return nil
}
