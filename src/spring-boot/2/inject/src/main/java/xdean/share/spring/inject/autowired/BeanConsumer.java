package xdean.share.spring.inject.autowired;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.stereotype.Component;
import xdean.share.spring.inject.autowired.Beans.*;

import java.util.List;
import java.util.Optional;

@Component
public class BeanConsumer {
    @Autowired
    BeanB b;

    @Autowired(required = false)
    private BeanC c;

    @Autowired
    public BeanConsumer(BeanA a) {
        System.out.println("Construct with: " + a);
    }

    @Autowired
    private void inject(Optional<BeanD> d, List<BeanE> es) {
        System.out.println("Inject with d: " + d);
        System.out.println("Inject with es: " + es);
        System.out.println("Now b:" + b);
        System.out.println("Now c:" + c);
    }
}
