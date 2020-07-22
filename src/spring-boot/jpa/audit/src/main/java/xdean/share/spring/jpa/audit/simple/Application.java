package xdean.share.spring.jpa.audit.simple;

import org.springframework.beans.factory.InitializingBean;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class Application implements InitializingBean {
  public static void main(String[] args) {
    SpringApplication.run(Application.class, args);
  }

  @Autowired UserRepo repo;

  @Override
  public void afterPropertiesSet() throws Exception {
    UserEntity user = new UserEntity();
    user.setName("Dean");
    user.setAge(25);
    user = repo.save(user);

    user.setAge(26);
    user = repo.save(user);

    repo.delete(user);
  }
}
