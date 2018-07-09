package xdean.share.lambda;

import static xdean.share.lambda.Util.printLambda;
import static xdean.share.lambda.Util.startTest;

import org.junit.FixMethodOrder;
import org.junit.Test;
import org.junit.runners.MethodSorters;

@FixMethodOrder(MethodSorters.JVM)
public class TestMethodReference extends TestBase {
  @Test
  public void testRefField() throws Exception {
    startTest();
    Runnable ref = field::print;
    Runnable lam = () -> field.print();
    printLambda(ref);
    printLambda(lam);

    ref.run();
    lam.run();
    field = new Bean(100);
    ref.run();
    lam.run();
  }
}
