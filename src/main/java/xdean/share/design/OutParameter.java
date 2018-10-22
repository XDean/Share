package xdean.share.design;

import java.util.List;

public interface OutParameter {

  List<Integer> parseInts(List<String> strs, List<String> errors);

  ParseIntResult parseInts(List<String> strs);

  public class ParseIntResult {
    List<Integer> result;
    List<String> errors;
  }
}
