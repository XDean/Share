package xdean.share.spring.inject.demo1;

import org.springframework.stereotype.Component;

@Component("New B")
public class B {
    public B() {
        System.err.println("B construct");
    }
}
