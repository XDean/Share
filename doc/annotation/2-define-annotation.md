# Define Annotation

```java
@interface① MyAnno② {
  int③ value④();
  
  String comment() default ""⑤;
}
```

- ① Use `@interface` to indicate this is an annotation type
    - In fact it's `@` and `interface`. So it's just a (special) interface.
    - It implies `implements java.lang.annotation.Annotation`
- ② Annotation type name
- ③ Attribute type, which must be following type
    - A primitive type
    - String
    - Class or an invocation of Class
    - An enum type
    - An annotation type (No self reference)
    - An array type whose component type is one of the preceding types
- ④ Attribute name
    - `value` is a special name that you can declared it without `value=`
- ⑤ Attribute can have a default value
    - Must be compile period constant
    
# Predefined Annotations

## Meta Annotations

- `@Target`, define legal target element types of the annotations
    - See `java.lang.annotation.ElementType`
    - If `@Target` is not defined, it equals defined wtih all values except `TYPE_PARAMETER` and `TYPE_USE`
- `@Retention`, define the annotation's retention, by default it's `CLASS`
    - `SOURCE`
    - `CLASS`
    - `RUNTIME`
- `@Inherited`
- `@Repeatable`
- `@Documented`

## Other Annotations

- `@Override`
- `@SuppressWarnings`
- `@Deprecated`
- `@SafeVarargs`
- `@FunctionalInterface`