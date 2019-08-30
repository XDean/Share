package xdean.share.spring.inject.outside;

import org.springframework.stereotype.Component;

@Component
public class Outside {
    public Outside(){
        System.err.println("Outside construct");
    }
}
