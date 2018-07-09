# Hints for lambda

1. Only use input arguments as far as possible
2. Do not use method ref on non-final instance. (`this` is final)
3. Be careful of `this` reference

# Samples

## `this` escape 

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

## Capture wrong instance

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


## Avoid capture `this`

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