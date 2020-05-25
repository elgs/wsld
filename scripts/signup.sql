-- username, email, password

set @safe_id := REPLACE(UUID(),'-','');
set @salt := SHA2(RAND(), 512);
set @now_utc := CONVERT_TZ(NOW(),'System','+0:0');

insert INTO USER SET ID=@safe_id, USERNAME=?, EMAIL=?, 
PASSWORD=ENCRYPT(?, CONCAT('\$6\$rounds=5000$',@salt)), 
STATUS=JSON_OBJECT('pfv','$pending-for-verification$'),
MODE='',TIME_CREATED=@now_utc;

SELECT EMAIL,JSON_UNQUOTE(JSON_EXTRACT(STATUS,'$.pfv')) AS V_CODE 
FROM USER WHERE ID=@safe_id;