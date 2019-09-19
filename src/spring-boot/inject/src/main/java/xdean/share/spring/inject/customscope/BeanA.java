package xdean.share.spring.inject.customscope;

public class BeanA {
    String world = WorldScope.WORLD_ID;

    @Override
    public String toString() {
        return "BeanA{world='" + world + '\'' + '}';
    }
}
