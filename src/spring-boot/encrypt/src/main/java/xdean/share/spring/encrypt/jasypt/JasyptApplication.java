package xdean.share.spring.encrypt.jasypt;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class JasyptApplication {
    public static void main(String[] args) {
        args = new String[]{"--jasypt.encryptor.password=xdean"}; // Mock从命令行传入
        SpringApplication.run(JasyptApplication.class, args);
    }

    @Autowired
    public void init(@Value("${jasypt-data}") String data) {
        System.out.println(data);
    }
}
