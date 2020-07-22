package xdean.share.spring.jpa.audit.auth;

import org.hibernate.envers.RevisionListener;
import org.springframework.stereotype.Component;

@Component
public class AuthRevisionListener implements RevisionListener {
  public static String user = "foo";

  @Override
  public void newRevision(Object revisionEntity) {
    AuthRevisionEntity e = (AuthRevisionEntity) revisionEntity;
    // If you are using spring-security, you can set it as authentication
    // Authentication auth = SecurityContextHolder.getContext().getAuthentication();
    e.setUsername(user);
  }
}
