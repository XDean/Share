package xdean.share.spring.inject.customscope;

import org.springframework.beans.factory.config.ConfigurableBeanFactory;
import org.springframework.beans.factory.config.CustomScopeConfigurer;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.ConfigurableApplicationContext;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Scope;

import javax.inject.Inject;

@SpringBootApplication
public class Application {
    public static void main(String[] args) {
        ConfigurableApplicationContext ctx = SpringApplication.run(Application.class, args);

        System.out.println(ctx.getBean(BeanA.class));
        WorldScope.WORLD_ID = "Another World";
        System.out.println(ctx.getBean(BeanA.class));
        WorldScope.WORLD_ID = "The World";
        System.out.println(ctx.getBean(BeanA.class));
    }

    @Inject
    public void config(ConfigurableBeanFactory beanFactory) {
        beanFactory.registerScope("world", new WorldScope());
    }

    @Bean
    @Scope("world")
    public static BeanA beanA() {
        return new BeanA();
    }
}
