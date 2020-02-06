package xdean.share.spring.encrypt.custom;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.env.EnvironmentPostProcessor;
import org.springframework.core.env.ConfigurableEnvironment;
import org.springframework.core.env.EnumerablePropertySource;
import org.springframework.core.env.MapPropertySource;
import org.springframework.core.env.PropertySource;
import org.springframework.stereotype.Component;
import xdean.spring.auto.AutoSpringFactories;

import java.util.Base64;
import java.util.HashMap;
import java.util.LinkedHashMap;
import java.util.Map;

@AutoSpringFactories(EnvironmentPostProcessor.class)
public class CustomEncryptProcessor implements EnvironmentPostProcessor {
    @Override
    public void postProcessEnvironment(ConfigurableEnvironment environment, SpringApplication application) {
        HashMap<String, Object> map = new HashMap<>();
        for (PropertySource<?> ps : environment.getPropertySources()) {
            if (ps instanceof EnumerablePropertySource) {
                EnumerablePropertySource eps = (EnumerablePropertySource) ps;
                for (String name : eps.getPropertyNames()) {
                    Object value = eps.getProperty(name);
                    if (value instanceof String) {
                        String str = (String) value;
                        if (str.startsWith("base64:")) {
                            String decode = new String(Base64.getDecoder().decode(str.substring(7)));
                            map.put(name, decode);
                        }
                    }
                }
            }
        }
        PropertySource newPs = new MapPropertySource("custom-encrypt", map);
        environment.getPropertySources().addFirst(newPs);
    }
}
