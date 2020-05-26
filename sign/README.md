

## NewSignatureWithNonce

用户在配置页面完成配置并点击“提交”时，OneNET平台会向填写URL地址发送HTTP **GET**请求进行URL验证，请求形式示例如下：

```
http://url?msg=xxx&nonce=xxx&signature=xxx
```

其中，*url*为用户在页面配置时填写的URL，*nonce*、*msg*、*signature*用于URL及token的验证

token验证过程如下：

1. 将配置页面中配置的*token*与*nonce*、*msg*的值计算MD5，并且编码为Base64字符串值
2. 将上一步中Base64字符串值通过URL Decode计算后的值与请求参数*signature*的值进行对比，如果相等则表示token验证成功

如果token验证成功，返回*msg*参数值，表示URL验证通过。

```
如果用户不想验证token，可以选择跳过MD5计算过程，直接返回msg参数值
```



## NewSignatureWithTimestamp
### token 组成与算法

token由多个参数构成，如下表：

| 名称    | 类型   | 是否必须 | 参数说明                                                     | 参数示例                                                     |
| :------ | :----- | :------- | :----------------------------------------------------------- | :----------------------------------------------------------- |
| version | string | 是       | 参数组版本号，日期格式，目前仅支持"2018-10-31"               | 2018-10-31                                                   |
| res     | string | 是       | 访问资源 resource 格式为：父资源类/父资源ID/子资源类/子资源ID 见res使用场景说明 | products/123123 products/123123/devices/78329710 mqs/osndf09nand9f21390 |
| et      | int    | 是       | 访问过期时间 expirationTime，unix时间 当一次访问参数中的et时间小于当前时间时，平台会认为访问参数过期从而拒绝该访问 | 1537255523 表示：北京时间 2018-09-18 15:25:23                |
| method  | string | 是       | 签名方法 signatureMethod 支持md5、sha1、sha256               | sha256                                                       |
| sign    | string | 是       | 签名结果字符串 signature                                     |                                                              |

关于token参数的特别说明：

#### res使用场景说明

使用场景如下表：

| 场景           | res参数格式                          | 示例                          | 说明                       |
| :------------- | :----------------------------------- | :---------------------------- | :------------------------- |
| API访问        | products/{pid}                       | products/123123               |                            |
| 消息队列MQ连接 | mqs/{MQ_ID}                          | mqs/osndf09nand9f21390        | 消息队列MQ作为独立资源访问 |
| MQTTS设备连接  | products/{pid}/devices/{device_name} | products/123123/devices/mydev | 需使用设备级密钥           |

#### sign签名算法

参数sign的生成算法为：

```
sign = base64(hmac_<method>(base64decode(accessKey), utf-8(StringForSignature))) 
```

其中：

- accessKey为OneNET为独立资源（例如，产品）分配的唯一访问密钥，其作为签名算法参数之一参与签名计算，为保证访问安全，请妥善保管
- accessKey参与计算前应先进行base64decode操作
- 用于计算签名的字符串 StringForSignature的组成顺序按照**参数名称进行字符串排序**，以'/n'作为参数分隔，当前版本中按照如下顺序进行排序：et、method、res、version

StringForSignature组成示例如下：

```
StringForSignature = et + '\n' + method + '\n' + res+ '\n' + version
```

注意：每个参数均为key=value格式组成，但是仅参数中的value参与计算签名的字符串 StringForSignature的组成，若token参数如下

```
et = 1537255523
method = sha1
res = products/123123
version = 2018-10-31
```

则用于计算签名的字符串StringForSignature为（按照et、method、res、version的顺序）

```
StringForSignature = "1537255523" + "\n" + "sha1"+ "\n" + "products/123123"+ "\n" + "2018-10-31"
```

计算出sign后，将每个参数均采用key=value的形式表示，并用'&'作为分隔符，示例如下：

```
version=2018-10-31&res=products/123123&et=1537255523&method=sha1&sign=ZjA1NzZlMmMxYzIOTg3MjBzNjYTI2MjA4Yw=
```

#### 参数编码

token中key=value的形式的value部分需要经过**URL编码**，需要进行编码的特殊符号如下：

| 序号 | 符号 | 编码 |
| :--: | :--: | :--: |
|  1   |  +   | %2B  |
|  2   | 空格 | %20  |
|  3   |  /   | %2F  |
|  4   |  ?   | %3F  |
|  5   |  %   | %25  |
|  6   |  #   | %23  |
|  7   |  &   | %26  |
|  8   |  =   | %3D  |

