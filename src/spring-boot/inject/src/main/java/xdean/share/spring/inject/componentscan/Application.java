package xdean.share.spring.inject.componentscan;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.ConfigurableApplicationContext;
import xdean.share.spring.inject.outside.Outside;

@SpringBootApplication(scanBasePackages = "xdean.share.spring.inject.outside")
public class Application {
    public static void main(String[] args) {
        ConfigurableApplicationContext ctx = SpringApplication.run(Application.class, args);
        System.out.println(ctx.getBean(Outside.class));
    }
}
