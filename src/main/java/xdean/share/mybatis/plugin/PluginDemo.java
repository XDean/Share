package xdean.share.mybatis.plugin;

import org.apache.ibatis.mapping.MappedStatement;
import org.apache.ibatis.plugin.Interceptor;
import org.apache.ibatis.session.Configuration;

import xdean.mybatis.extension.plugin.PluginBuilder;
import xdean.mybatis.extension.plugin.PluginEnum.ExecutorPlugin;

public class PluginDemo {
  public static void main(String[] command) {
    Configuration config = new Configuration();
    Interceptor plugin = PluginBuilder.create()
        .with(ExecutorPlugin.QUERY_2, ExecutorPlugin.UPDATE)
        .before(in -> {
          Object[] args = in.getArgs();
          MappedStatement statement = (MappedStatement) args[0];
          Object param = args[1];
          System.out.println("sql: " + statement.getBoundSql(param).getSql());
        })
        .build();
    config.addInterceptor(plugin);
  }
}
