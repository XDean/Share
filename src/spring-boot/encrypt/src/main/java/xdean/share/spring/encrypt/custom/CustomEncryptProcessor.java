package xdean.share.spring.encrypt.custom;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.env.EnvironmentPostProcessor;
import org.springframework.core.env.ConfigurableEnvironment;
import org.springframework.core.env.EnumerablePropertySource;
import org.springframework.core.env.MapPropertySource;
import org.springframework.core.env.PropertySource;
import xdean.spring.auto.AutoSpringFactories;

import java.util.Base64;
import java.util.HashMap;

@AutoSpringFactories(EnvironmentPostProcessor.class)
public class CustomEncryptProcessor implements EnvironmentPostProcessor {
    @Override
    public void postProcessEnvironment(ConfigurableEnvironment environment, SpringApplication application) {
        HashMap<String, Object> map = new HashMap<>(); // 准备一个Map存储解密后的数据
        for (PropertySource<?> ps : environment.getPropertySources()) { // 遍历所有的PropertySource
            if (ps instanceof EnumerablePropertySource) { // 对于每一个可以遍历的PropertySource
                EnumerablePropertySource eps = (EnumerablePropertySource) ps;
                for (String name : eps.getPropertyNames()) { // 遍历所有的属性
                    Object value = eps.getProperty(name);
                    if (value instanceof String) { // 对于值是字符串的属性
                        String str = (String) value;
                        if (str.startsWith("base64:")) { // 如果以 base64: 开头
                            String decode = new String(Base64.getDecoder().decode(str.substring(7)));
                            map.put(name, decode); // 解码并放入Map里
                        }
                    }
                }
            }
        }
        PropertySource newPs = new MapPropertySource("custom-encrypt", map);
        environment.getPropertySources().addFirst(newPs); // 将解密的数据放入环境变量，并处于第一优先级上
    }
}
