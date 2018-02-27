package xdean.share.junit.param;

import static java.lang.annotation.ElementType.TYPE;
import static java.lang.annotation.RetentionPolicy.RUNTIME;

import java.lang.annotation.Documented;
import java.lang.annotation.Retention;
import java.lang.annotation.Target;

import org.junit.experimental.theories.Theories;
import org.junit.runner.RunWith;

import xdean.share.junit.param.RunWithThories.Template;
import xdean.annotation.Aggregation;

@Documented
@Retention(RUNTIME)
@Target(TYPE)
@Aggregation(template = Template.class)
public @interface RunWithThories {
  @RunWith(Theories.class)
  class Template {
  }
}
