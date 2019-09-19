package xdean.share.spring.inject.autowired;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import xdean.share.spring.inject.autowired.Beans.*;

@Configuration
public class BeanConfiguration {

    BeanConfiguration(){

    }

    @Bean
    public BeanB b(BeanA a) {
        System.out.println("Create b");
        return new BeanB(a);
    }

    @Bean
    public BeanA a() {
        System.out.println("Create a");
        return new BeanA();
    }

    @Bean
    public BeanE e1() {
        System.out.println("Create e1");
        return new BeanE();
    }

    @Bean
    public BeanE e2() {
        System.out.println("Create e2");
        return new BeanE();
    }
}
