package xdean.share.pieces;

public class TryResourceVsFinally {
  public static void main(String[] args) throws Exception {
    try (AutoCloseable a = () -> System.out.println("in-try-resource")) {
      System.out.println("in-try");
      throw new Exception();
    } catch (Exception e) {
      System.out.println("in-catch");
    } finally {
      System.out.println("in-finally");
    }
  }
}
