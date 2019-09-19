package xdean.share.spring.inject.customcondition;

import org.springframework.beans.factory.config.ConfigurableBeanFactory;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.ConfigurableApplicationContext;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Scope;

import javax.inject.Inject;

@SpringBootApplication
public class Application {
    public static void main(String[] args) {
        ConfigurableApplicationContext ctx1 = SpringApplication.run(Application.class, args);
        System.out.println(ctx1.getBean(BeanA.class));
        ctx1.close();

        WorldCondition.WORLD_ID = "Another World";

        ConfigurableApplicationContext ctx2 = SpringApplication.run(Application.class, args);
        System.out.println(ctx2.getBean(BeanA.class));
    }

    @Bean
    @OnWorld("The World")
    public static BeanA beanA1() {
        System.out.println("The world bean construct");
        return new BeanA();
    }

    @Bean
    @OnWorld("Another World")
    public static BeanA beanA2() {
        System.out.println("Another world bean construct");
        return new BeanA();
    }
}
