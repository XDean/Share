package xdean.share.junit.param;

import static org.junit.Assert.assertEquals;

import java.util.Arrays;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.Parameterized;
import org.junit.runners.Parameterized.Parameters;

@RunWith(Parameterized.class)
public class ParameterizedDemo {
  @Parameters
  public static Iterable<Object[]> data() {
    return Arrays.asList(new Object[][] {
        { 0 },
        { 1 },
        { 2 },
        { 3 },
        { 4 }
    });
  }

  static int[] ADD = { 1, 2, 3, 4, 5 };
  static int[] SQURE = { 0, 1, 4, 9, 16 };

  private int input;

  public ParameterizedDemo(int input) {
    this.input = input;
  }

  @Test
  public void testAdd() {
    assertEquals(ADD[input], input + 1);
  }

  @Test
  public void testSqure() {
    assertEquals(SQURE[input], input * input);
  }

  @Test
  public void testOther() {
    assertEquals(ADD.length, SQURE.length);
  }
}
