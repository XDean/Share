package xdean.share.spring.jpa.audit.modify;

import lombok.Data;
import org.hibernate.envers.Audited;

import javax.persistence.*;

@Data
@Entity
@Audited(withModifiedFlag = true)
@Table(name = "users")
public class UserEntity {
  @Id
  @Column(name = "id")
  @GeneratedValue(strategy = GenerationType.IDENTITY)
  int id;

  @Column(name = "name")
  String name;

  @Column(name = "age")
  int age;
}
