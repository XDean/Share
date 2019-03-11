package xdean.share.pieces;

import lombok.Builder;

@Builder
public class BuilderDemo {

  public static void main(String[] args) {
    builder()
        .id(1)
        .name("dean")
        .age(24)
        .male(true)
        .build();
  }

  int id;
  String name;
  int age;
  boolean male;
}
