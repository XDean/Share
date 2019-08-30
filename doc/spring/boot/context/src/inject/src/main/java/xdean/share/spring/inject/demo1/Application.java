package xdean.share.spring.inject.demo1;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.ConfigurableApplicationContext;
import xdean.share.spring.inject.outside.Outside;

@SpringBootApplication
public class Application {
    public static void main(String[] args) {
        ConfigurableApplicationContext ctx = SpringApplication.run(Application.class, args);

        System.err.println(ctx.getBean(A.class));
        System.err.println(ctx.getBean("a"));
        System.err.println(ctx.getBean(B.class));
        System.err.println(ctx.getBean("New B"));

        System.err.println(ctx.getBean(Outside.class));
    }
}
