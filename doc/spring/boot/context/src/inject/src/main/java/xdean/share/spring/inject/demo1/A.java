package xdean.share.spring.inject.demo1;

import org.springframework.stereotype.Component;

@Component
public class A {
    public A() {
        System.err.println("A construct");
    }
}
