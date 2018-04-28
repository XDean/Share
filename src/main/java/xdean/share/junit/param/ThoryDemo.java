package xdean.share.junit.param;

import static org.junit.Assert.assertEquals;

import org.junit.Test;
import org.junit.experimental.theories.DataPoints;
import org.junit.experimental.theories.Theories;
import org.junit.experimental.theories.Theory;
import org.junit.runner.JUnitCore;
import org.junit.runner.Result;
import org.junit.runner.RunWith;

@RunWith(Theories.class)
public class ThoryDemo {
  
  public static void main(String[] args) {
    Result result = JUnitCore.runClasses(ThoryDemo.class);
    System.out.println(result.getRunCount());
    System.out.println(result.getFailureCount());
  }
  
  @DataPoints
  public static final int[] PARAM = { 1, 2, 3 };

  static int[] ADD = { 1, 2, 3, 4, 5 };
  static int[] SQURE = { 0, 1, 4, 9, 16 };

  @Theory
  public void testAdd(int i) {
    assertEquals(ADD[i], i + 1);
  }

  @Theory
  public void testSqure(int i) {
    assertEquals(SQURE[i], i * i);
  }

  @Test
  public void testOther() {
    assertEquals(ADD.length, SQURE.length);
  }
}
