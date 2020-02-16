package xdean.share.spring.cache.simple;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cache.annotation.*;
import org.springframework.context.ConfigurableApplicationContext;

@SpringBootApplication
@EnableCaching
public class Application {
    public static void main(String[] args) {
        ConfigurableApplicationContext ctx = SpringApplication.run(Application.class, args);
        Application app = ctx.getBean(Application.class);

        System.out.println("Test Simple Cache");
        System.out.println(app.hello("World"));
        System.out.println(app.hello("World"));
        System.out.println(app.hello("XDean"));
        System.out.println(app.hello("XDean"));

        System.out.println();
        System.out.println("Test Multiple Cache Scope");
        System.out.println(app.scope1("World"));
        System.out.println(app.scope2("World"));
        System.out.println(app.scopeBoth("World"));

        System.out.println(app.scopeBoth("XDean"));
        System.out.println(app.scope1("XDean"));
        System.out.println(app.scope2("XDean"));

        System.out.println();
        System.out.println("Test Key");
        System.out.println(app.keyAll("World", 0));
        System.out.println(app.keyAll("World", 1));
        System.out.println(app.keyWho("World", 0));
        System.out.println(app.keyWho("World", 1));

        System.out.println();
        System.out.println("Test Put");
        System.out.println(app.put("World"));
        System.out.println(app.put("World"));
        System.out.println(app.getPut("World"));

        System.out.println();
        System.out.println("Test Evict");
        System.out.println(app.getEvict("World"));
        System.out.println(app.getEvict("World"));
        app.evict("World");
        System.out.println(app.getEvict("World"));
    }

    @Cacheable(cacheNames = "test")
    public String hello(String who) {
        System.out.println("Calculating Hello: " + who);
        return "Hello " + who;
    }

    @Cacheable(cacheNames = "scope1")
    public String scope1(String who) {
        System.out.println("Calculating Scope1: " + who);
        return "Hello " + who;
    }

    @Cacheable(cacheNames = "scope2")
    public String scope2(String who) {
        System.out.println("Calculating Scope2: " + who);
        return "Hello " + who;
    }

    @Cacheable(cacheNames = {"scope2", "scope1"})
    public String scopeBoth(String who) {
        System.out.println("Calculating ScopeBoth: " + who);
        return "Hello " + who;
    }

    @Cacheable(cacheNames = "keyAll")
    public String keyAll(String who, int i) {
        System.out.println("Calculating KeyAll: " + who + ", " + i);
        return "Hello " + who;
    }

    @Cacheable(cacheNames = "keyWho", key = "#root.args[0]")
    public String keyWho(String who, int i) {
        System.out.println("Calculating KeyWho: " + who + ", " + i);
        return "Hello " + who;
    }

    @CachePut(cacheNames = "put")
    public String put(String who) {
        System.out.println("Calculating Put: " + who);
        return "Hello " + who;
    }

    @Cacheable(cacheNames = "put")
    public String getPut(String who) {
        System.out.println("Calculating GetPut: " + who);
        return "Hello " + who;
    }

    @CacheEvict(cacheNames = "evict")
    public void evict(String who) {
        System.out.println("Evict: " + who);
    }

    @Cacheable(cacheNames = "evict")
    public String getEvict(String who) {
        System.out.println("Calculating GetEvict: " + who);
        return "Hello " + who;
    }
}
