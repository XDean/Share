package xdean.share.spring.inject.customcondition;

import org.springframework.beans.factory.ObjectFactory;
import org.springframework.beans.factory.config.Scope;
import org.springframework.context.annotation.Condition;
import org.springframework.context.annotation.ConditionContext;
import org.springframework.core.type.AnnotatedTypeMetadata;

import java.util.HashMap;
import java.util.Map;

public class WorldCondition implements Condition {

    static String WORLD_ID = "The World";

    @Override
    public boolean matches(ConditionContext context, AnnotatedTypeMetadata metadata) {
        String world = metadata.getAnnotationAttributes(OnWorld.class.getName()).get("value").toString();
        return WORLD_ID.equals(world);
    }
}
