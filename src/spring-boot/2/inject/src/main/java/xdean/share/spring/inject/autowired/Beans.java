package xdean.share.spring.inject.autowired;

public class Beans {
    public static class BeanA {
    }

    public static class BeanB {
        BeanA a;

        public BeanB(BeanA a) {
            this.a = a;
        }
    }

    public static class BeanC {
    }

    public static class BeanD {
    }

    public static class BeanE {
    }
}
