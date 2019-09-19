package xdean.share.spring.inject.component;

import org.springframework.stereotype.Component;

@Component
public class BeanNoName {
    public BeanNoName() {
        System.err.println("BeanNoName construct");
    }
}
