## Hints

1. Only use method ref on final variable. (NPE in method reference will be confused in debug)
2. With rule 1, use method ref rather than lambda
3. Only use input arguments in lambda as far as possible
4. Be careful of `this` reference, in both lambda and method reference
5. With rule 1-4, use lambda rather than anonymous class

## Samples

### `this` escape 

```java
class Domain{
  StringProperty name = ...;
}

class UI {
  void init () {
    domain.name.addListener((ob, o, n) -> updateNameColumn()); // Don't do this, domain will hold this reference
  }
}
```

```java
domain.name.addListener(weak(this, (obj, o, n) -> obj.updateNameColumn())) // Do this, it only capture weak reference of this
```

### Capture wrong instance

```java
class UI {
  Session session;

  void initEvent(){
    EventBus.observe(MaskEvent.class)
            .filterSource(session::containsMask) // Don't do this, session is changeable, unless you really only care about session in this moment.
  }
  
  void bind(Session newSession){
    session = newSession;
  }
}
```

```java
.filterSource(source -> session.containsMask(source)) // Do this, it will check in current session
```


### Avoid capture `this`

```java
class GetLineHelper {
  final File file;
  final ByteBuffer buffer;
  final List<String> result;
  
  IntFunction<String> lineGetter(){
    return i -> result.get(i); // Dont' do this, it capture `this`.
  }
}
```

```java
  IntFunction<String> lineGetter(){
    return result::get; // Do this, it capture result only.
  }
```

```java
  // If it can't represent as method reference, you can let the final field be a local variable and then capture it.
  IntFunction<String> lineGetter(){
    List<String> result = this.result;
    return i -> {
      LOGGER.debug("get line " + i);
      return result.get(i);
    };
  }
```


| Previous | Next |
| --- | --- |
| [Lambda VS Method Reference](2-lambda-vs-method-reference.md) |   |