package contacts

import (
	"errors"
	"strings"
)

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
	ErrContactIdRequired = errors.New("contact id is required")
)

type ContactService interface {
	CreateContact(c Contact) error
	ListContact(contactId string) (Contact, error)
	ListAll() ([]Contact, error)
	UpdateContact(c Contact) error
	DeleteContact(contactId string) error
}

type contactService struct { // lo dejo en minuscula asi es privado y dejo solo publica la interfaz
	repo ContactRepository
}

func NewContactService(repo ContactRepository) ContactService { // no devuelve pointer porque es una interface
	return &contactService{repo: repo}
}

func (serv *contactService) CreateContact(c Contact) error {
	return nil
}

func (serv *contactService) ListContact(contactId string) (Contact, error) {

	if strings.TrimSpace(contactId) == "" {
		return Contact{}, ErrContactIdRequired
	}

	return serv.repo.FindByID(contactId)
}

func (serv *contactService) ListAll() ([]Contact, error) {

	//var contacts []Contact
	contacts, err := serv.repo.FindAll()

	if err != nil {
		// implement
	}

	return contacts, err
}

func (serv *contactService) UpdateContact(c Contact) error {
	return nil
}

func (serv *contactService) DeleteContact(contactId string) error {

	if strings.TrimSpace(contactId) == "" {
		return ErrContactIdRequired
	}

	return serv.repo.Delete(contactId)
}
