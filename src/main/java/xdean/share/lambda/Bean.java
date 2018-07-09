package xdean.share.lambda;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class Bean {
  int i = 0;

  public void print() {
    System.out.println("print: " + this);
  }

  public void print(Object o) {
    System.out.println("print: " + o + "\t" + this);
  }
}
