package xdean.share.spring.inject.customcondition;

import org.springframework.context.annotation.Conditional;

import java.lang.annotation.*;

@Target({ ElementType.TYPE, ElementType.METHOD })
@Retention(RetentionPolicy.RUNTIME)
@Conditional(WorldCondition.class)
public @interface OnWorld {
    String value();
}
