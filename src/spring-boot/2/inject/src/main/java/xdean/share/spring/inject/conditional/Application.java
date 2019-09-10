package xdean.share.spring.inject.conditional;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.condition.ConditionalOnBean;
import org.springframework.boot.autoconfigure.condition.ConditionalOnClass;
import org.springframework.boot.autoconfigure.condition.ConditionalOnMissingClass;
import org.springframework.context.ConfigurableApplicationContext;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Conditional;
import org.springframework.context.annotation.Scope;

@SpringBootApplication
public class Application {
    public static void main(String[] args) {
        ConfigurableApplicationContext ctx = SpringApplication.run(Application.class, args);
        System.out.println(ctx.getBean(BeanA.class).i);
    }

    @Bean
    @ConditionalOnClass(name = "xdean.share.spring.inject.conditional.SomeDriver")
    public static BeanA beanA1() {
        return new BeanA(1);
    }

    @Bean
    @ConditionalOnBean
    @ConditionalOnMissingClass("xdean.share.spring.inject.conditional.SomeDriver")
    public static BeanA beanA2() {
        return new BeanA(2);
    }
}
