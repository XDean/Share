package xdean.share.lambda;

import static xdean.jex.util.lang.ExceptionUtil.uncheck;

import java.lang.reflect.Field;
import java.lang.reflect.Modifier;
import java.util.Arrays;

import xdean.jex.util.reflect.ReflectUtil;
import xdean.jex.util.string.StringUtil;

public interface Util {
  public static void printLambda(Object lambda) {
    Class<?> clz = lambda.getClass();
    System.out.println(lambda);
    Arrays.stream(ReflectUtil.getAllFields(clz, false))
        .forEach(
            f -> System.out.printf("\tname: %s \ttype: %s \tvalue: %s\n", f.getName(), f.getType().getSimpleName(),
                uncheck(() -> {
                  f.setAccessible(true);
                  return f.get(lambda);
                })));
  }

  public static void startTest() {
    title(ReflectUtil.getCaller(1, false).getMethodName());
  }

  public static void title(String title) {
    printLine(50);
    System.out.println("- " + title);
    printLine(50);
  }

  public static void subTitle(String title) {
    printLine(25);
    System.out.println("- " + title);
    printLine(25);
  }

  public static void printLine() {
    printLine(50);
  }

  public static void printLine(int i) {
    System.out.println(StringUtil.repeat("-", i));
  }

  public static void setFinalStatic(Field field, Object owner, Object newValue)
      throws IllegalArgumentException, IllegalAccessException {
    field.setAccessible(true);
    // remove final modifier from field
    Field modifiersField = uncheck(() -> Field.class.getDeclaredField("modifiers"));
    modifiersField.setAccessible(true);
    modifiersField.setInt(field, field.getModifiers() & ~Modifier.FINAL);
    field.set(owner, newValue);
  }
}
