package xdean.share.lambda;

import lombok.ToString;

@ToString
public class TestBase {
  public static final Bean staticFinalField = new Bean(1);
  public static Bean staticField = new Bean(2);
  public final Bean finalField = new Bean(3);
  public Bean field = new Bean(4);
}
