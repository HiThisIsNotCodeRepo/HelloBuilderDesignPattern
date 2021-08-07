# HelloBuilderDesignPattern

> [Source](https://golangbyexample.com/builder-pattern-golang/)

## Flow

1. Generate `Builder`.
2. Generate `Director` from `Builder`.
3. `Builder` create object.

## Similarity with AbstractFactory

1. Use `Director` to generate object , like `Factory`.

## Difference with AbstractFactory

1. `Director` encapsulate the building process.
2. `Director` will return actual object not interface.

```
func (d *director) buildHouse() house {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}
```

## Interface vs Concrete Type

1. When each object's implementation is different from each other use interface. Like `iShoe`:

```
// The implementations in iShoe should different from objects.
type iShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}
```

2. Otherwise, use `struct`. Like `buildHouse()`

```
// Every house from buildHouse() is the same
func (d *director) buildHouse() house {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}
```