package xdean.share.spring.inject.importoutside;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.ConfigurableApplicationContext;
import org.springframework.context.annotation.Import;
import xdean.share.spring.inject.outside.Outside;

@SpringBootApplication
@Import(Outside.class)
public class Application {
    public static void main(String[] args) {
        ConfigurableApplicationContext ctx = SpringApplication.run(Application.class, args);
        System.out.println(ctx.getBean(Outside.class));
    }
}
