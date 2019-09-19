package xdean.share.spring.inject.customcondition;

public class BeanA {
    String world = WorldCondition.WORLD_ID;

    @Override
    public String toString() {
        return "BeanA{world='" + world + '\'' + '}';
    }
}
