package xdean.share.annotation;

import java.util.List;

import org.springframework.boot.autoconfigure.EnableAutoConfiguration;

import lombok.Builder;
import lombok.Singular;
import lombok.SneakyThrows;
import lombok.Value;
import xdean.auto.message.AutoMessage;
import xdean.spring.auto.AutoSpringFactories;

@AutoSpringFactories(EnableAutoConfiguration.class) // check your /target/classes
@AutoMessage(path = "messages.properties") // check your /target/generated-sources
public class PlayAnnotation {

  // @Test // comment out this line
  @SneakyThrows
  public int test() {
    Bean.builder()
        .name("PFM")
        .alias("PWO")
        .alias("PFC")
        .alias("PFM")
        .build();
    return 0;
  }

  @Value
  @Builder
  static class Bean {
    String name;
    @Singular
    List<String> aliases;
  }
}
