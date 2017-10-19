package xdean.share.mybatis.resultmap;

import java.util.List;

import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.type.JdbcType;

import xdean.mybatis.extension.model.Column;
import xdean.mybatis.extension.model.Column.ColumnBuilder;
import xdean.mybatis.extension.model.Table;
import xdean.mybatis.extension.resultmap.InitResultMap;
import xdean.share.mybatis.resultmap.model.User;

@Mapper
public interface MyBatisDemo {
  @Mapper
  InitResultMap<User> USER_MAP = InitResultMap
      .create(User.class)
      .namespace()
      .id(User.class.getName())
      .resultMap(b -> b
          .stringFree()
          .mapping(t_user.id, User::setId)
          .mapping(t_user.userName, User::setUsername)
          .mapping(t_user.password, User::setHashedPassword)
          .build())
      .build();

  List<User> getAll();
}

interface t_user {
  Table table = Table.create("t_user");
  Column id = ColumnBuilder.create().table(table).column("user_id").jdbcType(JdbcType.INTEGER).build();
  Column userName = ColumnBuilder.create().table(table).column("user_name").jdbcType(JdbcType.VARCHAR).build();
  Column password = ColumnBuilder.create().table(table).column("hashed_password").jdbcType(JdbcType.VARCHAR).build();
}