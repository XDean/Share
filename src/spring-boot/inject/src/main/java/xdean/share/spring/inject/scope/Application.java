package xdean.share.spring.inject.scope;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.ConfigurableApplicationContext;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Scope;

@SpringBootApplication
public class Application {
    public static void main(String[] args) {
        ConfigurableApplicationContext ctx = SpringApplication.run(Application.class, args);

        System.out.println(ctx.getBean(BeanA.class));
        System.out.println(ctx.getBean("beanA"));

        System.out.println(ctx.getBean(CompA.class));
        System.out.println(ctx.getBean("compA"));
    }

    @Bean
    @Scope("prototype")
    public static BeanA beanA() {
        return new BeanA();
    }
}
