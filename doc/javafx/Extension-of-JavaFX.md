# `PropertyEX`

- `addTransformer`, transform the coming value before actual `setValue`.
- `addVerifier`, verify the coming (transformed) value before actual `setValue`. If verify failed, the value will be rejected.
- `in`, constraint the value only be in the given `ObservableList`
- `defaultForNull`, set to default value when null coming.
- `nonNull`, reject null value.
- `setWhile`, set the value only in a period of time.
- `is`, create `BooleanBinding` by a `Predicate`.
- `softBind`/`weakBind`, bind the `Property` without be 'bound', also see javadoc.
- `getSafe`/`orElse`, get value as `Optional` conveniently.
- `convert`/`convertToString`, See `BeanConvertUtil`.
- `addListenerAndInvoke`, add listener and invoke it immediately to reduce duplicate code.
- `on`, convenient method to add listener, See `ListenerUtil.on`.
- `XXXPropertyEX.normalize`, convert `XXXPropertyEX` to `XXXProperty`.
- `BooleanPropertyEX.and/or`, constraint the property always be false(true) if the observable is false(true).
- `DoublePropertyEX.precision`, set precision of value.
- `StringPropertyEX.emptryForNull`, set to empty string when null comes.
- `MapPropertyEX`
  - `propertyAt`, get property reflect the value of given key in this map.
  - `keyIn`/`valueIn`, constraint the maps' key/value in the `ObservableList`.
  - `bijection`, constraint the map as bijection.

# `BeanUtil`

## Select `Property`(`ObservableValue`) in `ObservableValue`

Use `nestProp`(`BeanUtil.nestValue`) to select a `Property`(`ObservableValue`) from a `ObservableValue`'s value. Any change of the returned property will appear to the origin property. Vice versa. e.g.

```java
class Owner {
  Property<String> name = new SimpleStringProperty();
}
Owner a = new Owner();
a.name.setValue("a1");
Owner b = new Owner();
b.name.setValue("b1");
Property<Owner> owner = new SimpleObjectProperty<>(a);
Property<String> currentName = nestProp(owner, o -> o.name); // "a1"
a.name.setValue("a2"); // "a2"
b.name.setValue("b2"); // "a2"
owner.setValue(b); // "b2"
a.name.setValue("a3"); // "b2"
b.name.setValue("b3"); // "b3"
```

## Map `ObservableValue` by `Function`

Use `map` to create new `ObservableValue` which always calculated by a given `Function`. e.g.

```java
Property<Integer> p = new SimpleIntegerProperty(0);
ObservableValue<Integer> ov = map(p, i -> i + 1); // ov = 1
p.setValue(1); // ov = 2
```

# `BeanConvertUtil`

- `toXXX`, convert `Property<XXX>` to `XXXProperty`.
- `toXXXBinding`, convert `ObservableValue<XXX>` to `XXXBinding`.
- `convert`, convert `Property<F>` to `Property<T>` with forward and backward converters.
- `convertList`, convert `ObservableList<F>` to `ObservableList<T>` with forward and backward converters.

# `BindingUtil`

- `createListBinding`/`createMapBinding`, create that binding.
- `observeProperty`, observe list items' property.
- `nestBind`, bind nested property (Use `BeanUtil.nestProp(...).bind` will not actually bind the original `Property`).
- `autoValid`, valid the binding when it goes to invalid. See its API note.

# `ListenerUtil`

- `weak`, create listener only hold a weak reference of the object.
- `on`, create listener to do the action when meet the condition.
- `list`, create `ListChangeListener` with fluent API.
- `set`, create `SetChangeListener` with fluent API.
- Note that the returned `ChangeListenerEX` also have fluent API.