编码后，上例中实际传输token为：

```
version=2018-10-31&res=products%2F123123&et=1537255523&method=sha1&sign=ZjA1NzZlMmMxYzIOTg3MjBzNjYTI2MjA4Yw%3D
```

#### token生成示例

python代码示例

```python
import base64
import hmac
import time
from urllib.parse import quote

def token(id,access_key):

    version = '2018-10-31'

    res = 'mqs/%s' % id  # 通过MQ_ID访问MQ
    # res = 'products/%s' % id  # 通过产品ID访问产品API

    # 用户自定义token过期时间
    et = str(int(time.time()) + 3600)

    # 签名方法，支持md5、sha1、sha256
    method = 'sha1'

    # 对access_key进行decode
    key = base64.b64decode(access_key)

    # 计算sign
    org = et + '\n' + method + '\n' + res + '\n' + version
    sign_b = hmac.new(key=key, msg=org.encode(), digestmod=method)
    sign = base64.b64encode(sign_b.digest()).decode()

    # value 部分进行url编码，method/res/version值较为简单无需编码
    sign = quote(sign, safe='')
    res = quote(res, safe='')

    # token参数拼接
    token = 'version=%s&res=%s&et=%s&method=%s&sign=%s' % (version, res, et, method, sign)

    return token


if __name__ == '__main__':
    id = 'A1EB10110CFA9E06D6209E40C4A6D7976'
    access_key = 'KuF3NT/jUBJ62LNBB/A8XZA9CqS3Cu79B/ABmfA1UCw='

    print(token(id,access_key))
```

Java代码示例

```java
import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import java.io.UnsupportedEncodingException;
import java.net.URLEncoder;
import java.security.InvalidKeyException;
import java.security.NoSuchAlgorithmException;
import java.util.Base64;


public class Token {

    public static String assembleToken(String version, String resourceName, String expirationTime, String signatureMethod, String accessKey)
            throws UnsupportedEncodingException, NoSuchAlgorithmException, InvalidKeyException {
        StringBuilder sb = new StringBuilder();
        String res = URLEncoder.encode(resourceName, "UTF-8");
        String sig = URLEncoder.encode(generatorSignature(version, resourceName, expirationTime
                , accessKey, signatureMethod), "UTF-8");
        sb.append("version=")
                .append(version)
                .append("&res=")
                .append(res)
                .append("&et=")
                .append(expirationTime)
                .append("&method=")
                .append(signatureMethod)
                .append("&sign=")
                .append(sig);
        return sb.toString();
    }

    public static String generatorSignature(String version, String resourceName, String expirationTime, String accessKey, String signatureMethod) 
            throws NoSuchAlgorithmException, InvalidKeyException {
        String encryptText = expirationTime + "\n" + signatureMethod + "\n" + resourceName + "\n" + version;
        String signature;
        byte[] bytes = HmacEncrypt(encryptText, accessKey, signatureMethod);
        signature = Base64.getEncoder().encodeToString(bytes);
        return signature;
    }

    public static byte[] HmacEncrypt(String data, String key, String signatureMethod)
            throws NoSuchAlgorithmException, InvalidKeyException {
        //根据给定的字节数组构造一个密钥,第二参数指定一个密钥算法的名称
        SecretKeySpec signinKey = null;
        signinKey = new SecretKeySpec(Base64.getDecoder().decode(key),
                "Hmac" + signatureMethod.toUpperCase());

        //生成一个指定 Mac 算法 的 Mac 对象
        Mac mac = null;
        mac = Mac.getInstance("Hmac" + signatureMethod.toUpperCase());

        //用给定密钥初始化 Mac 对象
        mac.init(signinKey);

        //完成 Mac 操作
        return mac.doFinal(data.getBytes());
    }

    public enum SignatureMethod {
        SHA1, MD5, SHA256;
    }

    public static void main(String[] args) throws UnsupportedEncodingException, NoSuchAlgorithmException, InvalidKeyException {
        String version = "2018-10-31";
        String resourceName = "mqs/A1EB10110CFA9E06D6209E40C4A6D7976";
        String expirationTime = System.currentTimeMillis() / 1000 + 100 * 24 * 60 * 60 + "";
        String signatureMethod = SignatureMethod.SHA1.name().toLowerCase();
        String accessKey = "KuF3NT/jUBJ62LNBB/A8XZA9CqS3Cu79B/ABmfA1UCw=";
        String token = assembleToken(version, resourceName, expirationTime, signatureMethod, accessKey);
        System.out.println("Authorization:" + token);
    }
}
```