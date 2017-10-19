package xdean.share.mybatis.resultmap.model;

//t_user
public class User {
  private int id; // user_id
  private String username;// user_name
  private String hashedPassword;// hashed_password

  public int getId() {
    return id;
  }

  public void setId(int id) {
    this.id = id;
  }

  public String getUsername() {
    return username;
  }

  public void setUsername(String username) {
    this.username = username;
  }

  public String getHashedPassword() {
    return hashedPassword;
  }

  public void setHashedPassword(String hashedPassword) {
    this.hashedPassword = hashedPassword;
  }
}