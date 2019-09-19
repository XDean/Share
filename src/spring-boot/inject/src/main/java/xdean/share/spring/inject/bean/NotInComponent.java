package xdean.share.spring.inject.bean;

import org.springframework.context.annotation.Bean;

public class NotInComponent {
    @Bean
    public double aDouble() {
        return 42.0;
    }
}
