package xdean.share.spring.inject.lookup;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Scope;
import org.springframework.stereotype.Component;

@Component
@Scope("prototype")
public class BeanA {
    int i;
    public BeanA(int i) {
        this.i = i;
    }
    @Override
    public String toString() {
        return "BeanA{" +
                "i=" + i +
                '}';
    }
}
