package config

type Configuration interface {
	GetString(name string) (configValue string, found bool)
	GetInt(name string) (configValue int, found bool)
	GetBool(name string) (configValue bool, found bool)
	GetFloat(name string) (configValue float64, found bool)
	GetStringDefault(name, defVal string) (configValue string)
	GetIntDefault(name, defVal int) (configValue int)
	GetBoolDefault(name, defVal bool) (configValue bool)
	GetFloatDefault(name, defVal float64) (configValue float64)
	GetSection(sectionName string) (section Configuration, found bool)
}
