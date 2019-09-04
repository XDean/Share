package xdean.share.spring.inject.bean;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.ConfigurableApplicationContext;
import org.springframework.context.annotation.Bean;

@SpringBootApplication
public class Application {
    public static void main(String[] args) {
        ConfigurableApplicationContext ctx = SpringApplication.run(Application.class, args);

        System.out.println(ctx.getBean(String.class));
        System.out.println(ctx.getBean("string"));

        System.out.println(ctx.getBean(int.class));
        System.out.println(ctx.getBean("a int"));

        System.out.println(ctx.getBean(boolean.class));
        System.out.println(ctx.getBean("bool"));

        System.out.println(ctx.getBean(double.class));
    }

    @Bean
    public static String string() {
        return "a string";
    }

    @Bean("a int")
    public int intBean() {
        return 1024;
    }
}
