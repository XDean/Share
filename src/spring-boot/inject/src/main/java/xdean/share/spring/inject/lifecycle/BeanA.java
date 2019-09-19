package xdean.share.spring.inject.lifecycle;

import org.springframework.beans.factory.DisposableBean;
import org.springframework.beans.factory.InitializingBean;

import javax.annotation.PostConstruct;
import javax.annotation.PreDestroy;

public class BeanA implements InitializingBean, DisposableBean {
    @Override
    public void afterPropertiesSet() throws Exception {
        System.out.println("InitializingBean");
    }

    @Override
    public void destroy() throws Exception {
        System.out.println("DisposableBean");
    }

    @PostConstruct
    private void postConstruct(){
        System.out.println("PostConstruct");
    }

    @PreDestroy
    private void preDestroy(){
        System.out.println("PreDestroy");
    }

    public void close(){
        System.out.println("Inferred Close");
    }
}
