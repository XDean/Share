package xdean.share.junit.param;

import static org.junit.Assert.assertEquals;

import org.junit.Test;
import org.junit.runner.RunWith;

import xdean.junit.ex.param.ParamTestRunner;
import xdean.junit.ex.param.annotation.GroupBy;
import xdean.junit.ex.param.annotation.GroupBy.Group;
import xdean.junit.ex.param.annotation.Param;
import xdean.junit.ex.param.annotation.ParamTest;

@RunWith(ParamTestRunner.class)
@GroupBy(Group.PARAM)
public class ParamDemo {
  @Param
  public static final int[] PARAM = { 1, 2, 3 };

  static int[] ADD = { 1, 2, 3, 4, 5 };
  static int[] SQURE = { 0, 1, 4, 9, 16 };

  @ParamTest
  public void testAdd(int i) {
    assertEquals(ADD[i], i + 1);
  }

  @ParamTest
  public void testSqure(int i) {
    assertEquals(SQURE[i], i * i);
  }

  @Test
  public void testOther() {
    assertEquals(ADD.length, SQURE.length);
  }
}