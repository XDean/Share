package xdean.share.spring.inject.lookup;

import org.springframework.beans.factory.annotation.Lookup;
import org.springframework.context.annotation.Configuration;
import org.springframework.stereotype.Component;

@Configuration
public abstract class UseLookUp {
    @Lookup
    public abstract BeanA beanA();
}
