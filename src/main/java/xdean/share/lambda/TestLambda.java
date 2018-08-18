package xdean.share.lambda;

import static xdean.share.lambda.Util.printLambda;
import static xdean.share.lambda.Util.setFinalStatic;
import static xdean.share.lambda.Util.startTest;
import static xdean.share.lambda.Util.subTitle;

import java.util.function.Supplier;

import org.junit.FixMethodOrder;
import org.junit.Test;
import org.junit.runners.MethodSorters;

@FixMethodOrder(MethodSorters.JVM)
public class TestLambda extends TestBase {

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
    setFinalStatic(this.getClass().getField("finalField"), this, new Bean(1000));
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

    field.i = 100;
    printLambda(r);

    subTitle("change field value");

    field = new Bean(200);
    printLambda(r);
  }

  @Test
  public void testCaptureFinalFieldAsLocal() throws Exception {
    Bean local = finalField;
    startTest();
    Supplier<Object> r = () -> local;
    printLambda(r);
  }
}
