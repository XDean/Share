package xdean.share.spring.inject.component;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.ConfigurableApplicationContext;
import xdean.share.spring.inject.outside.Outside;

@SpringBootApplication
public class Application {
    public static void main(String[] args) {
        ConfigurableApplicationContext ctx = SpringApplication.run(Application.class, args);

        System.err.println(ctx.getBean(Application.class));
        System.err.println(ctx.getBean("application"));
        System.err.println(ctx.getBean(BeanNoName.class));
        System.err.println(ctx.getBean("beanNoName"));
        System.err.println(ctx.getBean(BeanHasName.class));
        System.err.println(ctx.getBean("New name"));
        System.err.println(ctx.getBean("beanHasName"));

        System.err.println(ctx.getBean(Outside.class));
    }

    public Application() {
        System.err.println("Application construct");
    }
}
