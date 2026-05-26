package contacts

import (
	"errors"
	"sync"
)

// Me conviene tener definod los errores aca, asi en el handler se que codigo devolver
var (
	ErrContactAlreadyExists = errors.New("contact already exists")
	ErrContactNotFound      = errors.New("contact not found")
	ErrNoContactsFound      = errors.New("no contacts created")
)

type ContactRepository interface {
	Create(c Contact) error
	FindByID(id string) (Contact, error)
	FindAll() ([]Contact, error)
	Update(c Contact) error
	Delete(id string) error
}

// Cuando quiera pasar a SQL, solamente creo
// una nueva struct tipo SQLRepo
// e implemento los metodos de "ContactRepository"
type InMemoryRepository struct {
	contacts map[string]Contact
	mu       sync.RWMutex
	/*
		Campos en minuscula. contacts y mu van sin exportar.
		Nadie de afuera del package debe tocar el map ni el mutex directamente —solo a través de los métodos.
		Eso es justamente lo que protege la concurrencia
	*/
}

// Debo hacer un constructor para el repository
func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		contacts: make(map[string]Contact),
		// mutex no se inicializa explicitamente
		// porque  su zero value ya es valido
	}
}

func (repo *InMemoryRepository) Create(c Contact) error {
	// El lock debe ir antes de consultar si existe contacto
	// y antes de intentar crear.
	repo.mu.Lock()
	defer repo.mu.Unlock()

	// Controlo que no exista el usuario a crear
	_, duplicateContact := repo.contacts[c.ID]
	if duplicateContact {
		return ErrContactAlreadyExists
	}

	// Creo contacto
	repo.contacts[c.ID] = c

	return nil
}

func (repo *InMemoryRepository) FindByID(id string) (Contact, error) {
	// Aca uso RLock ya que no es exclusivo, solamente queiro leer
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	contact, found := repo.contacts[id]

	if !found {
		return Contact{}, ErrContactNotFound
	}

	return contact, nil
}

func (repo *InMemoryRepository) FindAll() ([]Contact, error) { // por mas que aca no use error, lo dejo asi ya que mas adelante paso a SQL y ahi si necesito error
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	result := make([]Contact, 0, len(repo.contacts))
	// Convierto el map a slice
	for _, c := range repo.contacts {
		result = append(result, c)
	}

	return result, nil // retorno todos
}

func (repo *InMemoryRepository) Update(c Contact) error {
	// Simil create
	repo.mu.Lock()
	defer repo.mu.Unlock()

	// Controlo que exista el usuario a actualizar
	_, found := repo.contacts[c.ID]
	if !found {
		return ErrContactNotFound
	}

	// Actualizo
	repo.contacts[c.ID] = c

	return nil
}

func (repo *InMemoryRepository) Delete(id string) error {

	repo.mu.Lock()
	defer repo.mu.Unlock()

	// Controlo que exista el usuario a borrar
	_, found := repo.contacts[id]
	if !found {
		return ErrContactNotFound
	}

	// Elimino contacto
	delete(repo.contacts, id)

	return nil
}
