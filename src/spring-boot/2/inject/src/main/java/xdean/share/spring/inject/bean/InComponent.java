package xdean.share.spring.inject.bean;

import org.springframework.context.annotation.Bean;
import org.springframework.stereotype.Component;

@Component
public class InComponent {
    @Bean
    public boolean bool() {
        return true;
    }
}
