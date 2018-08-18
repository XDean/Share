package xdean.share.memorymodel;

public class ThisEscape {
  public static void main(String[] args) {
    new ThisEscape();
  }

  public static void print(ThisEscape t) {
    System.out.println(t.a);
  }

  private final int a;

  public ThisEscape() {
    print(this);
    this.a = 100;
    print(this);
  }
}
