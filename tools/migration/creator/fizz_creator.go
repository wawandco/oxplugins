package creator

// FizzCreator model struct for fizz generation files
type FizzCreator struct{}

// Name is the name of the migration type
func (f FizzCreator) Name() string {
	return "fizz"
}

// Create will create 2 .fizz files for the migration
func (f FizzCreator) Create(dir string, args []string) error {
	return nil
}
