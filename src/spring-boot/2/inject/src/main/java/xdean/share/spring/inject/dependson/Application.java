package xdean.share.spring.inject.dependson;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.condition.ConditionalOnBean;
import org.springframework.boot.autoconfigure.condition.ConditionalOnClass;
import org.springframework.boot.autoconfigure.condition.ConditionalOnMissingClass;
import org.springframework.context.ConfigurableApplicationContext;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.DependsOn;

@SpringBootApplication
public class Application {
    public static void main(String[] args) {
        SpringApplication.run(Application.class, args);
    }

    @Bean
    @DependsOn("beanA2")
    public static BeanA beanA1() {
        System.out.println("Bean 1");
        return new BeanA(1);
    }

    @Bean
    public static BeanA beanA2() {
        System.out.println("Bean 2");
        return new BeanA(2);
    }
}
