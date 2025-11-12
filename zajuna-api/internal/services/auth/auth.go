package auth

// Interfaz común para los plugins de autenticación
type AuthPlugin interface {
	PreventLocalPasswords() bool
	Login(username, password string) (bool, error)
}

// Registro global de plugins disponibles
var registry = map[string]AuthPlugin{}

// Registrar un nuevo plugin
func Register(name string, plugin AuthPlugin) {
	registry[name] = plugin
}

// Obtener un plugin por nombre
func Get(name string) (AuthPlugin, bool) {
	p, ok := registry[name]
	return p, ok
}
