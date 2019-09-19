package xdean.share.spring.inject.component;

import org.springframework.stereotype.Component;

@Component("New name")
public class BeanHasName {
    public BeanHasName() {
        System.err.println("BeanHasName construct");
    }
}
