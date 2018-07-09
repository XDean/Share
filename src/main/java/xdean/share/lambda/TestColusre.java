package xdean.share.lambda;

import static xdean.share.lambda.Util.*;
import static xdean.share.lambda.Util.startTest;

import java.util.function.Supplier;

import org.junit.FixMethodOrder;
import org.junit.Test;
import org.junit.runners.MethodSorters;

import lombok.ToString;

@ToString
@FixMethodOrder(MethodSorters.JVM)
public class TestColusre {

  static final Bean staticFinalField = new Bean(1);
  static Bean staticField = new Bean(2);
  final Bean finalField = new Bean(3);
  Bean field = new Bean(4);

  @Test
  public void testCaptureField() throws Exception {
    startTest();
    Supplier<Object> r = () -> field;
    printLambda(r);
  }

  @Test
  public void testCaptureFinalField() throws Exception {
    startTest();
    Supplier<Object> r = () -> finalField;
    printLambda(r);
  }

  @Test
  public void testCaptureStaticField() throws Exception {
    startTest();
    Supplier<Object> r = () -> staticField;
    printLambda(r);
  }

  @Test
  public void testCaptureStaticFinalField() throws Exception {
    startTest();
    Supplier<Object> r = () -> staticFinalField;
    printLambda(r);
  }

  @Test
  public void testCaptureFinalFieldModifyByReflect() throws Exception {
    startTest();
    Supplier<Object> r = () -> finalField;
    printLambda(r);
    setFinalStatic(this.getClass().getDeclaredField("finalField"), this, new Bean(1000));
    printLambda(r);
  }

  @Test
  public void testCaptureLocalVar() throws Exception {
    startTest();
    int localVar = 0;
    Supplier<Object> r = () -> localVar;
    printLambda(r);
  }

  @Test
  public void testCaptureLocalField() throws Exception {
    startTest();
    Bean localField = field;
    Supplier<Object> r = () -> localField;
    printLambda(r);

    subTitle("change local field data");

    field.i = 1;
    printLambda(r);

    subTitle("change local field value");

    field = new Bean(2);
    printLambda(r);
  }
}
