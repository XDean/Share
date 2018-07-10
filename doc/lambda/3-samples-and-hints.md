## Hints

1. Only use method ref on final variable.
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
    domain.name.addListener((ob, o, n) -> updateNameColumn());
  }
}
```

```java
domain.name.addListener(weak(this, (obj, o, n) -> obj.updateNameColumn()))
```

### Capture wrong instance

```java
class UI {
  Session session;

  void initEvent(){
    EventBus.observe(MaskEvent.class)
            .filterSource(session::containsMask)
  }
  
  void bind(Session newSession){
    session = newSession;
  }
}
```

```java
.filterSource(source -> session.containsMask(source))
```


### Avoid capture `this`

```java
class GetLineHelper {
  final File file;
  final ByteBuffer buffer;
  final List<String> result;
  
  IntFunction<String> lineGetter(){
    return i -> result.get(i);
  }
}
```

```java
  IntFunction<String> lineGetter(){
    return result::get;
  }
```

```java
  IntFunction<String> lineGetter(){
    List<String> result = this.result;
    return i -> {
      LOGGER.debug("get line " + i);
      return result.get(i);
    };
  }
```