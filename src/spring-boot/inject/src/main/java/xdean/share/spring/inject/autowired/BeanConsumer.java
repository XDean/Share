package xdean.share.spring.inject.autowired;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.stereotype.Component;
import xdean.share.spring.inject.autowired.Beans.*;

import javax.inject.Provider;
import java.util.List;
import java.util.Map;
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
    private void inject(Provider<BeanB> b, Optional<BeanD> d, List<BeanE> list, Map<String, BeanE> map) {
        System.out.println("Inject with b provider: " + b);
        System.out.println("Inject with b: " + b.get());
        System.out.println("Inject with d: " + d);
        System.out.println("Inject with list: " + list);
        System.out.println("Inject with map: " + map);
        System.out.println("Now b:" + b);
        System.out.println("Now c:" + c);
    }

    @Autowired
    public void injectFail(BeanD d) {
        System.out.println("never happen");
    }
}
