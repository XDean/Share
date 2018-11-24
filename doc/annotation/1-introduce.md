# All about Annotation

*What's annotation?*

Meta Data of Java Elements.

---
*How to define annotation?*

Use `@interface` keyword to define annotation with [0, n) attributes.

---
*How to use annotation?*

Read via reflection at runtime OR Read via annotation processing framework at compile time.

---
*What's the usually usage of annotation?*

- Provide meta information for runtime. e.g. JUnit `@Test`, Spring Boot `@Bean`
- Code generation. e.g. Lombok `@Value`, Google Auto `@AutoService`
- Compile time check. e.g. JDK `@Override`, Deannotation `@CheckXXX`