package xdean.share.spring.jpa.audit.auth;

import lombok.Getter;
import lombok.Setter;
import org.hibernate.envers.DefaultRevisionEntity;
import org.hibernate.envers.RevisionEntity;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Table;

@Entity
@RevisionEntity(AuthRevisionListener.class)
@Table(name = "REVINFO")
public class AuthRevisionEntity extends DefaultRevisionEntity {
  @Getter
  @Setter
  @Column(name = "username")
  private String username;
}
