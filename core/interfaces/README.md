# Interfaces

Interfaces package contains all the common interfaces implementations and their generated mock methods.

## Generate mock methods

- `mockgen --source core/interfaces/interfaces.go --destination core/interfaces/mocks/mocks.go`
- NOTE: If individual interfaces are long then split them into separate files and then generate mock files separately.