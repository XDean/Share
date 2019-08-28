package xdean.share.spring.hello;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class WebHandler {
    @GetMapping
    public String helloWorld() {
        return "Hello World";
    }
}
