package xdean.share.spring.inject.customscope;

import org.springframework.beans.factory.ObjectFactory;
import org.springframework.beans.factory.config.Scope;

import java.util.HashMap;
import java.util.Map;

public class WorldScope implements Scope {

    public static String WORLD_ID = "The World";

    private final Map<String, Map<String, Object>> beans = new HashMap<>();

    @Override
    public Object get(String name, ObjectFactory<?> objectFactory) {
        return beans.computeIfAbsent(WORLD_ID, k -> new HashMap<>())
                .computeIfAbsent(name, k -> objectFactory.getObject());
    }

    @Override
    public Object remove(String name) {
        return beans.computeIfAbsent(WORLD_ID, k -> new HashMap<>()).remove(name);
    }

    @Override
    public void registerDestructionCallback(String name, Runnable callback) {
    }

    @Override
    public Object resolveContextualObject(String key) {
        return null;
    }

    @Override
    public String getConversationId() {
        return WORLD_ID;
    }
}
