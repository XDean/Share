package xdean.share.spring.inject.lookup;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.ConfigurableApplicationContext;

@SpringBootApplication
public class Application {
    public static void main(String[] args) {
        ConfigurableApplicationContext ctx = SpringApplication.run(Application.class, args);
        System.err.println(ctx.getBean(UseLookUp.class));
        System.err.println(ctx.getBean(UseLookUp.class).beanA(1));
        System.err.println(ctx.getBean(UseLookUp.class).beanA(2));
    }
}
