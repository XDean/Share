package xdean.share.spring.encrypt.jasypt;

import javax.crypto.Cipher;
import javax.crypto.NoSuchPaddingException;
import java.security.NoSuchAlgorithmException;

public class JasyptGenerator {

    public static void main(String[] args) throws NoSuchPaddingException, NoSuchAlgorithmException {
        Cipher c = Cipher.getInstance("PBEWITHHMACSHA512ANDAES_256");
    }
}
