package xdean.share.spring.helloweb;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class WebHandler {
    @GetMapping("/hello")
    public String helloWorld() {
        return "Hello World";
    }
}
