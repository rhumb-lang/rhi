package vm

func (vm *VM) opResolve() error {
	version := vm.pop()
	path := vm.pop()
	resolver := vm.pop()

	lib, err := vm.Loader.Load(resolver.Str, path.Str, version)
	if err != nil {
		return err
	}
	vm.push(lib)
	return nil
}
